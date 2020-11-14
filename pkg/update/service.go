package update

import (
	"errors"
)

var ErrUserNotFound = errors.New("Erro: Usuário não encontrado")
var ErrUserDuplicate = errors.New("Erro: Nome de usuário já existe")
var ErrEventNotFound = errors.New("Error: Evento não encontrado")

type Service interface {
	UpdateUser(User) error
	UpdateEvent(Event) error
}

type Repository interface {
	UpdateUser(User) error
	UpdateEvent(Event) error
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

func (s *service) UpdateEvent(e Event) error {
	return s.r.UpdateEvent(e)
}
