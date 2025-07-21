package auth

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func InitGoogleProvider(clientId, clientSecret, redirectUrl string) *oauth2.Config {
	var oauthConfig *oauth2.Config
	oauthConfig = &oauth2.Config{
		RedirectURL:  redirectUrl,
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
	return oauthConfig
}
