package wedding

import (
	"g/x/math"
	"wedding-api/x/mongodb"
)

type Wedding struct {
	mongodb.Model `bson:",inline"`
	Phone         string        `bson:"phone" json:"phone"`
	Address       Address       `bson:"address" json:"address"`
	HTime         string        `bson:"htime" json:"htime"`
	RestaurantID  string        `bson:"restaurant_id" json:"restaurant_id"`
	CreatedBy     string        `bson:"created_by" json:"created_by"`
	Students      StudentStatus `bson:"student_status" json:"student_status"`
	Status        Status        `bson:"status" json:"status"`
	VerifyCode    string        `bson:"verify_code" json:"verify_code"`
}
type Status string
type StudentStatus map[string]struct {
	Name   string
	Status string
}
type Address struct {
	HomeNumber string `bson:"home_number" json:"home_number"`
	Street     string `bson:"street" json:"street"`
	District   string `bson:"district" json:"district"`
}

var weddingTable = mongodb.NewTable("wedding", "wed")

func (w *Wedding) Create() error {
	w.VerifyCode = math.RandNumString(5)
	return weddingTable.Create(w)
}
