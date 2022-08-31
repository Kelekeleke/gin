package Models

import (
	"gin/Databases/Mysql"
)

type Test struct {
	Id      int
	Testcol string `gorm:"column:name"`
}

// 设置表名
func (Test) TableName() string {
	return "app"
}

func (this *Test) Insert() (id int, err error) {
	result := Mysql.DB.Create(&this)
	id = this.Id
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

func (this *Test) Select() interface{} {
	var testArr []Test
	Mysql.DB.Find(&testArr)
	return testArr
}
