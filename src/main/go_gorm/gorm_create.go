package go_gorm

import (
	"fmt"
	"go_web/src/main/model"
	"go_web/src/main/utils"
	"gorm.io/gorm/clause"
	"net/http"
	"strconv"
	"strings"
)

func init() {
	http.HandleFunc("/createUser", CreateUser)
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
