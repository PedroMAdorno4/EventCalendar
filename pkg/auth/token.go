package auth

import (
	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Token struct {
	ID   primitive.ObjectID
	Name string
	jwt.StandardClaims
}

type ContextKey string

func (c ContextKey) String() string {
	return "Context Key " + string(c)
}
