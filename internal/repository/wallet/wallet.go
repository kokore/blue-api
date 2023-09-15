package wallet

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
}

func InitWalletRepository(connection database.Connection) Repo {
	return &repoImpl{coll: connection.Database().Collection("wallet")}
}

func (r repoImpl) InsertOne(ctx context.Context, filter Filter) error {
	_, err := r.coll.InsertOne(ctx, filter)
	if err != nil {
		return errorinternal.NewError(errorinternal.ErrorCodeWalletCantInsert, "can't insert wallet")
	}
	return nil
}
