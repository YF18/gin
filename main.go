package main

import (
    "github.com/gin-gonic/gin"
    _ "github.com/go-sql-driver/mysql"
    "net/http"
)

func main() {
    //创建一个路由Handler
    router:=gin.Default()

    // 
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "It works")
		// c.JSON(200, gin.H{
		// 	"message": "It works",
		// })
	})
	//
	router.GET("/welcome", func(c *gin.Context) {
        name := c.Query("name")
        sex := c.DefaultQuery("sex", "10")
        c.String(http.StatusOK, "My name is %s, sex is %s.", name, sex)
    })
    router.Run(":88")
}