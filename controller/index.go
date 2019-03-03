package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func Index(c *gin.Context) {
	c.String(http.StatusOK, "It works")
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{"ret": 0, "msg": "ok"})
	// type H map[string]interface{}
	//c.JSON(200, map[string]interface{}{"ret": 0, "msg": "pong"})
}

func Welcome(c *gin.Context) {
	name := c.Query("name")
	sex := c.DefaultQuery("sex", "18")
	c.String(http.StatusOK, "My name is %s, sex is %s.", name, sex)
}

func Async(c *gin.Context) {
	var copy = c.Copy() // Notice: copy context
	// async
	go func() {
		time.Sleep(5 * time.Second)
		log.Println("Done! in path" + copy.Request.URL.Path)
	}()
	c.String(http.StatusOK, "ok")
}
