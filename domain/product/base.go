package product

import (
	categoryRepository "github.com/ecommerce/domain/category/repository"
	merchantRepository "github.com/ecommerce/domain/merchant/repository"
	productRepository "github.com/ecommerce/domain/product/repository"
	"github.com/ecommerce/infra/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	Dbx *sqlx.DB
}

func RegisterServiceProduct(router fiber.Router, db DB) {
	productRepository := productRepository.NewProductRepository(db.Dbx)
	merchantRepository := merchantRepository.NewMerchantRepository(db.Dbx)
	categoryRepository := categoryRepository.NewCategoryRepository(db.Dbx)
	service := NewProductService(productRepository, merchantRepository, categoryRepository)
	handler := NewProductHandler(service)

	var productRouter = router.Group("/v1/products")
	{
		productRouter.Post("/", middleware.AuthMiddleware(), handler.CreateProduc)
		productRouter.Get("/", middleware.AuthMiddleware(), handler.GetListProduct)
		productRouter.Get("/id/:product_id", middleware.AuthMiddleware(), handler.GetDetailProduct)
		productRouter.Put("/id/:product_id", middleware.AuthMiddleware(), handler.UpdateProduct)
		productRouter.Get("/detail/:sku", middleware.AuthMiddleware(), handler.GetDetailProductUserPerspective)
	}
}
