package api

import (
	"wedding-api/api/admin"
	"wedding-api/api/auth"
	"wedding-api/api/guest"
	"wedding-api/api/manager"
	"wedding-api/api/public"
	"wedding-api/api/student"
	"wedding-api/api/supervisor"

	"github.com/gin-gonic/gin"
)

func NewApiServer(root *gin.RouterGroup) {
	admin.NewAdminServer(root, "admin")
	auth.NewAuthServer(root, "auth")
	public.NewPublicServer(root, "public")
	guest.NewGuestServer(root, "guest")
	supervisor.NewSupervisorServer(root, "supervisor")
	manager.NewManagerServer(root, "manager")
	student.NewStudentServer(root, "student")
}
