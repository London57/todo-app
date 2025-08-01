package jwtutil

import (
	"github.com/golang-jwt/jwt/v4"
)

type (
	JwtCustomClaims struct {
		jwt.RegisteredClaims
		ID       string `json:"id"`
		Username string `json:"username"`
	}
	JwtCustomRefreshClaims struct {
		jwt.RegisteredClaims
		ID string `json:"id"`
	}
)
