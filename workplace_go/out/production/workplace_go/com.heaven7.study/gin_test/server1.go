package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


func func1 (c *gin.Context)  {
	// 回复一个200OK,在client的http-get的resp的body中获取数据
	c.String(http.StatusOK, "test1 OK")
}
// func2: 处理最基本的POST
func func2 (c *gin.Context) {
	// 回复一个200 OK, 在client的http-post的resp的body中获取数据
	c.String(http.StatusOK, "test2 OK")
}
func main(){
	// 注册一个默认的路由器
	router := gin.Default()
	// 最基本的用法
	router.GET("/test1", func1)
	router.POST("/test2", func2)
	// 绑定端口是8888
	router.Run(":8888")
}
