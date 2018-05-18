package wedding

import (
	"gopkg.in/mgo.v2/bson"
	"wedding-api/x/mongodb"
)

type GeneralReport struct {
	All     int `json:"all"`
	Finish  int `json:"finish"`
	Missing int `json:"missing"`
}

func GetGeneralReportBySupervisor(userID string, role string) (*GeneralReport, error) {
	var res *GeneralReport

	var match = mongodb.Match(bson.M{"restaurant_id": userID})
	var project = bson.M{
		"$project": bson.M{
			"item": "1",
			"finish": bson.M{
				"$cond": []interface{}{bson.M{"$eq": []string{"$status", "finish"}}, 1, 0},
			},
			"missing": bson.M{
				"$cond": []interface{}{bson.M{"$eq": []string{"$status", "missing"}}, 1, 0},
			},
			"all": bson.M{
				"$cond": []interface{}{bson.M{"$eq": []int{1, 1}}, 1, 0},
			},
		},
	}
	var group = bson.M{
		"$group": bson.M{
			"_id":     "$item",
			"finish":  bson.M{"$sum": "$finish"},
			"missing": bson.M{"$sum": "$missing"},
			"all":     bson.M{"$sum": "$all"},
		},
	}
	var pipe = []bson.M{}
	if role == "super-admin" {
		pipe = []bson.M{project, group}
	} else {
		pipe = []bson.M{match, project, group}
	}
	err := weddingTable.Pipe(pipe).One(&res)
	return res, err
}
