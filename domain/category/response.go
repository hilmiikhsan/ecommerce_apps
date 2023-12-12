package category

import (
	"net/http"

	"github.com/ecommerce/entity"
	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
)

func WriteError(c *fiber.Ctx, err error) error {
	switch {
	case err == entity.ErrCategoryNameIsRequired:
		return write(c, http.StatusBadRequest, "bad request", err.Error(), "40001", nil)
	default:
		if iSSQLIntegrityConstraintViolation(err) {
			return write(c, http.StatusInternalServerError, "internal server error", "error repository", "50001", nil)
		}
		return write(c, http.StatusInternalServerError, "internal server error", "unknown error", "99999", nil)
	}
}

func WriteSuccess(c *fiber.Ctx, message string, payload interface{}, statusCode int) error {
	resp := response{
		Success: true,
		Message: message,
		Payload: payload,
	}
	c = c.Status(statusCode)
	return c.JSON(resp)
}

type response struct {
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
	Payload   interface{} `json:"payload,omitempty"`
	Error     *string     `json:"error,omitempty"`
	ErrorCode *string     `json:"error_code,omitempty"`
}

func write(c *fiber.Ctx, statusCode int, message, errorMessage, errorCode string, payload interface{}) error {
	c = c.Status(statusCode)
	isSuccess := statusCode >= 200 && statusCode < 300

	if isSuccess {
		return c.JSON(response{
			Success: true,
			Message: message,
			Payload: payload,
		})
	}

	return c.JSON(response{
		Success:   false,
		Message:   message,
		Error:     &errorMessage,
		ErrorCode: &errorCode,
	})
}

func iSSQLIntegrityConstraintViolation(err error) bool {
	if err, ok := err.(*pq.Error); ok && err.Code == "42601" {
		return true
	}
	return false
}
