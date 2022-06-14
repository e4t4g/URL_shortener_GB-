package app

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/e4t4g/URL_shortener_GB-/cmd/app/config"
	"github.com/e4t4g/URL_shortener_GB-/internal/app/delivery/ginrouter"
	"github.com/e4t4g/URL_shortener_GB-/internal/app/usecase"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestApp(t *testing.T) {
	router := setupRouter(t)

	w := httptest.NewRecorder()
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "/stat/1", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func setupRouter(t *testing.T) *gin.Engine {
	router := gin.Default()
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	sugar := logger.Sugar()
	defer func(logger *zap.Logger) {
		err = logger.Sync()
		if err != nil {
			sugar.Infof("%s", err)
		}
	}(logger)
	cfg := config.Config{}
	if err = cfg.ReadFromFile(sugar); err != nil {
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

	URL := &usecase.URLData{ID: 1, Counter: 5}

	businessLogic := usecase.NewUseCaseMock(t)
	businessLogic.GetStatMock.Return(URL, nil)
	deliveryLayer := ginrouter.New(businessLogic, sugar)

	router.GET("/stat/:id", deliveryLayer.GetStat())
	return router
}
