package dao

import (
	"apiserverv2/setting"
	"fmt"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func InitMysql(conf *setting.MySQLConfig) (err error) {
	// user:password@tcp(host:port)/dbname
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.User, conf.Password, conf.Host, conf.Port, conf.DB)
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return err
	}
	db.SetMaxOpenConns(conf.MaxOpenConns)
	db.SetMaxIdleConns(conf.MaxIdleConns)
	return err
}

// Close 关闭MySQL连接
func Close() {
	_ = db.Close()
}
