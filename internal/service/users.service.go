package service

import (
	"context"
	"errors"

	"github.com/otaxhu/clean-code-project/common/encryption"
	"github.com/otaxhu/clean-code-project/internal/models"
)

var (
	errUserAlreadyExists  = errors.New("user already exists")
	errInvalidCredentials = errors.New("invalid credentials")
)

func (s *serv) RegisterUser(ctx context.Context, email, name, password string) error {
	u, _ := s.repo.GetUserByEmail(ctx, email)
	if u != nil {
		return errUserAlreadyExists
	}
	hash, err := encryption.EncryptPassword([]byte(password))
	if err != nil {
		return err
	}
	password = encryption.FromHashToBase64(hash)
	return s.repo.SaveUser(ctx, email, name, password)
}

func (s *serv) LoginUser(ctx context.Context, email, password string) (*models.User, error) {
	u, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, errInvalidCredentials
	}
	hash, err := encryption.FromBase64ToHash(u.Password)
	if err != nil {
		return nil, err
	}
	if err := encryption.CompareHashAndPassword(hash, []byte(password)); err != nil {
		return nil, errInvalidCredentials
	}
	return &models.User{Id: u.Id, Name: u.Name, Email: u.Email}, nil
}
