package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/otaxhu/clean-code-project/internal/entity"
)

const (
	qryInsertUser     = "INSERT INTO users (id, email, name, password) VALUES (?, ?, ?, ?);"
	qryGetUserByEmail = "SELECT id, email, name, password FROM users WHERE email = ?"
)

func (r *repo) SaveUser(ctx context.Context, email, name, password string) error {
	_, err := r.db.ExecContext(ctx, qryInsertUser, uuid.NewString(), email, name, password)
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	u := &entity.User{}
	err := r.db.GetContext(ctx, u, qryGetUserByEmail, email)
	if err != nil {
		return nil, err
	}
	return u, nil
}
