package user

import (
	"ecommerce/internal/controller/account"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (pr *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	// public router
	// Using WIRE
	// userController, _ := wire.InitUserRouterHandler()
	userRouterPublic := router.Group("/user")
	{
		// userRouterPublic.POST("/register", userController.Register)
		userRouterPublic.POST("/register", account.Login.Register)
		userRouterPublic.POST("/verify_account", account.Login.VerifyOTP)
		userRouterPublic.POST("/update_pass_register", account.Login.UpdatePasswordRegister)
		userRouterPublic.POST("/login", account.Login.Login)
	}

	// private router
	userRouterPrivate := router.Group("/user")
	// adminRouterPrivate.Use(limiter())
	// adminRouterPrivate.Use(Authen())
	// adminRouterPrivate.Use(Permission())
	{
		userRouterPrivate.POST("/getInfo")
	}
}
