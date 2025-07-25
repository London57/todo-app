package auth

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/London57/todo-app/internal/controller/http/error"
	"github.com/London57/todo-app/internal/domain"
	"github.com/London57/todo-app/internal/domain/signup"
	"golang.org/x/oauth2"

	"github.com/gin-gonic/gin"
)

func (c *AuthController) SignUp(r *gin.Context) {
	var user signup.SignUpRequest
	if err := r.ShouldBindJSON(&user); err != nil {
		error.ErrorResponse(r, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}
	user_from_bd, err := c.GetUserByEmail(r, user.Email)
	if err == nil && (user_from_bd != domain.User{}) {
		error.ErrorResponse(r, http.StatusConflict, "user with this email already exists", "")
		return
	}

	id, err := c.CreateUser(r, signup.SignUpRequest{
		Name:     user.Name,
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
	})
	if err != nil {
		error.ErrorResponse(r, http.StatusInternalServerError, "server create user problems", err.Error())
		return
	}
	domain_user := domain.User{
		ID:       id,
		Name:     user.Name,
		Username: user.Username,
	}
	access_token, err := c.CreateAccessToken(
		domain_user, c.env.JWT.AccessTokenSecret, c.env.JWT.AccessTokenExpiryHour,
	)
	if err != nil {
		error.ErrorResponse(r, http.StatusInternalServerError, "create access jwt token error", err.Error())
		return
	}

	refresh_token, err := c.CreateRefreshToken(
		domain_user, c.env.JWT.RefreshTokenSecret, c.env.JWT.RefreshTokenExpiryHour,
	)
	if err != nil {
		error.ErrorResponse(r, http.StatusInternalServerError, "create refresh jwt token error", err.Error())
		return
	}
	signupResponse := signup.SignUpResponse{
		AccessToken:  access_token,
		RefreshToken: refresh_token,
	}
	r.JSON(http.StatusCreated, signupResponse)

}

func (c *AuthController) SignIn(r *gin.Context) {}

var oauthConfig *oauth2.Config

func (c *AuthController) OAuth2(r *gin.Context) {
	provider, ok := r.Params.Get("provider")

	if !ok {
		error.ErrorResponse(r, http.StatusBadRequest, "not specified provider param", "")
		return
	}

	if strings.ToLower(provider) == "google" {
		oauthConfig = InitGoogleProvider(
			c.env.OAuth2.Google.GoogleClientId,
			c.env.OAuth2.Google.GoogleClientSecret,
			c.env.API.Schema+c.env.API.Host+":"+strconv.Itoa(c.env.API.Port)+"/api/v1/auth/google/callback",
		)
	}
	url := oauthConfig.AuthCodeURL(c.env.OAuth2.OAuthStateString)
	r.Redirect(http.StatusFound, url)

}

func (c *AuthController) OAuth2Callback(r *gin.Context) {
	provider, ok := r.Params.Get("provider")
	if !ok {
		error.ErrorResponse(r, http.StatusBadRequest, "not specified provider param", "")
		return
	}

	state := r.Query("state")
	if state != c.env.OAuth2.OAuthStateString {
		error.ErrorResponse(r, http.StatusBadRequest, "invalid oauth state", "")
		return
	}
	code := r.Query("code")
	token, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		error.ErrorResponse(r, http.StatusBadRequest, "failed to exchange token", err.Error())
		return
	}

	var resp *http.Response

	client := oauthConfig.Client(context.Background(), token)
	if strings.ToLower(provider) == "google" {
		resp, err = client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
		if err != nil {
			error.ErrorResponse(r, http.StatusInternalServerError, "failed to get user info", err.Error())
			return
		}
	}
	defer resp.Body.Close()

	userOAuth2 := signup.SignUpOAuth2Request{}
	if err = json.NewDecoder(resp.Body).Decode(&userOAuth2); err != nil {
		error.ErrorResponse(r, http.StatusInternalServerError, "failedto decode user info", err.Error())
		return
	}

	user := signup.SignUpRequest{
		Name:     userOAuth2.Name,
		Email:    userOAuth2.Email,
		Username: "",
		Password: "",
	}

	var domain_user domain.User
	database_user, err := c.GetUserByEmail(r, user.Email)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		id, err := c.CreateUser(r, user)
		if err != nil {
			error.ErrorResponse(r, http.StatusInternalServerError, "server create user problems", err.Error())
			return
		}

		domain_user = domain.User{
			ID:       id,
			Name:     user.Name,
			Username: user.Username,
			Password: user.Password,
		}
	} else if err != nil && !errors.Is(err, sql.ErrNoRows) {
		error.ErrorResponse(r, http.StatusInternalServerError, "database error", err.Error())
		return
	} else { //if user with this email exists
		domain_user = domain.User{
			ID:       database_user.ID,
			Email:    database_user.Email,
			Username: database_user.Username,
			Password: database_user.Password,
		}
	}

	access_token, err := c.CreateAccessToken(
		domain_user, c.env.JWT.AccessTokenSecret, c.env.JWT.AccessTokenExpiryHour,
	)
	if err != nil {
		error.ErrorResponse(r, http.StatusInternalServerError, "create access jwt token error", err.Error())
		return
	}

	refresh_token, err := c.CreateRefreshToken(
		domain_user, c.env.JWT.RefreshTokenSecret, c.env.JWT.RefreshTokenExpiryHour,
	)
	if err != nil {
		error.ErrorResponse(r, http.StatusInternalServerError, "create refresh jwt token error", err.Error())
		return
	}
	signupResponse := signup.SignUpResponse{
		AccessToken:  access_token,
		RefreshToken: refresh_token,
	}
	r.JSON(http.StatusCreated, signupResponse)
}
