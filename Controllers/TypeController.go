package Controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Type struct{}

func (t *Type) GetType(c *gin.Context) {
	fmt.Println("test")
}
