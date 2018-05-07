package cache

import (
	"g/x/web"
	"github.com/gin-gonic/gin"
	"wedding-api/o/auth"
	"wedding-api/o/user"
	"wedding-api/x/logger"
)

var cacheLog = logger.NewLogger("cache")

func init() {
	loadAuthCache()
	loadStudentCache()
}

var authCache = map[string]*auth.Auth{}
var studentCache = map[string]*user.User{}

func loadAuthCache() {
	var res, err = auth.GetTokens()
	if err == nil && res != nil {
		for _, item := range res {
			authCache[item.ID] = item
		}
	}
	cacheLog.Infof("cache %d token", len(res))
}

func RemoveToken(userID string) {
	for _, item := range authCache {
		if item.UserID == userID {
			delete(authCache, item.ID)
		}
	}
	auth.RemoveToken(userID)
}

func loadStudentCache() {
	var res, err = user.GetUsers("student")
	if err == nil && res != nil {
		for _, item := range res {
			studentCache[item.ID] = item
		}
	}
	cacheLog.Infof("cache %d students", len(res))
}

func MustGetStudent(c *gin.Context) *user.User {
	var token = web.GetToken(c.Request)
	var authen, err = MustGetAuth(token)
	web.AssertNil(err)
	if student, ok := studentCache[authen.UserID]; ok {
		return student
	}
	student, _ := user.GetByID(authen.UserID)
	if student != nil {
		studentCache[student.ID] = student
		return student
	}
	return nil
}
