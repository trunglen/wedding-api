package manager

import (
	"g/x/web"
	"github.com/gin-gonic/gin"
	"wedding-api/middleware"
	"wedding-api/o/user"
)

type ManagerServer struct {
	*gin.RouterGroup
	web.JsonRender
}

func NewManagerServer(parent *gin.RouterGroup, name string) *ManagerServer {
	var s = ManagerServer{
		RouterGroup: parent.Group(name),
	}
	s.Use(middleware.MustBeManager)
	s.GET("list", s.getUsers)
	s.POST("create", s.createUser)
	s.POST("update", s.updateUser)
	s.GET("delete", s.deleteUser)
	s.POST("wedding/create", s.createWedding)
	s.GET("wedding/list", s.getWeddings)
	s.GET("wedding/detail", s.getWedding)
	return &s
}

func (s *ManagerServer) getUsers(c *gin.Context) {
	var users, err = user.GetManagers()
	web.AssertNil(err)
	s.SendData(c, users)
}

func (s *ManagerServer) createUser(c *gin.Context) {
	var user *user.User
	user.Role = "manager"
	web.AssertNil(c.BindJSON(&user))
	web.AssertNil(user.Create())
	s.SendData(c, user)
}

func (s *ManagerServer) updateUser(c *gin.Context) {
	var user *user.User
	web.AssertNil(c.BindJSON(&user))
	web.AssertNil(user.Create())
	s.SendData(c, user)
}

func (s *ManagerServer) deleteUser(c *gin.Context) {
	var id = c.Query("id")
	web.AssertNil(user.DeleteUserByID(id))
	s.Success(c)
}
