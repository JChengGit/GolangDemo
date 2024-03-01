package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 创建一个默认的 Gin 引擎
	router := gin.Default()

	// 定义一个 GET 请求的路由，当访问根路径时返回 "Hello, World!"
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	// 启动服务，监听端口 8080
	router.Run(":8080")
}
