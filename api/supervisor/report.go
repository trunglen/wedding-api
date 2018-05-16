package supervisor

import (
	"g/x/web"
	"github.com/gin-gonic/gin"
	"wedding-api/o/auth"
	"wedding-api/o/user"
	"wedding-api/o/wedding"
)

func (s *SupervisorServer) getGeneralReport(c *gin.Context) {
	var au, err = auth.GetByID(web.GetToken(c.Request))
	web.AssertNil(err)
	var res = struct {
		*wedding.GeneralReport
		Cashback int `json:"cashback"`
	}{}
	general, err := wedding.GetGeneralReportBySupervisor(au.UserID, au.Role)
	cashback := user.GetStudentCashback(au.UserID, "supervisor")
	web.AssertNil(err)
	res.GeneralReport = general
	res.Cashback = cashback
	s.SendData(c, res)
}
