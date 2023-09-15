package wallet

import "go.mongodb.org/mongo-driver/bson/primitive"

type Wallet struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Coins     map[int]int        `json:"coins" bson:"coins"`
	Banknotes map[int]int        `json:"banknotes" bson:"banknotes"`
	Total     uint               `json:"total" bson:"total"`
}
