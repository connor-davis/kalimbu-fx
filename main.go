package main

import (
	"github.com/connor-davis/kalimbu-fx/cmd/api"
	"github.com/connor-davis/kalimbu-fx/cmd/migrate"
	"github.com/connor-davis/kalimbu-fx/internal/postgres"
	"github.com/connor-davis/kalimbu-fx/internal/services/system_services"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		// Provide the services as dependencies.
		system_services.ProvideSystemServices(),
		// Provide the gorm database connection as a dependency.
		fx.Provide(
			postgres.NewPostgresDatabase,
			migrate.RunMigrations,
		),
		// Provide the cmdlets and the zap logger as dependencies.
		fx.Provide(
			api.RunFiberApi,
			zap.NewProduction,
		),
	).Run()
}
