package student

import (
	"fmt"
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
	s.GET("logout", s.logout)
	s.POST("profile/update", s.updateProfile)
	s.GET("profile/get", s.getProfile)
	s.POST("password/change", s.changePassword)
	s.POST("avatar", s.uploadAvatar)
	s.POST("portrait", s.uploadPortrait)
	s.GET("wedding/get", s.getWedding)
	s.POST("wedding/join", s.joinWedding)
	s.POST("wedding/move", s.moveToWedding)
	s.POST("wedding/finish", s.finishWedding)
	s.POST("wedding/cancel", s.cancelWedding)
	s.GET("wedding/list/mine", s.listWedding)
	s.GET("wedding/list/missing", s.listMissingWedding)
	s.POST("push_token/register", s.registerPushToken)
	return &s
}

func (s *StudentServer) login(c *gin.Context) {
	var loginInfo = struct {
		Phone    string
		Password string
	}{}
	web.AssertNil(c.BindJSON(&loginInfo))
	user, err := user.LoginByStudent(loginInfo.Phone, loginInfo.Password)
	{
		user.Information.Avatar = "http://" + c.Request.Host + "/static/student/avatar/" + user.ID
		user.Information.Portrait = "http://" + c.Request.Host + "/static/student/portrait/" + user.ID
	}
	web.AssertNil(err)
	cache.RemoveToken(user.ID)
	var auth = cache.CreateAuth(user.ID, user.Role)
	s.SendData(c, map[string]interface{}{
		"token":     auth.ID,
		"user_info": user,
	})
}

func (s *StudentServer) logout(c *gin.Context) {
	var user = cache.MustGetStudent(c)
	cache.RemoveToken(user.ID)
	fmt.Println(c.Request.Host)
	s.Success(c)
}

func (s *StudentServer) uploadAvatar(c *gin.Context) {
	var user = cache.MustGetStudent(c)
	var file, err = c.FormFile("avatar")
	web.AssertNil(err)
	web.AssertNil(c.SaveUploadedFile(file, "./upload/student/avatar/"+user.ID))
	s.SendData(c, map[string]interface{}{
		"url": "http://" + c.Request.Host + "/static/student/avatar/" + user.ID,
	})
}
func (s *StudentServer) uploadPortrait(c *gin.Context) {
	var user = cache.MustGetStudent(c)
	var file, err = c.FormFile("portrait")
	web.AssertNil(err)
	web.AssertNil(c.SaveUploadedFile(file, "./upload/student/portrait/"+user.ID))
	s.SendData(c, map[string]interface{}{
		"url": "http://" + c.Request.Host + "/static/student/portrait/" + user.ID,
	})
}

func (s *StudentServer) updateProfile(c *gin.Context) {
	var usr = cache.MustGetStudent(c)
	var u *user.User
	web.AssertNil(c.BindJSON(&u))
	u.ID = usr.ID
	web.AssertNil(u.UpdateProfile())
	u.SetAvatarAndUpload(c.Request)
	s.SendData(c, u)
}

func (s *StudentServer) getProfile(c *gin.Context) {
	var usr = cache.MustGetStudent(c)
	usr, err := user.GetByID(usr.ID)
	web.AssertNil(err)
	usr.SetAvatarAndUpload(c.Request)
	s.SendData(c, usr)
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
