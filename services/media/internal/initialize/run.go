package initialize

import (
	"github.com/gin-gonic/gin"
)

func Run() *gin.Engine {
	// LoadConfig
	LoadConfig("./config/config.yaml", "./.env")
	InitLogger()
	//InitMysql()
	//InitRedis()
	//i18n.Init()
	InitS3()

	r := InitRouter()
	return r
	//r.Run(":8002")
}
