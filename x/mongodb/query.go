package mongodb

import (
	"gopkg.in/mgo.v2/bson"
)

func Join(table, localField, foreignField, as string) bson.M {
	return bson.M{
		"$lookup": bson.M{
			"from":         table,
			"localField":   localField,
			"foreignField": foreignField,
			"as":           as,
		},
	}
}

func Match(match bson.M) bson.M {
	return bson.M{
		"$match": match,
	}
}
func Project(match bson.M) bson.M {
	return bson.M{
		"$project": match,
	}
}
func Unwind(field string) bson.M {
	return bson.M{
		"$unwind": bson.M{
			"path": "$" + field,
		},
	}
}
func OrderBy(field string, order int) bson.M {
	return bson.M{
		"$sort": bson.M{
			field: order,
		},
	}
}
