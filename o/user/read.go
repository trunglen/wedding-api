package user

import (
	"g/x/web"
	"gopkg.in/mgo.v2/bson"
	"wedding-api/x/mongodb"
)

var ErrMismatchedHashAndPassword = "crypto/bcrypt: hashedPassword is not the hash of the given password"

func GetByID(id string) (*User, error) {
	var user *User
	return user, userTable.FindID(id, &user)
}
func GetUsers(role string) ([]*User, error) {
	var users []*User
	var query = bson.M{
		"role": role,
		"dtime": bson.M{
			"$ne": 0,
		},
	}
	err := userTable.Find(query).Sort("-ctime").All(&users)
	return users, err
}
func GetUsersByRole(role, userID, userRole string) ([]*User, error) {
	var users []*User
	var query = bson.M{
		"role": role,
		"dtime": bson.M{
			"$ne": 0,
		},
	}
	if userRole != "super-admin" {
		query["restaurant_id"] = userID
	}
	err := userTable.Find(query).Sort("-ctime").All(&users)
	return users, err
}

func GetManagers() ([]*User, error) {
	var users []*User
	var match = mongodb.Match(bson.M{
		"role": "manager",
		"dtime": bson.M{
			"$ne": 0,
		},
	})
	var selfJoin = mongodb.Join("user", "restaurant_id", "_id", "restaurant")
	var unwind = mongodb.Unwind("restaurant")
	var project = mongodb.Project(bson.M{
		"_id":             "$_id",
		"name":            "$name",
		"phone":           "$phone",
		"restaurant_id":   "$restaurant._id",
		"restaurant_name": "$restaurant.restaurant_name",
		"ctime":           "$ctime",
	})
	var orderby = mongodb.OrderBy("ctime", -1)
	var pipe = []bson.M{match, selfJoin, unwind, project, orderby}
	err := userTable.Pipe(pipe).All(&users)
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

func GetUserByRestaurantID(restaurantID string) ([]*User, error) {
	var user []*User
	err := userTable.FindWhere(bson.M{"restaurant_id": restaurantID}, &user)
	return user, err
}

func GetUserIDByRestaurantID(restaurantID string) []string {
	var users, _ = GetUserByRestaurantID(restaurantID)
	var userIDS = make([]string, 0)
	if users != nil {
		for _, item := range users {
			userIDS = append(userIDS, item.ID)
		}
	}
	return userIDS
}
