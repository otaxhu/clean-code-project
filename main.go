package main

import (
	"context"

	"github.com/otaxhu/clean-code-project/database"
	"github.com/otaxhu/clean-code-project/internal/users/repository"
	"github.com/otaxhu/clean-code-project/internal/users/service"
	"github.com/otaxhu/clean-code-project/settings"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			context.Background,
			settings.New,
			database.NewMysqlConection,
			repository.NewUsersRepoMysql,
			service.New,
		),
		fx.Invoke(),
	)
	app.Run()
}
