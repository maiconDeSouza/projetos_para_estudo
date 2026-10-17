package service

import (
	"fmt"
	"golink/internal/domain"
	"golink/internal/infra/env"
	"golink/internal/infra/errapp"
	"golink/internal/repository"
	"net/http"
	"time"
	"uuid"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type ServiceInterface interface {
	CreateUserAdmin(id uuid.UUID, userRequest domain.UserRequest) (*domain.UserResponse, *errapp.ErrResponse)
	CreateUser(userRequest domain.UserRequest) (*domain.UserResponse, *errapp.ErrResponse)
	Login(login domain.Login) (*http.Cookie, *domain.LoginResponse, *errapp.ErrResponse)
	GetAllUser(id uuid.UUID) ([]*domain.User, *errapp.ErrResponse)
	AddLinkAdmin(userID uuid.UUID, linkResquest domain.LinkRequest) (*domain.UserResponse, *errapp.ErrResponse)
	TitleLink(title string) (*domain.Link, *errapp.ErrResponse)
}

type Service struct {
	repo     repository.RepositoryInterface
	env      *env.ENV
	validate *validator.Validate
}

func NewService(repo repository.RepositoryInterface, env *env.ENV) *Service {
	return &Service{repo: repo, env: env, validate: validator.New()}
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

func generateJWT(id uuid.UUID, isAdmin bool, env *env.ENV) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   id,
		"admin": isAdmin,
		"exp":   time.Now().Add(time.Hour).Unix(),
		"iat":   time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte(env.JWT_SECRET_KEY))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidarToken(tokenString string, secretKey string) (string, bool, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de assinatura inesperado: %v", t.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		return "", false, fmt.Errorf("token inválido ou expirado: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", false, fmt.Errorf("erro ao mapear os claims do token")
	}

	sub, ok := claims["sub"].(string)
	if !ok {
		return "", false, fmt.Errorf("campo 'sub' não encontrado ou inválido no token")
	}

	isAdmin, ok := claims["admin"].(bool)
	if !ok {
		isAdmin = false
	}

	return sub, isAdmin, nil
}

func generateCookie(jwt string) *http.Cookie {
	cookie := &http.Cookie{
		Name:     "session_token_user",
		Value:    jwt,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	return cookie
}
