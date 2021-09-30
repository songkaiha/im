package routers

import (
	"github.com/gin-gonic/gin"
	interfaces "github.com/songkaiha/im/interface"
	models "github.com/songkaiha/im/model"
	"net/http"
	"os"
)

func v1init(App *gin.Engine) {
	v1Cmd()
	App.GET("/hello", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": "hello"})
	})

	//v1Router := App.Group("/v1", interfaces.Validator)

}

//可分离项目做单台部署，若维护麻烦，多台部署需加锁，保证单一执行
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
	models.Scheduler.Start()
}
