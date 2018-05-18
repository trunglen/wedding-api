package manager

import (
	"g/x/web"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"wedding-api/middleware"
	"wedding-api/o/auth"
	"wedding-api/o/push_token"
	"wedding-api/o/user"
	"wedding-api/o/wedding"
	"wedding-api/x/fcm"
)

type ManagerServer struct {
	*gin.RouterGroup
	web.JsonRender
}

func NewManagerServer(parent *gin.RouterGroup, name string) *ManagerServer {
	var s = ManagerServer{
		RouterGroup: parent.Group(name),
	}
	s.Use(middleware.MustBeManager)
	s.POST("wedding/create", s.createWedding)
	s.GET("wedding/list", s.getWeddings)
	s.GET("wedding/get", s.getWedding)
	s.GET("wedding/detail", s.getWedding)
	s.POST("wedding/update", s.updateWedding)
	s.GET("wedding/delete", s.deleteWedding)
	return &s
}

func (s *ManagerServer) createWedding(c *gin.Context) {
	var wedding *wedding.Wedding
	web.AssertNil(c.BindJSON(&wedding))
	web.AssertNil(wedding.Create())
	var userIDS = user.GetUserIDByRestaurantID(wedding.RestaurantID)
	var pushTokens, _ = push_token.GetAllPushToken(userIDS)
	glog.Info(pushTokens)
	go func() {
		fcm.SendToMany(pushTokens, fcm.FmcMessage{Title: "Có một đám cưới mới", Body: "Có một đám cưới mới", Data: map[string]interface{}{
			"title": "Có một đám cưới mới",
			"id":    wedding.ID,
		}})
	}()
	s.SendData(c, wedding)
}

func (s *ManagerServer) getWeddings(c *gin.Context) {
	var au, err = auth.GetByID(web.GetToken(c.Request))
	web.AssertNil(err)
	user, err := user.GetByID(au.UserID)
	web.AssertNil(err)
	result, err := wedding.GetWeddingsByRole(user.RestaurantID, au.Role)
	web.AssertNil(err)
	s.SendData(c, result)
}

func (s *ManagerServer) getWedding(c *gin.Context) {
	var id = c.Query("id")
	var result, err = wedding.GetWedding(id)
	web.AssertNil(err)
	s.SendData(c, result)
}

func (s *ManagerServer) deleteWedding(c *gin.Context) {
	var id = c.Query("id")
	web.AssertNil(wedding.DeleteWeddingByID(id))
	s.Success(c)
}

func (s *ManagerServer) updateWedding(c *gin.Context) {
	var wed *wedding.WeddingInfo
	c.BindJSON(&wed)
	web.AssertNil(wed.UpdateWedding())
	s.Success(c)
}
