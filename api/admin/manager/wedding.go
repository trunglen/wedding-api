package manager

import (
	"g/x/web"
	"github.com/gin-gonic/gin"
	"wedding-api/o/wedding"
)

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
