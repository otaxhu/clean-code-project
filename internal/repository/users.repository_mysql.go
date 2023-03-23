package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/otaxhu/clean-code-project/internal/entity"
)

const (
	qryInsertUser     = "INSERT INTO users (id, email, name, password) VALUES (?, ?, ?, ?)"
	qryGetUserByEmail = "SELECT id, email, name, password FROM users WHERE email = ?"
	qryGetUserById    = "SELECT id, email, name, password FROM users WHERE id = ?"
	qryInsertUserRole = "INSERT INTO user_roles (id, user_id, role_id) VALUES (?, ?, ?)"
	qryDeleteUserRole = "DELETE FROM user_roles WHERE user_id = ? AND role_id = ?"
	qryDeleteUser     = "DELETE FROM users WHERE id = ?"
	qryGetUserRoles   = "SELECT id, user_id, role_id FROM user_roles WHERE user_id = ?"
)

func (r *repoMysql) SaveUser(ctx context.Context, email, name, password string) error {
	_, err := r.db.ExecContext(ctx, qryInsertUser, uuid.NewString(), email, name, password)
	if err != nil {
		return err
	}
	return nil
}

func (r *repoMysql) DeleteUser(ctx context.Context, userId string) error {
	_, err := r.db.ExecContext(ctx, qryDeleteUser, userId)
	if err != nil {
		return err
	}
	return nil
}

func (r *repoMysql) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	u := &entity.User{}
	err := r.db.GetContext(ctx, u, qryGetUserByEmail, email)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *repoMysql) GetUserById(ctx context.Context, userId string) (*entity.User, error) {
	u := &entity.User{}
	err := r.db.GetContext(ctx, u, qryGetUserById)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *repoMysql) SaveUserRole(ctx context.Context, userId string, roleId int) error {
	_, err := r.db.ExecContext(ctx, qryInsertUserRole, uuid.NewString(), userId, roleId)
	if err != nil {
		return err
	}
	return nil
}

func (r *repoMysql) DeleteUserRole(ctx context.Context, userId string, roleId int) error {
	_, err := r.db.ExecContext(ctx, qryDeleteUserRole, userId, roleId)
	if err != nil {
		return err
	}
	return nil
}

func (r *repoMysql) GetUserRoles(ctx context.Context, userId string) ([]entity.UserRole, error) {
	roles := []entity.UserRole{}
	err := r.db.SelectContext(ctx, &roles, qryGetUserRoles, userId)
	if err != nil {
		return nil, err
	}
	return roles, nil
}
