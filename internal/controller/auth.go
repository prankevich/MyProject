package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/prankevich/MyProject/internal/errs"
	"github.com/prankevich/MyProject/internal/models"
	"github.com/prankevich/MyProject/pkg"
	"net/http"
)

type SignUpRequest struct {
	FullName string `json:"full_name"`
	Username string `json:"user_name"`
	Password string `json:"password"`
}

// SignUp
// @Summary Регистрация
// @Description Создание нового аккаунта
// @Tags Auth
// @Consume json
// @Produce json
// @Param request_body body SignUpRequest true "Информация о новом аккаунте"
// @Failure 201 {object} CommonResponse
// @Failure 400 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /auth/sign-up [post]
func (ctrl *Controller) SignUp(c *gin.Context) {
	var input SignUpRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		ctrl.handleError(c, errors.Join(errs.ErrInvalidRequestBody, err))
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

	c.JSON(http.StatusCreated, CommonResponse{Message: "User created successfully!"})
}

type SignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenPairResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// SignIn
// @Summary Вход
// @Description Вход в аккаунт
// @Tags Auth
// @Consume json
// @Produce json
// @Param request_body body SignInRequest true "Данные для входа пользователя"
// @Success 200 {object} TokenPairResponse
// @Failure 400 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /auth/sign-in [post]
func (ctrl *Controller) SignIn(c *gin.Context) {
	var input SignInRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		ctrl.handleError(c, errors.Join(errs.ErrInvalidRequestBody, err))
		return
	}

	userID, userRole, err := ctrl.service.Authenticate(c, models.User{
		Username: input.Username,
		Password: input.Password,
	})
	if err != nil {
		ctrl.handleError(c, err)
		return
	}

	accessToken, refreshToken, err := ctrl.generateNewTokenPair(userID, userRole)
	if err != nil {
		ctrl.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, TokenPairResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

const (
	refreshTokenHeader = "Token"
)

// RefreshTokenPair
// @Summary Обновить пару токенов
// @Description Обновить пару токенов
// @Tags Auth
// @Produce json
// @Param Token header string true "вставьте refresh token"
// @Success 200 {object} TokenPairResponse
// @Failure 400 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /auth/refresh [get]
func (ctrl *Controller) RefreshTokenPair(c *gin.Context) {
	token, err := ctrl.extractTokenFromHeader(c, refreshTokenHeader)
	if err != nil {
		c.JSON(http.StatusUnauthorized, CommonError{Error: err.Error()})
		return
	}

	userID, isRefresh, userRole, err := pkg.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, CommonError{Error: err.Error()})
		return
	}

	if !isRefresh {
		c.JSON(http.StatusUnauthorized, CommonError{Error: "inappropriate token"})
		return
	}

	accessToken, refreshToken, err := ctrl.generateNewTokenPair(userID, userRole)
	if err != nil {
		ctrl.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, TokenPairResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}
