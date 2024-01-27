package main

import (
	"net/http"

	"github.com/eduardogomesf/go-api/configs"
	"github.com/eduardogomesf/go-api/internal/entity"
	"github.com/eduardogomesf/go-api/internal/infra/database"
	"github.com/eduardogomesf/go-api/internal/infra/webserver/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.User{}, &entity.Product{})

	r := chi.NewRouter()

	// middlewares
	r.Use(middleware.Logger)

	// handlers
	productDB := database.NewProductDB(db)
	productHandler := handlers.NewProductHandler(productDB)

	userDB := database.NewUserDB(db)
	userHandler := handlers.NewUserHandler(userDB)

	// products routes
	r.Post("/products", productHandler.CreateProduct)
	r.Get("/products/{id}", productHandler.GetProduct)
	r.Put("/products/{id}", productHandler.UpdateProduct)
	r.Delete("/products/{id}", productHandler.DeleteProduct)
	r.Get("/products", productHandler.GetProducts)

	// users routes
	r.Post("/users", userHandler.CreateUser)

	http.ListenAndServe(":8080", r)
}
