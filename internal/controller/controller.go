package controller

import (
	"myChat/internal/repository"
)

type Controller struct {
	repo repository.Repository
}

func NewController(repo repository.Repository) *Controller {
	return &Controller{repo: repo}
}
