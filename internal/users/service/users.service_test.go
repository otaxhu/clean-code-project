package service

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/otaxhu/clean-code-project/common/encryption"
	"github.com/otaxhu/clean-code-project/internal/users/entity"
	"github.com/otaxhu/clean-code-project/internal/users/repository/mocks"
	"github.com/stretchr/testify/mock"
)

var repo *mocks.Repository = &mocks.Repository{}
var s UsersService = New(repo)
var tooLongPassword = "A continuación, colocaré más de 72 caracteres, que es el máximo permitido por la encriptación, para asegurarme de que la prueba sea robusta y cubra todos los posibles casos de uso"
var existentRolesUserId = uuid.NewString()
var validUserIdButNoRoles = uuid.NewString()
var validUserId = uuid.NewString()

var (
	errRepositoryFKConstraintFails = errors.New("foreign key constraint fails")
	errRepositoryUserNotFound      = errors.New("user not found")
)

func TestMain(m *testing.M) {
	validPassword, _ := encryption.EncryptPassword([]byte("validPassword"))
	encryptedPassword := encryption.FromHashToBase64(validPassword)
	validUser := &entity.User{Id: validUserId, Email: "test@exists.com", Password: encryptedPassword}
	validUserRoles := []entity.UserRole{{Id: uuid.NewString(), UserId: existentRolesUserId, RoleId: 1}}

	repo.On("SaveUser", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	repo.On("GetUserByEmail", mock.Anything, "test@valid.com").Return(nil, nil)

	repo.On("GetUserByEmail", mock.Anything, "test@notExists.com").Return(nil, errInvalidCredentials)
	repo.On("GetUserByEmail", mock.Anything, "test@exists.com").Return(validUser, nil)

	repo.On("GetUserById", mock.Anything, validUserId).Return(validUser, nil)
	repo.On("GetUserById", mock.Anything, "inexistentUserId").Return(nil, errRepositoryUserNotFound)

	repo.On("GetUserRoles", mock.Anything, existentRolesUserId).Return(validUserRoles, nil)
	repo.On("GetUserRoles", mock.Anything, "inexistentUserId").Return(nil, nil)
	repo.On("GetUserRoles", mock.Anything, validUserIdButNoRoles).Return([]entity.UserRole{}, nil)

	repo.On("SaveUserRole", mock.Anything, existentRolesUserId, 1).Return(nil)
	repo.On("SaveUserRole", mock.Anything, validUserIdButNoRoles, 1).Return(nil)
	repo.On("SaveUserRole", mock.Anything, "inexistentUserId", mock.Anything).Return(errRepositoryFKConstraintFails)

	repo.On("DeleteUserRole", mock.Anything, existentRolesUserId, 1).Return(nil)
	repo.On("DeleteUserRole", mock.Anything, "inexistentUserId", mock.Anything).Return(errUserRoleNotFound)

	repo.On("DeleteUser", mock.Anything, validUserId).Return(nil)

	code := m.Run()
	os.Exit(code)
}

func TestRegisterUser(t *testing.T) {
	testCases := []struct {
		name          string
		email         string
		userName      string
		password      string
		expectedError error
	}{
		{
			"RegisterUser_Success",
			"test@valid.com",
			"validUserName",
			"validPassword",
			nil,
		},
		{
			"UserAlreadyExists",
			"test@exists.com",
			"validUserName",
			"validPassword",
			errUserAlreadyExists,
		},
		{
			"TooLongPassword",
			"test@valid.com",
			"validUserName",
			tooLongPassword,
			encryption.ErrPasswordTooLong,
		},
	}
	ctx := context.Background()
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			repo.Mock.Test(t)
			if err := s.RegisterUser(ctx, tc.email, tc.userName, tc.password); err != tc.expectedError {
				t.Errorf("Expected error %v, got %v", tc.expectedError, err)
			}
		})
	}
}

func TestLoginUser(t *testing.T) {
	testCases := []struct {
		name          string
		email         string
		password      string
		expectedError error
	}{
		{
			"LoginUser_Success",
			"test@exists.com",
			"validPassword",
			nil,
		},
		{
			"InvalidCredentialsEmail",
			"test@notExists.com",
			"validPassword",
			errInvalidCredentials,
		},
		{
			"InvalidCredentialsPassword",
			"test@exists.com",
			"invalidPassword",
			errInvalidCredentials,
		},
	}
	ctx := context.Background()
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			repo.Mock.Test(t)
			if _, err := s.LoginUser(ctx, tc.email, tc.password); err != tc.expectedError {
				t.Errorf("Expected error %v, got %v", tc.expectedError, err)
			}

		})
	}
}

func TestAddUserRole(t *testing.T) {
	testCases := []struct {
		name          string
		userId        string
		roleId        int
		expectedError error
	}{
		{
			"AddUserRole_Success",
			validUserIdButNoRoles,
			1,
			nil,
		},
		{
			"InvalidUserId",
			"inexistentUserId",
			1,
			errRepositoryFKConstraintFails,
		},
		{
			"RoleAlreadyExist",
			existentRolesUserId,
			1,
			errUserRoleAdded,
		},
	}
	ctx := context.Background()
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			repo.Mock.Test(t)
			if err := s.AddUserRole(ctx, tc.userId, tc.roleId); err != tc.expectedError {
				t.Errorf("Expected error %v, got %v", tc.expectedError, err)
			}
		})
	}
}

func TestRemoveUserRole(t *testing.T) {
	testCases := []struct {
		name          string
		userId        string
		roleId        int
		expectedError error
	}{
		{
			"RemoveUserRole_Success",
			existentRolesUserId,
			1,
			nil,
		},
		{
			"InvalidUserId",
			"inexistentUserId",
			1,
			errUserRoleNotFound,
		},
	}
	ctx := context.Background()
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			repo.Mock.Test(t)
			if err := s.RemoveUserRole(ctx, tc.userId, tc.roleId); err != tc.expectedError {
				t.Errorf("Expected error %v, got %v", tc.expectedError, err)
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	testCases := []struct {
		name          string
		userId        string
		password      string
		expectedError error
	}{
		{
			"DeleteUser_Success",
			validUserId,
			"validPassword",
			nil,
		},
		{
			"InvalidUserId",
			"inexistentUserId",
			"validPassword",
			errInvalidCredentials,
		},
		{
			"InvalidPassword",
			validUserId,
			"invalidPassword",
			errInvalidCredentials,
		},
	}
	ctx := context.Background()
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			repo.Mock.Test(t)
			if err := s.DeleteUser(ctx, tc.userId, tc.password); err != tc.expectedError {
				t.Errorf("Expected error %v, got %v", tc.expectedError, err)
			}
		})
	}
}
