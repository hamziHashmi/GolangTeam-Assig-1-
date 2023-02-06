package models



type User struct {

	Id int `json:"id" bson:"id"`
	Date string `json:"date" bson:"date"`
	Time string `json:"time" bson:"time"`
	DoctorName    string  `json:"doctorname" bson:"doctorname"`
}
