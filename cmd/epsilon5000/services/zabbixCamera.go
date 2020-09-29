package services

import "awesomeProject/cmd/epsilon5000/models"

type zabbixCameraDAO interface {
	Get(id models.Id) ([]models.СameraIncidents ,error)
}

type zabbixCamera struct {
	dao zabbixCameraDAO
}

func NewZabbixCameraDAO(dao zabbixCameraDAO) *zabbixCamera {

	return &zabbixCamera{dao}
}

func (s zabbixCamera) Get(id models.Id) ([]models.СameraIncidents, error) {

	return s.dao.Get(id)
}
