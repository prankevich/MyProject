package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/prankevich/MyProject/internal/config"
	"github.com/prankevich/MyProject/internal/models"
	"github.com/prankevich/MyProject/pkg"
	"strings"
)

func (ctrl *Controller) extractTokenFromHeader(c *gin.Context, headerKey string) (string, error) {
	header := c.GetHeader(headerKey)

	if header == "" {
		return "", errors.New("empty authorization header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		return "", errors.New("invalid authorization header")
	}

	if len(headerParts[1]) == 0 {
		return "", errors.New("empty token")
	}

	return headerParts[1], nil
}

func (ctrl *Controller) generateNewTokenPair(userID int, userRole models.Role) (string, string, error) {
	accessToken, err := pkg.GenerateToken(userID,
		config.AppSettings.AuthParams.TtlMinutes,
		userRole, false)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := pkg.GenerateToken(userID,
		config.AppSettings.AuthParams.RefreshTTLdays,
		userRole, true)
	if err != nil {

		return "", "", err
	}

	return accessToken, refreshToken, nil
}
