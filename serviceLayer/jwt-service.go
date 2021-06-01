package service

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	million   = 1000000
	secretKey = "s5v8y/B?E(H+MbQeThWmZq4t6w9z$C&F"
	issuer    = "Chetan"
)

type jwtCustomeClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

type jwtService struct {
	SecretKey, Issuer string
}

func NewJwtService() *jwtService {
	return &jwtService{SecretKey: secretKey, Issuer: issuer}
}

func (js *jwtService) GenerateToken(name string) (string, error) {
	claims := jwtCustomeClaims{
		Name: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
			Issuer:    js.Issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, encodingError := token.SignedString([]byte(js.SecretKey))
	if encodingError != nil {
		log.Printf("Error encoding token: %v\n", encodingError.Error())

		return "", encodingError
	}

	fmt.Println("Token generated for user: ", name)
	return t, nil
}

func (js *jwtService) ValidateToken() {
	cwd, _ := os.Getwd()
	fmt.Println(cwd, time.Now())
}
