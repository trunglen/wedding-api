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
	s.SendData(c, wedding)
}

func (s *ManagerServer) getWeddings(c *gin.Context) {
	var result, err = wedding.GetWeddings()
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
