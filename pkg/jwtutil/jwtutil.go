package jwtutil

import (
	"errors"
	"fmt"
	"time"

	"github.com/London57/todo-app/internal/domain"
	"github.com/golang-jwt/jwt/v4"
)

func CreateAccessToken(user *domain.User, secret string, expiry int) (string, error) {
	exp := time.Now().Add(time.Hour * time.Duration(expiry))
	claims := &domain.JwtCustomClaims{
		ID: user.ID.String(),
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("failed to signed jwt: %w", err)
	}
	return t, nil
}

func CreateRefreshToken(user *domain.User, secret string, expiry int) (string, error) {
	exp := time.Now().Add(time.Hour * time.Duration(expiry))
	claimsRefresh := &domain.JwtCustomRefreshClaims{
		ID: user.ID.String(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("failed to signed jwt: %w", err)
	}
	return t, nil
}

func IsAuthorized(requestToken string, secret string) (bool, error) {
	_, err := jwt.Parse(requestToken, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected sighing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		jwt_error := err.(jwt.ValidationError)

		if jwt_error.Errors & jwt.ValidationErrorExpired != 0 {
			return false, errors.New("token expired")
		}
		return false, fmt.Errorf("failed to parse jwt: %v", err)
	}
	return true, nil
}

func ExtractIDFromToken(requestToken string, secret string) (string, error) {
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected sighing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return "", fmt.Errorf("failed to parse jwt: %v", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", fmt.Errorf("token is invalid")
	} 
	return claims["id"].(string), nil
}
