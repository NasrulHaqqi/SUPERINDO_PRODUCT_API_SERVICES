package products

import (
	"errors"
	"strings"
	"superindo-product-api/features/models"
)

type ProductService struct {
	products []models.Product
}

func NewProductService() *ProductService {
	return &ProductService{
		products: make([]models.Product, 0),
	}
}

func (s *ProductService) AddProduct(product *models.Product) error {
	s.products = append(s.products, *product)
	return nil
}

func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	return s.products, nil
}

func (s *ProductService) GetProductByID(id int) (*models.Product, error) {
	for _, p := range s.products {
		if p.ID == id {
			return &p, nil
		}
	}
	return nil, errors.New("product not found")
}

func (s *ProductService) SearchProduct(query string) ([]models.Product, error) {
	var result []models.Product
	for _, p := range s.products {
		if containsIgnoreCase(p.Name, query) {
			result = append(result, p)
		}
	}
	return result, nil
}

func (s *ProductService) FilterProductByType(productType string) ([]models.Product, error) {
	var result []models.Product
	for _, p := range s.products {
		if p.Type == productType {
			result = append(result, p)
		}
	}
	return result, nil
}

func (s *ProductService) SortProductsBy(field string) ([]models.Product, error) {
	return s.products, nil
}

func containsIgnoreCase(s, substr string) bool {
	s, substr = strings.ToLower(s), strings.ToLower(substr)
	return strings.Contains(s, substr)
}