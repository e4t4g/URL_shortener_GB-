package usecase

import (
	"github.com/e4t4g/URL_shortener_GB-/internal/pkg/repository"
	"github.com/e4t4g/URL_shortener_GB-/internal/pkg/repository/models"
	"math/rand"
	"time"
)

type UrlData struct {
	Id int
	FullUrl  string `yaml:"full_url"`
	ShortUrl string `yaml:"short_url"`
	Counter int64 `yaml:"counter"`
}

type UseCase interface {
	Create(fullUrl *models.UrlData) (*models.UrlData, error)
	//GetStat(id *models.UrlData) *models.UrlData
}

type useCaseUrl struct {
	repository repository.Repository
}

func New(repo repository.Repository) *useCaseUrl {
	return &useCaseUrl{repository: repo}
}

func (usc *useCaseUrl) Create(fullUrl *models.UrlData) (*models.UrlData, error) {
	panic("implement me")
}












//func (usc *useCaseUrl) GetStat(id *models.UrlData) *models.UrlData {
//	panic("implement me")
//}


func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

