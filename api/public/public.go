package public

import (
	"fmt"
	"g/x/web"
	"github.com/gin-gonic/gin"
	"wedding-api/o/push_token"
	"wedding-api/o/wedding"
	"wedding-api/x/fcm"
)

type PublicServer struct {
	*gin.RouterGroup
	web.JsonRender
}

func NewPublicServer(parent *gin.RouterGroup, name string) *PublicServer {
	var s = PublicServer{
		RouterGroup: parent.Group(name),
	}
	s.POST("push_token/create", s.createPush)
	s.GET("push_token/test", s.test)
	s.GET("test", s.testPublic)
	return &s
}

func (s *PublicServer) test(c *gin.Context) {
	var token = c.Query("token")
	var err, str = fcm.SendToOne(token, fcm.FmcMessage{Title: "Hello", Body: "Anh"})
	fmt.Println(err)
	fmt.Println(str)
	s.Success(c)
}

func (s *PublicServer) testPublic(c *gin.Context) {
	var res, err = wedding.GetMisingWarningWedding("usr_CUdK1clg5JWyCC5FbDHD")
	web.AssertNil(err)
	s.SendData(c, res)
}

func (s *PublicServer) createPush(c *gin.Context) {
	var push *push_token.PushToken
	c.BindJSON(&push)
	web.AssertNil(push.Create())
	s.Success(c)
}
