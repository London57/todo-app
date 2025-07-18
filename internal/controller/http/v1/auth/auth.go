package auth

import (
	"net/http"

	"github.com/London57/todo-app/internal/controller/http/v1/error"
	"github.com/gin-gonic/gin"
)

func (c *AuthController) SignUp(r *gin.Context) {
	var user User
	if err := r.ShouldBindJSON(&user); err != nil {
		error.ErrorResponse(r, http.StatusBadRequest, err.Error())
		return
	}

}

func (c *AuthController) SignIn(r *gin.Context) {}
