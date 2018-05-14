package wedding

import (
	"gopkg.in/mgo.v2/bson"
	"wedding-api/x/logger"
	"wedding-api/x/mongodb"
)

var weddingLog = logger.NewLogger("wedding")

func GetWeddings() ([]*Wedding, error) {
	var result []*Wedding
	var err = weddingTable.FindAll(&result)
	return result, err
}

func GetMyWeddingByStatus(userID string, status string, page int) ([]*Wedding, error) {
	var result []*Wedding
	var query = bson.M{
		"students.id": userID,
	}
	if status == STATUS_FINISH {
		query["status"] = STATUS_FINISH
	} else {
		query["status"] = bson.M{
			"$ne": STATUS_FINISH,
		}
	}
	// if status == "missing" {
	// 	query["$where"] = "this.students.filter(x=>x.sex==true).length<5"
	// }
	var err = weddingTable.Find(query).Sort("-$ctime").Skip((page - 1) * LIMIT).Limit(LIMIT).All(&result)
	return result, err
}

func GetMisingWedding(restaurantID, userID string, sex bool, page int) ([]*Wedding, error) {
	var result []*Wedding
	var query = bson.M{
		"students.id": bson.M{
			"$ne": userID,
		},
		"restaurant_id": restaurantID,
		"status":        STATUS_NEW,
		"$where":        "this.students.filter(x=>x.sex==true).length<5",
	}

	var err = weddingTable.Find(query).Skip((page - 1) * LIMIT).Limit(LIMIT).All(&result)
	return result, err
}

func GetWedding(id string) (*Wedding, error) {
	var result *Wedding
	var err = weddingTable.FindID(id, &result)
	return result, err
}

func GetWeddingDetail(id string) (*WeddingDetail, error) {
	var result *WeddingDetail
	var match = mongodb.Match(bson.M{"_id": id})
	var joinRestaurant = mongodb.Join("user", "restaurant_id", "_id", "restaurant")
	var unwindRestaurant = mongodb.Unwind("restaurant")
	var joinManager = mongodb.Join("user", "created_by", "_id", "manager")
	var unwindManager = mongodb.Unwind("manager")
	var addField = mongodb.AddField(bson.M{
		"manager_phone":   "$manager.phone",
		"manager_name":    "$manager.name",
		"restaurant_name": "$restaurant.name",
	})
	var err = weddingTable.Pipe([]bson.M{match, joinRestaurant, unwindRestaurant, joinManager, unwindManager, addField}).One(&result)
	if err != nil {
		weddingLog.Error(err)
		return nil, err
	}
	return result, err
}

type WeddingDetail struct {
	Wedding        `bson:",inline"`
	ManagerName    string `bson:"manager_name" json:"manager_name"`
	RestaurantName string `bson:"restaurant_name" json:"restaurant_name"`
	ManagerPhone   string `bson:"manager_phone" json:"manager_phone"`
}

func DeleteWeddingByID(id string) error {
	return weddingTable.RemoveId(id)
}

const LIMIT = 10
