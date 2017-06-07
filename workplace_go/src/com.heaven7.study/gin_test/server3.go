package main

import (
	"github.com/gin-gonic/gin"
)

// 参数是form中获得,即从Body中获得,忽略URL中的参数
func func8(c *gin.Context)  {
	message := c.PostForm("message")
	extra := c.PostForm("extra")
	nick := c.DefaultPostForm("nick", "anonymous")

	c.JSON(200, gin.H{
		"status":  "test8:posted",
		"message": message,
		"nick":    nick,
		"extra": extra,
	})
}

func main(){
	router := gin.Default()
	// 使用post_form形式,注意必须要设置Post的type,
	// 同时此方法中忽略URL中带的参数,所有的参数需要从Body中获得
	router.POST("/test8", func8)

	router.Run(":8888")
}
