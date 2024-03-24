package main

import (
	"database/sql"
	"fmt"
	"log"
	"superindo-product-api/config"
	"superindo-product-api/features/models"
	"superindo-product-api/features/products"
	"time"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func seedProducts(db *sql.DB) error {
    products := []models.Product{
        {Name: "Product 1", Type: "Type A", Price: 19.99, CreatedAt: time.Now()},
        {Name: "Product 2", Type: "Type B", Price: 29.99, CreatedAt: time.Now()},
        {Name: "Product 3", Type: "Type C", Price: 39.99, CreatedAt: time.Now()},
    }

    for _, p := range products {
        _, err := db.Exec("INSERT INTO products (name, type, price, created_at) VALUES ($1, $2, $3, $4)", p.Name, p.Type, p.Price, p.CreatedAt)
        if err != nil {
            return err
        }
    }

    return nil
}

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

	if err := seedProducts(db); err != nil {
        log.Fatal("Failed to seed products table:", err)
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
