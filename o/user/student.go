package user

import (
	"gopkg.in/mgo.v2/bson"
)

//using fir update
type StudentProfile struct {
	ID          string      `json:"id"`
	Name        string      `bson:"name" json:"name"`
	Information Information `bson:"information" json:"information,omitempty"`
}

func (u *StudentProfile) UpdateStudent() error {
	return userTable.UpdateId(u.ID, bson.M{"$set": u})
}
