package service

import (
	"context"

	"github.com/otaxhu/clean-code-project/internal/models"
	"github.com/otaxhu/clean-code-project/internal/repository"
)

// Service representa a la logica de negocios de la aplicacion
//
//go:generate mockery --name=Service
type Service interface {
	RegisterUser(ctx context.Context, email, name, password string) error
	LoginUser(ctx context.Context, email, password string) (*models.User, error)
	DeleteUser(ctx context.Context, userId, password string) error

	AddUserRole(ctx context.Context, userId string, roleId int) error
	RemoveUserRole(ctx context.Context, userId string, roleId int) error
}

type serv struct {
	repo repository.Repository
}

func New(repo repository.Repository) Service {
	return &serv{repo: repo}
}
