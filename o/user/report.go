package user

import (
	"gopkg.in/mgo.v2/bson"
	"wedding-api/x/mongodb"
)

func GetStudentCashback(userID string, role string) int {
	var res = &struct {
		Cashback int `json:"cashback"`
	}{}
	var match = mongodb.Match(bson.M{"restaurant_id": userID, "role": role})
	var group = bson.M{
		"$group": bson.M{
			"_id":      "$restaurant_id",
			"cashback": bson.M{"$sum": "$balance"},
		},
	}
	userTable.Pipe([]bson.M{match, group}).One(res)
	if res != nil {
		return res.Cashback
	}
	return 0
}
