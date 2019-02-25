package steam

import (
	"time"

	"github.com/battlesrv/p42/controllers"

	"github.com/battlesrv/go-gsstat/steam"
	"github.com/gin-gonic/gin"
)

// Info ..
func Info(c *gin.Context) {
	var req controllers.RequestStat
	if err := controllers.ReadRequest(c, &req); err != nil {
		controllers.ResponseError(c, err)
		return
	}

	if steamInfo, err := steam.GetInfo(req.Addr, time.Second*5); err != nil {
		controllers.ResponseError(c, err)
	} else {
		controllers.ResponseSuccess(c, steamInfo)
	}
	return
}
