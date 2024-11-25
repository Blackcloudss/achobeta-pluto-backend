package model

import "time"

type Sign struct {
	CommonModel
	Issuer     string    `gorm:"column:issuer;type:char(19);unique;not null;comment:'签发标识'"`
	UserId     int64     `gorm:"column:user_id;type:bigint;not null;comment:'用户id'"`
	OnlineTime time.Time `gorm:"column:online_time;not null;comment:'上线时间'"`
	UserAgent  string    `gorm:"column:user_agent;type:varchar(255);not null;comment:'用户代理'"`
	IP         string    `gorm:"column:ip;type:VARCHAR(45);not null;comment:'ip地址'"`
	DeviceName string    `gorm:"column:device_name;type:VARCHAR(50);comment:'设备名称''"`
}

func (t *Sign) TableName() string {
	return "sign"
}
