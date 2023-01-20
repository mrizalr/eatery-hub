package mysql

import (
	"fmt"

	"github.com/mrizalr/eatery-hub/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysqlDB(c *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.Mysql.User,
		c.Mysql.Password,
		c.Mysql.Port,
		c.Mysql.DBname,
	)

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
