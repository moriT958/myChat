package controller

import (
	"myChat/internal/service"
)

type Controller struct {
	Service *service.AppService
}

func NewController(s *service.AppService) *Controller {
	return &Controller{Service: s}
}
