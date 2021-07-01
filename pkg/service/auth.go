package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/PutskouDzmitry/golang-training-Library/pkg/entity"
	"github.com/PutskouDzmitry/golang-training-Library/pkg/repository"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	salt            = "hjqrhjqw124617ajfhajs"
	signingKey      = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenAccessTTL  = 30 * time.Minute
	tokenRefreshTTL = 12 * time.Hour
)

type AuthService struct {
	repo repository.Authorization
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId string `json:"user_id"`
}

func (a *AuthService) GenerateTokenAccessToken(id primitive.ObjectID, username string, password string) (string, error) {
	user, err := a.repo.GetUser(id, username, a.generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenAccessTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})
	tokenStr, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", err
	}
	err = a.repo.SetAccessToken(tokenStr, id.Hex())
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func (a *AuthService) GenerateTokenRefreshToken(id primitive.ObjectID, username string, password string) (string, error) {
	_, err := a.repo.GetUser(id, username, a.generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenRefreshTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		a.generatePasswordHash(password),
	})
	tokenStr, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", err
	}
	err = a.repo.SetRefreshToken(tokenStr, id.Hex())
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func (s *AuthService) ParseAccessToken(token string) (string, error) {
	tokens, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(accessToken *jwt.Token) (interface{}, error) {
		if _, ok := accessToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid singing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := tokens.Claims.(*tokenClaims)
	if !ok {
		return "", errors.New("token claims are not of type tokenClaims")
	}
	_, err = s.repo.GetAccessToken(token, claims.UserId)
	if err != nil {
		return "", err
	}
	return claims.UserId, nil
}

func (s *AuthService) ParseRefreshToken(token string) (string, error) {
	refreshToken, err := s.repo.GetRefreshToken(token, "qwe")
	if err != nil {
		return "", err
	}
	tokens, err := jwt.ParseWithClaims(refreshToken, &tokenClaims{}, func(refreshToken *jwt.Token) (interface{}, error) {
		if _, ok := refreshToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid singing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := tokens.Claims.(*tokenClaims)
	if !ok {
		return "", errors.New("token claims are not of type tokenClaims")
	}
	return claims.UserId, nil
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (a *AuthService) CreateUser(user entity.User) (string, error) {
	user.Password = a.generatePasswordHash(user.Password)
	return a.repo.CreateUser(user)
}

func (a *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
