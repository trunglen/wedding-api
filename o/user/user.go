package user

import (
	"g/x/web"
	"gopkg.in/mgo.v2/bson"
	"wedding-api/o/wedding"
	"wedding-api/x/logger"
	"wedding-api/x/mongodb"
	"wedding-api/x/validator"
)

var userTable = mongodb.NewTable("user", "usr")
var userLog = logger.NewLogger("user")

const (
	GUEST = "guest"
	ADMIN = "admin"
)

type BaseUser struct {
	mongodb.Model  `bson:",inline"`
	Name           string   `bson:"name" json:"name"`
	Phone          string   `bson:"phone" json:"phone"`
	HashedPassword string   `bson:"password" json:"-"`
	Password       Password `bson:"-" json:"password,omitempty"`
}

type User struct {
	mongodb.Model     `bson:",inline"`
	Name              string      `bson:"name" json:"name"`
	Phone             string      `bson:"phone" json:"phone"`
	HashedPassword    string      `bson:"password" json:"-"`
	Password          Password    `bson:"-" json:"password,omitempty"`
	RestaurantName    string      `bson:"restaurant_name" json:"restaurant_name,omitempty"`
	RestaurantID      string      `bson:"restaurant_id" json:"restaurant_id,omitempty"`
	RestaurantAddress string      `bson:"restaurant_address" json:"restaurant_address,omitempty"`
	Role              string      `bson:"role" json:"role"`
	Information       Information `bson:"information" json:"information,omitempty"`
}

type Information struct {
	Weight       float32 `bson:"weight" json:"weight"`
	Height       float32 `bson:"height" json:"height"`
	Sex          bool    `bson:"sex" json:"sex"`
	Rating       float32 `bson:"rating" json:"rating"`
	Balance      float32 `bson:"balance" json:"balance"`
	RestaurantID string  `bson:"restaurant_id" json:"restaurant_id,omitempty"`
	Avatar       string  `bson:"-" json:"avatar"`
	Portrait     string  `bson:"-" json:"portrait"`
	BirthYear    int     `bson:"birth_year" json:"birth_year"`
}

const (
	errExists           = "user exists!"
	errPhoneExists      = "Số điện thoại đã tồn tại trong hệ thống!"
	errMisMatchUNamePwd = "username or password is incorect!"
)

func (u *User) CreateAccount() error {
	if user, _ := GetUserByPhone(u.Phone); user != nil {
		return web.BadRequest(errExists)
	}
	return userTable.Create(u)
}

func GetUserByPhone(phone string) (*User, error) {
	var user *User
	var err = userTable.FindOne(bson.M{
		"phone": phone,
	}, &user)
	return user, err
}

func (u *User) Create() error {
	var err = validator.Struct(u)
	if err != nil {
		userLog.Error(err)
		return web.WrapBadRequest(err, "")
	}
	user, err := GetUserByPhone(u.Phone)
	if user != nil {
		return web.BadRequest(errPhoneExists)
	}
	hashed, _ := u.Password.Hash()
	u.HashedPassword = hashed
	return userTable.Create(u)
}

func GetAdmin(uname string, role string) (*User, error) {
	var user *User
	var err = userTable.FindOne(bson.M{
		"uname": uname,
		"role":  role,
	}, &user)
	return user, err
}

func GetGuestByEmail(email string) (*User, error) {
	var user *User
	var err = userTable.FindOne(bson.M{
		"email": email,
		"role":  GUEST,
	}, &user)
	return user, err
}

func DeleteUserByID(id string) error {
	return userTable.DeleteByID(id)
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

func ChangePassword(id, oldPwd, newPwd string) error {
	var user *User
	userTable.FindID(id, &user)
	if user != nil {
		if Password(oldPwd).ComparePassword(user.HashedPassword) != nil {
			return web.BadRequest("Mật khẩu cũ không chính xác")
		}
	} else {
		return web.BadRequest("user_id not found")
	}
	var newHashPwd, _ = Password(newPwd).Hash()
	return userTable.UpdateId(id, bson.M{
		"$set": bson.M{
			"password": string(newHashPwd),
		},
	})
}

func (user *User) ToStudent(status string) wedding.Student {
	return wedding.Student{
		ID:     user.ID,
		Name:   user.Name,
		Sex:    user.Information.Sex,
		Status: status,
		Phone:  user.Phone,
	}
}
