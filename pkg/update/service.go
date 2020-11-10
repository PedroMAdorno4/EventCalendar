package update

import (
	"errors"
)

var ErrUserNotFound = errors.New("User was not found")
var ErrUserDuplicate = errors.New("Username already exists")

type Service interface {
	UpdateUser(User) error
}

type Repository interface {
	UpdateUser(User) error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) UpdateUser(u User) error {
	return s.r.UpdateUser(u)
}
