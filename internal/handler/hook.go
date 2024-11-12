package handler

import (
	"gorm.io/gorm"
	"reflect"
	"tgwp/log/zlog"
	"tgwp/util/snowflake"
)

// 默认节点 (目前仅这里需要用到)
const Default_User_ID = 2
const Default_Team_ID = 3
const Default_Node_ID = 4

// 注册 gorm 钩子
func RegisterUser(db *gorm.DB) {
	db.Callback().Create().Before("gorm:Create").Register("before_create_User", BeforeCreateUser)
}

// 在操作数据库前 创建 用户 ID
func BeforeCreateUser(tx *gorm.DB) {
	node, err := snowflake.NewNode(Default_User_ID)
	if err != nil {
		zlog.Errorf("生成 Node 出错")
		return
	}

	// 如果是指针，返回指针指向的值；如果是非指针，直接返回
	tx.Statement.ReflectValue = reflect.Indirect(tx.Statement.ReflectValue)

	// 确认 user_id 字段是否存在，并且是 int64 类型
	if field := tx.Statement.ReflectValue.FieldByName("user_id"); field.IsValid() && field.CanSet() {
		// 设置生成的唯一 ID
		field.SetInt(node.Generate().Int64())
	}
}

func RegisterTeam(db *gorm.DB) {
	db.Callback().Create().Before("gorm:Create").Register("before_create_Team", BeforeCreateTeam)
}

// 创建 团队 ID
func BeforeCreateTeam(tx *gorm.DB) {
	node, err := snowflake.NewNode(Default_Team_ID)
	if err != nil {
		zlog.Errorf("生成 Node 出错")
		return
	}
	tx.Statement.ReflectValue = reflect.Indirect(tx.Statement.ReflectValue)
	if field := tx.Statement.ReflectValue.FieldByName("team_id"); field.IsValid() && field.CanSet() {
		field.SetInt(node.Generate().Int64())
	}
}

func RegisterNode(db *gorm.DB) {
	db.Callback().Create().Before("gorm:Create").Register("before_create_Node", BeforeCreateNode)
}

// 创建 节点 ID
func BeforeCreateNode(tx *gorm.DB) {
	node, err := snowflake.NewNode(Default_Node_ID)
	if err != nil {
		zlog.Errorf("生成 Node 出错")
		return
	}
	tx.Statement.ReflectValue = reflect.Indirect(tx.Statement.ReflectValue)
	if field := tx.Statement.ReflectValue.FieldByName("myself_id"); field.IsValid() && field.CanSet() {
		field.SetInt(node.Generate().Int64())
	}
}
