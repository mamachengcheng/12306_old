package utils

import (
	"github.com/mamachengcheng/12306/app/static"
	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	MysqlDB    *gorm.DB
	MysqlDBErr error
)

func init() {
	cfg, _ := ini.Load(static.ConfFilePath)

	mysqlCfg := cfg.Section("mysql")
	username := mysqlCfg.Key("username").String()
	password := mysqlCfg.Key("password").String()
	address := mysqlCfg.Key("host").String() + ":" + mysqlCfg.Key("port").String()
	dbname := mysqlCfg.Key("dbname").String()

	dsn := username + ":" + password + "@tcp(" + address + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	MysqlDB, MysqlDBErr = gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
