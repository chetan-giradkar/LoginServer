package service

import (
	"LoginServer/dao"
	"LoginServer/models"

	"github.com/gin-gonic/gin"
)

type LoginService struct {
	DataStore  dao.DataStore
	JwtService JwtService
}

func NewService(ds dao.DataStore, js JwtService) *LoginService {
	return &LoginService{DataStore: ds, JwtService: js}
}

func (s *LoginService) Login(c *gin.Context, creds models.Credentials) (string, error) {
	passwordError := s.DataStore.CheckCreds(c, creds)
	if passwordError != nil {
		return "", passwordError
	}

	token, tokenError := s.JwtService.GenerateToken(creds.UserName)

	if tokenError != nil {
		return "", tokenError
	}

	return token, nil
}

func (s *LoginService) Register(c *gin.Context, creds models.Credentials) error {
	passwordError := s.DataStore.Register(c, creds)
	if passwordError != nil {
		return passwordError
	}

	return nil
}
