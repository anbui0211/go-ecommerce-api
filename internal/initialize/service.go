package initialize

import (
	"ecommerce/global"
	"ecommerce/internal/database"
	"ecommerce/internal/service"
	"ecommerce/internal/service/impl"
)

func InitServiceInterface() {
	queries := database.New(global.Mdbc)
	// User service interface
	service.InitUserLogin(impl.NewUserLoginImpl(queries))
}
