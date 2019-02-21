## 
https://gin-gonic.github.io/gin/

## 安装gin
> go get github.com/gin-gonic/gin
> go get gopkg.in/gin-gonic/gin.v1

## 路由  => 采用的路由库是 httprouter
// 创建带有默认中间件的路由: 日志与恢复中间件
router := gin.Default()
//创建不带中间件的路由：
//r := gin.New()

router.GET("/someGet", getting)
router.POST("/somePost", posting)
router.PUT("/somePut", putting)
router.DELETE("/someDelete", deleting)
router.PATCH("/somePatch", patching)
router.HEAD("/someHead", head)
router.OPTIONS("/someOptions", options)

#### 路由参数 => 参数通过Context的Param方法来获取
// 获取路由参数，假设有路由为"/user/:name"
c.Params.ByName("name")

// 获取query参数
c.Query("name")
c.DefaultQuery("name", "Guest")

// 获取表单参数
c.PostForm("name")
c.DefaultPostForm("name")

#### 路由响应
// 返回简单的字符串
c.String(200, "pong")

// 返回JSON数据
c.JSON(200, gin.H{
    "message": "pong",
})

// 重定向
c.Redirect(http.StatusMovedPermanently, "https://google.com")




