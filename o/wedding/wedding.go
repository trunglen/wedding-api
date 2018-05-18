package wedding

import (
	"g/x/math"
	"g/x/web"
	"gopkg.in/mgo.v2/bson"
	"wedding-api/x/mongodb"
)

const (
	STATUS_NEW            = "missing"
	STATUS_FULL           = "full"
	STATUS_FINISH         = "finish"
	STATUS_SAFE           = "safe"
	STATUS_STUDENT_MOVE   = "move"
	STATUS_STUDENT_JOIN   = "join"
	STATUS_STUDENT_FINISH = "finish"
)

type Wedding struct {
	mongodb.Model    `bson:",inline"`
	Phone            string    `bson:"phone" json:"phone"`
	Address          Address   `bson:"address" json:"address"`
	HTime            int64     `bson:"htime" json:"htime"`
	RestaurantID     string    `bson:"restaurant_id" json:"restaurant_id"`
	Price            int64     `bson:"price" json:"price"`
	CreatedBy        string    `bson:"created_by" json:"created_by"`
	Students         []Student `bson:"students" json:"students"`
	Status           Status    `bson:"status" json:"status"`
	VerifyCode       string    `bson:"verify_code" json:"verify_code"`
	NumberOfStudents int       `bson:"number_of_students" json:"number_of_students"`
	Clothes          string    `bson:"clothes" json:"clothes"`
	Note             string    `bson:"note" json:"note"`
	PaidSystem       bool      `bson:"paid_system" json:"paid_system"`
	PaidStudent      bool      `bson:"paid_student" json:"paid_student"`
}
type Status string
type Student struct {
	ID     string `bson:"id" json:"id"`
	Phone  string `bson:"phone" json:"phone"`
	Sex    bool   `bson:"sex" json:"sex"`
	Name   string `bson:"name" json:"name"`
	Status string `bson:"status" json:"status"`
}
type Address struct {
	HomeNumber string `bson:"home_number" json:"home_number"`
	Street     string `bson:"street" json:"street"`
	District   string `bson:"district" json:"district"`
}

var weddingTable = mongodb.NewTable("wedding", "wed")

func (w *Wedding) Create() error {
	w.VerifyCode = math.RandNumString(5)
	w.Students = []Student{}
	w.Status = STATUS_NEW
	return weddingTable.Create(w)
}

func (w *Wedding) AcceptStudent(s Student) error {
	if err := w.CheckExistStudent(s); err != nil {
		return err
	}
	male, felmale := w.CountNumberStudentAccept()
	if s.Sex {
		if male == w.NumberOfStudents/2 {
			return web.BadRequest("Đám này đã đủ số lượng nam")
		}
	} else {
		if felmale == w.NumberOfStudents/2 {
			return web.BadRequest("Đám này đã đủ số lượng nữ")
		}
	}
	w.Students = append(w.Students, s)
	var update = bson.M{"$push": bson.M{
		"students": s,
	}}
	if len(w.Students) == w.NumberOfStudents {
		w.Status = "full"
		weddingTable.UpdateId(w.ID, bson.M{"$set": bson.M{"status": "full"}})
	}
	return weddingTable.UpdateId(w.ID, update)
}

func (w *Wedding) UpdateStudentStatus(s Student, verifyCode string) error {
	var countStatus = 0
	if s.Status == STATUS_FINISH {
		if w.VerifyCode != verifyCode {
			return web.BadRequest("Mã xác nhận không chính xác")
		}
	}
	for i, item := range w.Students {
		if item.ID == s.ID {
			w.Students[i] = s
			item.Status = s.Status
		}
		if item.Status == s.Status {
			countStatus++
		}
	}
	var update = bson.M{}
	if countStatus == len(w.Students) && len(w.Students) == w.NumberOfStudents {
		update["status"] = s.Status
		w.Status = Status(s.Status)
	}
	update["students"] = w.Students
	return weddingTable.UpdateId(w.ID, bson.M{"$set": update})
}

func (w *Wedding) MoveToWedding(s Student) error {
	if err := w.CheckExistStudent(s); err != nil {
		return err
	}
	if male, felmale := w.CountNumberStudentAccept(); male == w.NumberOfStudents/2 || felmale == w.NumberOfStudents/2 {
		return web.BadRequest("Đám này đã đủ số lượng")
	}
	w.Students = append(w.Students, s)
	return weddingTable.UpdateId(w.ID, bson.M{"$push": bson.M{
		"students": s,
	}})
}

func (w *Wedding) CountNumberStudentAccept() (int, int) {
	var countMale = 0
	var countFelmale = 0
	for _, item := range w.Students {
		if item.Sex {
			countMale++
		} else {
			countFelmale++
		}
	}
	return countMale, countFelmale
}

func (w *Wedding) CheckExistStudent(s Student) error {
	for _, item := range w.Students {
		if item.ID == s.ID {
			return web.BadRequest("Bạn đã tham gia đám này")
		}
	}
	return nil
}

// GetWeddingsByRole GetWeddingsByRole
func GetWeddingsByRole(userID string, role string) ([]*Wedding, error) {
	var result []*Wedding
	var query = bson.M{}
	if role != "super-admin" {
		// if role == "manager" {
		// 	query["created_by"] = userID
		// } else {
		// 	query["restaurant_id"] = userID
		// }
		query["restaurant_id"] = userID
	}
	var err = weddingTable.FindWhere(query, &result)
	return result, err
}
