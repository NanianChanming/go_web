package model

import "time"

type HrUserWorkEmail struct {
	id            int64
	UserCode      string
	UserWorkEmail string
	MailName      string
	SuffixNum     int8
	InUse         bool
	CreateTime    time.Time
	CreateUserId  string
	UpdateTime    time.Time
	UpdateUserId  string
}

func (HrUserWorkEmail) tableName() string {
	return "hr_user_work_email"
}
