package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-ecommerce-backend-api/global"
	"go-ecommerce-backend-api/pkg/i18n"
)

func Run() *gin.Engine {
	// LoadConfig
	LoadConfig("./config/config.yaml", "./.env")
	fmt.Println("Loading configuration mysql", global.Config.Mysql.Username)
	InitLogger()
	InitMysql()
	InitRedis()
	i18n.Init()

	r := InitRouter()
	return r
	//r.Run(":8002")
}
