package router

import "github.com/gin-gonic/gin"
import . "../controller" // Notice: ".", 点操作 导入包后 调用这个包的函数时 可以省略包名

func Init() *gin.Engine {
	// router
	router := gin.Default()
	// index.go
	router.GET("/", Index)
	router.GET("/index", Index)
	router.GET("/ping", Ping)
	router.GET("/welcome", Welcome)
	router.GET("/async", Async)
	// person.go
	router.POST("/person", AddPerson)
	router.GET("/person/:id", GetPerson)
	router.GET("/person", GetAllPerson)
	router.DELETE("/person/:id", DelPerson)
	router.PUT("/person/:id", UpdatePerson)
	router.StaticFile("/favicon.ico", "./resource/favicon.ico")
	return router
}
