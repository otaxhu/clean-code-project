package database

import (
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/otaxhu/clean-code-project/settings"
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
