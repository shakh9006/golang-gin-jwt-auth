package services

import "github.com/shakh9006/golang-gin-jwt-auth/internal/apiserver/models"

type AuthService interface {
	SignUpUser(*models.SignUpInput) (*models.DBResponse, error)
	SignInUser(*models.SignInInput) (*models.DBResponse, error)
}
