package utils

import (
	"fmt"
	"log"
	"os"
	"ruoyi-go/config"
	"ruoyi-go/utils/R"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var once = sync.Once{}

type connect struct {
	client *gorm.DB
}

// 设置一个常量
var _connect *connect

// 在上一篇文章写过https://blog.csdn.net/bei_FengBoby/article/details/124736603?spm=1001.2014.3001.5501 查询，直接拿过用
func connectMysql() {
	//启用打印日志
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level: Silent、Error、Warn、Info
			Colorful:      false,       // 禁用彩色打印
		},
	)
	//换成自己的
	dsn := config.Database.UserName + ":" + config.Database.Password + "@tcp(" + config.Database.Host + ":" + config.Database.Port + ")/" + config.Database.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		fmt.Println("数据库连接失败")
		panic(R.ReturnFailMsg("数据库连接失败"))
		return
	}

	_connect = &connect{
		client: db,
	}

}

func Client() *gorm.DB {
	if _connect == nil {
		once.Do(func() {
			connectMysql()
		})
	}
	return _connect.client
}
