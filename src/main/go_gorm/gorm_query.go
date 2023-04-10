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
