package category

import (
	"github.com/ecommerce/domain/category/repository"
	"github.com/ecommerce/infra/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	Dbx *sqlx.DB
}

func RegisterServiceCategory(router fiber.Router, db DB) {
	categoryRepository := repository.NewCategoryRepository(db.Dbx)
	service := NewCategoryService(categoryRepository)
	handler := NewCategoryHandler(service)

	var categoryRouter = router.Group("/v1/categories")
	{
		categoryRouter.Post("/", middleware.AuthMiddleware(), handler.CreateCategory)
		categoryRouter.Get("/", middleware.AuthMiddleware(), handler.GetListCategory)
	}
}
