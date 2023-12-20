package user

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/ecommerce/dto"
	"github.com/ecommerce/entity"
	logs "github.com/ecommerce/infra/logger"
)

type AuthHandler struct {
	service UserService
}

func NewAuthHandler(service UserService) AuthHandler {
	return AuthHandler{
		service: service,
	}
}

func (a AuthHandler) CreateUserProfile(c *fiber.Ctx) error {
	var req dto.CreateOrUpdateUserRequest
	id := c.Locals("id").(string)
	email := c.Locals("email").(string)

	if err := c.BodyParser(&req); err != nil {
		return WriteError(c, err)
	}

	model, err := entity.NewUser().Validate(req, id)
	if err != nil {
		logs.Logger(logs.GetFunctionPath(), logs.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return WriteError(c, err)
	}

	if err := a.service.CreateUserProfile(c.UserContext(), model, email); err != nil {
		logs.Logger(logs.GetFunctionPath(), logs.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return WriteError(c, err)
	}
	return WriteSuccess(c, "registration success", nil, fiber.StatusCreated)
}

func (a AuthHandler) UpdateUserProfile(c *fiber.Ctx) error {
	var req dto.CreateOrUpdateUserRequest
	id := c.Locals("id").(string)

	// name := c.Locals("name").(string)
	// dateOfBirth := c.Locals("name").(string)
	// phoneNumber := c.Locals("name").(string)
	// gender := c.Locals("name").(string)
	// address := c.Locals("name").(string)
	// imageUrl := c.Locals("name").(string)

	if err := c.BodyParser(&req); err != nil {
		return WriteError(c, err)
	}

	model, err := entity.NewUser().Validate(req, id)
	if err != nil {
		logs.Logger(logs.GetFunctionPath(), logs.LoggerLevelError, fmt.Sprintf("Error : %", err.Error()))
		return WriteError(c, err)
	}

	if err := a.service.repository.UpdateUser(c.UserContext(), id, model); err != nil {
		logs.Logger(logs.GetFunctionPath(), logs.LoggerLevelError, fmt.Sprintf("Error : %", err.Error()))
		return WriteError(c, err)
	}
	return WriteSuccess(c, "update profile success", nil, fiber.StatusOK)
}
