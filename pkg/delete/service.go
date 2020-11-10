package delete

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ErrUserNotFound = errors.New("User was not found")
var ErrEventNotFound = errors.New("Event was not found")

type Service interface {
	DeleteUser(primitive.ObjectID) error
	DeleteEvent(primitive.ObjectID) error
}

type Repository interface {
	DeleteUser(primitive.ObjectID) error
	DeleteEvent(primitive.ObjectID) error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) DeleteUser(id primitive.ObjectID) error {
	return s.r.DeleteUser(id)
}

func (s *service) DeleteEvent(id primitive.ObjectID) error {
	return s.r.DeleteEvent(id)
}
