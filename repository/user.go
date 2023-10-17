package repository

import (
	"context"

	"github.com/dilyara4949/clean-project/domain"
	"github.com/dilyara4949/clean-project/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepository struct {
	database mongo.Database
	collection string
}

func NewUserRepository(db mongo.Database, collection string) domain.UserRepository {
	return &userRepository{
		database:   db,
		collection: collection,
	}
}

func (u *userRepository) Create(c context.Context, user *domain.User) error {
	collection := u.database.Collection(u.collection)
	_, err := collection.InsertOne(c, u)
	return err
}

func (u *userRepository) Fetch(c context.Context) ([]domain.User, error) {
	collection := u.database.Collection(u.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}
	var users []domain.User

	err = cursor.All(c, &users)
	if users == nil {
		return []domain.User{}, err
	}

	return users, err

}

func (u *userRepository) GetByEmail(c context.Context, email string) (domain.User, error) {
	collection := u.database.Collection(u.collection)
	var user domain.User

	err := collection.FindOne(c, bson.M{"email":email}).Decode(&user)

	return user, err
}

func (u *userRepository) GetByID(c context.Context, id string) (domain.User, error) {
	collection := u.database.Collection(u.collection)

	var user domain.User

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, err
	}

	err = collection.FindOne(c, bson.M{"_id":idHex}).Decode(&user)
	return user, err
}





