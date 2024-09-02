package route

import (
	"gemini-care/api/handler"
	"gemini-care/bootstrap"
	"gemini-care/helper"
	"gemini-care/repository"
	"gemini-care/usecase"
	"github.com/labstack/echo/v4"
)

func AuthRouter(r *echo.Group) {
	user_repo := repository.NewUserRepository(bootstrap.DB)
	passwordHelper := helper.NewPasswordHelper()
	validationHelper := helper.NewValidationHelper()
	tokenHelper := helper.NewJWTToken()
	user_uc := usecase.NewUserUsecase(user_repo, passwordHelper, tokenHelper)
	user_handler := handler.NewUserHandler(validationHelper, user_uc)
	r.POST("/register", user_handler.Create)
	r.GET("/login/google", user_handler.GoogleLogin)
	r.GET("/callback/google", user_handler.GoogleCallback)
}
