package admin

import (
	"g/x/web"
	"github.com/gin-gonic/gin"
	"wedding-api/api/admin/category"
	"wedding-api/api/admin/post"
	"wedding-api/api/admin/user"
	"wedding-api/middleware"
)

type AdminServer struct {
	*gin.RouterGroup
	web.JsonRender
}

func NewAdminServer(parent *gin.RouterGroup, name string) *AdminServer {
	var s = AdminServer{
		RouterGroup: parent.Group(name),
	}
	s.Use(middleware.MustBeAdmin)
	post.NewPostServer(s.RouterGroup, "post")
	category.NewCategoryServer(s.RouterGroup, "category")
	user.NewUserServer(s.RouterGroup, "user")
	return &s
}
