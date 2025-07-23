package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"strings"

	"github.com/London57/todo-app/internal/controller/http/error"
	"github.com/London57/todo-app/pkg/logger"
	"github.com/gin-gonic/gin"
)

func buildPanic(c *gin.Context, err any) string {
	var b strings.Builder
	b.WriteString(c.ClientIP())
	b.WriteString(" - ")
	b.WriteString(c.Request.Method)
	b.WriteString(" ")
	b.WriteString(c.Request.RequestURI)
	b.WriteString(" PANIC!! ")
	b.WriteString(fmt.Sprintf("%v\n%s", err, debug.Stack()))
	return b.String()
}

func RecoveryMiddleware(l logger.Interface) gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, err any) {
		panic := buildPanic(c, err)
		l.Error(panic)
		error.ErrorResponse(c, http.StatusInternalServerError, "internal server error")
	})
}
