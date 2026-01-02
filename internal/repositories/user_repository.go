package repositories

import (
	"github.com/Zain0205/cf-stunting-backend-go/internal/database"
	"github.com/Zain0205/cf-stunting-backend-go/internal/models"
)

func CreateUser(user *models.User) error {
	return database.DB.Create(user).Error
}

func GetUserByPhone(phone string) (*models.User, error) {
	var user models.User
	err := database.DB.Where("phone_number = ?", phone).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
