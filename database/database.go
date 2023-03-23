package database

import (
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/otaxhu/clean-code-project/settings"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMysqlConection(ctx context.Context, s *settings.Settings) (*sqlx.DB, error) {
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		s.MysqlConfig.User,
		s.MysqlConfig.Password,
		s.MysqlConfig.Host,
		s.MysqlConfig.Port,
		s.MysqlConfig.Name,
	)
	return sqlx.ConnectContext(ctx, "mysql", connectionString)
}

func NewMongoConnection(ctx context.Context, s *settings.Settings) (*mongo.Client, error) {
	connectionString := fmt.Sprintf(
		"mongodb://%s:%s@%s:%d/%s",
		s.MongoConfig.User,
		s.MongoConfig.Password,
		s.MongoConfig.Host,
		s.MongoConfig.Port,
		s.MongoConfig.Name,
	)
	options := options.Client().ApplyURI(connectionString)
	return mongo.Connect(ctx, options)
}
