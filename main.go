package main

import (
	"github.com/connor-davis/kalimbu-fx/cmd/api"
	"github.com/connor-davis/kalimbu-fx/cmd/migrate"
	"github.com/connor-davis/kalimbu-fx/internal/postgres"
	"github.com/connor-davis/kalimbu-fx/internal/services"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		// Provide the gorm database connection as a dependency.
		fx.Provide(
			postgres.NewPostgresDatabase,
			migrate.RunMigrations,
		),
		// Provide the services as dependencies.
		services.ProvideServices(),
		// Provide the cmdlets and the zap logger as dependencies.
		fx.Provide(
			api.RunFiberApi,
			zap.NewProduction,
		),
		// Run the migrations and the Fiber API server.
		fx.Invoke(func(*migrate.Migrations) {}),
		fx.Invoke(func(*fiber.App) {}),
	).Run()
}
