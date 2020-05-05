package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default() //创建路由
	// ▼解析模板方法1：这样解析之后 templates 之内的 *.tmpl 才可被调用
	// router.LoadHTMLFiles("templates/index.tmpl") //接收多个参数
	// ▼解析模板方法2：常用参数：templates/*	templates/**/*
	router.LoadHTMLGlob("templates/**/*")

	// ▼如果用户通过get请求到/，用2参的函数返回内容
	router.GET("/", func(ctx *gin.Context) {
		// ctx.JSON(200, gin.H{ "message": "欢迎" })
		// ctx.String(http.StatusOK, "<h1>Hello world</h1>")
		ctx.HTML(http.StatusOK, "index01", gin.H{
			"title": "这是首页",
		})
	})
	router.Run(":8099") // 监听并在 0.0.0.0:8080 上启动服务
}

