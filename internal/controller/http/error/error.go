package error

import "github.com/gin-gonic/gin"

type error struct {
	Message string            `json:"message"`
	Details map[string]string `json:"details"`
}

func ErrorResponse(r *gin.Context, statusCode int, message string) {
	r.AbortWithStatusJSON(statusCode, error{Message: message})
}
