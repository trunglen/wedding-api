package user

import (
	"wedding-api/x/logger"
	"wedding-api/x/mongodb"
)

var transactionTable = mongodb.NewTable("transaction", "tran")
var transactionLog = logger.NewLogger("transaction")

type Transaction struct {
	mongodb.Model `bson:",inline"`
	RestaurantID  string `bson:"restaurant_id" json:"restaurant_id"`
	Total         int64  `bson:"total" json:"total"`
}

func (t *Transaction) Create() error {
	return transactionTable.Create(t)
}
