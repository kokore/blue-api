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

func (f Filter) SetQuantity(quantity uint) Filter {
	f["quantity"] = quantity
	return f
}

func (f Filter) SetImage(image string) Filter {
	f["image"] = image
	return f
}

func (f Filter) SetCreatedAt() Filter {
	f["created_at"] = time.Now()
	return f
}

func (f Filter) SetUpdatedAt() Filter {
	f["updated_at"] = time.Now()
	return f
}

type Update bson.M

func NewUpdate() Update {
	return Update{
		"$set": bson.M{
			"updated_at": time.Now(),
		},
	}
}

func (u Update) set(field string, value any) Update {
	u["$set"].(bson.M)[field] = value
	return u
}

func (u Update) SetName(s string) Update {
	u.set("name", s)
	return u
}

func (u Update) SetPrice(s uint) Update {
	u.set("price", s)
	return u
}

func (u Update) SetQuantity(s uint) Update {
	u.set("quantity", s)
	return u
}

func (u Update) SetImage(s string) Update {
	u.set("image", s)
	return u
}
