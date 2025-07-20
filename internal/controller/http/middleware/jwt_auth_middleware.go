package middleware

import (
	"net/http"
	"strings"

	"github.com/London57/todo-app/internal/controller/http/error"
	"github.com/London57/todo-app/pkg/jwtutil"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		s := strings.Split(authHeader, " ")
		if len(s) == 2 {
			authToken := s[1]
			authorized, err := jwtutil.IsAuthorized(authToken, secret)
			if authorized {
				userID, err := jwtutil.ExtractIDFromToken(authToken, secret)
				if err != nil {
					error.ErrorResponse(c, http.StatusUnauthorized, err.Error())
					return
				}
				c.Set("userID", userID)
				c.Next()
				return
			}
			error.ErrorResponse(c, http.StatusUnauthorized, err.Error())
			return
		}
		error.ErrorResponse(c, http.StatusUnauthorized, "Not authorized")
	}
}