package repository

import (
	"context"
	"log"

	"github.com/otaxhu/clean-code-project/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *repoMongo) SaveUser(ctx context.Context, email, name, password string) error {
	_, err := r.db.Collection("users").InsertOne(ctx, bson.D{{Key: "email", Value: email}, {Key: "name", Value: name}, {Key: "password", Value: password}})
	if err != nil {
		return err
	}
	return nil
}

func (r *repoMongo) DeleteUser(ctx context.Context, userId string) error {
	_, err := r.db.Collection("users").DeleteOne(ctx, bson.D{{Key: "_id", Value: userId}})
	if err != nil {
		return err
	}
	return nil
}

func (r *repoMongo) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	u := &entity.User{}
	if err := r.db.Collection("users").FindOne(ctx, bson.D{{Key: "email", Value: email}}).Decode(u); err != nil {
		return nil, err
	}
	return u, nil
}

func (r *repoMongo) GetUserById(ctx context.Context, userId string) (*entity.User, error) {
	u := &entity.User{}
	if err := r.db.Collection("users").FindOne(ctx, bson.D{{Key: "_id", Value: userId}}).Decode(u); err != nil {
		return nil, err
	}
	return u, nil
}

func (r *repoMongo) SaveUserRole(ctx context.Context, userId string, roleId int) error {
	_, err := r.db.Collection("users").UpdateByID(ctx, userId, bson.D{{Key: "$set", Value: bson.E{Key: "roles", Value: bson.A{roleId}}}})
	if err != nil {
		return err
	}
	return nil
}

func (r *repoMongo) DeleteUserRole(ctx context.Context, userId string, roleId int) error {
	_, err := r.db.Collection("users").DeleteOne(ctx, bson.D{{Key: "_id", Value: userId}})
	if err != nil {
		return err
	}
	return nil
}

func (r *repoMongo) GetUserRoles(ctx context.Context, userId string) ([]entity.UserRole, error) {
	user := &entity.User{}
	if err := r.db.Collection("users").FindOne(ctx, bson.D{{Key: "_id", Value: userId}}).Decode(user); err != nil {
		log.Println(user)
		return nil, err
	}
	return user.Roles, nil
}
