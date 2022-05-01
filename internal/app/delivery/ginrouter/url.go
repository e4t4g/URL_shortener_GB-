package ginrouter

import (
	"fmt"
	"github.com/e4t4g/URL_shortener_GB-/cmd/app/config"
	"github.com/e4t4g/URL_shortener_GB-/internal/app/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Delivery interface {
	Create() gin.HandlerFunc
	Redirect() gin.HandlerFunc
	GetStat() gin.HandlerFunc
}

type delivery struct {
	url usecase.UseCase
}

type URLData struct {
	ID       int    `json:"id" yaml:"id"`
	FullURL  string `json:"full_url" yaml:"full_url"`
	ShortURL string `json:"short_url" yaml:"short_url"`
	Counter  int64  `json:"counter" yaml:"counter"`
}

func New(url usecase.UseCase) Delivery {
	return delivery{
		url: url,
	}
}

func (d delivery) Create() gin.HandlerFunc {
	return func(c *gin.Context) {

		cfg := config.Config{}

		err := cfg.ReadFromFile()
		if err != nil {
			return
		}

		//host := c.ClientIP()
		host := cfg.Host
		port := cfg.Port

		var newURL *URLData
		if err = c.ShouldBind(&newURL); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
		}

		result, err := d.url.Create((*usecase.URLData)(newURL))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "unable to create"})
		}

		shortURL := fmt.Sprintf("%s:%d/%s", host, port, result.ShortURL)
		statURL := fmt.Sprintf("%s:%d/stat/%d", host, port, result.ID)

		c.JSON(http.StatusOK,
			gin.H{
				"short_url": shortURL,
				"stat_link": statURL,
			})
	}
}

func (d delivery) Redirect() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.Param("token")

		redirectStruct, err := d.url.Redirect(c.Request.Context(), token)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "incorrect token format"})
		}

		c.Redirect(http.StatusMovedPermanently, redirectStruct.FullURL)

	}
}

func (d delivery) GetStat() gin.HandlerFunc {
	return func(c *gin.Context) {

		ID, err := strconv.Atoi(c.Param("ID"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "incorrect ID"})
		}

		redirectStruct, err := d.url.GetStat(ID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"status": "not found"})
		}

		c.JSON(http.StatusOK,
			gin.H{"Stat": redirectStruct.Counter})

	}
}
