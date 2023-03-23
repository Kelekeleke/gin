package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gohouse/converter"
)

// 生成表结构用 ./Models/model.go 是初始的 不要动这个
func Converter(c *gin.Context) {
	err := converter.NewTable2Struct().
		SavePath("./Models/tagModel.go").
		Dsn("fbad:password@tcp(alphaonline.cpnupr9lfkqf.us-east-1.rds.amazonaws.com:3306)/poto?charset=utf8").
		TagKey("gorm").
		EnableJsonTag(true).
		PackageName("Models").
		Table("tag").
		Run()
	fmt.Println(err)
}
