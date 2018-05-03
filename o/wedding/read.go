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
		"status":      status,
	}
	// if status == "missing" {
	// 	query["$where"] = "this.students.filter(x=>x.sex==true).length<5"
	// }
	var err = weddingTable.Find(query).Skip((page - 1) * LIMIT).Limit(LIMIT).All(&result)
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
