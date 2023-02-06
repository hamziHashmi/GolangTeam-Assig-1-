package models

type Address struct {
	State   string `json:"state" bson:"state"`
	City    string `json:"city" bson:"city"`
	Pincode int    `json:"pincode" bson:"pincode"`
}

type Patient struct {
	Id      int     `json:"id" bson:"patient_id"`
	Name    string  `json:"name" bson:"patient_name"`
	Age     int     `json:"age" bson:"patient_age"`
	Address Address `json:"address" bson:"patient_address"`
}
