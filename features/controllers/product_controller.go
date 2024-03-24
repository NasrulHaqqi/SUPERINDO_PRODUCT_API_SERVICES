package controllers

import (
	"superindo-product-api/features/services"
)

type ProductController struct {
	ProductService services.ProductService
}

func NewProductController(productService services.ProductService) *ProductController {
    return &ProductController{ProductService: productService}
}

