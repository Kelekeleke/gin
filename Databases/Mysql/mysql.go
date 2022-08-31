package Mysql

import (
	"fmt"
	"gin/pkg/setting"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func init() {
	var (
		err                                  error
		dbType, dbName, user, password, host string
	)

	if lambdaTaskRoot := os.Getenv("LAMBDA_TASK_ROOT"); lambdaTaskRoot != "" {

		dbType = os.Getenv("DB_POTO_TYPE")
		dbName = os.Getenv("DB_POTO_NAME")
		user = os.Getenv("DB_POTO_USER")
		password = os.Getenv("DB_POTO_PASSWORD")
		host = os.Getenv("DB_POTO_HOST")

	} else {
		sec, err := setting.Cfg.GetSection("database")
		if err != nil {
			log.Fatal(2, "Fail to get section 'database': %v", err)
		}

		dbType = sec.Key("TYPE").String()
		dbName = sec.Key("NAME").String()
		user = sec.Key("USER").String()
		password = sec.Key("PASSWORD").String()
		host = sec.Key("HOST").String()
	}

	DB, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}
	if DB.Error != nil {
		fmt.Printf("database error %v", DB.Error)
	}
}
