package student

import (
	"g/x/web"
	"github.com/gin-gonic/gin"
	"strconv"
	"wedding-api/cache"
	"wedding-api/o/wedding"
)

func (s *StudentServer) getWedding(c *gin.Context) {
	cache.MustGetStudent(c)
	var id = c.Query("id")
	var wed, err = wedding.GetWeddingDetail(id)
	web.AssertNil(err)
	s.SendData(c, wed)
}

func (s *StudentServer) listMissingWedding(c *gin.Context) {
	var user = cache.MustGetStudent(c)
	var page, _ = strconv.ParseInt(c.Query("page"), 10, 32)
	var wed, err = wedding.GetMisingWedding(user.RestaurantID, user.ID, user.Information.Sex, int(page))
	web.AssertNil(err)
	s.SendData(c, wed)
}

func (s *StudentServer) joinWedding(c *gin.Context) {
	var user = cache.MustGetStudent(c)
	var weddingID = c.Query("wedding_id")
	var wed, _ = wedding.GetWedding(weddingID)
	web.AssertNil(wed.AcceptStudent(wedding.Student{
		ID:     user.ID,
		Name:   user.Name,
		Sex:    user.Information.Sex,
		Status: wedding.STATUS_STUDENT_JOIN,
		Phone:  user.Phone,
	}))
	s.SendData(c, wed)
}

func (s *StudentServer) moveToWedding(c *gin.Context) {
	var user = cache.MustGetStudent(c)
	var weddingID = c.Query("wedding_id")
	var wed, _ = wedding.GetWedding(weddingID)
	web.AssertNil(wed.UpdateStudentStatus(user.ToStudent(wedding.STATUS_STUDENT_MOVE), ""))
	s.SendData(c, wed)
}

func (s *StudentServer) finishWedding(c *gin.Context) {
	var user = cache.MustGetStudent(c)
	var body = struct {
		WeddingID  string `json:"wedding_id"`
		VerifyCode string `json:"verify_code"`
	}{}
	web.AssertNil(c.BindJSON(&body))
	var wed, _ = wedding.GetWedding(body.WeddingID)
	web.AssertNil(wed.UpdateStudentStatus(user.ToStudent(wedding.STATUS_STUDENT_FINISH), body.VerifyCode))
	web.AssertNil(user.UpBalance(wed.Price))
	s.SendData(c, wed)
}

func (s *StudentServer) cancelWedding(c *gin.Context) {
	var user = cache.MustGetStudent(c)
	var body = struct {
		WeddingID  string `json:"wedding_id"`
		VerifyCode string `json:"verify_code"`
	}{}
	web.AssertNil(c.BindJSON(&body))
	var wed, _ = wedding.GetWedding(body.WeddingID)
	web.AssertNil(wed.UpdateStudentStatus(user.ToStudent(wedding.STATUS_STUDENT_FINISH), body.VerifyCode))
	web.AssertNil(user.UpBalance(wed.Price))
	s.SendData(c, wed)
}

func (s *StudentServer) listWedding(c *gin.Context) {
	var user = cache.MustGetStudent(c)
	var status = c.Query("status")
	var page, _ = strconv.ParseInt(c.Query("page"), 10, 32)
	var wed, err = wedding.GetMyWeddingByStatus(user.ID, status, int(page))
	web.AssertNil(err)
	s.SendData(c, wed)
}
