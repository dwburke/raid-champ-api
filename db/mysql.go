package db

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var conn *gorm.DB

func init() {
	viper.SetDefault("db.conn.maxidleconnections", 5)
	viper.SetDefault("db.conn.maxopenconnections", 20)
	viper.SetDefault("db.conn.host", "127.0.0.1")
	viper.SetDefault("db.conn.port", 3306)
	viper.SetDefault("db.conn.user", "")
	viper.SetDefault("db.conn.name", "")
	viper.SetDefault("db.conn.pass", "")
}

func Open() *gorm.DB {

	if conn != nil {
		return conn
	}

	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%port)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("db.conn.user"),
		viper.GetString("db.conn.pass"),
		viper.GetString("db.conn.host"),
		viper.GetInt("db.conn.port"),
		viper.GetString("db.conn.name"),
	)

	conn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	conn = conn.Set("gorm:auto_preload", true)

	if dbh, err := conn.DB(); err != nil {
		panic(err)
	} else {
		dbh.SetMaxIdleConns(viper.GetInt("db.conn.maxidleconnections"))
		dbh.SetMaxOpenConns(viper.GetInt("db.conn.maxopenconnections"))
	}

	return conn
}

//func Close() {
//if conn != nil {
//conn.Close()
//}
//}
