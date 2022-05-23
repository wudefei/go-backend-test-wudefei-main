package db

import (
	"time"

	"user-api/internal/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var dbs map[string]*gorm.DB

type DbConf struct {
	Driver      string        `json:"driver"`
	Uri         string        `json:"uri"`
	Username    string        `json:"user_name"`
	Password    string        `json:"password"`
	MaxOpenConn int           `json:"max_open_conn"`
	MaxIdleConn int           `json:"max_idle_conn"`
	MaxLifetime time.Duration `json:"max_life_time"`
}

func Init() error {
	var err error
	configDb := config.Config.Db
	dbs = make(map[string]*gorm.DB)

	for name, cfg := range configDb {
		dbs[name], err = gorm.Open(cfg.Driver, cfg.Uri)
		if err != nil {
			panic("连接数据库失败:" + err.Error())
		}

		dbs[name].DB().SetMaxIdleConns(cfg.MaxIdleConn)
		dbs[name].DB().SetMaxOpenConns(cfg.MaxOpenConn)
		dbs[name].DB().SetConnMaxLifetime(cfg.MaxLifetime)
		dbs[name].LogMode(config.Config.DbLogEnable)
	}

	return nil
}

func GetGinBase() *gorm.DB {
	return dbs["dp_user"]
}
