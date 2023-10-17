package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Article struct {
	ID     primitive.ObjectID `bson:"_id" json:"-"`
	Title  string             `bson:"title" form:"title" binding:"required" json:"title"`
	UserID primitive.ObjectID `bson:"userID" json:"-"`
	Body string `bson:"body"`
}

type ArticleRepository interface {
	Create(c context.Context, art *Article) error
	FetchByUserID(c context.Context, userID string) ([]Article, error)

}

type ArticleUserCase interface {
	Create(c context.Context, task *Article) error
	FetchByUserID(c context.Context, userID string) ([]Article, error)
}