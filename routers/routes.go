package routers

import (
	"github.com/battlesrv/p42/controllers/minecraft"
	"github.com/battlesrv/p42/controllers/steam"
	"github.com/battlesrv/p42/middleware"

	"github.com/gin-gonic/gin"
)

// Init ..
func Init() (r *gin.Engine) {
	r = gin.Default()
	apiv1 := r.Group("/v1")

	gSteamAPIv1 := apiv1.Group("/steam").Use(middleware.CheckCredits())
	{
		gSteamAPIv1.POST("/info", steam.Info)
		gSteamAPIv1.POST("/players", steam.Players)
		gSteamAPIv1.POST("/rules", steam.Rules)
	}

	gMinecraftAPIv1 := apiv1.Group("/minecraft").Use(middleware.CheckCredits())
	{
		gMinecraftAPIv1.POST("/fullstat", minecraft.FullStat)
	}

	return
}
