package wedding

import (
	"gopkg.in/mgo.v2/bson"
)

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

func GetMisingWedding(userID string, sex bool, page int) ([]*Wedding, error) {
	var result []*Wedding
	var query = bson.M{
		"students.id": bson.M{
			"$ne": userID,
		},
		"status": STATUS_NEW,
		"$where": "this.students.filter(x=>x.sex==true).length<5",
	}

	var err = weddingTable.Find(query).Skip((page - 1) * LIMIT).Limit(LIMIT).All(&result)
	return result, err
}

func GetWedding(id string) (*Wedding, error) {
	var result *Wedding
	var err = weddingTable.FindID(id, &result)
	return result, err
}

func DeleteWeddingByID(id string) error {
	return weddingTable.RemoveId(id)
}

const LIMIT = 10
