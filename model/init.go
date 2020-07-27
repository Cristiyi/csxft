package model

import (
	"csxft/util"
	"time"

	"github.com/jinzhu/gorm"

	//
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB 数据库链接单例
var DB *gorm.DB

// Database 在中间件中初始化mysql链接
func Database(connString string) {
	db, err := gorm.Open("mysql", connString)
	db.LogMode(true)
	// Error
	if err != nil {
		util.Log().Panic("连接数据库不成功", err)
	}
	//设置连接池
	//空闲
	db.DB().SetMaxIdleConns(2000)
	//打开
	db.DB().SetMaxOpenConns(2000)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)
	//单数表名
	//db.SingularTable(true)
	DB = db
	gorm.DefaultTableNameHandler= func(db *gorm.DB, defaultTableName string) string {
		return "xft_" +defaultTableName
	}


	migration()


}
