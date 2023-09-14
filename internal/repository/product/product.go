package product

import (
	"blue-api/internal/database"
	"blue-api/internal/errorinternal"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type repoImpl struct {
	coll *mongo.Collection
}

type Repo interface {
	InsertOne(ctx context.Context, f Filter) error
	Find(ctx context.Context) ([]Product, error)
	FindAndUpdate(ctx context.Context, productId primitive.ObjectID, u Update) (*Product, error)
	DeleteOneById(ctx context.Context, productId primitive.ObjectID) error
}

func InitUserRepository(connection database.Connection) Repo {
	return &repoImpl{coll: connection.Database().Collection("product")}
}

func (r repoImpl) InsertOne(ctx context.Context, filter Filter) error {
	_, err := r.coll.InsertOne(ctx, filter)
	if err != nil {
		return errorinternal.NewError(errorinternal.ErrorCodeProductCantInsert, "can't insert product")
	}
	return nil
}

func (r repoImpl) Find(ctx context.Context) ([]Product, error) {
	var products []Product

	cursor, err := r.coll.Find(ctx, NewFilter())
	if err != nil {
		return nil, errorinternal.NewError(errorinternal.ErrorCodeProductNotFound, "can't find product.")
	}

	for cursor.Next(ctx) {
		var product Product
		if err := cursor.Decode(&product); err != nil {
			return nil, errorinternal.NewError(errorinternal.ErrorCodeProductNotFound, "can't find product.")
		}
		products = append(products, product)
	}
	return products, nil
}

func (r repoImpl) FindAndUpdate(ctx context.Context, productId primitive.ObjectID, u Update) (*Product, error) {
	f := NewFilter().SetID(productId)

	result := r.coll.FindOneAndUpdate(ctx, f, u, options.FindOneAndUpdate().SetReturnDocument(options.After))
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return nil, errorinternal.NewError(errorinternal.ErrorCodeProductNotFound, "can't find product.")
		}
		return nil, result.Err()
	}
	var product Product
	err := result.Decode(&product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r repoImpl) DeleteOneById(ctx context.Context, productId primitive.ObjectID) error {
	f := NewFilter().SetID(productId)
	_, err := r.coll.DeleteOne(ctx, f)
	if err != nil {
		return err
	}
	return err
}
