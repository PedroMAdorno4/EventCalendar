package create

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ErrUserDuplicate = errors.New("Erro: Nome de usuário já existe")
var ErrEventOverlap = errors.New("Erro: Evento está sobrescrevendo um já existente")

type Service interface {
	CreateUser(User) (primitive.ObjectID, error)
	CreateEvent(Event) (primitive.ObjectID, error)
	// CreateGroup(Group) error
	// CreateGame(Game) (*mongo.InsertOneResult, error)
	// CreateAct(Act) error
	// CreateScenes(Game) error
}

type Repository interface {
	CreateUser(User) (primitive.ObjectID, error)
	CreateEvent(Event) (primitive.ObjectID, error)
	// CreateGroup(Group) error
	// CreateGame(Game) (*mongo.InsertOneResult, error)
	// CreateAct(Act) error
	// CreateScene(Scene) error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) CreateUser(u User) (primitive.ObjectID, error) {
	var err error
	id, err := s.r.CreateUser(u)
	if err != nil {
		return id, err
	}
	return id, err
}

func (s *service) CreateEvent(e Event) (primitive.ObjectID, error) {
	var err error
	id, err := s.r.CreateEvent(e)
	if err != nil {
		return id, err
	}
	return id, err
}
