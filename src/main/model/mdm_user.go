package model

type MdmUser struct {
	UserId       string
	UserCode     string
	UserName     string
	UserPassword string
	UserNickname string
	UserEmail    string
	UserMobile   string
	UserAvatar   string
	WxAccount    string
	QqAccount    string
	WxUserId     string
	WxOpenId     string
	Gender       string
}

func (MdmUser) TableName() string {
	return "mdm_user"
}

type UserCode struct {
	UserCode string
}

func (UserCode) TableName() string {
	return "mdm_user"
}
