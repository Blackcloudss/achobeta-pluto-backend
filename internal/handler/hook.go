package handler

import (
	"gorm.io/gorm"
	"reflect"
	"tgwp/global"
	"tgwp/log/zlog"
	"tgwp/util/snowflake"
)

// 注册 gorm 钩子
func RegisterHook(db *gorm.DB) {
	db.Callback().Create().Before("gorm:Create").Register("before_create_Node", BeforeCreateNode)
}

// 在操作数据库前 创建 用户 ID
func BeforeCreateNode(db *gorm.DB) {
	node, err := snowflake.NewNode(global.DEFAULT_NODE_ID)
	if err != nil {
		zlog.Errorf("生成 Node 出错")
		return
	}

	// 如果是指针，返回指针指向的值；如果是非指针，直接返回
	db.Statement.ReflectValue = reflect.Indirect(db.Statement.ReflectValue)

	// 确认 id 字段是否存在，并且是 int64 类型
	if field := db.Statement.ReflectValue.FieldByName("id"); field.IsValid() && field.CanSet() {
		// 设置生成的唯一 ID
		field.SetInt(node.Generate().Int64())
	}
}
