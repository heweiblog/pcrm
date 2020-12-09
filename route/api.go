package route

import (
	"github.com/gin-gonic/gin"
	//"io"
	//"os"
	"pcrm/controller"
)

func initEngine() *gin.Engine {
	// gin日志
	//f, _ := os.Create("/tmp/pcrm.log")
	//gin.DefaultWriter = io.MultiWriter(f)

	// 如果需要将日志同时写入文件和控制台，请使用以下代码
	//gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// 引擎
	return gin.Default()
}

func initApiRoute(g *gin.Engine) {
	v1 := g.Group("/api/v1.0/internal")
	{
		v1.GET("/status", controller.Heartbeat)
		v1.POST("/status", controller.ResetMsId)
		v1.GET("/configs", controller.GetConfigs)
		v1.POST("/configs", controller.PostConfig)
		v1.POST("/all-configs", controller.RegisterConf)
		v1.GET("/oplog", controller.GetOplog)
		v1.GET("/tasks", controller.GetTask)
		v1.POST("/tasks", controller.PostTask)
		v1.DELETE("/tasks", controller.DeleteTask)
	}
}

func Server(addr string) {
	g := initEngine()
	initApiRoute(g)
	g.Run(addr)
}
