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

func (f Filter) SetTotal(total int) Filter {
	f["total"] = total
	return f
}

func (f Filter) SetCoins(coins map[int]int) Filter {
	f["coins"] = coins
	return f
}

func (f Filter) SetBanknotes(banknotes map[int]int) Filter {
	f["banknotes"] = banknotes
	return f
}

type Update bson.M

func NewUpdate() Update {
	return Update{
		"$set": bson.M{},
	}
}

func (u Update) set(field string, value any) Update {
	u["$set"].(bson.M)[field] = value
	return u
}

func (u Update) SetCoins(s map[int]int) Update {
	u.set("coins", s)
	return u
}

func (u Update) SetBanknotes(s map[int]int) Update {
	u.set("banknotes", s)
	return u
}

func (u Update) SetTotal(s int) Update {
	u.set("total", s)
	return u
}
