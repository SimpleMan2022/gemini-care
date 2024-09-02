package oauth

import (
	"gemini-care/bootstrap"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func NewGoogleOauth() *oauth2.Config {
	env := bootstrap.NewEnv()
	return &oauth2.Config{
		ClientID:     env.GOOGLE_CLIENT_ID,
		ClientSecret: env.GOOGLE_CLIENT_SECRET,
		Endpoint:     google.Endpoint,
		RedirectURL:  env.GOOGLE_REDIRECT_URL,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
	}
}
