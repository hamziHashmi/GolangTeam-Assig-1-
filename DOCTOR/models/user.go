package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Address struct {
	State   string `json:"state" bson:"state"`
	City    string `json:"city" bson:"city"`
	Pincode int    `json:"pincode" bson:"pincode"`
}

type User struct {
	ID             primitive.ObjectID `json:"_id " bson:"_id,omitempty"`
	Name           string             `json:"name" bson:"doctor_name"`
	Specialization string             `json:"specialization" bson:"specialization"`
	Department     string             `json:"department" bson:"department"`
	Experience     string             `json:"experience" bson:"experience"`
	Address        Address            `json:"address" bson:"address"`
}
