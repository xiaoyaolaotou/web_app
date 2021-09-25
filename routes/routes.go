package routes

import (
	"web_app/controllers"
	"web_app/logger"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 注册业务路由
	r.POST("/signup", controllers.SignUpHandler)

	r.GET("/", controllers.QueryBackupHandler)

	r.POST("/v1/backup", controllers.BackupHandler)

	return r
}
