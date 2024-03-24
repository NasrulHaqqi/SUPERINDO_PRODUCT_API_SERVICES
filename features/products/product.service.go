package products

import (
	"database/sql"
	"errors"
	"fmt"
	"superindo-product-api/features/models"
)

type ProductService struct {
	products []models.Product
	db       *sql.DB
}

func NewProductService() *ProductService {
    return &ProductService{
        products: make([]models.Product, 0),
    }
}

func (s *ProductService) SetDB(db *sql.DB) {
	s.db = db
}

func (s *ProductService) AddProduct(product *models.Product) error {
	if s.db == nil {
        return errors.New("database connection is not set")
    }

	query := "INSERT INTO products (name, type, price) VALUES ($1, $2, $3)"
    _, err := s.db.Exec(query, product.Name, product.Type, product.Price)
    if err != nil {
        return fmt.Errorf("failed to add product to database: %v", err)
    }

    return nil
}

func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	query := "SELECT id, name, type, price, created_at FROM products"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Type, &p.Price, &p.CreatedAt); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (s *ProductService) GetProductByID(id int) (*models.Product, error) {
	query := "SELECT id, name, type, price, created_at FROM products WHERE id = $1"
	row := s.db.QueryRow(query, id)
	var p models.Product
	err := row.Scan(&p.ID, &p.Name, &p.Type, &p.Price, &p.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}
	return &p, nil
}

func (s *ProductService) SearchProduct(query string) ([]models.Product, error) {
	query = "SELECT id, name, type, price, created_at FROM products WHERE name ILIKE '%' || $1 || '%'"
	rows, err := s.db.Query(query, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Type, &p.Price, &p.CreatedAt); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (s *ProductService) FilterProductByType(productType string) ([]models.Product, error) {
	query := "SELECT id, name, type, price, created_at FROM products WHERE type = $1"
	rows, err := s.db.Query(query, productType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Type, &p.Price, &p.CreatedAt); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (s *ProductService) SortProductsBy(field string) ([]models.Product, error) {
	var orderBy string
	switch field {
	case "date":
		orderBy = "created_at"
	case "price":
		orderBy = "price"
	case "name":
		orderBy = "name"
	default:
		return nil, fmt.Errorf("unsupported field for sorting: %s", field)
	}

	query := fmt.Sprintf("SELECT id, name, type, price, created_at FROM products ORDER BY %s", orderBy)
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Type, &p.Price, &p.CreatedAt); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}