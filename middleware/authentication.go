package middleware

import (
	"g/x/web"
	"github.com/gin-gonic/gin"
	"wedding-api/o/auth"
)

const errNotPermision = "you are not enough permision"
const errUnauthorize = "you are not enough permision"

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func Authenticate(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var token = web.GetToken(c.Request)
		var res, err = auth.GetByID(token)
		if err == nil && res != nil {
			if !contains(roles, res.Role) {
				c.AbortWithStatusJSON(401, map[string]interface{}{
					"error":  errNotPermision,
					"status": "error",
				})
				return
			}
		} else {
			if err.Error() == "not found" {
				c.AbortWithStatusJSON(401, map[string]interface{}{
					"error":  errUnauthorize,
					"status": "error",
				})
				return
			}
		}
		c.Next()
	}
}

var MustBeAdmin = Authenticate("super-admin")
var MustBeSupervisor = Authenticate("supervisor", "super-admin")
var MustBeManager = Authenticate("manager", "supervisor", "super-admin")
var MustBeStudent = Authenticate("student")
