package app

import (
	"context"
	"errors"
	"flag"
	"fmt"

	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/e4t4g/URL_shortener_GB-/cmd/app/config"
	"github.com/e4t4g/URL_shortener_GB-/cmd/app/logger"
	"github.com/e4t4g/URL_shortener_GB-/internal/app/delivery/ginrouter"
	"github.com/e4t4g/URL_shortener_GB-/internal/app/repository/sqlite"
	"github.com/e4t4g/URL_shortener_GB-/internal/app/usecase"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/mxk/go-sqlite/sqlite3"
)

type URLData struct {
	ID       int    `json:"id" yaml:"id"`
	FullURL  string `json:"full_url" yaml:"full_url"`
	ShortURL string `json:"short_url" yaml:"short_url"`
	Counter  int64  `json:"counter" yaml:"counter"`
}

func App() {
	gin.SetMode(gin.ReleaseMode)

	sugar := logger.NewLogger()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	router := gin.Default()

	flag.Parse()

	cfg := config.Config{}
	if err := cfg.ReadFromFile(sugar); err != nil {
		cfg.REadFromEnv(sugar)
	}

	db, err := sqlx.Open("sqlite3", cfg.DBconfig.DBurl)
	if err != nil {
		sugar.Fatalf("can not connect to DB: %s", err)
	}

	defer func(db *sqlx.DB) {
		err = db.Close()
		if err != nil {
			sugar.Fatalf("can not close DB: %s", err)
		}
	}(db)

	router.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("[%s] %s ",
			params.TimeStamp.Format(time.RFC3339),
			params.Request.Proto,
		)
	}))
	router.StaticFile("/favicon.ico", "./favicon.ico")

	repository := sqlite.New(db, sugar)
	businessLogic := usecase.New(repository, sugar)
	deliveryLayer := ginrouter.New(businessLogic, sugar)

	router.POST("/linkUrl", deliveryLayer.Create())
	router.GET("/stat/:id", deliveryLayer.GetStat())
	router.GET("/:token", deliveryLayer.Redirect())
	router.GET("/", deliveryLayer.WebGenerate())

	srv := &http.Server{Addr: fmt.Sprintf(":%d", cfg.Port), Handler: router}

	sugar.Info("[shutting down gracefully, press Ctrl+C again to force]")

	go func(ctx context.Context, srv *http.Server) {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err = srv.Shutdown(ctx); err != nil {
			sugar.Fatalf("Server forced to shutdown: %s", err)
		}
	}(ctx, srv)

	sugar.Infof("[[Listening and serving HTTP on: %d]]", cfg.Port)

	if err = srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		sugar.Fatalf("listen: %s\n", err.Error())
	}
	sugar.Info("Server exiting")
}
