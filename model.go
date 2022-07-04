package main

import (
	"errors"
	"github.com/cristalhq/jwt/v3"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"os"
	"time"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Password string
}

type Register struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) HashPassword(plain string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plain), 14)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func (u *User) CheckPassword(plain, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))

	return err == nil
}

func (u *User) GenerateJWT(email string) (string, error) {
	expirationTime := time.Now().Add(time.Minute * 60)
	signer, err := jwt.NewSignerHS(jwt.HS256, []byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	claims := &jwt.RegisteredClaims{
		Issuer:    "jwt-example",
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(expirationTime),
		Subject:   email,
	}
	builder := jwt.NewBuilder(signer)
	token, _ := builder.Build(claims)
	if token == nil {
		return "", errors.New("failed creating token")
	}
	return token.String(), nil
}
