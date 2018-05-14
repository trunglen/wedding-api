package user

import (
	"gopkg.in/mgo.v2/bson"
)

func (u *User) UpBalance(money int64) error {
	return userTable.UpdateId(u.ID, bson.M{
		"$inc": bson.M{
			"information.balance":        money,
			"information.finish_wedding": 1,
		},
	})
}
func (u *User) UpdateProfile() error {
	return userTable.UpdateId(u.ID, bson.M{
		"$set": bson.M{
			"name": u.Name,
			"information": bson.M{
				"weight":     u.Information.Weight,
				"height":     u.Information.Height,
				"sex":        u.Information.Sex,
				"birth_year": u.Information.BirthYear,
			},
		},
	})
}
