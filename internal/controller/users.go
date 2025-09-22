package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/prankevich/MyProject/internal/errs"
	"github.com/prankevich/MyProject/internal/models"
	"net/http"
	"strconv"
)

// GetAllUsers
// @Summary Получение данных пользователя
// @Description Получения списка всех пользователей
// @Tags User
// @Produce json
// @Success 200 {array} CommonResponse
// @Failure 500 {object} CommonError
// @Router /users [get]
func (ctrl *Controller) GetAllUsers(c *gin.Context) {
	users, err := ctrl.service.GetAllUsers()
	if err != nil {
		ctrl.handleError(c, err)
	}

	c.JSON(http.StatusOK, users)
}

// GetUsersByID
// @Summary Получение данных пользователя по ID
// @Description Получения информации о пользователе по ID
// @Tags User
// @Produce json
// @Param id path int true "id пользователя"
// @Success 200  {object} CommonResponse
// @Failure 400 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /users/{id} [get]
func (ctrl *Controller) GetUsersByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctrl.handleError(c, errs.ErrInvalidUserID)
		return
	}
	users, err := ctrl.service.GetUsersByID(id)
	if err != nil || id < 1 {
		ctrl.handleError(c, errs.ErrInvalidUserID)
		return
	}
	c.JSON(http.StatusOK, users)
}

type CreateUsersRequest struct {
	ID    int    `json:"id"`
	Age   int    `json:"age"`
	Name  string `son:"name"`
	Email string `json:"email"`
}

// CreateUsersByID
// @Summary  Создание пользователя
// @Description Создание карточки пользователя
// @Tags User
// @Consume json
// @Produce json
// @Param request_body body CreateUsersRequest true "информация о новом пользователе"
// @Success 200 {object} CommonResponse
// @Failure 400 {object} CommonError
// @Failure 404 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /users/{id} [put]
func (ctrl *Controller) CreateUsersByID(c *gin.Context) {
	var users models.User
	if err := c.ShouldBindJSON(&users); err != nil {
		ctrl.handleError(c, errs.ErrInvalidRequestBody)
		return
	}
	if users.Name == "" || users.Email == "" || users.Age < 0 || users.ID < 0 {
		ctrl.handleError(c, errs.ErrInvalidFieldValue)
		return
	}
	if err := ctrl.service.CreateUsersByID(users); err != nil {
		ctrl.handleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, CommonError{"User created"})
}

// UpdateUsersByID
// @Summary  Обновление данных пользователя
// @Description  Обновление данных пользователя по ID
// @Tags User
// @Consume json
// @Produce json
// @Param request_body body CreateUsersRequest true "информация о пользователе"
// @Success 200 {object} CommonResponse
// @Failure 400 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /users/{id} [put]
func (ctrl *Controller) UpdateUsersByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctrl.handleError(c, errs.ErrInvalidUserID)
		return
	}

	var users models.User
	if err = c.ShouldBindJSON(&users); err != nil {
		ctrl.handleError(c, errs.ErrInvalidRequestBody)
		return
	}
	if users.Name == "" || users.Email == "" || users.Age < 0 || users.ID < 0 {
		ctrl.handleError(c, errs.ErrInvalidFieldValue)
		return
	}

	users.ID = id

	if err = ctrl.service.UpdateUsersByID(users); err != nil {
		ctrl.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Users updated",
	})
}

// DeleteUserByID
// @Summary  Удаление
// @Description   Удаление данных пользователя по ID
// @Tags User
// @Produce json
// @Param id path int true "id пользователя"
// @Success 200  {object} CommonResponse
// @Failure 400 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /users/{id} [DELETE]
func (ctrl *Controller) DeleteUserByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, CommonError{err.Error()})
		return
	}
	if id <= 0 {
		c.JSON(http.StatusBadRequest, CommonError{"user id must be positive"})
		return
	}
	if err = ctrl.service.DeleteUsersByID(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Users deleted",
	})
}
