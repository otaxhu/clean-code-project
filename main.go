package main

import (
	"context"

	"github.com/otaxhu/clean-code-project/database"
	"github.com/otaxhu/clean-code-project/internal/repository"
	"github.com/otaxhu/clean-code-project/internal/service"
	"github.com/otaxhu/clean-code-project/settings"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			context.Background,
			settings.New,
			database.New,
			repository.New,
			service.New,
		),
		fx.Invoke(),
	)
	app.Run()
}
