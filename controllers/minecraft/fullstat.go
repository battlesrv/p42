package minecraft

import (
	"time"

	"github.com/battlesrv/p42/controllers"

	"github.com/battlesrv/go-gsstat/minecraft"
	"github.com/gin-gonic/gin"
)

// FullStat ..
func FullStat(c *gin.Context) {
	var req controllers.RequestStat
	if err := controllers.ReadRequest(c, &req); err != nil {
		controllers.ResponseError(c, err)
		return
	}

	if minecraftFullStat, err := minecraft.GetStats(req.Addr, time.Second*5); err != nil {
		controllers.ResponseError(c, err)
	} else {
		controllers.ResponseSuccess(c, minecraftFullStat)
	}
	return
}
