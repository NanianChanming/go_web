package go_gorm

import (
	"fmt"
	"go_web/src/main/deploy"
	"go_web/src/main/model"
	"go_web/src/main/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/http"
	"strconv"
	"strings"
)

var GormDB *gorm.DB

func init() {
	connect()
	http.HandleFunc("/createUser", CreateUser)
}

/**
gorm 操作数据库
注意: 想要正确的处理time.Time,需要带上parseTime参数，要支持完整的UTF—8编码，需要将charset=utf8更改为utf8mb4
*/
func connectMysql() {
	dsn := "root:1234567890@tcp(127.0.0.1:3306)/mdm_data_1124?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// mysql驱动提供了一些高级配置可以在初始化过程中使用，例如：
	gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,   //string类型字段的默认长度
		DisableDatetimePrecision:  true,  //禁用datetime精度，MySQL5.6之前的数据库不支持
		DontSupportRenameIndex:    true,  //重命名索引时采用删除并新建的方式，5.7之前不支持
		DontSupportRenameColumn:   true,  //用‘change’重命名列，MySQL 8之前的数据库和MariaDB不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前MySQL版本自动配置
	}), &gorm.Config{})
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

func CreateUser(w http.ResponseWriter, r *http.Request) {
	worker, err := utils.NewWorker(1)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		return
	}
	id := worker.GetId()
	userId := strconv.AppendInt([]byte("USER"), id, 10)
	user := model.MdmUser{
		UserId:   string(userId),
		UserCode: getUserCode(),
		UserName: "张三",
	}
	GormDB.Create(&user)
	fmt.Fprintln(w, user)
}

func getUserCode() string {
	m := new(model.UserCode)
	c := clause.Column{Table: "mdm_user", Name: "user_code"}
	column := clause.OrderByColumn{Column: c, Desc: true}
	GormDB.Model(&model.UserCode{}).Order(column).First(&m)
	userCode := strings.Replace(m.UserCode, "YL", "", 1)
	// ParseInt参数(str, 进制, int类型长度：类似int int64)
	i, _ := strconv.ParseInt(userCode, 10, 64)
	// FormatInt参数（int, 进制）
	code := strconv.AppendInt([]byte("YL"), i+1, 10)
	return string(code)
}
