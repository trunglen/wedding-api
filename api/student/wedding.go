package student

import (
	"g/x/web"
	"github.com/gin-gonic/gin"
	"strconv"
	"wedding-api/cache"
	"wedding-api/o/wedding"
)

func (s *StudentServer) listMissingWedding(c *gin.Context) {
	var user = cache.MustGetStudent(c)
	var page, _ = strconv.ParseInt(c.Query("page"), 10, 32)
	var wed, err = wedding.GetMisingWedding(user.ID, user.Information.Sex, int(page))
	web.AssertNil(err)
	s.SendData(c, wed)
}
