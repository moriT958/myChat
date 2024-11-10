package service

type AppService struct {
	Auth  *AuthService
	Forum *ForumService
}

func NewAppService(auth *AuthService, forum *ForumService) *AppService {
	return &AppService{
		Auth:  auth,
		Forum: forum,
	}
}
