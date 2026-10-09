package user

import (
	"errors"
	"net/http"

	"goticket/internal/httpresponse"
	"goticket/internal/user/dto"

	"github.com/labstack/echo/v5"
)

type handler struct {
	service *service
}

func NewHandler(service *service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) CreateUser(c *echo.Context) error {
	var req dto.CreateUserRequest

	// Bind request body
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid request payload",
			Details: err.Error(),
		})
	}

	// Validate request
	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.Error{
			Code:    http.StatusBadRequest,
			Message: "Validation failed",
			Details: err.Error(),
		})
	}

	// Create user
	response, err := h.service.CreateUser(req)

	if err != nil {
		if errors.Is(err, ErrorAlreadyExist) {
			return c.JSON(http.StatusConflict, httpresponse.Error{
				Code:    http.StatusConflict,
				Message: "Failed to create user",
				Details: err.Error(),
			
			})
		}

		return c.JSON(http.StatusInternalServerError, httpresponse.Error{
			Code:    http.StatusInternalServerError,
			Message: "Failed to create user",
			Details: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, response)
}

func (h *handler) LoginUser(c *echo.Context) error {
	var req dto.LoginRequest // input

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid request payload",
			Details: err.Error(),
		})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.Error{
			Code:    http.StatusBadRequest,
			Message: "Validation failed",
			Details: err.Error(),
		})
	}

	response, err := h.service.LoginUser(req)

	if err != nil {
		if errors.Is(err, ErrInvalidCredentials) {
			return c.JSON(http.StatusUnauthorized, httpresponse.Error{
				Code:    http.StatusUnauthorized,
				Message: "Cannot login user",
				Details: err.Error(),
			})
		}

		return c.JSON(http.StatusInternalServerError, httpresponse.Error{
			Code:    http.StatusInternalServerError,
			Message: "Failed to login user",
			Details: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response)

}