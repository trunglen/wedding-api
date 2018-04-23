package auth

import (
	"g/x/web"
	"github.com/gin-gonic/gin"
	"wedding-api/o/auth"
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
	s.POST("login", s.login)
	s.POST("avatar", s.uploadAvatar)
	s.POST("portrait", s.uploadPortrait)
	return &s
}

func (s *StudentServer) login(c *gin.Context) {
	var loginInfo = struct {
		Phone    string
		Password string
	}{}
	c.BindJSON(&loginInfo)
	user, err := user.LoginByStudent(loginInfo.Phone, loginInfo.Password)
	web.AssertNil(err)
	var auth = auth.Create(user.ID, user.Role)
	s.SendData(c, map[string]interface{}{
		"token":     auth.ID,
		"user_info": user,
	})
}

func (s *StudentServer) uploadAvatar(c *gin.Context) {
	var userID = c.Query("user_id")
	var file, err = c.FormFile("avatar")
	web.AssertNil(err)
	web.AssertNil(c.SaveUploadedFile(file, "./upload/student/avatar/"+userID))
	s.Success(c)
}
func (s *StudentServer) uploadPortrait(c *gin.Context) {
	var userID = c.Query("user_id")
	var file, err = c.FormFile("avatar")
	web.AssertNil(err)
	web.AssertNil(c.SaveUploadedFile(file, "./upload/student/portrait/"+userID))
	s.Success(c)
}
