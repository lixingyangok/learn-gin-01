package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default() //创建路由
	router.LoadHTMLFiles("templates/index.tmpl")
	// ▼如果用户通过get请求到/，按用2参的函数返回内容
	router.GET("/", func(ctx *gin.Context) {
		// ctx.JSON(200, gin.H{ "message": "欢迎" })
		// ctx.String(http.StatusOK, "<h1>Hello world</h1>")
		ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "这是首页",
		})
	})
	router.Run(":8099") // 监听并在 0.0.0.0:8080 上启动服务
}
