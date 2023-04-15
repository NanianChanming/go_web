package go_gorm

import (
	"database/sql"
	"encoding/json"
	"fmt"
	log "github.com/cihub/seelog"
	"go_web/src/main/constants"
	"go_web/src/main/model"
	"go_web/src/main/utils"
	"net/http"
	"strconv"
)

func init() {
	http.HandleFunc("/queryOneUser", queryOneUser)
	http.HandleFunc("/queryUserCode", queryUserCode)
	http.HandleFunc("/subQuery", subQuery)
	http.HandleFunc("/fromSbuQuery", fromSbuQuery)
	http.HandleFunc("/groupCondition", groupCondition)
	http.HandleFunc("/inColumns", inColumns)
	http.HandleFunc("/nameArgs", nameArgs)
	http.HandleFunc("/findToMap", findToMap)
	http.HandleFunc("/firstOrInit", firstOrInit)
	http.HandleFunc("/firstOrCreate", firstOrCreate)
	http.HandleFunc("/firstOrCreateAttrs", firstOrCreateAttrs)
	http.HandleFunc("/firstOrCreateAssign", firstOrCreateAssign)
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

/*
带有多个列的in查询
可以使用二维数组
*/
func inColumns(w http.ResponseWriter, r *http.Request) {
	defer log.Flush()
	var users []model.MdmUser
	GormDB.Table("mdm_user").Where("(user_name, user_code) in ?", [][]interface{}{{"周曼雪", "YL063851"}, {"徐英瀚", "YL063849"}}).Find(&users)
	log.Info(users)
}

/*
命名参数
GORM支持sql.NameArg和map[string]interface{}{}形式的命名参数，例如
*/
func nameArgs(w http.ResponseWriter, r *http.Request) {
	defer log.Flush()
	var user []model.MdmUser
	GormDB.Where("user_code = @code or user_name = @name", sql.Named("code", "YL900098"), sql.Named("name", "周曼雪")).Find(&user)
	log.Info(user)

	GormDB.Where("user_code = @code or user_name = @name", map[string]interface{}{"code": "YL063850", "name": "徐英瀚"}).Find(&user)
	log.Info(user)
}

/*
gorm 允许扫描结果至map[string]interface{}或[]map[string]interface{}, 此时别忘了指定Model或Table
*/
func findToMap(w http.ResponseWriter, r *http.Request) {
	defer log.Flush()
	result := map[string]interface{}{}
	GormDB.Model(&model.MdmUser{}).First(&result, "user_code = ?", "YL063850")
	log.Info(result)
}

/*
获取第一条匹配的记录或者根据给定的条件初始化一个实例
仅支持struct和map条件
*/
func firstOrInit(w http.ResponseWriter, r *http.Request) {
	defer log.Flush()
	var user model.MdmUser
	//GormDB.FirstOrInit(&user, model.MdmUser{UserName: "Default", UserCode: "Default"})
	//log.Info(user)

	//GormDB.Where("user_code = ?", "YL063850").FirstOrInit(&user)
	//log.Info(user)

	GormDB.FirstOrInit(&user, map[string]interface{}{"user_name": "Default", "user_code": "Default"})
	log.Info(user)

	// 如果没有找到记录，可以使用包含更多属性的结构体初始化，Attrs不会被用于生成查询sql
	GormDB.Where("user_code = ?", "YLjalkgja").Attrs(model.MdmUser{UserCode: "YL900098", UserName: "华硕（ASUS）DUAL GeForce RTX4070"}).FirstOrInit(&user)
	log.Info(user)

	// 不管是否找到记录，Assign都会将属性赋值给struct，但这些属性不会被用于生成查询sql，也不会被保存到数据库
	GormDB.Where("user_code = ?", "YL900098").Assign("user_code", "华硕（ASUS）DUAL GeForce RTX4070").FirstOrInit(&user)
	log.Info(user)

}

/*
获取匹配的第一条记录或者根据给定条件创建一条新纪录(仅struct，map条件有效), RowsAffected返回创建、更新的记录数
*/
func firstOrCreate(w http.ResponseWriter, r *http.Request) {
	user := model.MdmUser{UserCode: "YL900099"}
	worker, _ := utils.NextWorker()
	// firstOrCreate会被拼接入sql条件
	create := GormDB.Where(&user).FirstOrCreate(&user, model.MdmUser{
		UserId:   string(strconv.AppendInt([]byte("USER"), worker.GetId(), 10)),
		UserCode: GetUserCode(),
		UserName: utils.GenerateName(),
	})
	log.Info("affected = ", create.RowsAffected)
	log.Info(user)
}

/*
如果没有找到记录，可以使用包含更多属性的结构体创建记录，Attrs不会被用于生成查询sql
*/
func firstOrCreateAttrs(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user := model.MdmUser{UserName: utils.GenerateName()}
	tx := GormDB.Where("user_code = ?", r.Form.Get("user_code")).Attrs(model.MdmUser{
		UserId:   string(strconv.AppendInt([]byte(constants.UserIdPrefix), utils.NextId(), constants.BaseInt)),
		UserName: utils.GenerateName()}).FirstOrCreate(&user)
	log.Info("affected = ", tx.RowsAffected)
	log.Info(user)
}

/*
不管是否找到记录，Assign都会将属性赋值给struct，并将结果写入到数据库
*/
func firstOrCreateAssign(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user := model.MdmUser{UserId: r.Form.Get("user_id")}
	tx := GormDB.Where(&user).Assign(model.MdmUser{
		UserName: utils.GenerateName(),
	}).FirstOrCreate(&user)
	log.Info("affected = ", tx.RowsAffected)
	log.Info(user)
}
