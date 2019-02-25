package steam

import (
	"time"

	"github.com/battlesrv/p42/controllers"

	"github.com/battlesrv/go-gsstat/steam"
	"github.com/gin-gonic/gin"
)

// Players ..
func Players(c *gin.Context) {
	var req controllers.RequestStat
	if err := controllers.ReadRequest(c, &req); err != nil {
		controllers.ResponseError(c, err)
		return
	}

	if steamPlayers, err := steam.GetPlayers(req.Addr, time.Second*5); err != nil {
		controllers.ResponseError(c, err)
	} else {
		controllers.ResponseSuccess(c, steamPlayers)
	}
	return
}
