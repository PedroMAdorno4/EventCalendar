package read

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrUserNotFound  = errors.New("User was not found")
	ErrEventNotFound = errors.New("Group was not found")
)

type Service interface {
	GetUser(primitive.ObjectID) (User, error)
	GetIDByUsername(Username string) (primitive.ObjectID, error)
}

type Repository interface {
	GetUser(primitive.ObjectID) (User, error)
	GetIDByUsername(Username string) (primitive.ObjectID, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetUser(id primitive.ObjectID) (User, error) {
	return s.r.GetUser(id)
}

func (s *service) GetIDByUsername(username string) (primitive.ObjectID, error) {
	return s.r.GetIDByUsername(username)
}