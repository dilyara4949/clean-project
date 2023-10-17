package repository

import (
	"context"

	"github.com/dilyara4949/clean-project/domain"
	"github.com/dilyara4949/clean-project/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ArticleRepository struct {
	database mongo.Database
	collection string
}

func NewArticleRepository (db mongo.Database, collection string) domain.ArticleRepository {
	return &ArticleRepository{
		database: db,
		collection: collection,
	}
}

func (ar *ArticleRepository) Create(c context.Context, article *domain.Article) error {
	collection :=ar.database.Collection(ar.collection)

	_, err := collection.InsertOne(c, article)
	return err
}

func (ar *ArticleRepository) FetchByUserID(c context.Context, userID string) ([]domain.Article, error) {
	collection := ar.database.Collection(ar.collection)

	var articles []domain.Article

	idHex, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return articles, err
	}

	cursor, err := collection.Find(c, bson.M{"_id":idHex})
	if err != nil {
		return articles, err
	}

	err = cursor.All(c, &articles)
	if articles == nil {
		return []domain.Article{}, err
	}
	return articles, err
}