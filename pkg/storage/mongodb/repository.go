package mongodb

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/PedroMAdorno4/Desafio/pkg/auth"
	"github.com/PedroMAdorno4/Desafio/pkg/create"
	"github.com/PedroMAdorno4/Desafio/pkg/http/rest"
	"github.com/PedroMAdorno4/Desafio/pkg/read"
	"github.com/PedroMAdorno4/Desafio/pkg/update"
	"github.com/PedroMAdorno4/Desafio/pkg/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbName     = "calendar"
	dbURL      = "mongodb://localhost:27017"
	collUsers  = "users"
	collEvents = "events"
)

type Storage struct {
	db *mongo.Database
}

var ctx context.Context

func NewStorage() (*Storage, error) {
	var err error

	st := new(Storage)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://api:rOhWJ3WIwrGoKKyf@cluster0.ssffs.mongodb.net/calendar?retryWrites=true&w=majority",
	))
	if err != nil {
		log.Fatal(err)
	}

	st.db = client.Database(dbName)

	return st, nil
}

func (st *Storage) checkUsernameExists(username string) bool {
	coll := st.db.Collection(collUsers)
	result, _ := coll.CountDocuments(ctx, bson.M{"username": username})
	return result > 0
}

//For auth service
func (st *Storage) CheckUserCredentials(username string, password string) (bool, error) {
	coll := st.db.Collection(collUsers)
	var loginInfo rest.UserLogin

	result := coll.FindOne(ctx, bson.M{"username": username})
	err := result.Decode(&loginInfo)
	if err != nil {
		return false, err
	}
	if utils.CheckPasswordHash(password, loginInfo.Password) {
		return true, nil
	}
	return false, auth.ErrCredentialsInvalid
}

func (st *Storage) GetUserTokenInfo(id primitive.ObjectID) (auth.Token, error) {
	user, err := st.GetUser(id)
	var token auth.Token
	token.ID = user.ID
	token.Name = user.Firstname
	// token.Permissions = user.Permissions
	// token.GroupID = user.GroupID
	token.ExpiresAt = time.Now().Add(time.Hour*24*30 + time.Hour).Unix()
	token.IssuedAt = time.Now().Unix()

	return token, err
}

//For create Service
func (st *Storage) CreateUser(u create.User) (primitive.ObjectID, error) {
	coll := st.db.Collection(collUsers)
	u.Password = utils.HashPassword(u.Password)

	if st.checkUsernameExists(u.Username) {
		return primitive.NilObjectID, create.ErrUserDuplicate
	}

	result, err := coll.InsertOne(ctx, u)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return result.InsertedID.(primitive.ObjectID), err
}

//For delete service
func (st *Storage) DeleteUser(id primitive.ObjectID) error {
	coll := st.db.Collection(collUsers)
	_, err := coll.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

//For read service
func (st *Storage) GetUser(id primitive.ObjectID) (read.User, error) {
	coll := st.db.Collection(collUsers)
	var user read.User

	result := coll.FindOne(ctx, bson.M{"_id": id})
	err := result.Decode(&user)
	return user, err

}

func (st *Storage) GetIDByUsername(username string) (primitive.ObjectID, error) {
	coll := st.db.Collection(collUsers)
	var user read.User

	result := coll.FindOne(ctx, bson.M{"username": username})
	err := result.Decode(&user)
	return user.ID, err
}

//For update service
func (st *Storage) UpdateUser(u update.User) error {
	coll := st.db.Collection(collUsers)

	if val, ok := u["username"].(string); ok {
		newUsername := val
		json.Unmarshal([]byte(val), &newUsername)
		if st.checkUsernameExists(newUsername) {
			return update.ErrUserDuplicate
		}
		fmt.Println(newUsername)
	}

	id, _ := primitive.ObjectIDFromHex(u["_id"].(string))

	for k, v := range u {
		if k != "_id" {
			var newValue interface{}
			if k == "password" {
				newValue = utils.HashPassword(v.(string))
			} else {
				newValue = v
			}
			_, err := coll.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{k: newValue}})
			if err != nil {
				return err
			}
		}
	}

	return nil
}
