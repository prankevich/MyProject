package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/prankevich/MyProject/internal/errs"
	"github.com/prankevich/MyProject/internal/models"
	"net/http"
)

type SignInRequest struct {
	Username string `json:"user_name"`
	Password string `json:"password"`
}
type SignInResponse struct {
	Token string `json:"token"`
}
type SignUpRequest struct {
	FullName string `json:"full_name" db:"full_name"`
	Username string `json:"user_name" db:"user_name"`
	Password string `json:"password" db:"password"`
}

func (ctrl *Controller) SignUp(c *gin.Context) {
	var input SignUpRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.service.CreateUser(c, models.User{
		FullName: input.FullName,
		Username: input.Username,
		Password: input.Password,
	}); err != nil {
		ctrl.handleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, CommonResponse{Message: "User created successfully"})

}
func (ctrl *Controller) SignIn(c *gin.Context) {
	var input SignInRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		ctrl.handleError(c, errors.Join(errs.ErrInvalidRequestBody, err))
		return
	}
	token, err := ctrl.service.Authenticate(c, models.User{
		Username: input.Username,
		Password: input.Password,
	})
	if err != nil {
		ctrl.handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, SignInResponse{Token: token})

}
