package repository

import "github.com/e4t4g/URL_shortener_GB-/internal/pkg/repository/models"

type Repository interface {
	Create(url *models.UrlData) (*models.UrlData, error)
	GetStat(id int) *models.UrlData
}
