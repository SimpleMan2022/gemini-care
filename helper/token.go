package helper

import (
	"github.com/golang-jwt/jwt"
	"os"
	"time"
)

type JWTToken interface {
	GenerateAccessToken(userID int) (string, error)
	GenerateRefreshToken(userID int) (string, error)
}

type jwtToken struct{}

func NewJWTToken() *jwtToken {
	return &jwtToken{}
}

type JWTClaims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

func (j *jwtToken) GenerateAccessToken(userID int) (string, error) {
	secret := []byte(os.Getenv("ACCESS_TOKEN_SECRET"))

	claims := JWTClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
			IssuedAt:  time.Now().Unix(),
			NotBefore: time.Now().Unix(),
		},
	}

	withClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := withClaims.SignedString(secret)

	if err != nil {
		return "", err
	}

	return signedString, nil
}

func (j *jwtToken) GenerateRefreshToken(userID int) (string, error) {
	secret := []byte(os.Getenv("REFRESH_TOKEN_SECRET"))

	claims := JWTClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
			NotBefore: time.Now().Unix(),
		},
	}

	withClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := withClaims.SignedString(secret)

	if err != nil {
		return "", err
	}

	return signedString, nil
}
