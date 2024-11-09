package controller

import (
	"myChat/internal/repository"
)

type Controller struct {
	uRepo repository.UserRepository
	sRepo repository.SessionRepository
	tRepo repository.ThreadRepository
	pRepo repository.PostRepository
}

func NewController(
	uRepo repository.UserRepository,
	sRepo repository.SessionRepository,
	tRepo repository.ThreadRepository,
	pRepo repository.PostRepository,
) *Controller {
	return &Controller{
		uRepo: uRepo,
		sRepo: sRepo,
		tRepo: tRepo,
		pRepo: pRepo,
	}
}
