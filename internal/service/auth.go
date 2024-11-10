package service

import (
	"log"
	"myChat/internal/domain/model"
	"myChat/internal/domain/repository"
	"myChat/pkg/utils"
	"time"
)

type AuthService struct {
	uRepo repository.UserRepository
	sRepo repository.SessionRepository
}

func NewAuthService(u repository.UserRepository, s repository.SessionRepository) *AuthService {
	return &AuthService{
		uRepo: u,
		sRepo: s,
	}
}

// Check if user logged in.
func (as *AuthService) CheckSession(uuid string) (model.Session, error) {

	// return empty session and error,
	// if session doesnt exits in database.
	session, err := as.sRepo.FindByUuid(uuid)
	if (err != nil || session == model.Session{}) {
		log.Println("auth failed, cannot get user session", err)
		return model.Session{}, err
	}

	return session, nil
}

func (as *AuthService) Login(email string, password string) (string, error) {
	user, err := as.uRepo.FindByEmail(email)
	if err != nil {
		return "", err
	}
	if user.Password == utils.Encrypt(password) {
		session := model.Session{
			Uuid:      utils.CreateUUID(),
			Email:     user.Email,
			UserId:    user.Id,
			CreatedAt: time.Now(),
		}
		if err := as.sRepo.Save(session); err != nil {
			return "", err
		}
		return session.Uuid, nil
	}
	return "", err
}

func (as *AuthService) CreateUser(name string, email string, password string) error {
	user := model.User{
		Uuid:      utils.CreateUUID(),
		Name:      name,
		Email:     email,
		Password:  utils.Encrypt(password),
		CreatedAt: time.Now(),
	}
	if err := as.uRepo.Save(user); err != nil {
		return err
	}
	return nil
}

func (as *AuthService) Logout(uuid string) error {
	session, err := as.sRepo.FindByUuid(uuid)
	if err != nil {
		return err
	}
	if err := as.sRepo.DeleteByUserId(session.UserId); err != nil {
		return err
	}
	return nil
}
