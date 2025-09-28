package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/prankevich/MyProject/internal/errs"
	"github.com/prankevich/MyProject/internal/models"
	"net/http"
	"strconv"
)

// GetAllEmployees
// @Summary Получение данных пользователя
// @Description Получения списка всех пользователей
// @Tags User
// @Produce json
// @Success 200 {array} CommonResponse
// @Failure 500 {object} CommonError
// @Router /users [get]
func (ctrl *Controller) GetAllEmployees(c *gin.Context) {
	employees, err := ctrl.service.GetAllEmployees()
	if err != nil {
		ctrl.handleError(c, err)
	}

	c.JSON(http.StatusOK, employees)
}

// GetEmployeesByID
// @Summary Получение данных пользователя по ID
// @Description Получения информации о пользователе по ID
// @Tags User
// @Produce json
// @Param id path int true "id пользователя"
// @Success 200  {object} CommonResponse
// @Failure 400 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /users/{id} [get]
func (ctrl *Controller) GetEmployeesByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctrl.handleError(c, errs.ErrInvalidEmployeesID)
		return
	}
	employees, err := ctrl.service.GetEmployeesByID(id)
	if err != nil || id < 1 {
		ctrl.handleError(c, errs.ErrInvalidEmployeesID)
		return
	}
	c.JSON(http.StatusOK, employees)
}

type CreateEmployeesRequest struct {
	ID    int    `json:"id"`
	Age   int    `json:"age"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// CreateEmployees
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
func (ctrl *Controller) CreateEmployees(c *gin.Context) {
	var employees models.Employees
	if err := c.ShouldBindJSON(&employees); err != nil {
		ctrl.handleError(c, errs.ErrInvalidRequestBody)
		return
	}
	if employees.Name == "" || employees.Email == "" || employees.Age < 0 || employees.ID < 0 {
		ctrl.handleError(c, errs.ErrInvalidFieldValue)
		return
	}
	if err := ctrl.service.CreateEmployees(employees); err != nil {
		ctrl.handleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, CommonError{"Employees created"})
}

// UpdateEmployeesByID
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
func (ctrl *Controller) UpdateEmployeesByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctrl.handleError(c, errs.ErrInvalidEmployeesID)
		return
	}

	var employees models.Employees
	if err = c.ShouldBindJSON(&employees); err != nil {
		ctrl.handleError(c, errs.ErrInvalidRequestBody)
		return
	}
	if employees.Name == "" || employees.Email == "" || employees.Age < 0 || employees.ID < 0 {
		ctrl.handleError(c, errs.ErrInvalidFieldValue)
		return
	}

	employees.ID = id

	if err = ctrl.service.UpdateEmployeesByID(employees); err != nil {
		ctrl.handleError(c, err)
		return
	}
	ctrl.logger.Info().Str("func", "controller.UpdateEmployeesByID").Int("employee_id", employees.ID).
		Msg("Сотрудник успешно обновлён")

	c.JSON(http.StatusOK, gin.H{
		"message": "Employees updated",
	})

}

// DeleteEmployeesByID
// @Summary  Удаление
// @Description   Удаление данных пользователя по ID
// @Tags User
// @Produce json
// @Param id path int true "id пользователя"
// @Success 200  {object} CommonResponse
// @Failure 400 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /users/{id} [DELETE]
func (ctrl *Controller) DeleteEmployeesByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, CommonError{err.Error()})
		return
	}
	if id <= 0 {
		c.JSON(http.StatusBadRequest, CommonError{"Employees id must be positive"})
		return
	}
	if err = ctrl.service.DeleteEmployeesByID(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctrl.logger.Info().Str("func", "controller.DeleteEmployeesByID").Int("id", id).
		Msg("Сотрудник успешно удален")
	c.JSON(http.StatusOK, gin.H{
		"message": "Employees deleted",
	})
}
