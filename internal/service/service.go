package service

import (
	"context"

	"github.com/otaxhu/clean-code-project/internal/models"
	"github.com/otaxhu/clean-code-project/internal/repository"
)

// Service representa a la logica de negocios de la aplicacion
//
//go:generate mockery --name=Service --output=service --inpackage=true
type Service interface {
	RegisterUser(ctx context.Context, email, name, password string) error
	LoginUser(ctx context.Context, email, password string) (*models.User, error)
}

type serv struct {
	repo repository.Repository
}

func New(repo repository.Repository) Service {
	return &serv{repo: repo}
}
