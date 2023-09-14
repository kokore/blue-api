package product

import (
	"blue-api/internal/database"
	"blue-api/internal/errorinternal"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type repoImpl struct {
	coll *mongo.Collection
}

type Repo interface {
	InsertOne(ctx context.Context, f Filter) error
	Find(ctx context.Context) ([]Product, error)
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
		return nil, errorinternal.NewError(errorinternal.ErrorCodeProductNotFound, "can't find product")
	}

	for cursor.Next(ctx) {
		var product Product
		if err := cursor.Decode(&product); err != nil {
			return nil, errorinternal.NewError(errorinternal.ErrorCodeProductNotFound, "can't find product")
		}
		products = append(products, product)
	}

	return products, nil
}
