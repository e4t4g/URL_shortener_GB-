package ginrouter

import (
	"context"
	"github.com/e4t4g/URL_shortener_GB-/cmd/app/config"
	"github.com/e4t4g/URL_shortener_GB-/internal/app/repository/sqlite"
	"github.com/e4t4g/URL_shortener_GB-/internal/app/usecase"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/require"
)

func TestDelivery_Create(t *testing.T) {

}

func TestDelivery_GetStat(t *testing.T) {


	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	sugar := logger.Sugar()

	cfg := config.Config{}
	err = cfg.ReadFromFile(sugar)
	if err != nil {
		sugar.Fatalf("can not read config file %s", err)
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

	repository := sqlite.New(db, sugar)
	businessLogic := usecase.New(repository, sugar)
	deliveryLayer := New(businessLogic, sugar)

	stat := deliveryLayer.GetStat()
	rr := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/stat/:id", nil)
	require.NoError(t, err)



	c, _ := gin.CreateTestContext(rr)




}

func TestDelivery_Redirect(t *testing.T) {

}