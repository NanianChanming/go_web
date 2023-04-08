package text

import (
	"fmt"
	log "github.com/cihub/seelog"
	"github.com/xuri/excelize/v2"
	"net/http"
	"strings"
)

func init() {
	log.Info("-- excel parse load --")
}

func ParseExcel(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Error("不支持的请求方式")
		fmt.Fprintln(w, "method not support")
	}
	r.ParseForm()
	file, err := excelize.OpenFile(r.Form.Get("filePath"))
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		return
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Error("文件关闭失败")
		}
	}()
	list := file.GetSheetList()
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		return
	}
	file.RemoveRow(list[0], 1)
	rows, err := file.Rows(list[0])
	sql := "select * from hr_user_edu_ex where user_id in ( select user_id from mdm_prod.mdm_user where user_code in ("
	var build strings.Builder
	for rows.Next() {
		row, err := rows.Columns()
		if row[1] == "" {
			break
		}
		if err != nil {
			fmt.Println("Fatal error ", err.Error())
			return
		}
		//build.WriteString(sql)
		build.WriteString("'")
		build.WriteString(row[0])
		build.WriteString("'")
		build.WriteString(",")
	}
	build.WriteString("))")
	sql = sql + build.String()
	index := strings.LastIndex(sql, ",)")
	if index > 0 {
		sql = strings.Replace(sql, ",)", ")", index)
	}
	fmt.Fprintln(w, sql)
}
