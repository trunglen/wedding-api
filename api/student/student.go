package student

import (
	"g/x/web"
	"github.com/gin-gonic/gin"
	"wedding-api/cache"
	"wedding-api/o/push_token"
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
	s.POST("profile/update", s.updateProfile)
	s.POST("password/change", s.changePassword)
	s.POST("avatar", s.uploadAvatar)
	s.POST("portrait", s.uploadPortrait)
	s.POST("wedding/join", s.joinWedding)
	s.POST("wedding/move", s.moveToWedding)
	s.POST("wedding/finish", s.finishWedding)
	s.GET("wedding/list/mine", s.listWedding)
	s.GET("wedding/list/missing", s.listMissingWedding)
	s.GET("pust_token/register", s.registerPushToken)
	return &s
}

func (s *StudentServer) login(c *gin.Context) {
	var loginInfo = struct {
		Phone    string
		Password string
	}{}
	web.AssertNil(c.BindJSON(&loginInfo))
	user, err := user.LoginByStudent(loginInfo.Phone, loginInfo.Password)
	web.AssertNil(err)
	var auth = cache.CreateAuth(user.ID, user.Role)
	s.SendData(c, map[string]interface{}{
		"token":     auth.ID,
		"user_info": user,
	})
}

func (s *StudentServer) uploadAvatar(c *gin.Context) {
	var user = cache.MustGetStudent(c)
	var file, err = c.FormFile("avatar")
	web.AssertNil(err)
	web.AssertNil(c.SaveUploadedFile(file, "./upload/student/avatar/"+user.ID))
	s.Success(c)
}
func (s *StudentServer) uploadPortrait(c *gin.Context) {
	var user = cache.MustGetStudent(c)
	var file, err = c.FormFile("avatar")
	web.AssertNil(err)
	web.AssertNil(c.SaveUploadedFile(file, "./upload/student/portrait/"+user.ID))
	s.Success(c)
}

func (s *StudentServer) updateProfile(c *gin.Context) {
	var usr = cache.MustGetStudent(c)
	var u *user.User
	u.ID = usr.ID
	web.AssertNil(c.BindJSON(&u))
	web.AssertNil(u.UpdateProfile())
	s.Success(c)
}

func (s *StudentServer) changePassword(c *gin.Context) {
	var u = cache.MustGetStudent(c)
	var body = struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}{}
	web.AssertNil(c.BindJSON(&body))
	web.AssertNil(user.ChangePassword(u.ID, body.OldPassword, body.NewPassword))
	s.Success(c)
}

func (s *StudentServer) registerPushToken(c *gin.Context) {
	var u = cache.MustGetStudent(c)
	var body = push_token.PushToken{}
	body.UserID = u.ID
	web.AssertNil(c.BindJSON(&body))
	web.AssertNil(body.Create())
	s.Success(c)
}
