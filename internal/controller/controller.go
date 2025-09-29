package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/prankevich/MyProject/internal/contracts"
	"github.com/prankevich/MyProject/internal/errs"
	"github.com/rs/zerolog"
	"net/http"
)

type Controller struct {
	service contracts.ServiceI
	router  *gin.Engine
	logger  zerolog.Logger

}

func New(service contracts.ServiceI, logger zerolog.Logger) *Controller {
	return &Controller{
		service: service,
		router:  gin.Default(),
		logger:  logger,
	}

}

func (ctrl *Controller) handleError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, errs.ErrEmployeesNotfound) || errors.Is(err, errs.ErrNotfound) ||
		errors.Is(err, errs.ErrUserNotFound):
		c.JSON(http.StatusNotFound, CommonError{Error: err.Error()})
	case errors.Is(err, errs.ErrIncorrectUsernameOrPassword):
		c.JSON(http.StatusUnauthorized, CommonError{Error: err.Error()})
	case errors.Is(err, errs.ErrInvalidEmployeesID) || errors.Is(err, errs.ErrInvalidRequestBody):
		c.JSON(http.StatusBadRequest, CommonError{Error: err.Error()})
	case errors.Is(err, errs.ErrInvalidFieldValue):
		c.JSON(http.StatusUnprocessableEntity, CommonError{Error: err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, CommonError{Error: err.Error()})
	}
}
