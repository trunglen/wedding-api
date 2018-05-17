package manager

import (
	"g/x/web"
	"github.com/gin-gonic/gin"
	"wedding-api/middleware"
	"wedding-api/o/auth"
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
	s.GET("wedding/warning_move/list", s.getWedding)
	s.GET("wedding/warning_missing/list", s.getWedding)
	return &s
}

func (s *ManagerServer) getUsers(c *gin.Context) {
	var au, err = auth.GetByID(web.GetToken(c.Request))
	web.AssertNil(err)
	var users []*user.User
	if au.Role == "super-admin" {
		users, err = user.GetManagers()
	} else {
		users, err = user.GetUsersByRole("manager", au.UserID, au.Role)
	}
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
