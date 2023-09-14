package product

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Filter bson.M

func NewFilter() Filter {
	return Filter{}
}

func (f Filter) SetID(ids ...primitive.ObjectID) Filter {
	f["_id"] = bson.M{
		"$in": ids,
	}
	return f
}

func (f Filter) SetName(name string) Filter {
	f["name"] = name
	return f
}

func (f Filter) SetPrice(price uint) Filter {
	f["price"] = price
	return f
}

func (f Filter) SetCurrentStock(currentstock uint) Filter {
	f["current_stock"] = currentstock
	return f
}

func (f Filter) SetImage(image string) Filter {
	f["Image"] = image
	return f
}

func (f Filter) SetCreatedAt() Filter {
	f["created_at"] = time.Now().UTC()
	return f
}

func (f Filter) SetUpdatedAt() Filter {
	f["updated_at"] = time.Now().UTC()
	return f
}
