package auth

import (
	"errors"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrTokenMissing       = errors.New("Erro: Missing auth token")
	ErrTokenMalformed     = errors.New("Erro: Malformed auth token")
	ErrTokenInvalid       = errors.New("Erro: Invalid auth token")
	ErrUserNotFound       = errors.New("Erro: User was not found")
	ErrUserUnauthorized   = errors.New("Erro: User does not have permission")
	ErrCredentialsInvalid = errors.New("Erro: Credentials invalid")
	ErrCredentialsMissing = errors.New("Erro: Credentials not present")
)

type Service interface {
	CreateToken(primitive.ObjectID) (string, error)
	GetInfoFromToken(r *http.Request) (*Token, error)
	CheckCredentials(username string, password string) (bool, error)
}

type Repository interface {
	GetUserTokenInfo(primitive.ObjectID) (Token, error)
	CheckUserCredentials(string, string) (bool, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) CreateToken(id primitive.ObjectID) (string, error) {
	tk, err := s.r.GetUserTokenInfo(id)

	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	//TODO: add toke password to env file
	tokenString, _ := token.SignedString([]byte("tokenPassword"))

	return tokenString, nil
}

func (s *service) GetInfoFromToken(r *http.Request) (*Token, error) {
	return ValidateToken(r)
}

func (s *service) CheckCredentials(username string, password string) (bool, error) {
	if username != "" && password != "" {
		return s.r.CheckUserCredentials(username, password)
	}
	return false, ErrCredentialsMissing
}
