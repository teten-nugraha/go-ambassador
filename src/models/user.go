package models

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

type User struct {
	Id           uint
	FirstName    string
	LastName     string
	Email        string
	Password     []byte
	IsAmbassador bool
}

func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	user.Password = hashedPassword
}

func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}

func GeneratePayload(userId uint) (string, error) {

	payload := jwt.StandardClaims{
		Subject:   strconv.Itoa(int(userId)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString([]byte("secret"))

	return token, err
}