package product

import (
	"blue-api/internal/database"

	"go.mongodb.org/mongo-driver/mongo"
)

type repoImpl struct {
	coll *mongo.Collection
}

type Repo interface {
}

func InitUserRepository(connection database.Connection) Repo {
	return &repoImpl{coll: connection.Database().Collection("product")}
}
