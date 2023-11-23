package go_gorm

import (
	log "github.com/cihub/seelog"
	"github.com/xuri/excelize/v2"
	"go_web/src/main/model"
	"regexp"
	"strings"
	"time"
)

func InitWorkEmail() {
	emails := parseEmailExcel()
	for index, email := range emails {
		split := strings.Split(email.UserWorkEmail, "@")
		emails[index].MailName = split[0]
		/*length := len(mailName)
		if isNumber(mailName) || !containsDigit(mailName) {
			continue
		}
		var index = 0
		for index = length - 1; index >= 0; index-- {
			if !isNumber(string(mailName[index])) {
				break
			}
		}
		strNum := mailName[index+1 : length-1]
		atoi, err := strconv.Atoi(strNum)
		if err != nil {
			log.Errorf("数字转换错误，user_code = %s ", email.UserCode)
			continue
		}
		email.SuffixNum = int8(atoi)*/
	}

	log.Info("excel parse completed")
	GormDB.CreateInBatches(emails, 300)
	log.Info("insert db completed")
}

func containsDigit(s string) bool {
	for _, char := range s {
		if char >= '0' && char <= '9' {
			return true
		}
	}
	return false
}

func isNumber(character string) bool {
	numericRegex := regexp.MustCompile(`^\d$`)
	return numericRegex.MatchString(character)
}

func parseEmailExcel() []model.HrUserWorkEmail {
	defer log.Flush()
	path := "C:/Users/Administrator/Desktop/xxx/xxxx.xlsx"
	file, err := excelize.OpenFile(path)
	if err != nil {
		log.Error(err.Error())
		return nil
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Error("文件关闭失败")
		}
	}()
	sheetList := file.GetSheetList()
	file.RemoveRow(sheetList[0], 1)
	rows, err := file.Rows(sheetList[0])
	currentRow := 0
	var emails []model.HrUserWorkEmail
	for rows.Next() {
		currentRow++
		columns, err := rows.Columns()
		if err != nil {
			log.Errorf("第 %d 行读取失败,", currentRow)
			continue
		}
		wxUserId := columns[1]
		user := model.MdmUser{}
		GormDB.Model(&user).Select("user_code, wx_user_id").Where("wx_user_id = (?)", wxUserId).First(&user)
		if user.UserCode == "" {
			log.Errorf("第 %d 行查询员工编码失败, wx_user_id = %s ", currentRow, wxUserId)
			continue
		}
		userWorkEmail := model.HrUserWorkEmail{
			UserCode:      user.UserCode,
			UserWorkEmail: columns[3],
			InUse:         true,
			CreateTime:    time.Now(),
			UpdateTime:    time.Now(),
		}
		emails = append(emails, userWorkEmail)
	}
	return emails
}
