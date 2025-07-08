package system_services

import (
	"github.com/connor-davis/kalimbu-fx/internal/services/system_services/permissions_service"
	"github.com/connor-davis/kalimbu-fx/internal/services/system_services/roles_service"
	"github.com/connor-davis/kalimbu-fx/internal/services/system_services/users_service"
	"go.uber.org/fx"
)

func ProvideSystemServices() fx.Option {
	return fx.Provide(
		permissions_service.NewPermissionsService,
		roles_service.NewRolesService,
		users_service.NewUsersService,
	)
}
