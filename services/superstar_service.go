package services

import "github.com/yz124/superstar/models"

type SuperstarService interface {
	GetAll() []models.StarInfo
	Search(country string) []models.StarInfo
	Get(id int) *models.StarInfo
	Delete(id int) error
	Update(user *models.StarInfo, columns []string) error
	Create(info *models.StarInfo) error
}
