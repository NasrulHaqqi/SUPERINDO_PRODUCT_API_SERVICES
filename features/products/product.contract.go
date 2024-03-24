package products

import (
	"superindo-product-api/features/models"
	"superindo-product-api/helper"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Request struct {

}

var Validator = validator.New()

func Validate(c *fiber.Ctx) error {
	var errors []*helper.ValidatorError
	body := Request{}
	err := c.BodyParser(&body)
	if err != nil {
		helper.Exception(err)
		return err
	}

	err = Validator.Struct(body)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var el helper.ValidatorError
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			errors = append(errors, &el)
		}
		return c.Status(fiber.StatusBadRequest).JSON(helper.Response{
			Code:    fiber.StatusBadRequest,
			Success: false,
			Message: "Validation error",
			Data:    errors,
		})
	}

	return c.Next()
}

type Service interface {
    AddProduct(product *models.Product) error
    GetAllProducts() ([]models.Product, error)
    GetProductByID(id int) (*models.Product, error)
    SearchProduct(query string) ([]models.Product, error)
    FilterProductByType(productType string) ([]models.Product, error)
    SortProductsBy(field string) ([]models.Product, error)
}

type Controller interface {
    AddProduct(ctx *fiber.Ctx) error
    GetAllProducts(ctx *fiber.Ctx) error
    GetProductByID(ctx *fiber.Ctx) error
    SearchProduct(ctx *fiber.Ctx) error
    FilterProductByType(ctx *fiber.Ctx) error
    SortProductsBy(ctx *fiber.Ctx) error
}