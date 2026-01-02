package services

import (
	"errors"
	"time"

	"github.com/Zain0205/cf-stunting-backend-go/internal/config"
	"github.com/Zain0205/cf-stunting-backend-go/internal/models"
	"github.com/Zain0205/cf-stunting-backend-go/internal/repositories"
	"github.com/Zain0205/cf-stunting-backend-go/internal/utils"

	"github.com/golang-jwt/jwt/v5"
)

func Register(name, phone, password string, category models.UserCategory) error {
	hashed, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	user := &models.User{
		Name:        name,
		PhoneNumber: phone,
		Password:    hashed,
		Category:    category,
	}

	return repositories.CreateUser(user)
}

func Login(phone, password string) (string, error) {
	user, err := repositories.GetUserByPhone(phone)
	if err != nil {
		return "", errors.New("invalid phone or password")
	}

	if err := utils.CheckPassword(user.Password, password); err != nil {
		return "", errors.New("invalid phone or password")
	}

	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"category": user.Category,
		"exp":      time.Now().Add(config.JWTExpire()).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.JWTSecret()))
}
