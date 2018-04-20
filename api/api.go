package api

import (
	"wedding-api/api/admin"
	"wedding-api/api/auth"
	"wedding-api/api/guest"
	"wedding-api/api/public"

	"github.com/gin-gonic/gin"
)

func NewApiServer(root *gin.RouterGroup) {
	admin.NewAdminServer(root, "admin")
	auth.NewAuthServer(root, "auth")
	public.NewPublicServer(root, "public")
	guest.NewGuestServer(root, "guest")
}
