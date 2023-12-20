package category

import (
	"fmt"

	"github.com/ecommerce/dto"
	"github.com/ecommerce/entity"
	logs "github.com/ecommerce/infra/logger"
	"github.com/gofiber/fiber/v2"
)

type CategoryHandler struct {
	service Service
}

func NewCategoryHandler(service Service) CategoryHandler {
	return CategoryHandler{
		service: service,
	}
}

func (ca CategoryHandler) CreateCategory(c *fiber.Ctx) error {
	var req dto.CreateCategoryRequest
	id := c.Locals("id").(string)

	if err := c.BodyParser(&req); err != nil {
		return WriteError(c, err)
	}

	model, err := entity.NewCategory().Validate(req, id)
	if err != nil {
		logs.Logger(logs.GetFunctionPath(), logs.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return WriteError(c, err)
	}

	if err := ca.service.CreateCategory(c.UserContext(), model); err != nil {
		logs.Logger(logs.GetFunctionPath(), logs.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return WriteError(c, err)
	}

	return WriteSuccess(c, "create category success", nil, fiber.StatusCreated)
}

func (ca CategoryHandler) GetListCategory(c *fiber.Ctx) error {
	responses, err := ca.service.GetListCategory(c.UserContext())
	if err != nil {
		logs.Logger(logs.GetFunctionPath(), logs.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return WriteError(c, err)
	}

	return WriteSuccess(c, "get categories success", responses, fiber.StatusOK)
}
