package Models

import (
	Mysql "gin/Databases/Mysql/Famey"
)

type Config struct {
	Id  int
	Key string `gorm:"column:key"`
}

// 设置表名
func (Config) TableName() string {
	return "config"
}

func (this *Config) Insert() (id int, err error) {
	result := Mysql.DB.Create(&this)
	id = this.Id
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

func (this *Config) Select() []Config {
	var testArr []Config
	Mysql.DB.Find(&testArr)
	return testArr
}
