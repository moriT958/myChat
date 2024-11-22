package service

import (
	"errors"
	"myChat/internal/domain/model"
	"myChat/internal/domain/repository"
	"myChat/pkg/apperrors"
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
		err = apperrors.NoSessionFound.Wrap(err, "failed to get user session")
		return model.Session{}, err
	}

	return session, nil
}

func (as *AuthService) Login(email string, password string) (string, error) {
	user, err := as.uRepo.FindByEmail(email)
	if err != nil {
		err = apperrors.NoUserFound.Wrap(err, "failed to get user by email")
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
			err = apperrors.CreateSessionFailed.Wrap(err, "failed to save session")
			return "", err
		}
		return session.Uuid, nil
	} else {
		err := errors.New("failed to login")
		err = apperrors.CreateSessionFailed.Wrap(err, "password does't match")
		return "", err
	}
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
		err = apperrors.CreateUserFailed.Wrap(err, "failed to save user data")
		return err
	}

	return nil
}

func (as *AuthService) Logout(uuid string) error {
	session, err := as.sRepo.FindByUuid(uuid)
	if err != nil {
		err = apperrors.NoSessionFound.Wrap(err, "no session found by uuid")
		return err
	}

	if err := as.sRepo.DeleteByUserId(session.UserId); err != nil {
		err = apperrors.DeleteSessionFailed.Wrap(err, "failed to delete session by user_id")
		return err
	}

	return nil
}
