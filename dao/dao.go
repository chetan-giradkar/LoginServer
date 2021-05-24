package dao

import (
	"LoginServer/models"
	"LoginServer/store"
	"fmt"
	"net/http"

	"crypto/sha256"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	queryCheckPassword = "SELECT password FROM credentials WHERE userid = ?"
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

	var res []string

	queryError := ds.Store.Select(&res, queryCheckPassword, creds.UserName)
	if queryError != nil {
		c.JSON(http.StatusInternalServerError, queryError)
		c.Abort()

		return fmt.Errorf("internal server error: %v", queryError.Error())
	}

	sum := res[0]

	if sum != passHash {
		return fmt.Errorf("error: wrong password")
	}

	return nil
}
