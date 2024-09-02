package handler

import (
	"encoding/json"
	"fmt"
	"gemini-care/dto"
	errorHandler "gemini-care/error"
	"gemini-care/external/oauth"
	"gemini-care/helper"
	"gemini-care/usecase"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"net/http"
	"strings"
)

type userHandler struct {
	validate    helper.ValidationHelper
	userUsecase usecase.UserUsecase
	oauthGoogle *oauth2.Config
}

func NewUserHandler(validate helper.ValidationHelper, userUsecase usecase.UserUsecase) *userHandler {
	return &userHandler{validate, userUsecase, oauth.NewGoogleOauth()}
}

func (h *userHandler) Create(ctx echo.Context) error {
	var request dto.CreateRequest

	if err := ctx.Bind(&request); err != nil {
		return errorHandler.HandleError(ctx, &errorHandler.BadRequestError{Message: err.Error()})
	}

	if err := h.validate.ValidateRequest(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	if err := h.userUsecase.Create(&request); err != nil {
		return errorHandler.HandleError(ctx, err)
	}

	response := dto.ResponseParam{
		StatusCode: http.StatusCreated,
		Message:    "User created successfully",
	}

	return ctx.JSON(http.StatusCreated, helper.Response(response))
}

func (h *userHandler) GoogleLogin(ctx echo.Context) error {
	googleOauth := oauth.NewGoogleOauth()
	url := googleOauth.AuthCodeURL("state")
	return ctx.Redirect(http.StatusTemporaryRedirect, url)
}

func (h *userHandler) GoogleCallback(ctx echo.Context) error {
	value := ctx.FormValue("state")
	if value != "state" {
		return ctx.JSON(http.StatusBadRequest, "Invalid state")
	}

	code := ctx.FormValue("code")
	token, err := h.oauthGoogle.Exchange(ctx.Request().Context(), code)
	if err != nil {
		return errorHandler.HandleError(ctx, &errorHandler.InternalServerError{Message: err.Error()})
	}
	resp, err := http.Get(fmt.Sprintf("https://www.googleapis.com/oauth2/v2/userinfo?access_token=%s", token.AccessToken))

	if err != nil {
		return errorHandler.HandleError(ctx, &errorHandler.InternalServerError{Message: err.Error()})
	}
	defer resp.Body.Close()
	var userInfo map[string]interface{}
	if err = json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to decode user info")
	}

	request := dto.GoogleLoginRequest{
		Name:  strings.ReplaceAll(userInfo["name"].(string), " ", ""),
		Email: userInfo["email"].(string),
	}

	res, err := h.userUsecase.LoginOrRegisterGoogle(&request)

	if err != nil {
		return errorHandler.HandleError(ctx, err)
	}

	ctx.SetCookie(&http.Cookie{
		Name:     "refresh_token",
		Value:    res.RefreshToken,
		Path:     "/",
		Domain:   "",
		MaxAge:   24 * 60 * 60,
		Secure:   true,
		HttpOnly: true,
	})

	response := dto.ToLoginResponse(res.AccessToken)
	return ctx.JSON(http.StatusOK, helper.Response(dto.ResponseParam{
		StatusCode: http.StatusOK,
		Message:    "Login with google success",
		Data:       response,
	}))
}
