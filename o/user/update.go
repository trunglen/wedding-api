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
