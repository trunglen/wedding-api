package manager

import (
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
	fcm.SendToMany(pushTokens, fcm.FmcMessage{Title: "Có một đám cưới mới", Body: "Có một đám cưới mới", Data: map[string]interface{}{
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

func (s *ManagerServer) getWedding(c *gin.Context) {
	var id = c.Query("id")
	var result, err = wedding.GetWedding(id)
	web.AssertNil(err)
	s.SendData(c, result)
}
