package manager

import (
	"g/x/web"
	"github.com/gin-gonic/gin"
	"wedding-api/middleware"
	"wedding-api/o/wedding"
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
	return &s
}

func (s *ManagerServer) createWedding(c *gin.Context) {
	var wedding *wedding.Wedding
	web.AssertNil(c.BindJSON(&wedding))
	web.AssertNil(wedding.Create())
	s.SendData(c, wedding)
}

func (s *ManagerServer) getWeddings(c *gin.Context) {
	var result, err = wedding.GetWeddings()
	web.AssertNil(err)
	s.SendData(c, result)
}
