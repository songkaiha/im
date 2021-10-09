package routers

import (
	"github.com/gin-gonic/gin"
	interfaces "github.com/songkaiha/im/interface"
	v1 "github.com/songkaiha/im/interface/v1"
	models "github.com/songkaiha/im/model"
	"net/http"
	"os"
)

func v1init(App *gin.Engine) {
	v1Cmd()
	App.GET("/hello", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": "hello"})
	})
	App.LoadHTMLGlob("./html/*")
	App.GET("client.html", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "client.html", nil)
	})
	v1Router := App.Group("/v1")
	v1Router.GET("conn", v1.Conn)
	v1Router.POST("send", v1.Send)
	v1Router.POST("isOnline", v1.IsOnline)
	v1Router.GET("total", v1.Total)

}

func v1Cmd() {
	//记录下系统参数
	minutes := "10"
	switch os.Getenv("ENV") {
	case "dev":
		minutes = "1"
	case "test":
		minutes = "5"
	case "prod":
		minutes = "10"
	}
	_, _ = models.Scheduler.AddFunc("*/"+minutes+" * * * *", func() {
		interfaces.PrintMemStats()
	})

	//可分离项目做单台部署，若维护麻烦，多台部署需加锁，保证单一执行
	models.Scheduler.Start()
}
