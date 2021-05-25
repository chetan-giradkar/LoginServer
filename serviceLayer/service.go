package service

import (
	"LoginServer/dao"
	"LoginServer/models"

	"github.com/gin-gonic/gin"
)

type Service struct {
	DataStore dao.DataStore
}

func NewService(ds dao.DataStore) *Service {
	return &Service{DataStore: ds}
}

func (s *Service) Login(c *gin.Context, creds models.Credentials) error {
	passwordError := s.DataStore.CheckCreds(c, creds)
	if passwordError != nil {
		return passwordError
	}

	return nil
}

func (s *Service) Register(c *gin.Context, creds models.Credentials) error {
	passwordError := s.DataStore.Register(c, creds)
	if passwordError != nil {
		return passwordError
	}

	return nil
}
