package main

import (
	"github.com/connor-davis/kalimbu-fx/cmd/api"
	"github.com/connor-davis/kalimbu-fx/cmd/migrate"
	"github.com/connor-davis/kalimbu-fx/internal/postgres"
	permissionsService "github.com/connor-davis/kalimbu-fx/internal/services/permissions"
	rolesService "github.com/connor-davis/kalimbu-fx/internal/services/roles"
	usersService "github.com/connor-davis/kalimbu-fx/internal/services/users"
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
		fx.Provide(
			rolesService.NewRolesService,
			permissionsService.NewPermissionsService,
			usersService.NewUsersService,
		),

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
