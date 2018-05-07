package user

import (
	"g/x/web"
	"github.com/gin-gonic/gin"
	"wedding-api/middleware"
	"wedding-api/o/auth"
	"wedding-api/o/user"
)

type UserServer struct {
	*gin.RouterGroup
	web.JsonRender
}

func NewUserServer(parent *gin.RouterGroup, name string) *UserServer {
	var s = UserServer{
		RouterGroup: parent.Group(name),
	}
	s.Use(middleware.MustBeManager)
	s.GET("list", s.list)
	s.POST("create", s.create)
	s.GET("delete", s.delete)
	s.GET("get", s.get)
	s.POST("password/change", s.changePassword)
	s.POST("supervisor/update", s.updateSupervisor)
	s.POST("student/update", s.updateStudent)

	return &s
}

func (s *UserServer) changePassword(c *gin.Context) {
	var body = struct {
		UserID      string `json:"user_id"`
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}{}
	web.AssertNil(c.BindJSON(&body))
	web.AssertNil(user.ChangePassword(body.UserID, body.OldPassword, body.NewPassword))
	s.Success(c)
}

func (s *UserServer) list(c *gin.Context) {
	var role = c.Query("role")
	var au, err = auth.GetByID(web.GetToken(c.Request))
	web.AssertNil(err)
	users, err := user.GetUsersByRole(role, au.UserID, au.Role)
	web.AssertNil(err)
	s.SendData(c, users)
}
func (s *UserServer) get(c *gin.Context) {
	var id = c.Query("id")
	var user, err = user.GetByID(id)
	web.AssertNil(err)
	s.SendData(c, user)
}
func (s *UserServer) create(c *gin.Context) {
	var user *user.User
	web.AssertNil(c.BindJSON(&user))
	web.AssertNil(user.Create())
	s.SendData(c, user)
}

func (s *UserServer) delete(c *gin.Context) {
	var id = c.Query("id")
	web.AssertNil(user.DeleteUserByID(id))
	s.Success(c)
}

func (s *UserServer) updateSupervisor(c *gin.Context) {
	var supervisor *user.SupervisorProfile
	web.AssertNil(c.BindJSON(&supervisor))
	web.AssertNil(supervisor.UpdateSupervisor())
	s.Success(c)
}

func (s *UserServer) updateStudent(c *gin.Context) {
	var student *user.StudentProfile
	web.AssertNil(c.BindJSON(&student))
	web.AssertNil(student.UpdateStudent())
	s.Success(c)
}
