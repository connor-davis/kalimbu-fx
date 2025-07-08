package main

import (
	"github.com/connor-davis/kalimbu-fx/cmd/api"
	"github.com/connor-davis/kalimbu-fx/internal/postgres"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		// Provide the gorm database connection as a dependency.
		fx.Provide(
			postgres.NewPostgresDatabase,
		),
		// Provide the cmdlets and the zap logger as dependencies.
		fx.Provide(
			api.NewFiberApi,
			zap.NewProduction,
		),
		fx.Invoke(func(*fiber.App) {}),
	).Run()
}
