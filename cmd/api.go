package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Abir-Zayn/kotoNilo/internal/products"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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
	productsService := products.NewService()
	productsHandler := products.NewHandler(productsService)
	r.Get("/products", productsHandler.ListProducts)
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
	// db driver

}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dsn string // user = password
}
