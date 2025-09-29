package pkg

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/prankevich/MyProject/internal/models"
	"os"
	"time"
)

type CustomClaims struct {
	jwt.StandardClaims
	UserID    int         `json:"user_id"`
	Role      models.Role `json:"role"`
	IsRefresh bool        `json:"is_refresh"`
}

func GenerateToken(userID, ttl int, Role models.Role, isRefresh bool) (string, error) {
	claims := CustomClaims{
		StandardClaims: jwt.StandardClaims{},
		UserID:         userID,
		IsRefresh:      isRefresh,
		Role:           Role,
	}

	if isRefresh {
		claims.StandardClaims.ExpiresAt = int64(time.Duration(ttl) * 24 * time.Hour)
	} else {
		claims.StandardClaims.ExpiresAt = int64(time.Duration(ttl) * time.Minute)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func ParseToken(tokenString string) (int, bool, models.Role, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return 0, false, "", err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims.UserID, claims.IsRefresh, claims.Role, nil
	}

	return 0, false, "", fmt.Errorf("invalid token")
}
