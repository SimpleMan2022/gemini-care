package usecase

import (
	"gemini-care/dto"
	"gemini-care/entity"
	errorHandler "gemini-care/error"
	"gemini-care/helper"
	"gemini-care/repository"
)

type UserUsecase interface {
	Create(request *dto.CreateRequest) error
	LoginOrRegisterGoogle(request *dto.GoogleLoginRequest) (*dto.LoginResponse, error)
}
type userUsecase struct {
	userrepo       repository.UserRepository
	passwordHelper helper.PasswordHelper
	tokenHelper    helper.JWTToken
}

func NewUserUsecase(userrepo repository.UserRepository, passwordHelper helper.PasswordHelper, tokenHelper helper.JWTToken) *userUsecase {
	return &userUsecase{userrepo, passwordHelper, tokenHelper}
}

func (uc *userUsecase) Create(request *dto.CreateRequest) error {
	user, _ := uc.userrepo.FindOneByUsername(request.Username)
	if user != nil {
		return &errorHandler.BadRequestError{Message: "Username already exists"}
	}
	password, err := uc.passwordHelper.HashPassword(request.Password)
	if err != nil {
		return err
	}
	newUser := &entity.User{
		Username: request.Username,
		Email:    request.Email,
		Password: password,
		Provider: "local",
	}
	if err = uc.userrepo.Create(newUser); err != nil {
		return &errorHandler.InternalServerError{err.Error()}
	}
	return nil
}

func (uc *userUsecase) LoginOrRegisterGoogle(request *dto.GoogleLoginRequest) (*dto.LoginResponse, error) {
	var id int
	user, _ := uc.userrepo.FindOneByEmail(request.Email)
	if user == nil {
		newUser := &entity.User{
			Username: request.Name,
			Email:    request.Email,
			Provider: "google",
		}

		if err := uc.userrepo.Create(newUser); err != nil {
			return nil, &errorHandler.InternalServerError{err.Error()}
		}
		user = newUser
	}

	id = user.Id
	accessToken, err := uc.tokenHelper.GenerateAccessToken(id)
	if err != nil {
		return nil, &errorHandler.InternalServerError{err.Error()}
	}
	refreshToken, err := uc.tokenHelper.GenerateRefreshToken(id)
	if err != nil {
		return nil, &errorHandler.InternalServerError{err.Error()}
	}

	response := &dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return response, nil
}
