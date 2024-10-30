package initialize

import (
	"fmt"

	"ecommerce/global"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Run() *gin.Engine{
	// Load configuration
	LoadConfig()
	m := global.Config.Mysql
	fmt.Println("Loading configuration mysql ", m.Username, m.Password)
	InitLogger()
	global.Logger.Info("config log ok !!!", zap.String("oke", "success"))
	InitMySQL()
	InitMySQLC()
	InitServiceInterface()
	InitRedis()
	InitKafka()

	r := InitRouter()
	return r
	// r.Run(":8002")
}
