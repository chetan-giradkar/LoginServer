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
	queryAddUser       = "INSERT INTO credentials (userid, password) VALUES (?, ?)"
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

	if len(res) < 1 || res[0] != passHash {
		return fmt.Errorf("error: wrong password")
	}

	return nil
}

func (ds *DataStore) Register(c *gin.Context, creds models.Credentials) error {
	h := sha256.New()
	h.Write([]byte(creds.Password))
	passHash := fmt.Sprintf("%x", h.Sum(nil))

	creds.Password = passHash

	res := ds.Store.MustExec(queryAddUser, creds.UserName, creds.Password)
	num, queryError := res.RowsAffected()
	if queryError != nil {
		c.JSON(http.StatusInternalServerError, queryError)
		c.Abort()

		return fmt.Errorf("internal server error: %v", queryError.Error())
	}

	if num != 1 {
		errorNum := fmt.Errorf("internal server error: User could not be registered")
		c.JSON(http.StatusInternalServerError, errorNum)
		c.Abort()

		return errorNum
	}

	return nil
}
