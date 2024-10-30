package initialize

import (
	"ecommerce/global"
	"ecommerce/internal/routers"

	"github.com/gin-gonic/gin"
)

func InitRouter() (r *gin.Engine) {
	// Check if development
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New() // gin.New() sẽ không ghi nhật kí
	}

	// Middleware
	r.Use() // logging
	r.Use() // cross
	r.Use() // limiter global

	userRouter := routers.RouterGroupApp.User
	manageRouter := routers.RouterGroupApp.Manage

	mainGroup := r.Group("/v1/2024")
	{
		mainGroup.GET("/checkStatus", func(ctx *gin.Context) { ctx.JSON(200, "OK") }) // tracking monitor	status
	}
	{
		userRouter.InitUserRouter(mainGroup)
		userRouter.InitProductRouter(mainGroup)
	}
	{
		manageRouter.InitUserRouter(mainGroup)
		manageRouter.InitAdminRouter(mainGroup)
	}

	return
}
