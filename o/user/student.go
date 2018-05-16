package user

import (
	"gopkg.in/mgo.v2/bson"
)

//using fir update
type StudentProfile struct {
	ID          string `json:"id"`
	Name        string `bson:"name" json:"name"`
	Information struct {
		Weight   float32 `bson:"weight" json:"weight"`
		Height   float32 `bson:"height" json:"height"`
		Sex      bool    `bson:"sex" json:"sex"`
		BirthDay string  `bson:"birth_day" json:"birth_day"`
	} `bson:"information" json:"information,omitempty"`
}

func (u *StudentProfile) UpdateStudent() error {
	return userTable.UpdateId(u.ID, bson.M{"$set": u})
}
