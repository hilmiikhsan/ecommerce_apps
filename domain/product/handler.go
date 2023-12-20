package product

import (
	"fmt"
	"strconv"

	"github.com/ecommerce/dto"
	"github.com/ecommerce/entity"
	logs "github.com/ecommerce/infra/logger"
	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	service Service
}

func NewProductHandler(service Service) ProductHandler {
	return ProductHandler{
		service: service,
	}
}

func (p ProductHandler) CreateProducts(c *fiber.Ctx) error {
	var req dto.CreateOrUpdateProductRequest

	id := c.Locals("id").(string)

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	model, err := entity.NewProduct().Validate(req, id)
	if err != nil {
		logs.Logger(logs.GetFunctionPath(), logs.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return WriteError(c, err)
	}

	if err := p.service.CreateProduct(c.UserContext(), model, id); err != nil {
		logs.Logger(logs.GetFunctionPath(), logs.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return WriteError(c, err)
	}

	return WriteSuccess(c, "create product success", nil, nil, fiber.StatusCreated)
}

func (p ProductHandler) GetListProduct(c *fiber.Ctx) error {
	id := c.Locals("id").(string)
	queryParam := c.Query("query")

	limit := c.Query("limit", "10")
	limitValue, err := strconv.Atoi(limit)
	if err != nil {
		logs.Logger(logs.GetFunctionPath(), logs.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return WriteError(c, err)
	}

	page := c.Query("page", "1")
	pageValue, err := strconv.Atoi(page)
	if err != nil {
		logs.Logger(logs.GetFunctionPath(), logs.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return WriteError(c, err)
	}

	responses, totalData, err := p.service.GetListProduct(c.UserContext(), id, queryParam, limitValue, pageValue)
	if err != nil {
		logs.Logger(logs.GetFunctionPath(), logs.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return WriteError(c, err)
	}

	if totalData == 0 {
		return WriteSuccess(c, "get products success", responses, nil, fiber.StatusOK)
	}

	paginationResponse := dto.NewPaginationResponse(queryParam, limitValue, pageValue, totalData)

	return WriteSuccess(c, "get products success", responses, paginationResponse, fiber.StatusOK)
}

func (p ProductHandler) GetDetailProduct(c *fiber.Ctx) error {
	id := c.Locals("id").(string)
	productId := c.Params("product_id")

	productIdValue, err := strconv.Atoi(productId)
	if err != nil {
		logs.Logger(logs.GetFunctionPath(), logs.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return WriteError(c, err)
	}

	response, err := p.service.GetDetailProduct(c.UserContext(), productIdValue, id)
	if err != nil {
		logs.Logger(logs.GetFunctionPath(), logs.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return WriteError(c, err)
	}

	return WriteSuccess(c, "get product success", response, nil, fiber.StatusOK)
}

func (p ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	var req dto.CreateOrUpdateProductRequest
	id := c.Locals("id").(string)
	productId := c.Params("product_id")

	productIdValue, err := strconv.Atoi(productId)
	if err != nil {
		logs.Logger(logs.GetFunctionPath(), logs.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return WriteError(c, err)
	}

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	model, err := entity.NewProduct().Validate(req, id)
	if err != nil {
		logs.Logger(logs.GetFunctionPath(), logs.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return WriteError(c, err)
	}
	model.ID = productIdValue

	if err := p.service.UpdateProduct(c.UserContext(), model, id); err != nil {
		logs.Logger(logs.GetFunctionPath(), logs.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return WriteError(c, err)
	}

	return WriteSuccess(c, "update product success", nil, nil, fiber.StatusOK)
}

func (p ProductHandler) GetDetailProductUserPerspective(c *fiber.Ctx) error {
	productSku := c.Params("sku")

	response, err := p.service.GetDetailProductUserPerspective(c.UserContext(), productSku)
	if err != nil {
		logs.Logger(logs.GetFunctionPath(), logs.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return WriteError(c, err)
	}

	return WriteSuccess(c, "get products success", response, nil, fiber.StatusOK)
}
