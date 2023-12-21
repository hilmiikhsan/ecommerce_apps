package file

import (
	"github.com/ecommerce/infra/middleware"
	"github.com/ecommerce/infra/storage/images"
	"github.com/gofiber/fiber/v2"
)

func RegisterServiceFile(router fiber.Router, cloud images.Cloudinary) {
	service := NewFileService(cloud)
	handler := NewFileHandler(service)

	var fileRouter = router.Group("/v1/files")
	{
		fileRouter.Post("/upload", middleware.AuthMiddleware(), handler.Upload)
	}
}
