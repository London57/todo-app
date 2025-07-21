package auth

import (
	"github.com/London57/todo-app/config"
	"github.com/London57/todo-app/internal/controller/http/common"
	"github.com/London57/todo-app/internal/domain/signup"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
)

type AuthController struct {
	common.BaseController
	signup.SignUpUseCase
	env *config.Config
}

func init() {
	googleClientId := config.Config.OAuth2.Google.GoogleClientId
	googleClientSecret := config.Config.OAuth2.Google.GoogleClientSecret
	goth.UseProviders(
		google.New(
			googleClientId,
			googleClientSecret,
			config.Config.HTTP.Schema+config.Config.HTTP.IP+"api/v1/auth/google/callback",
			"email",
			"profile",
		),
	)
}

func NewAuthController(b common.BaseController, sauc signup.SignUpUseCase, env *config.Config) *AuthController {
	return &AuthController{b, sauc, env}
}
