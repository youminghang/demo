package initilize

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"github.com/youminghang/go-gin-example/pkg/setting"
)

func InitDB() {
	// dsn := "root:123@tcp(192.168.0.36:3306)/mxshop_goods_srv?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/mxshop_goods_srv?charset=utf8mb4&parseTime=True&loc=Local",
		setting.ServerConfig.MysqlInfo.User, setting.ServerConfig.MysqlInfo.Password, setting.ServerConfig.MysqlInfo.Host, setting.ServerConfig.MysqlInfo.Port)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, // 慢SQL阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // 禁用彩色打印
		},
	)
	// 全局模式
	var err error
	setting.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
}
