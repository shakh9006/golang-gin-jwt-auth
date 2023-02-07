package services

import "github.com/shakh9006/golang-gin-jwt-auth/internal/apiserver/models"

type UserService interface {
	FindUserById(string) (*models.DBResponse, error)
	FindUserByEmail(string) (*models.DBResponse, error)
}
