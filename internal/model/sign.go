package model

type Sign struct {
	CommonModel
	LoginId string `gorm:"column:login_id;type:char(19);not null;comment:'登录id'"`
	Issuer  string `gorm:"column:issuer;type:char(19);not null;comment:'签发标识'"`
	//下面这三个是必须的，我得处理login_id唯一，且对应一台设备
	UserId    string `gorm:"column:user_id;type:char(19);not null;comment:'用户id'"`
	UserAgent string `gorm:"column:user_agent;type:VARCHAR(50);not null;comment:'用户代理'"`
	IP        string `gorm:"column:ip;type:VARCHAR(45);not null;comment:'ip地址'"`
}

func (t *Sign) TableName() string {
	return "sign"
}
