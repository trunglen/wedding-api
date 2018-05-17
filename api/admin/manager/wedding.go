package manager

import (
	"fmt"
	"g/x/web"
	"github.com/gin-gonic/gin"
	"wedding-api/o/auth"
	"wedding-api/o/push_token"
	"wedding-api/o/user"
	"wedding-api/o/wedding"
	"wedding-api/x/fcm"
)

func (s *ManagerServer) createWedding(c *gin.Context) {
	var wedding *wedding.Wedding
	web.AssertNil(c.BindJSON(&wedding))
	web.AssertNil(wedding.Create())
	var userIDS = user.GetUserIDByRestaurantID(wedding.RestaurantID)
	var pushTokens, _ = push_token.GetAllPushToken(userIDS)
	var address = wedding.Address
	fcm.SendToMany(pushTokens, fcm.FmcMessage{
		Title: fmt.Sprintf("Đám cưới mới: Số nhà %s, đường %s, quận %s", address.HomeNumber, address.Street, address.District),
		Body:  fmt.Sprintf("%d", wedding.Price),
		Data: map[string]interface{}{
			"title": "Có một đám cưới mới",
			"id":    wedding.ID,
		}})
	s.SendData(c, wedding)
}

func (s *ManagerServer) getWeddings(c *gin.Context) {
	var au, err = auth.GetByID(web.GetToken(c.Request))
	web.AssertNil(err)
	result, err := wedding.GetWeddingsByRole(au.UserID, au.Role)
	web.AssertNil(err)
	s.SendData(c, result)
}

func (s *ManagerServer) getMoveWarningWeddings(c *gin.Context) {
	var au, err = auth.GetByID(web.GetToken(c.Request))
	web.AssertNil(err)
	result, err := wedding.GetWeddingsByRole(au.UserID, au.Role)
	web.AssertNil(err)
	s.SendData(c, result)
}

func (s *ManagerServer) getMissingWarningWeddings(c *gin.Context) {
	result, err := wedding.GetMisingWarningWedding(c.Query("restaurant_id"))
	web.AssertNil(err)
	s.SendData(c, result)
}

func (s *ManagerServer) getWedding(c *gin.Context) {
	var id = c.Query("id")
	var result, err = wedding.GetWedding(id)
	web.AssertNil(err)
	s.SendData(c, result)
}
