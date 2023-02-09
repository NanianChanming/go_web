package database

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"go_web/src/main/model"
	"net/http"
)

func init() {
	fmt.Println("初始化文件")
	// 注册驱动
	fmt.Println("注册驱动")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// 设置默认数据库
	orm.RegisterDataBase("default", "mysql", "root:1234567890@/mdm_data_20220921")
	// 注册定义的model 可以注册多个
	orm.RegisterModel(new(model.User))
	// 创建table
	orm.RunSyncdb("default", false, true)

}

func GoOrm(w http.ResponseWriter, r *http.Request) {
	ormer := orm.NewOrm()
	// 插入数据
	user := model.User{Id: 1, UserCode: "YL0000000001", UserName: "zhangsansan"}
	insert, _ := ormer.Insert(&user)
	fmt.Printf("UserCode: %d \n", insert)

	// 更新数据
	//user.UserName = "lisisi"
	//num, _ := ormer.Update(&user)
	//fmt.Printf("num: %d", num)

	// 查询数据-1 数据读取后赋值给传入user对象
	err := ormer.Read(&user)
	fmt.Printf("err: %v \n", err)

	// 查询数据-2 返回一个query_setter
	seter := ormer.QueryTable(user)
	seter.Filter("id", 1)

	// 删除数据
	//i, _ := ormer.Delete(&user)
	//fmt.Printf("num: %d", i)
}
