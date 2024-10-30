package manage

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (pr *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	// private router
	userRouterPrivate := router.Group("/admin/user")
	// adminRouterPrivate.Use(limiter())
	// adminRouterPrivate.Use(Authen())
	// adminRouterPrivate.Use(Permission())
	{
		userRouterPrivate.POST("/active_user")
	}
}
