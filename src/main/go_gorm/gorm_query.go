package go_gorm

import (
	"encoding/json"
	"fmt"
	log "github.com/cihub/seelog"
	"go_web/src/main/model"
	"net/http"
)

func init() {
	http.HandleFunc("/queryOneUser", queryOneUser)
	http.HandleFunc("/queryUserCode", queryUserCode)
	http.HandleFunc("/subQuery", subQuery)
	http.HandleFunc("/fromSbuQuery", fromSbuQuery)
	http.HandleFunc("/groupCondition", groupCondition)
}

/*
Gorm提供了First、Tak、Last方法, 以便从数据库中检索单个对象，当查询数据库时它添加了limit 1条件，
且没有找到记录时，它返回ErrRecordNotFound错误
·三个方法均可以添加条件
*/
func queryOneUser(w http.ResponseWriter, r *http.Request) {
	defer log.Flush()
	user := new(model.MdmUser)
	// 获取第一条记录，主键升序
	GormDB.First(&user)
	log.Info(user)

	// 获取一条记录，没有指定排序字段
	GormDB.Take(&user, "user_code = ?", "YL900097")
	log.Info(user)

	//获取最后一条记录，主键降序
	GormDB.Last(&user)
	log.Info(user)
}

/*
高级查询
gorm允许通过Select方法选择特定字段，如果经常使用此功能，可以定义一个小的结构体来实现调用Api时自动选择特定字段
*/
func queryUserCode(w http.ResponseWriter, r *http.Request) {
	defer log.Flush()
	var user model.MdmUser
	var codes []model.UserCode
	GormDB.Model(&user).Limit(10).Find(&codes)
	log.Info(codes)
	bytes, _ := json.Marshal(codes)
	fmt.Fprintf(w, string(bytes))
}

/*
子查询
子查询可以嵌套在查询中，gorm允许在使用 *gorm.DB对象 作为参数时生成子查询
*/
func subQuery(w http.ResponseWriter, r *http.Request) {
	defer log.Flush()
	var user = model.MdmUser{}
	GormDB.Where("user_code = (?)", GormDB.Table("mdm_user").Select("user_code").Where("user_id = ?", "USER649641597029322752")).Find(&user)
	bytes, _ := json.Marshal(user)
	log.Info(string(bytes))

	subQuery := GormDB.Table("mdm_user").Select("user_code").Where("user_id = ?", "USER1463160424244613120")
	GormDB.Select("*").Where("user_code = (?)", subQuery).Find(&user)
	bytes, _ = json.Marshal(user)
	log.Info(string(bytes))
}

/*
from 子查询
gorm允许在table方法中通过from子句使用子查询
*/
func fromSbuQuery(w http.ResponseWriter, r *http.Request) {
	defer log.Flush()
	var users []model.MdmUser
	var codes = []string{"YL900098", "YL063851"}
	GormDB.Table("(?) as u", GormDB.Model(&users).Select("user_name, user_code").Where("user_code in (?)", codes)).Find(&users)

	subQuery := GormDB.Model(&users).Select("user_code").Where("user_code in (?)", codes)
	var userCodes []model.UserCode
	GormDB.Table("(?) as c", subQuery).Find(&userCodes)
	log.Info(userCodes)
}

/*
Group 条件
非sql中group关键字，是gorm中查询条件分组
使用Group条件可以很轻松的编写复杂sql
*/
func groupCondition(w http.ResponseWriter, r *http.Request) {
	defer log.Flush()
	var users []model.MdmUser
	GormDB.Table("mdm_user").Where(
		GormDB.Where(GormDB.Where("user_name = ?", "周曼雪").Or("user_code = ?", "YL063851")),
	).Or(
		GormDB.Where("user_name = ? or user_code = ?", "容杰", "YL063848").Or("user_name = ?", "徐英瀚"),
	).Find(&users)
	log.Info(users)
}
