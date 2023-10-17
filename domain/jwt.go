package domain

import (
	"github.com/golang-jwt/jwt/v4"
)

type JwtClaims struct {
	Name string `json:"name"`
	ID string `json:"id"`
	jwt.StandardClaims
}

type JwtRefreshClaims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}