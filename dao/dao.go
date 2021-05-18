package dao

import (
	"LoginServer/models"
	"LoginServer/store"
	"fmt"
	"log"

	"crypto/sha256"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

const (
	queryCheckPassword = "SELECT password FROM login WHERE userid=$1"
)

type DataStore struct {
	Store *sqlx.DB
}

func NewDaoStore(db *store.DataStore) *DataStore {
	return &DataStore{Store: db.DB}
}

func (ds *DataStore) CheckCreds(c *gin.Context, creds models.Credentials) error {
	h := sha256.New()
	h.Write([]byte(creds.Password))
	passHash := fmt.Sprintf("%x", h.Sum(nil))
	log.Println(passHash)
	var sum string

	queryError := ds.Store.Get(&sum, queryCheckPassword, creds.UserName)
	if queryError != nil {
		return fmt.Errorf("error: no such user present: %v", queryError)
	}

	if sum != passHash {
		return fmt.Errorf("error: wrong password")
	}

	return nil
}
