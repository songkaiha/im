package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	models "github.com/songkaiha/im/model"
	"net/http"
	"time"
)

func Conn(ctx *gin.Context) {
	c, err := models.Upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		logrus.Error("client false", err)
		return
	}
	clientId := fmt.Sprintf("%v", &c)
	models.AllClients[clientId] = c
	logrus.Println("连接成功，clientid：" + clientId)
	logrus.Printf("当前终端设备数 %v", len(models.AllClients))
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			delete(models.AllClients, clientId)
			logrus.Printf("clientid %v 已下线，当前在线人数 %v", &c, len(models.AllClients))
			//logrus.Println("read:", err)
			break
		}
		logrus.Printf("mt: %v recv: %s", mt, message)
	}
}

type SendParams struct {
	Clientid string `form:"clientid" json:"clientid" binding:"required"`
	Msg      string `form:"msg" json:"msg" binding:"required"`
}

func Send(ctx *gin.Context) {
	CodeError := ""
	var (
		p SendParams
	)
	if err := ctx.Bind(&p); err != nil {
		CodeError = "必要参数不能为空" + err.Error()
		ctx.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "error": CodeError})
		return
	}
	if c, ok := models.AllClients[p.Clientid]; !ok {
		CodeError = "终端不存在或已下线"
		ctx.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "error": CodeError})
		return
	} else {
		err := c.WriteMessage(websocket.TextMessage, []byte(p.Msg))
		if err != nil {
			CodeError = "数据写入错误" + err.Error()
			ctx.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "error": CodeError})
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "error": CodeError})
}

type IsOnlineParams struct {
	Clientid string `form:"clientid" json:"clientid" binding:"required"`
}

func IsOnline(ctx *gin.Context) {
	CodeError := ""
	var (
		p IsOnlineParams
	)
	if err := ctx.Bind(&p); err != nil {
		CodeError = "必要参数不能为空" + err.Error()
		ctx.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "error": CodeError})
		return
	}
	if c, ok := models.AllClients[p.Clientid]; !ok {
		CodeError = "终端不存在或已下线"
		ctx.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "error": CodeError})
		return
	} else {
		err := c.WriteControl(websocket.PingMessage, []byte("hello"), time.Now().Add(time.Second*3))
		if err != nil {
			CodeError = "终端已下线"
			delete(models.AllClients, p.Clientid)
			ctx.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "error": CodeError})
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "error": CodeError})
}
