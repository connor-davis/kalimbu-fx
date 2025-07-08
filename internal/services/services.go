package services

import (
	permissionsService "github.com/connor-davis/kalimbu-fx/internal/services/permissions"
	rolesService "github.com/connor-davis/kalimbu-fx/internal/services/roles"
	usersService "github.com/connor-davis/kalimbu-fx/internal/services/users"
	"go.uber.org/fx"
)

func ProvideServices() fx.Option {
	return fx.Provide(
		usersService.NewUsersService,
		rolesService.NewRolesService,
		permissionsService.NewPermissionsService,
	)
}
