package main

import (
	"ecommerce/internal/initialize"

	_ "ecommerce/cmd/swag/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           API Documentation Ecommerce Backend SHOPDEV
// @version         1.0.0
// @description     This is a sample server celler server.
// @termsOfService  https://github.com/anbui0211/go-ecommerce-backend-api

// @contact.name   ANBUI
// @contact.url    https://github.com/anbui0211/go-ecommerce-backend-api
// @contact.email  anbui021100@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html2

// @host      localhost:8002
// @BasePath  /v1/2024
// @schema http
func main() {
	r := initialize.Run()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8002")
}
