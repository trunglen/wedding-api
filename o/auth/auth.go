package auth

import (
	"g/x/math"
	"gopkg.in/mgo.v2/bson"
	"wedding-api/x/mongodb"
)

var authTable = mongodb.NewTable("auth", "auth")

type Auth struct {
	mongodb.Model `bson:",inline"`
	UserID        string `bson:"user_id" json:"user_id"`
	Role          string `bson:"role" json:"role"`
	Revoked       bool   `bson:"revoked" json:"revoked"`
}

func Create(userID, role string) *Auth {
	var auth = &Auth{
		UserID: userID,
		Role:   role,
	}
	auth.SetID(math.RandString("auth", 80))
	if auth != nil {
		authTable.Remove(bson.M{
			"user_id": userID,
			"role":    role,
		})
	}
	authTable.Insert(auth)
	return auth
}

func GetByID(id string) (*Auth, error) {
	var auth *Auth
	return auth, authTable.FindID(id, &auth)
}

func GetTokens() ([]*Auth, error) {
	var auth []*Auth
	return auth, authTable.FindAll(&auth)
}

func RemoveToken(userID string) {
	authTable.Remove(bson.M{"user_id": userID})
}
