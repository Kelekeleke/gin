package famey

import (
	"fmt"
	u "gin/pkg/util"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func init() {
	var (
		err                                  error
		dbType, dbName, user, password, host string
	)

	dbType = u.GetEnv("DB_FAMEY_TYPE")
	dbName = u.GetEnv("DB_FAMEY_NAME")
	user = u.GetEnv("DB_FAMEY_USER")
	password = u.GetEnv("DB_FAMEY_PASSWORD")
	host = u.GetEnv("DB_FAMEY_HOST")

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
