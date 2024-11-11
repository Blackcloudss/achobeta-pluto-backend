package handler

import (
	"gorm.io/gorm"
	"reflect"
	"tgwp/log/zlog"
	"tgwp/util/snowflake"
)

// 默认节点
const DefaultNodeID = 1

// RegisterHooks 注册 gorm 钩子
func RegisterHooks(db *gorm.DB) {
	db.Callback().Create().Before("gorm:Create").Register("before_create_hook", BeforeCreateHook)
}

// 创建对象之前的钩子函数
func BeforeCreateHook(tx *gorm.DB) {
	node, err := snowflake.NewNode(DefaultNodeID)
	if err != nil {
		zlog.Fatalf("生成 Node 出错")
	}

	tx.Statement.ReflectValue = reflect.Indirect(tx.Statement.ReflectValue)
	// 确认 ID 字段是否存在，并且是 string 类型
	if field := tx.Statement.ReflectValue.FieldByName("ID"); field.IsValid() && field.CanSet() {
		// 设置生成的唯一 ID
		field.SetString(node.Generate().String())
	}
}
