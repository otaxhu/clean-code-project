package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/otaxhu/clean-code-project/internal/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

// Repository representa las operaciones CRUD de la aplicacion
//
//go:generate mockery --name=Repository
type Repository interface {
	SaveUser(ctx context.Context, email, name, password string) error
	DeleteUser(ctx context.Context, userId string) error
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	GetUserById(ctx context.Context, userId string) (*entity.User, error)

	SaveUserRole(ctx context.Context, userId string, roleId int) error
	DeleteUserRole(ctx context.Context, userId string, roleId int) error
	GetUserRoles(ctx context.Context, userId string) ([]entity.UserRole, error)
}

type repoMysql struct {
	db *sqlx.DB
}

type repoMongo struct {
	client *mongo.Client
}

func NewRepoMysql(db *sqlx.DB) Repository {
	return &repoMysql{db: db}
}

func NewRepoMongo(client *mongo.Client) Repository {
	return &repoMongo{client: client}
}
