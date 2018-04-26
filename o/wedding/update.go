package wedding

import (
	"gopkg.in/mgo.v2/bson"
)

type WeddingInfo struct {
	ID               string  `bson:"-" json:"id"`
	Phone            string  `bson:"phone" json:"phone"`
	HTime            string  `bson:"htime" json:"htime"`
	NumberOfStudents int     `bson:"number_of_students" json:"number_of_students"`
	Address          Address `bson:"address" json:"address"`
}

func (w *WeddingInfo) UpdateWedding() error {
	return weddingTable.UpdateId(w.ID, bson.M{
		"$set": w,
	})
}
