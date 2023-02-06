package services

import "github.com/HamzaImtiaz/Hashmi-gin-mongo-apis/models"


type UserService interface {
	CreateUser(*models.User) error
	GetUser(*int) (*models.User, error)
	GetAll() ([]*models.User, error)
	UpdateUser(*models.User) error
	DeleteUser(*int) error
}
