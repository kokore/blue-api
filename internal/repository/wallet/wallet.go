package wallet

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
	FindOne(ctx context.Context) (*Wallet, error)
	UpdateOne(ctx context.Context, walletId primitive.ObjectID, u Update) (*Wallet, error)
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

func (r repoImpl) FindOne(ctx context.Context) (*Wallet, error) {
	var result Wallet
	err := r.coll.FindOne(ctx, NewFilter()).Decode(&result)
	if err != nil {
		return nil, errorinternal.NewError(errorinternal.ErrorCodeWalletNotFound, "can't find wallet")
	}
	return &result, nil
}

func (r repoImpl) UpdateOne(ctx context.Context, walletId primitive.ObjectID, u Update) (*Wallet, error) {
	f := NewFilter().SetID(walletId)

	result := r.coll.FindOneAndUpdate(ctx, f, u, options.FindOneAndUpdate().SetReturnDocument(options.After))
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return nil, errorinternal.NewError(errorinternal.ErrorCodeWalletNotFound, "can't find wallet.")
		}
		return nil, result.Err()
	}
	var wallet Wallet
	err := result.Decode(&wallet)
	if err != nil {
		return nil, err
	}
	return &wallet, nil
}
