package middleware

import (
	"net/http"
	"time"

	"github.com/battlesrv/p42/common"
	"github.com/battlesrv/p42/db"

	"github.com/gin-gonic/gin"
)

var timeoutBetweenRequests = uint32(1)

// CheckCredits ..
func CheckCredits() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user db.User
		if err := db.Read(c.GetHeader("X-Email"), &user); err != nil {
			unauthorizedStatus(c)
			return
		}

		if common.Sha256Sum(c.GetHeader("X-Token")) != user.TokenHash {
			unauthorizedStatus(c)
			return
		}

		if user.NextAccess != 0 && user.NextAccess > uint32(time.Now().Unix()) {
			limitedAccessStatus(c)
			return
		}

		if user.NextAccess != 0 {
			user.NextAccess = uint32(time.Now().Unix()) + timeoutBetweenRequests
		}

		if err := db.Write(&user, c.GetHeader("X-Email")); err != nil {
			internalServerErrorStatus(c)
			return
		}

		c.Next()
	}
}

func unauthorizedStatus(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"ok":  false,
		"msg": "unauthorized",
	})
	c.Abort()
}

func internalServerErrorStatus(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"ok":  false,
		"msg": "internal server error",
	})
	c.Abort()
}

func limitedAccessStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ok":  false,
		"msg": "limit is reachable, change price",
	})
	c.Abort()
}
