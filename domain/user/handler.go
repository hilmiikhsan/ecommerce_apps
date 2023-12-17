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

// func (a AuthHandler) Login(c *fiber.Ctx) error {
// 	var req dto.AuthRequest

// 	if err := c.BodyParser(&req); err != nil {
// 		return WriteError(c, err)
// 	}

// 	model, err := entity.NewAuth().Validate(req)
// 	if err != nil {
// 		logs.Logger(logs.GetFunctionPath(), logs.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
// 		return WriteError(c, err)
// 	}

// 	response, accessToken, err := a.service.Login(c.UserContext(), model)
// 	if err != nil {
// 		logs.Logger(logs.GetFunctionPath(), logs.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
// 		return WriteError(c, err)
// 	}

// 	payload := dto.LoginResponse{
// 		AccessToken: accessToken,
// 		Role:        response.Role,
// 	}

// 	return WriteSuccess(c, "login success", payload, fiber.StatusOK)
// }

// func (a AuthHandler) UpdateRole(c *fiber.Ctx) error {
// 	email := c.Locals("email").(string)

// 	if err := a.service.UpdateRole(c.UserContext(), email); err != nil {
// 		logs.Logger(logs.GetFunctionPath(), logs.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
// 		return WriteError(c, err)
// 	}

// 	return WriteSuccess(c, "update role success", nil, fiber.StatusOK)
// }
