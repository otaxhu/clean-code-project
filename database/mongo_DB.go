package database

import (
	"context"
	"fmt"

	"github.com/otaxhu/clean-code-project/settings"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoConnection(ctx context.Context, s *settings.Settings) (*mongo.Database, error) {
	var connectionString string
	if s.MongoConfig.User == "" && s.MongoConfig.Password == "" {
		connectionString = fmt.Sprintf("mongodb://%s:%d/%s", s.MongoConfig.Host, s.MongoConfig.Port, s.MongoConfig.Name)
	} else {
		connectionString = fmt.Sprintf(
			"mongodb://%s:%s@%s:%d/%s",
			s.MongoConfig.User,
			s.MongoConfig.Password,
			s.MongoConfig.Host,
			s.MongoConfig.Port,
			s.MongoConfig.Name,
		)
	}
	options := options.Client().ApplyURI(connectionString)
	connection, err := mongo.Connect(ctx, options)
	if err != nil {
		return nil, err
	}
	return connection.Database(s.MongoConfig.Name), nil
}
