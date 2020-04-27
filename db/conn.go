package db

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"guthub.com/haibin4739/galileo-whistler/conf/setting"
)

var Conn *gorm.DB

func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)

	dbType = setting.Database.Type
	dbName = setting.Database.Name
	user = setting.Database.User
	password = setting.Database.Password
	host = setting.Database.Host
	tablePrefix = setting.Database.TablePrefix

	Conn, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	Conn.SingularTable(true)
	Conn.LogMode(true)
	Conn.DB().SetMaxIdleConns(10)
	Conn.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer Conn.Close()
}
