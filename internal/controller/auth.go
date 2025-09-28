package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/prankevich/MyProject/internal/models"
	"net/http"
)

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

}
