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
	tokenTTL   = 12 * time.Hour
)
type AuthService struct {
	repo repository.Authorization
}


type tokenClaims struct {
	jwt.StandardClaims
	userId string `json:"user_id"`
}

func (a *AuthService) GenerateToken(id primitive.ObjectID, username string, password string) (string, error) {
	user, err := a.repo.GetUser(id, username, a.generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})
	tokenStr, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", err
	}
	tokenFromRedis, err := a.repo.SetToken(tokenStr)
	if err != nil {
		return "", err
	}
	return tokenFromRedis, nil
}

func (s *AuthService) ParseToken(token string) (string, error) {
	tokenT, err := s.repo.GetToken(token)
	if err != nil {
		return "", err
	}
	tokens, err := jwt.ParseWithClaims(tokenT, &tokenClaims{}, func(tokenT *jwt.Token) (interface{}, error){
		if _, ok := tokenT.Method.(*jwt.SigningMethodHMAC); !ok{
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
