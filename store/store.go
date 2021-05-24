package store

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DataStore struct {
	DB *sqlx.DB
}

func InitDB() *DataStore {
	db, err := sqlx.Open("mysql", "root@tcp(127.0.0.1:3306)/login")
	if err != nil {
		log.Fatalln(err)
	}

	return &DataStore{DB: db}
}
