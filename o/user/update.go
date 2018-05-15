package user

import (
	"gopkg.in/mgo.v2/bson"
	"net/http"
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
				"weight":    u.Information.Weight,
				"height":    u.Information.Height,
				"sex":       u.Information.Sex,
				"birth_day": u.Information.BirthDay,
			},
		},
	})
}
func (u *User) SetAvatarAndUpload(r *http.Request) {
	u.Information.Avatar = "http://" + r.Host + "/static/student/avatar/" + u.ID
	u.Information.Portrait = "http://" + r.Host + "/static/student/portrait/" + u.ID
}
