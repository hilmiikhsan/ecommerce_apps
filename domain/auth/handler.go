package auth

import (
	"fmt"

	"github.com/ecommerce/dto"
	"github.com/ecommerce/entity"
	logs "github.com/ecommerce/infra/logger"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	service AuthService
}

func NewAuthHandler(service AuthService) AuthHandler {
	return AuthHandler{
		service: service,
	}
}

func (a AuthHandler) Register(c *fiber.Ctx) error {
	var req dto.AuthRequest

	if err := c.BodyParser(&req); err != nil {
		return WriteError(c, err)
	}

	model, err := entity.NewAuth().Validate(req)
	if err != nil {
		logs.Logger(logs.GetFunctionPath(), logs.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return WriteError(c, err)
	}

	if err := a.service.Register(c.UserContext(), model); err != nil {
		logs.Logger(logs.GetFunctionPath(), logs.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return WriteError(c, err)
	}

	return WriteSuccess(c, "registration success", nil, fiber.StatusCreated)
}

func (a AuthHandler) Login(c *fiber.Ctx) error {
	var req dto.AuthRequest

	if err := c.BodyParser(&req); err != nil {
		return WriteError(c, err)
	}

	model, err := entity.NewAuth().Validate(req)
	if err != nil {
		logs.Logger(logs.GetFunctionPath(), logs.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return WriteError(c, err)
	}

	response, accessToken, err := a.service.Login(c.UserContext(), model)
	if err != nil {
		logs.Logger(logs.GetFunctionPath(), logs.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return WriteError(c, err)
	}

	payload := dto.LoginResponse{
		AccessToken: accessToken,
		Role:        response.Role,
	}

	return WriteSuccess(c, "login success", payload, fiber.StatusOK)
}

func (a AuthHandler) UpdateRole(c *fiber.Ctx) error {
	id := c.Locals("id").(string)
	email := c.Locals("email").(string)

	if err := a.service.UpdateRole(c.UserContext(), id, email); err != nil {
		logs.Logger(logs.GetFunctionPath(), logs.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return WriteError(c, err)
	}

	return WriteSuccess(c, "update role success", nil, fiber.StatusOK)
}
