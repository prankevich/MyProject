package pkg

import "github.com/go-openapi/jsonpointer"

type CustomClaims struct {
	jwt.StandardClaims
	UserId    int  `json:"user_id"`
	IsRefresh bool `json:"is_refresh"`
}

func GenerateToken(userId int, isRefresh bool) (string, error) {
	claims := CustomClaims{
		StandardClaims: jwt.StandardClaims{},
		UserId:         userId,
	}
	if isRefresh {
		claims.StandardClaims.E
	}
}
