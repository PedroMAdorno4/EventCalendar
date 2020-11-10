package create

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ErrUserDuplicate = errors.New("Username already exists")
var ErrGroupDuplicate = errors.New("Group already exists")

type Service interface {
	CreateUser(User) (primitive.ObjectID, error)
	// CreateGroup(Group) error
	// CreateGame(Game) (*mongo.InsertOneResult, error)
	// CreateAct(Act) error
	// CreateScenes(Game) error
}

type Repository interface {
	CreateUser(User) (primitive.ObjectID, error)
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

//TODO: check error
/*
func (s *service) CreateGroup(g Group) error {
	return s.r.CreateGroup(g)
}

func (s *service) CreateGame(g Game) (*mongo.InsertOneResult, error) {
	result, err := s.r.CreateGame(g)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (s *service) CreateAct(a Act) error {
	err := s.r.CreateAct(a)
	return err
}

func (s *service) CreateScenes(g Game) error {
	var err error
	for _, a := range g.Acts {
		for _, scene := range a.Scenes {
			err := s.r.CreateScene(scene)
			if err != nil {
				return err
			}
		}
	}
	return err
}
*/
