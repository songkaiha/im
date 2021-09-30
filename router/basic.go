package routers

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

var App *gin.Engine

func init() {
	switch os.Getenv("ENV") {
	case "prod":
		gin.SetMode(gin.ReleaseMode)
		break
	case "dev":
		gin.SetMode(gin.DebugMode)
		break
	case "test":
		gin.SetMode(gin.TestMode)

	default:
		gin.SetMode(gin.DebugMode)
	}

	App = gin.Default()
	App.Use(SetHeader())
	App.Use(gzip.Gzip(gzip.DefaultCompression))

	App.HEAD("/", func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})
	v1init(App)
}

func SetHeader() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS,POST,PUT,DELETE")
		ctx.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization, AccessToken, Token")
		ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		ctx.Header("Access-Control-Allow-Credentials", "true")
		// 放行所有OPTIONS方法，因为有的模板是要请求两次的
		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
		}
		ctx.Next()
	}
}
