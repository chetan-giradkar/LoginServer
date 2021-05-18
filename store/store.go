package store

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type DataStore struct {
	DB *sqlx.DB
}

func InitDB() *DataStore {
	db, err := sqlx.Connect("mysql", "user=root dbname=login sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	return &DataStore{DB: db}
}
