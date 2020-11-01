package db

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var conn *gorm.DB

func init() {
	viper.SetDefault("db.mysql.maxidleconnections", 5)
	viper.SetDefault("db.mysql.maxopenconnections", 20)
	viper.SetDefault("db.mysql.host", "127.0.0.1")
	viper.SetDefault("db.mysql.port", 3306)
	viper.SetDefault("db.mysql.user", "")
	viper.SetDefault("db.mysql.name", "")
	viper.SetDefault("db.mysql.pass", "")
}

func Open() *gorm.DB {

	if conn != nil {
		return conn
	}

	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("db.mysql.user"),
		viper.GetString("db.mysql.pass"),
		viper.GetString("db.mysql.host"),
		viper.GetInt("db.mysql.port"),
		viper.GetString("db.mysql.name"),
	)

	conn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	conn = conn.Set("gorm:auto_preload", true)

	if dbh, err := conn.DB(); err != nil {
		panic(err)
	} else {
		dbh.SetMaxIdleConns(viper.GetInt("db.mysql.maxidleconnections"))
		dbh.SetMaxOpenConns(viper.GetInt("db.mysql.maxopenconnections"))
	}

	return conn
}

//func Close() {
//if conn != nil {
//conn.Close()
//}
//}
