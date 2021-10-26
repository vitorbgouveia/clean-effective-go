package entities

import "time"

// Sales entity
type Sales struct {
	Desc     string    `bson:"desc" json:"desc"`
	Value    float32   `bson:"value" json:"value"`
	Cpf      string    `bson:"cpf" json:"cpf"`
	Cashback float32   `bson:"cashback" json:"cashback"`
	CreateAt time.Time `bson:"create_at" json:"create_at"`
}
