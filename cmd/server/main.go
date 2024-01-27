package main

import (
	"net/http"

	"github.com/eduardogomesf/go-api/configs"
	"github.com/eduardogomesf/go-api/internal/entity"
	"github.com/eduardogomesf/go-api/internal/infra/database"
	"github.com/eduardogomesf/go-api/internal/infra/webserver/handlers"
	"github.com/go-chi/chi"
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

	productDB := database.NewProductDB(db)
	productHandler := handlers.NewProductHandler(productDB)

	r := chi.NewRouter()
	r.Post("/products", productHandler.CreateProduct)
	http.ListenAndServe(":8080", r)
}
