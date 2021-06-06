package service

import (
	"fmt"
	"log"
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

type JwtService struct {
	SecretKey, Issuer string
}

func NewJwtService() *JwtService {
	return &JwtService{SecretKey: secretKey, Issuer: issuer}
}

func (js *JwtService) GenerateToken(name string) (string, error) {
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

func (js *JwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signin method: %v", token.Header["alg"])
		}

		return []byte(js.SecretKey), nil
	})
}
