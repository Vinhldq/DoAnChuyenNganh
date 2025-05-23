package initializer

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/global"
	"go.uber.org/zap"
)

func Run() *gin.Engine {
	LoadConfig()
	fmt.Printf("Server: %d. Port: %s", global.Config.Server.Port, global.Config.Server.Mode)
	InitLogger()
	global.Logger.Info("Config", zap.String("ok", "success"))

	InitMysql()
	InitRedis()

	InitInterface()

	InitKafka()

	r := InitRouter()
	return r
}
