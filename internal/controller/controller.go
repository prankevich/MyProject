package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/prankevich/MyProject/internal/contracts"
	"github.com/prankevich/MyProject/internal/errs"
	"net/http"
)

type Controller struct {
	service contracts.ServiceI
	router  *gin.Engine
}

func New(service contracts.ServiceI) *Controller {
	return &Controller{
		service: service,
		router:  gin.Default(),
	}

}

func (ctrl *Controller) handleError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, errs.ErrUserNotfound) || errors.Is(err, errs.ErrNotfound):
		c.JSON(http.StatusNotFound, CommonError{Error: err.Error()})
	case errors.Is(err, errs.ErrInvalidUserID) || errors.Is(err, errs.ErrInvalidRequestBody):
		c.JSON(http.StatusBadRequest, CommonError{Error: err.Error()})
	case errors.Is(err, errs.ErrInvalidFieldValue):
		c.JSON(http.StatusUnprocessableEntity, CommonError{Error: err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, CommonError{Error: err.Error()})
	}
}
