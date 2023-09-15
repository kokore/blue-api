package wallet

import (
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

func (f Filter) SetTotal(total uint) Filter {
	f["total"] = total
	return f
}
