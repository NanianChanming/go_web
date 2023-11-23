package go_gorm

import (
	"fmt"
	"go_web/src/main/deploy"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

var GormDB *gorm.DB

func init() {
	connectMysql()
}

/**
gorm 操作数据库
注意: 想要正确的处理time.Time,需要带上parseTime参数，要支持完整的UTF—8编码，需要将charset=utf8更改为utf8mb4
*/
func connectMysql() {
	dsn := "root:1234567890@tcp(127.0.0.1:3306)/mdm_data_20231009?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// mysql驱动提供了一些高级配置可以在初始化过程中使用，例如：
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,   //string类型字段的默认长度
		DisableDatetimePrecision:  true,  //禁用datetime精度，MySQL5.6之前的数据库不支持
		DontSupportRenameIndex:    true,  //重命名索引时采用删除并新建的方式，5.7之前不支持
		DontSupportRenameColumn:   true,  //用‘change’重命名列，MySQL 8之前的数据库和MariaDB不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前MySQL版本自动配置
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 禁用复数表名
		},
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
			logger.Config{
				SlowThreshold:             time.Second, // 慢 SQL 阈值
				LogLevel:                  logger.Info, // 日志级别
				IgnoreRecordNotFoundError: false,       // 忽略ErrRecordNotFound（记录未找到）错误
				Colorful:                  false,       // 禁用彩色打印
			}),
	})
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		return
	}
	dbPool, err := db.DB()
	dbPool.SetMaxIdleConns(10)
	dbPool.SetMaxOpenConns(20)
	dbPool.SetConnMaxLifetime(time.Hour)
	GormDB = db
}

/*
连接池
*/
func connect() {
	dsn := "root:1234567890@tcp(127.0.0.1:3306)/mdm_data_1124?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		return
	}
	GormDB = db
	deploy.Logger.Info("db connect success")
}
