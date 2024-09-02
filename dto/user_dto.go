package dto

type CreateRequest struct {
	Username string `json:"username" validate:"required,min=3"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

type GoogleLoginRequest struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

func ToLoginResponse(accessToken string) *LoginResponse {
	return &LoginResponse{
		AccessToken: accessToken,
	}
}
