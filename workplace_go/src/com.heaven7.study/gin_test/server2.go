package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// func3: 处理带参数的path-GET
func func3(c *gin.Context)  {
	// 回复一个200 OK
	// 获取传入的参数
	name := c.Param("name")
	passwd := c.Param("passwd")
	c.String(http.StatusOK, "参数:%s %s  test3 OK", name, passwd)
}
// func4: 处理带参数的path-POST
func func4(c *gin.Context)  {
	// 回复一个200 OK
	// 获取传入的参数
	name := c.Param("name")
	passwd := c.Param("passwd")
	c.String(http.StatusOK, "参数:%s %s  test4 OK", name, passwd)
}
// func5: 注意':'和'*'的区别
func func5(c *gin.Context)  {
	// 回复一个200 OK
	// 获取传入的参数
	name := c.Param("name")
	passwd := c.Param("passwd")
	c.String(http.StatusOK, "参数:%s %s  test5 OK", name, passwd)
}

/**
transport param
 */
func main(){
	router := gin.Default()
	// TODO:注意':'必须要匹配,'*'选择匹配,即存在就匹配,否则可以不考虑
	router.GET("/test3/:name/:passwd", func3)
	router.POST("/test4/:name/:passwd", func4)
	router.GET("/test5/:name/*passwd", func5)

	router.Run(":8888")
}

// 使用gin.Context中的Query方法解析, 使用Query获取参数
func func6(c *gin.Context)  {
	// 回复一个200 OK
	// 获取传入的参数
	name := c.Query("name")
	passwd := c.Query("passwd")
	c.String(http.StatusOK, "参数:%s %s  test6 OK", name, passwd)
}
// 使用Query获取参数
func func7(c *gin.Context)  {
	// 回复一个200 OK
	// 获取传入的参数
	name := c.Query("name")
	passwd := c.Query("passwd")
	c.String(http.StatusOK, "参数:%s %s  test7 OK", name, passwd)
}
