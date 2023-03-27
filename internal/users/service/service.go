package service

import (
	"context"

	"github.com/otaxhu/clean-code-project/internal/users/models"
	"github.com/otaxhu/clean-code-project/internal/users/repository"
)

// Service representa a la logica de negocios de la aplicacion
//
//go:generate mockery --name=UsersService
type UsersService interface {
	RegisterUser(ctx context.Context, email, name, password string) error
	LoginUser(ctx context.Context, email, password string) (*models.User, error)
	DeleteUser(ctx context.Context, userId, password string) error

	AddUserRole(ctx context.Context, userId string, roleId int) error
	RemoveUserRole(ctx context.Context, userId string, roleId int) error
}

type userServ struct {
	repo repository.UsersRepository
}

func New(repo repository.UsersRepository) UsersService {
	return &userServ{repo: repo}
}
