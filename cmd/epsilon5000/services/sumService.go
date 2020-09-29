package services

import "awesomeProject/cmd/epsilon5000/models"

type sumServiceDAO interface {
	Get(inc string) ([]models.ServicesCount, error)
}

type sumService struct {
	dao sumServiceDAO
}

func NewSumServiceDAO(dao sumServiceDAO) *sumService  {

	return &sumService{dao}
}

func (s sumService) Get(inc string) ([]models.ServicesCount, error)  {
	return s.dao.Get(inc)
}