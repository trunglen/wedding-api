package student

import (
	"g/x/web"
	"github.com/gin-gonic/gin"
	"wedding-api/middleware"
	"wedding-api/o/user"
)

type StudentServer struct {
	*gin.RouterGroup
	web.JsonRender
}

func NewStudentServer(parent *gin.RouterGroup, name string) *StudentServer {
	var s = StudentServer{
		RouterGroup: parent.Group(name),
	}
	s.Use(middleware.MustBeManager)
	s.GET("list", s.getUsers)
	s.POST("create", s.createUser)
	s.POST("update", s.updateUser)
	s.GET("delete", s.deleteUser)
	return &s
}

func (s *StudentServer) getUsers(c *gin.Context) {
	var users, err = user.GetUsers("student")
	web.AssertNil(err)
	s.SendData(c, users)
}

func (s *StudentServer) createUser(c *gin.Context) {
	var user *user.User
	user.Role = "student"
	web.AssertNil(c.BindJSON(&user))
	web.AssertNil(user.Create())
	s.SendData(c, user)
}

func (s *StudentServer) updateUser(c *gin.Context) {
	var user *user.User
	web.AssertNil(c.BindJSON(&user))
	web.AssertNil(user.Create())
	s.SendData(c, user)
}

func (s *StudentServer) deleteUser(c *gin.Context) {
	var id = c.Query("id")
	web.AssertNil(user.DeleteUserByID(id))
	s.Success(c)
}
