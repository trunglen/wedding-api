package report

import (
	"g/x/web"
	"github.com/gin-gonic/gin"
	"wedding-api/o/wedding"
)

type ReportServer struct {
	*gin.RouterGroup
	web.JsonRender
}

func NewReportServer(parent *gin.RouterGroup, name string) *ReportServer {
	var s = ReportServer{
		RouterGroup: parent.Group(name),
	}
	s.GET("restaurant", s.getRestaurantReport)
	s.GET("wedding", s.getWeddingReport)
	return &s
}

func (s *ReportServer) getRestaurantReport(c *gin.Context) {
	result, err := wedding.GetRestaurantReport()
	web.AssertNil(err)
	s.SendData(c, result)
}

func (s *ReportServer) getWeddingReport(c *gin.Context) {
	result, err := wedding.GetWeddingReport("userID string", "role string")
	web.AssertNil(err)
	s.SendData(c, result)
}
