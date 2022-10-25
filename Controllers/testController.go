package Controllers

import (
	"gin/Services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Test struct{}

func (t *Test) TestInsert(c *gin.Context) {
	var testService Services.Test

	err := c.ShouldBindJSON(&testService)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := testService.Insert()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "Insert() error!",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "success",
		"data":    id,
	})

}

func (t *Test) TestSelect(c *gin.Context) {

	var test Services.Test
	data := test.Select()
	// Redis.RedisCli.Set("poto_test_redis", 2, 20)
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "success2",
		"data": data,
		// "redisData": Redis.RedisCli.Get("poto_test_redis"),
	})
}
