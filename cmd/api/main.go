package api

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func RunFiberApi(lc fx.Lifecycle, logger *zap.Logger, postgres *gorm.DB) *fiber.App {
	app := fiber.New(fiber.Config{
		ServerHeader: "Thusa Kalimbu API",
		AppName:      "Thusa Kalimbu API",
	})

	lc.Append(fx.Hook{
		// OnStart is called when the application starts.
		// It is executed after all constructors have been called and all
		// dependencies have been injected.
		OnStart: func(ctx context.Context) error {
			go func() {
				logger.Info("Starting Fiber API server", zap.String("address", ":6173"))

				if err := app.Listen(":6173"); err != nil {
					logger.Error("Failed to start Fiber API server", zap.Error(err))
					return
				}
			}()

			return nil
		},
		// OnStop is called when the application stops.
		OnStop: func(ctx context.Context) error {
			return app.Shutdown()
		},
	})

	return app
}
