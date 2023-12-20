package user

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"

	"github.com/ecommerce/entity"
)

func WriteError(c *fiber.Ctx, err error) error {
	switch {
	case err == entity.ErrAddressIsRequired:
		return write(c, http.StatusBadRequest, "bad request", err.Error(), "40005", nil)
	case err == entity.ErrNameIsRequired:
		return write(c, http.StatusBadRequest, "bad request", err.Error(), "40005", nil)
	case err == entity.ErrDateOfBirtIsRequired:
		return write(c, http.StatusBadRequest, "bad request", err.Error(), "40006", nil)
	case err == entity.ErrPhoneNumberIsRequired:
		return write(c, http.StatusBadRequest, "bad request", err.Error(), "40001", nil)
	case err == entity.ErrGenderIsInvalid:
		return write(c, http.StatusBadRequest, "bad request", err.Error(), "40002", nil)
	case err == entity.ErrAddressIsRequired:
		return write(c, http.StatusBadRequest, "bad request", err.Error(), "40003", nil)
	case err == entity.ErrInvalidRole:
		return write(c, http.StatusUnauthorized, "unauthorized", err.Error(), "40102", nil)
	default:
		if iSSQLIntegrityConstraintViolation(err) {
			return write(c, http.StatusInternalServerError, "internal server error", "error repository", "50001", nil)
		}
		return write(c, http.StatusInternalServerError, "internal server error", err.Error(), "99999", nil)
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
