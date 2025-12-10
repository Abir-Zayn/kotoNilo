package main

import (
	"log"
	"net/http"
	"time"

	repo "github.com/Abir-Zayn/kotoNilo/internal/adapters/postgresql/sqlc"
	"github.com/Abir-Zayn/kotoNilo/internal/products"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
)

// mount
func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	// A middlware Stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	// r.Use(app.metrics())
	// r.Use(middleware.URLFormat)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("doing good."))
	})

	// Products
	q := repo.New(app.db)
	productsService := products.NewService(q)
	productsHandler := products.NewHandler(productsService)
	r.Get("/products", productsHandler.ListProducts)
	r.Post("/products", productsHandler.CreateProduct)
	return r
}

// run
func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute * 1,
	}
	log.Printf("Starting server on %s\n", app.config.addr)
	return srv.ListenAndServe()
}

type application struct {
	config config
	// logger
	db *pgxpool.Pool
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dsn string // user = password
}
