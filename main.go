package main

import (
	mapper "LoginServer/app"
	"LoginServer/store"
)

func main() {
	db := store.InitDB()
	mapper.Map(db)
}
