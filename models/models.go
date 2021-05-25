package models

type Credentials struct {
	UserName string `json:"userid" db:"userid"`
	Password string `json:"password" db:"password"`
}
