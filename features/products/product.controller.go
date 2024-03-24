package products

import (
	"superindo-product-api/features/models"

	"github.com/gofiber/fiber/v2"
)

type ProductController struct {
	Service Service
}

func (controller *ProductController) AddProduct(ctx *fiber.Ctx) error {
	var product models.Product

	if err := ctx.BodyParser(&product); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := controller.Service.AddProduct(&product); err != nil {

        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to add product to database",
        })
    }
	
    return ctx.Status(fiber.StatusCreated).JSON(product)
}

func (controller *ProductController) GetAllProducts(ctx *fiber.Ctx) error {
	products, err := controller.Service.GetAllProducts()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get products",
		})
	}

	return ctx.JSON(products)
}

func (controller *ProductController) GetProductByID(ctx *fiber.Ctx) error {
	var requestData struct {
        ID int `json:"id"`
    }

	if err := ctx.BodyParser(&requestData); err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request format",
        })
    }

    product, err := controller.Service.GetProductByID(requestData.ID)
    if err != nil {
        return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Product not found",
        })
    }

	return ctx.JSON(product)
}

func (controller *ProductController) GetProductByName(ctx *fiber.Ctx) error {
	var requestData struct {
        Name string `json:"name"`
    }

	if err := ctx.BodyParser(&requestData); err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request format",
        })
    }

    product, err := controller.Service.GetProductByName(requestData.Name)
    if err != nil {
        return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Product not found",
        })
    }

	return ctx.JSON(product)
}

func (controller *ProductController) FilterProductByType(ctx *fiber.Ctx) error {
	productType := ctx.Params("type")

	products, err := controller.Service.FilterProductByType(productType)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to filter products by type",
		})
	}

	return ctx.JSON(products)
}

func (controller *ProductController) SortProductsBy(ctx *fiber.Ctx) error {
    field := ctx.Params("field")

    if field == "" {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Field parameter 'field' is required",
        })
    }

    products, err := controller.Service.SortProductsBy(field)
    if err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to sort products",
        })
    }

    return ctx.JSON(products)
}