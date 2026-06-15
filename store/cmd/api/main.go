package main

import (
	"context"
	"log"
	_ "store/docs"
	"store/internal/config"
	"store/internal/di"

	"go.uber.org/fx"
	//"go.uber.org/fx/fxevent"
)

// @title           Appliance Store API
// @version         1.0
// @description     API для магазина бытовой техники
// @host      localhost:8080
// @BasePath  /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer " followed by your token
func main() {
	fx.New(
		// fx.WithLogger(func() fxevent.Logger {
		// 	return fxevent.NopLogger
		// }),
		fx.Provide(context.Background),

		di.Module,

		fx.Invoke(func(lc fx.Lifecycle, cfg *config.Config) {
			log.Printf("Server will run on %s", cfg.ServerAddress())
			log.Printf("Swagger UI: http://%s/swagger/index.html", cfg.ServerAddress())
		}),
	).Run()
}
