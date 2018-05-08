package admin

import (
	"g/x/web"
	"github.com/gin-gonic/gin"
	"wedding-api/api/admin/manager"
	"wedding-api/api/admin/report"
	"wedding-api/api/admin/user"
)

type AdminServer struct {
	*gin.RouterGroup
	web.JsonRender
}

func NewAdminServer(parent *gin.RouterGroup, name string) *AdminServer {
	var s = AdminServer{
		RouterGroup: parent.Group(name),
	}
	user.NewUserServer(s.RouterGroup, "user")
	report.NewReportServer(s.RouterGroup, "report")
	manager.NewManagerServer(s.RouterGroup, "manager")
	return &s
}
