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
		v1.POST("/testinsert", new(Controllers.Test).TestInsert)
		v1.GET("/testSelect", new(Controllers.Test).TestSelect)
	}

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "111"})
	})
	router.GET("/go-famey", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "go-famey 111"})
	})
	router.GET("/go-famey/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "go-famey test 111"})
	})
	router.GET("/hello-world", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hello test completed successfully"})
	})

	return router
}
