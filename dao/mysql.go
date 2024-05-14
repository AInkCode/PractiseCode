package dao

import (
	"PratiseCode/setting"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitMysql(conf *setting.MySQLConfig) (err error) {
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.User, conf.Password, conf.Host, conf.Port, conf.DB)
	// fmt.Printf("%s, %s, %s, %d, %s", conf.User, conf.Password, conf.Host, conf.Port, conf.DB)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.User, conf.Password, conf.Host, conf.Port, conf.DB)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return err
}

func Close() {
	if DB == nil {
		return
	}
	if db, err := DB.DB(); err == nil {
		db.Close()
	}
}

// func Close() {
// 	// 通过 DB.DB() 获取 *sql.DB 对象，并关闭连接
// 	if db, err := DB.DB(); err == nil {
// 		db.Close()
// 	}
// }
