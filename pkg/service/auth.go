package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/extaleus/pmanager/pkg/repository"
	"github.com/extaleus/pmanager/user"
)

const (
	salt       = "fgsdktjweporijgsdrtwyergsdhhvcbncdvshtue56324rsgy2345154673576795780f0129345u32094itqae"
	signingKey = "rtw11yWERWLLKJiopYWSSSWQrx22xufg2sgidf-8#2#$IrgdfDGADGHSITIZBXNXHFDdritw4##"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user user.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
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

	return token.SignedString([]byte(signingKey))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}