package controller

import (
	"myChat/internal/repository"
)

type Controller struct {
	uRepo repository.UserRepository
	tRepo repository.ThreadRepository
}

func NewController(uRepo repository.UserRepository, tRepo repository.ThreadRepository) *Controller {
	return &Controller{
		uRepo: uRepo,
		tRepo: tRepo,
	}
}
