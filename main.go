package main

import (
	"log"
	"superindo-product-api/features/products"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	productService := products.NewProductService()
	productController := &products.ProductController{Service: productService}

	app.Post("/api/products", productController.AddProduct)
	app.Get("/api/products", productController.GetAllProducts)
	app.Get("/api/products/:id", productController.GetProductByID)
	app.Get("/api/products/search", productController.SearchProduct)
	app.Get("/api/products/type/:type", productController.FilterProductByType)
	app.Get("/api/products/sort", productController.SortProductsBy)

	log.Fatal(app.Listen(":3000"))
}
