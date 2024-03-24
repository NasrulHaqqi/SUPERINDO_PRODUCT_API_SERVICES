package repositories

import "superindo-product-api/features/models"

type ProductRepository interface {
    AddProduct(product *models.Product) error
    GetAllProducts() ([]models.Product, error)
    GetProductByID(id int) (*models.Product, error)
    SearchProduct(query string) ([]models.Product, error)
    FilterProductByType(productType string) ([]models.Product, error)
    SortProductsBy(field string) ([]models.Product, error)
}