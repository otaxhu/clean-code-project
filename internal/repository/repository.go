package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/otaxhu/clean-code-project/internal/entity"
)

// Repository representa las operaciones CRUD de la aplicacion
//
//go:generate mockery --name=Repository
type Repository interface {
	SaveUser(ctx context.Context, email, name, password string) error
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)

	SaveUserRole(ctx context.Context, userId string, roleId int) error
	DeleteUserRole(ctx context.Context, userId string, roleId int) error
	GetUserRoles(ctx context.Context, userId string) ([]entity.UserRole, error)
}

type repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return &repo{db: db}
}
