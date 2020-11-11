package db

import (
	"fmt"

	"github.com/dwburke/go-tools"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

var Conn *gorm.DB

func init() {
	viper.SetDefault("db.postgres.maxidleconnections", 2)
	viper.SetDefault("db.postgres.maxopenconnections", 12)
	viper.SetDefault("db.postgres.port", 5432)
	viper.SetDefault("db.postgres.user", "")
	viper.SetDefault("db.postgres.name", "")
	viper.SetDefault("db.postgres.pass", "")
}

func Open() *gorm.DB {

	if Conn != nil {
		return Conn
	}

	var err error

	connStr := PgConnectString()

	Conn, err = gorm.Open("postgres", connStr)
	tools.FatalError(err)

	Conn = Conn.Set("gorm:auto_preload", true)

	Conn.DB().SetMaxIdleConns(viper.GetInt("db.postgres.maxidleconnections"))
	Conn.DB().SetMaxOpenConns(viper.GetInt("db.postgres.maxopenconnections"))

	return Conn
}

func PgConnectString() string {
	//connStr := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
	connStr := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		viper.GetString("db.postgres.host"),
		viper.GetInt("db.postgres.port"),
		viper.GetString("db.postgres.user"),
		viper.GetString("db.postgres.name"),
		//viper.GetString("db.postgres.pass"),
	)

	return connStr
}
