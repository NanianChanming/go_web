package go_excel

import (
	"fmt"
	log "github.com/cihub/seelog"
	"github.com/xuri/excelize/v2"
	"strings"
)

func init() {
	defer log.Flush()
}

var filePath = "C:/Users/Administrator/Desktop/export_result.xlsx"

func ParseExcelExport() {
	defer log.Flush()
	file, err := excelize.OpenFile(filePath)
	if err != nil {
		log.Error("Fatal error ", err.Error())
		return
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Error("文件关闭失败")
		}
	}()
	list := file.GetSheetList()
	if err != nil {
		log.Error("Fatal error ", err.Error())
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
			log.Error("Fatal error ", err.Error())
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
	log.Info("-- 解析完成, 输出sql --")
	log.Info(sql)
}

func ParseExcelImport(filePath string) {
	defer log.Flush()
	file, err := excelize.OpenFile(filePath)
	if err != nil {
		log.Error("Fatal error ", err.Error())
		return
	}
	list := file.GetSheetList()
	file.RemoveRow(list[0], 1)
	rows, err := file.Rows(list[0])
	if err != nil {
		log.Error("Fatal error ", err.Error())
		return
	}

	current := 2
	builder := strings.Builder{}
	s := strings.Builder{}
	for rows.Next() {
		row, err := rows.Columns()
		if err != nil {
			fmt.Println("Fatal error ", err.Error())
			return
		}
		if len(row) < 11 || row[1] == "" || row[9] == "" {
			log.Warn("current row = ", current, "UserCode = ", row[0])
			continue
		}
		builder.WriteString("update hr_user_edu_ex set edu_cert_num = '")
		builder.WriteString(row[10])
		builder.WriteString("' where edu_ex_id = '")
		builder.WriteString(row[1])
		builder.WriteString("'; \n")
		s.WriteString("'")
		s.WriteString(row[1])
		s.WriteString("',")
		current++
		log.Info("current row = ", current)
	}
	log.Info(strings.ReplaceAll(builder.String(), "'`", "'"))
	log.Info(s.String())
}

func ParseExcelImportEntryDate(filePath string) {
	defer log.Flush()
	file, err := excelize.OpenFile(filePath)
	if err != nil {
		log.Error("Fatal error ", err.Error())
		return
	}
	list := file.GetSheetList()
	file.RemoveRow(list[0], 1)
	rows, err := file.Rows(list[0])
	if err != nil {
		log.Error("Fatal error ", err.Error())
		return
	}
	builder := strings.Builder{}
	for rows.Next() {
		row, err := rows.Columns()
		if err != nil {
			fmt.Println("Fatal error ", err.Error())
			return
		}
		if row[0] == "" || row[3] == "" {
			continue
		}
		builder.WriteString("update mdm_user_detail set entry_date = '")
		//builder.WriteString(row[3])
		date := row[3]
		date = strings.ReplaceAll(date, "年", "-")
		date = strings.ReplaceAll(date, "月", "-")
		date = strings.ReplaceAll(date, "日", "")
		builder.WriteString(date)
		builder.WriteString("' where user_id = (select user_id from mdm_user where user_code = '")
		builder.WriteString(row[0])
		builder.WriteString("');\n")
	}
	log.Info(builder.String())
}
