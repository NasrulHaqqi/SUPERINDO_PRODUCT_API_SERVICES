package wire

import (
	"superindo-product-api/config"
	"superindo-product-api/features/products"

	"github.com/google/wire"
)

func InitializeProductController(cfg config.Config) *products.ProductController {
	wire.Build(
        products.NewProductService,
        wire.Struct(new(products.ProductController), "*"),
    )
    return &products.ProductController{}
}