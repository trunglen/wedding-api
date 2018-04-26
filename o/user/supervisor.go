package user

import (
	"gopkg.in/mgo.v2/bson"
)

//using fir update
type SupervisorProfile struct {
	ID                string `json:"id"`
	Name              string `bson:"name" json:"name"`
	Phone             string `bson:"phone" json:"phone"`
	RestaurantName    string `bson:"restaurant_name" json:"restaurant_name,omitempty"`
	RestaurantAddress string `bson:"restaurant_address" json:"restaurant_address,omitempty"`
}

func (u *SupervisorProfile) UpdateSupervisor() error {
	return userTable.UpdateId(u.ID, bson.M{"$set": u})
}
