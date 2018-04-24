package wedding

import (
	"g/x/math"
	"gopkg.in/mgo.v2/bson"
	"wedding-api/x/mongodb"
)

const (
	STATUS_NEW    = "new"
	STATUS_MOVE   = "move"
	STATUS_FINISH = "finish"
	STATUS_JOIN   = "join"
)

type Wedding struct {
	mongodb.Model    `bson:",inline"`
	Phone            string    `bson:"phone" json:"phone"`
	Address          Address   `bson:"address" json:"address"`
	HTime            string    `bson:"htime" json:"htime"`
	RestaurantID     string    `bson:"restaurant_id" json:"restaurant_id"`
	CreatedBy        string    `bson:"created_by" json:"created_by"`
	Students         []Student `bson:"students" json:"students"`
	Status           Status    `bson:"status" json:"status"`
	VerifyCode       string    `bson:"verify_code" json:"verify_code"`
	NumberOfStudents int       `bson:"number_of_students" json:"number_of_students"`
}
type Status string
type Student struct {
	ID     string `bson:"id" json:"id"`
	Phone  string `bson:"phone" json:"phone"`
	Sex    bool   `bson:"sex" json:"sex"`
	Name   string `bson:"name" json:"name"`
	Status string `bson:"status" json:"status"`
}
type Address struct {
	HomeNumber string `bson:"home_number" json:"home_number"`
	Street     string `bson:"street" json:"street"`
	District   string `bson:"district" json:"district"`
}

var weddingTable = mongodb.NewTable("wedding", "wed")

func (w *Wedding) Create() error {
	w.VerifyCode = math.RandNumString(5)
	w.Students = []Student{}
	w.Status = STATUS_NEW
	return weddingTable.Create(w)
}

func (w *Wedding) AcceptStudent(s Student) error {
	w.Students = append(w.Students, s)
	return weddingTable.UpdateId(w.ID, bson.M{"$push": bson.M{
		"students": s,
	}})
}

func (w *Wedding) CountNumberStudentAccept() {

}
