package Router

import (
	"net/http"

	"gin/Controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(gin.DebugMode)

	router := gin.Default()
	// 要在路由组之前全局使用「跨域中间件」, 否则OPTIONS会返回404
	// router.Use(Middlewares.Cors())

	v1 := router.Group("v1")
	{
		v1.POST("/testinsert", Controllers.TestInsert)
		v1.GET("/testSelect", Controllers.TestSelect)
	}

	router.GET("/hello-world", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hello test completed successfully"})
	})

	return router
}
