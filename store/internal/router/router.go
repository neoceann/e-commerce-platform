package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/swaggo/http-swagger"
	"go.uber.org/fx"
	"net/http"
	"store/internal/handler"
)

type RouterParams struct {
	fx.In

	ClientHandler   *handler.ClientHandler
	SupplierHandler *handler.SupplierHandler
	ImageHandler    *handler.ImageHandler
	CategoryHandler *handler.CategoryHandler
	ProductHandler  *handler.ProductHandler
	AddressHandler  *handler.AddressHandler
	AuthHadler		*handler.AuthHandler
}

func NewRouter(p RouterParams) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodGet {
				w.Header().Set("Cache-Control", "max-age=30, public")
			} else {
				w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate")
			}
			next.ServeHTTP(w, r)
		})
	})

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
	}))

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/clients", func(r chi.Router) {
			r.Post("/", p.ClientHandler.CreateClient)
			r.Delete("/{id}", p.ClientHandler.DeleteClient)
			r.Get("/by_name", p.ClientHandler.GetClientsByFullName)
			r.Get("/by_page", p.ClientHandler.GetClientsWithPage)
			r.Patch("/{id}/address", p.ClientHandler.UpdateClientAddr)
		})
		r.Route("/suppliers", func(r chi.Router) {
			r.Post("/", p.SupplierHandler.CreateSupplier)
			r.Delete("/{id}", p.SupplierHandler.DeleteSupplier)
			r.Get("/", p.SupplierHandler.GetAllSuppliers)
			r.Get("/{id}", p.SupplierHandler.GetSupplierByID)
			r.Patch("/{id}/address", p.SupplierHandler.UpdateSupplierAddr)
		})
		r.Route("/images", func(r chi.Router) {
			r.Post("/", p.ImageHandler.CreateImage)
			r.Delete("/{id}", p.ImageHandler.DeleteImage)
			r.Patch("/{id}/update", p.ImageHandler.UpdateImage)
			r.Get("/{id}", p.ImageHandler.GetImageByID)
			r.Get("/product/{id}", p.ImageHandler.GetImagesByProductID)
		})
		r.Route("/categories", func(r chi.Router) {
			r.Post("/", p.CategoryHandler.CreateCategory)
			r.Get("/", p.CategoryHandler.GetAllCategories)
		})
		r.Route("/products", func(r chi.Router) {
			r.Post("/", p.ProductHandler.CreateProduct)
			r.Patch("/{id}/increase", p.ProductHandler.IncreaseProductStock)
			r.Patch("/{id}/decrease", p.ProductHandler.DecreaseProductStock)
			r.Get("/{id}", p.ProductHandler.GetProductByID)
			r.Get("/available", p.ProductHandler.GetAvailableProducts)
			r.Delete("/{id}", p.ProductHandler.DeleteProductByID)
		})
		r.Route("/addresses", func(r chi.Router) {
			r.Get("/{id}", p.AddressHandler.GetAddressByID)
		})
		r.Post("/register", p.AuthHadler.Register)
		r.Post("/auth", p.AuthHadler.Auth)
	})

	return r
}
