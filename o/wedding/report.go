package wedding

import (
	"gopkg.in/mgo.v2/bson"
	"wedding-api/x/mongodb"
)

type RestaurantReport struct {
	RestaurantName string `json:"restaurant_name" bson:"restaurant_name"`
	RestaurantID   string `json:"restaurant_id" bson:"restaurant_id"`
	Total          int64  `json:"total" bson:"total"`
	WeddingCount   int    `json:"wedding_count" bson:"wedding_count"`
}

func GetRestaurantReport() ([]*RestaurantReport, error) {
	var res []*RestaurantReport
	var group = bson.M{
		"$group": bson.M{
			"_id": bson.M{
				"restaurant": "$restaurant_id",
			},
			"wedding_count": bson.M{"$sum": 1},
			"total": bson.M{
				"$sum": bson.M{
					"$multiply": []interface{}{bson.M{"$size": "$students"}, "$price"},
				},
			},
		},
	}
	var joinRestaurant = mongodb.Join("user", "_id.restaurant", "_id", "restaurant")
	var unwindRestaurant = mongodb.Unwind("restaurant")
	var addField = mongodb.AddField(bson.M{
		"restaurant_name": "$restaurant.restaurant_name",
		"restaurant_id":   "$restaurant.restaurant_id",
	})
	err := weddingTable.Pipe([]bson.M{group, joinRestaurant, unwindRestaurant, addField}).All(&res)
	return res, err
}

type WeddingReport struct {
	WeddingAddress Address `json:"wedding_address" bson:"wedding_address"`
	WeddingID      string  `json:"wedding_id" bson:"wedding_id"`
	Total          int64   `json:"total" bson:"total"`
	PaidStudent    bool    `json:"paid_student" bson:"paid_student"`
}

func GetWeddingReport(userID, role string) ([]*WeddingReport, error) {
	var res []*WeddingReport
	// var match = bson.M{
	// 	"$match": bson.M{
	// 		"restaurant_id": userID,
	// 	},
	// }
	var project = bson.M{
		"$project": bson.M{
			"wedding_address": "$address",
			"wedding_id":      "$_id",
			"paid_student":    "$paid_student",
			"total": bson.M{
				"$sum": bson.M{
					"$multiply": []interface{}{bson.M{"$size": "$students"}, "$price"},
				},
			},
		},
	}
	err := weddingTable.Pipe([]bson.M{project}).All(&res)
	return res, err
}
