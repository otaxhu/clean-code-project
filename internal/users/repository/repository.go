package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/otaxhu/clean-code-project/internal/users/entity"
)

// Repository representa las operaciones CRUD de la aplicacion
//
//go:generate mockery --name=UsersRepository
type UsersRepository interface {
	SaveUser(ctx context.Context, email, name, password string) error
	DeleteUser(ctx context.Context, userId string) error
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	GetUserById(ctx context.Context, userId string) (*entity.User, error)

	SaveUserRole(ctx context.Context, userId string, roleId int) error
	DeleteUserRole(ctx context.Context, userId string, roleId int) error
	GetUserRoles(ctx context.Context, userId string) ([]entity.UserRole, error)
}

type usersRepoMysql struct {
	db *sqlx.DB
}

func NewUsersRepoMysql(db *sqlx.DB) UsersRepository {
	return &usersRepoMysql{db: db}
}
