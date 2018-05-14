package push_token

import (
	"gopkg.in/mgo.v2/bson"
	"wedding-api/x/mongodb"
)

var pushTable = mongodb.NewTable("push_token", "push")

type PushToken struct {
	PushToken string `bson:"push_token" json:"push_token" `
	UserID    string `bson:"user_id" json:"user_id" `
}

func (push *PushToken) Create() error {
	var _, err = pushTable.Upsert(bson.M{
		"user_id": push.UserID,
	}, push)
	return err
}

func GetAllPushToken(userID []string) ([]string, error) {
	var push []*PushToken
	var result = make([]string, 0)
	var err = pushTable.FindWhere(bson.M{"user_id": bson.M{"$in": userID}}, &push)
	if push != nil {
		for _, item := range push {
			result = append(result, item.PushToken)
		}
	}
	return result, err
}

// func GetByDeviceID(deviceID string) (*PushToken, error) {
// 	var push *PushToken
// 	return push, pushTable.Find(bson.M{
// 		"device_id": deviceID,
// 	}).One(&push)
// }
