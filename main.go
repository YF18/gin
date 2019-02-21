package main

import (
 "github.com/gin-gonic/gin"
 "net/http"
)

func main() {
	// 
	router := gin.Default()
	// 
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "It works")
		// c.JSON(200, gin.H{
		// 	"message": "It works",
		// })
	})
	//
	router.GET("/welcome:name", func(c *gin.Context) {
		name := c.Param("name")
        firstname := c.DefaultQuery("firstname", "Guest")
        lastname := c.Query("lastname")

        c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
    })
    // 
    router.GET("/async", func(c *gin.Context) {
	    cCp := c.Copy()
	    go func() {
	        time.Sleep(5 * time.Second)
	        log.Println("Done! in path" + cCp.Request.URL.Path)
	    }()
	})
	//
	router.Run(":8000")
}