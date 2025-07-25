package middleware

import (
	"net/http"
	"strings"

	"github.com/London57/todo-app/internal/controller/http/error"
	"github.com/London57/todo-app/internal/domain/jwtutil"
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
					error.ErrorResponse(c, http.StatusUnauthorized, "failed to get ID from token", err.Error())
					return
				}
				c.Set("userID", userID)
				c.Next()
				return
			}
			error.ErrorResponse(c, http.StatusUnauthorized, "len authorazion header 2, but error", err.Error())
			return
		}
		error.ErrorResponse(c, http.StatusUnauthorized, "Not authorized", "")
	}
}
