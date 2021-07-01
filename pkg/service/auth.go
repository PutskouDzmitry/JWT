package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/PutskouDzmitry/golang-training-Library/pkg/repository"
	"github.com/PutskouDzmitry/golang-training-Library/pkg/struct"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	salt       = "hjqrhjqw124617ajfhajs"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenAccessTTL   = 30 * time.Minute
	tokenRefreshTTL   = 12 * time.Hour
)
type AuthService struct {
	repo repository.Authorization
}

type tokenClaims struct {
	jwt.StandardClaims
	userId string `json:"user_id"`
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
	accessToken, err := a.repo.SetAccessToken(tokenStr)
	if err != nil {
		return "", err
	}
	return accessToken, nil
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
	refreshToken, err := a.repo.SetRefreshToken(tokenStr)
	if err != nil {
		return "", err
	}
	return refreshToken, nil
}


func (s *AuthService) ParseAccessToken(token string) (string, error) {
	accessToken, err := s.repo.GetAccessToken(token)
	if err != nil {
		return "", err
	}
	tokens, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(accessToken *jwt.Token) (interface{}, error){
		if _, ok := accessToken.Method.(*jwt.SigningMethodHMAC); !ok{
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
	return claims.userId, nil
}

func (s *AuthService) ParseRefreshToken(token string) (string, error) {
	refreshToken, err := s.repo.GetRefreshToken(token)
	if err != nil {
		return "", err
	}
	tokens, err := jwt.ParseWithClaims(refreshToken, &tokenClaims{}, func(refreshToken *jwt.Token) (interface{}, error){
		if _, ok := refreshToken.Method.(*jwt.SigningMethodHMAC); !ok{
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
	return claims.userId, nil
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (a *AuthService) CreateUser(user _struct.User) (string, error) {
	user.Password = a.generatePasswordHash(user.Password)
	return a.repo.CreateUser(user)
}

func (a *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
