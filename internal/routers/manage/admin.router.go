package manage

import (
	"github.com/gin-gonic/gin"
)

type AdminRouter struct{}

func (pr *AdminRouter) InitAdminRouter(router *gin.RouterGroup) {
	// public router
	adminRouterPublic := router.Group("/admin")
	{
		adminRouterPublic.POST("/login")
	}

	// private router
	adminRouterPrivate := router.Group("/admin/user")
	// adminRouterPrivate.Use(limiter())
	// adminRouterPrivate.Use(Authen())
	// adminRouterPrivate.Use(Permission())
	{
		adminRouterPrivate.POST("/active_user/abc")
	}
}
