package main

import (
	"database/sql"
	"fmt"
	"log"
	"superindo-product-api/config"
	"superindo-product-api/features/products"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.Config{
		DBHost:     "localhost",
		DBPort:     "5432",
		DBUser:     "postgres",
		DBPassword: "password",
		DBName:     "superindo_db",
		RedisHost:  "localhost",
		RedisPort:  "6379",
	}

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	app := fiber.New()

	productService := products.NewProductService()
	
	productService.SetDB(db)
	
	productController := &products.ProductController{Service: productService}

	app.Post("/api/products", productController.AddProduct)
	app.Get("/api/products", productController.GetAllProducts)
	app.Get("/api/products/:id", productController.GetProductByID)
	app.Get("/api/products/search", productController.SearchProduct)
	app.Get("/api/products/type/:type", productController.FilterProductByType)
	app.Get("/api/products/sort/:field", productController.SortProductsBy)

	log.Fatal(app.Listen(":3000"))
}
