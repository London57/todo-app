package auth

import (
	"context"
	"net"
	"net/http"

	"github.com/London57/todo-app/internal/controller/http/error"
	"github.com/London57/todo-app/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

func (c *AuthController) SignUp(r *gin.Context) {
	var user domain.UserRequest
	if err := r.ShouldBindJSON(&user); err != nil {
		error.ErrorResponse(r, http.StatusBadRequest, "invalid request body")
		return
	}

	id, err := c.RegisterUserByUsername(r.Request.Context(), domain.User{
		Name: user.Name,
		Username: user.Username,
		Password: user.Password,
	})
	if err != nil {
		error.ErrorResponse(r, http.StatusInternalServerError, "server create user problems")
	}
	_ = id
	return
}
func (c *AuthController) SignIn(r *gin.Context) {}


const (
	maxAge = 86400 * 30
	isProd = false
)

func (c *AuthController) OAuth2(r *gin.Context) {
	// store := sessions.NewCookieStore([]byte(c.extra_data["Key"]))
	googleClientId := c.extra_data["Google_Client_Id"] 
	googleClientSecret := c.extra_data["Google_Client_Secret"]
	// store.Options.Path = "/"
	// store.Options.HttpOnly = true
	// store.Options.Secure = isProd
	// store.MaxAge(maxAge)

	// gothic.Store = store

	var scheme string
	req := r.Request
	if req.TLS != nil {
		scheme = "https://"
	} else {			
		scheme = "http://"
	}
	goth.UseProviders(
		google.New(googleClientId, googleClientSecret, scheme + "api/v1/auth/google/callback", req.Context().Value(http.LocalAddrContextKey).(net.Addr).String()), 
	)
	gothic.BeginAuthHandler(r.Writer, r.Request)
}

func (c *AuthController) AuthCallback(r *gin.Context) {
	provider := r.Param("provider")
	req := r.Request
	req = req.WithContext(context.WithValue(req.Context(), "provider", provider))

	user_data, err := gothic.CompleteUserAuth(r.Writer, r.Request)
	if err != nil {
		error.ErrorResponse(r, http.StatusBadRequest, "failed to register")
		return
	}

	user := domain.User {
		Name: user_data.Name,
		Username: user_data.NickName,
		Password: "",
	}
	user.CreateUser(user)
	

	var scheme string
	if req.TLS != nil {
		scheme = "https://"
	} else { 
		scheme = "http://"
	}
	r.Redirect(http.StatusFound, scheme + req.Context().Value(http.LocalAddrContextKey).(net.Addr).String() + "api/v1/lists/")
}