package student

import (
	"g/x/web"
	"github.com/gin-gonic/gin"
	"wedding-api/middleware"
	"wedding-api/o/user"
)

type SupervisorServer struct {
	*gin.RouterGroup
	web.JsonRender
}

func NewSupervisorServer(parent *gin.RouterGroup, name string) *SupervisorServer {
	var s = SupervisorServer{
		RouterGroup: parent.Group(name),
	}
	s.Use(middleware.MustBeManager)
	s.GET("list", s.getUsers)
	s.POST("create", s.createUser)
	s.POST("update", s.updateUser)
	s.GET("delete", s.deleteUser)
	return &s
}

func (s *SupervisorServer) getUsers(c *gin.Context) {
	var users, err = user.GetUsers("supervisor")
	web.AssertNil(err)
	s.SendData(c, users)
}

func (s *SupervisorServer) createUser(c *gin.Context) {
	var user *user.User
	user.Role = "supervisor"
	web.AssertNil(c.BindJSON(&user))
	web.AssertNil(user.Create())
	s.SendData(c, user)
}

func (s *SupervisorServer) updateUser(c *gin.Context) {
	var user *user.User
	web.AssertNil(c.BindJSON(&user))
	web.AssertNil(user.Create())
	s.SendData(c, user)
}

func (s *SupervisorServer) deleteUser(c *gin.Context) {
	var id = c.Query("id")
	web.AssertNil(user.DeleteUserByID(id))
	s.Success(c)
}
