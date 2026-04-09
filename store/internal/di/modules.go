package di

import (
	"context"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/fx"

	"store/internal/config"
	"store/internal/handler"
	client_repo "store/internal/repository/client"
	supplier_repo "store/internal/repository/supplier"
	image_repo "store/internal/repository/image"
	category_repo "store/internal/repository/category"
	product_repo "store/internal/repository/product"
	address_repo "store/internal/repository/address"
	dbconn "store/internal/repository/db"
	"store/internal/router"
	client_service "store/internal/service/client"
	supplier_service "store/internal/service/supplier"
	image_service "store/internal/service/image"
	category_service "store/internal/service/category"
	product_service "store/internal/service/product"
	address_service "store/internal/service/address"
)

var Module = fx.Options(
	fx.Provide(config.LoadConfig),

	fx.Provide(config.NewDSN),
	fx.Provide(dbconn.NewConnection),
	fx.Provide(dbconn.NewDBTX),
	fx.Provide(dbconn.New),

	fx.Provide(client_repo.NewClientRepository),
	fx.Provide(supplier_repo.NewSupplierRepository),
	fx.Provide(image_repo.NewImageRepository),
	fx.Provide(category_repo.NewCategoryRepositry),
	fx.Provide(product_repo.NewProductRepository),
	fx.Provide(address_repo.NewAddressRepository),

	fx.Provide(client_service.NewClientService),
	fx.Provide(supplier_service.NewSupplierService),
	fx.Provide(image_service.NewImageService),
	fx.Provide(category_service.NewCategoryService),
	fx.Provide(product_service.NewProductService),
	fx.Provide(address_service.NewAddressService),

	fx.Provide(handler.NewClientHandler),
	fx.Provide(handler.NewSupplierHandler),
	fx.Provide(handler.NewImageHandler),
	fx.Provide(handler.NewCategoryHandler),
	fx.Provide(handler.NewProductHandler),
	fx.Provide(handler.NewAddressHandler),

	fx.Provide(router.NewRouter),

	fx.Invoke(startServer),
	fx.Invoke(registerShutdownHooks),
)

func startServer(
	lc fx.Lifecycle,
	cfg *config.Config,
	router http.Handler,
) {
	server := &http.Server{
		Addr:         cfg.ServerAddress(),
		Handler:      router,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Printf("Starting server on %s", server.Addr)
			go func() {
				if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					log.Fatalf("server failed: %v", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Shutting down server...")
			return server.Shutdown(ctx)
		},
	})
}

func registerShutdownHooks(
	lc fx.Lifecycle,
	pool *pgxpool.Pool,
) {
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			log.Println("Closing database connection...")
			pool.Close()
			return nil
		},
	})
}