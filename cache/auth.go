package cache

import (
	"g/x/web"
	"wedding-api/o/auth"
)

func MustGetAuth(token string) (*auth.Auth, error) {
	if authen, ok := authCache[token]; ok {
		return authen, nil
	}
	authen, err := auth.GetByID(token)
	if authen != nil {
		authCache[authen.ID] = authen
		return authen, nil
	}
	web.AssertNil(web.Unauthorized("access_token not found"))
	return nil, err
}

func CreateAuth(userID, role string) *auth.Auth {
	var authen = auth.Create(userID, role)
	authCache[authen.ID] = authen
	return authen
}
