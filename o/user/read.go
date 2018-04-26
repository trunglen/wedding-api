package user

import (
	"g/x/web"
	"gopkg.in/mgo.v2/bson"
)

var ErrMismatchedHashAndPassword = "crypto/bcrypt: hashedPassword is not the hash of the given password"

func GetByID(id string) (*User, error) {
	var user *User
	return user, userTable.FindID(id, &user)
}

func GetUsers(role string) ([]*User, error) {
	var users []*User
	err := userTable.Find(bson.M{
		"role": role,
		"dtime": bson.M{
			"$ne": 0,
		},
	}).Sort("-ctime").All(&users)
	return users, err
}

func Login(uname, pwd string) (*User, error) {
	var user *User
	var query = bson.M{"phone": uname}
	err := userTable.FindOne(query, &user)
	if err != nil {
		if err.Error() == "not found" {
			return nil, web.BadRequest("Sai tên đăng nhập hoặc mật khẩu")
		}
		return nil, err
	}
	if err := Password(pwd).ComparePassword(user.HashedPassword); err != nil {
		if err.Error() == ErrMismatchedHashAndPassword {
			userLog.Error(err)
			return nil, web.BadRequest("Sai tên đăng nhập hoặc mật khẩu")
		}
		return nil, err
	}
	return user, nil
}

func LoginByStudent(uname, pwd string) (*User, error) {
	var user *User
	var query = bson.M{"phone": uname, "role": "student"}
	err := userTable.FindOne(query, &user)
	if err != nil {
		if err.Error() == "not found" {
			return nil, web.BadRequest("Sai tên đăng nhập hoặc mật khẩu")
		}
		return nil, err
	}
	if err := Password(pwd).ComparePassword(user.HashedPassword); err != nil {
		if err.Error() == ErrMismatchedHashAndPassword {
			userLog.Error(err)
			return nil, web.BadRequest("Sai tên đăng nhập hoặc mật khẩu")
		}
		return nil, err
	}
	return user, nil
}
