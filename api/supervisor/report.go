package supervisor

import (
	"g/x/web"
	"github.com/gin-gonic/gin"
	"wedding-api/o/auth"
	"wedding-api/o/wedding"
)

func (s *SupervisorServer) getGeneralReport(c *gin.Context) {
	var au, err = auth.GetByID(web.GetToken(c.Request))
	web.AssertNil(err)
	var res = struct {
		*wedding.GeneralReport
	}{}
	general, err := wedding.GetGeneralReportBySupervisor(au.UserID, au.Role)
	web.AssertNil(err)
	s.SendData(c, generral)
}
