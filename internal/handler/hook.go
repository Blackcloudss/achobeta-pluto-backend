package handler

import (
	"gorm.io/gorm"
	"reflect"
	"tgwp/log/zlog"
	"tgwp/util/snowflake"
)

// 默认节点 (目前仅这里需要用到)
// 业务需求小，一个节点即可
const Default_Node_ID = 2

// 注册 gorm 钩子
func RegisterHook(db *gorm.DB) {
	db.Callback().Create().Before("gorm:Create").Register("before_create_Node", BeforeCreateNode)
}

// 在操作数据库前 创建 用户 ID
func BeforeCreateNode(tx *gorm.DB) {
	node, err := snowflake.NewNode(Default_Node_ID)
	if err != nil {
		zlog.Errorf("生成 Node 出错")
		return
	}

	// 如果是指针，返回指针指向的值；如果是非指针，直接返回
	tx.Statement.ReflectValue = reflect.Indirect(tx.Statement.ReflectValue)

	// 确认 id 字段是否存在，并且是 int64 类型
	if field := tx.Statement.ReflectValue.FieldByName("id"); field.IsValid() && field.CanSet() {
		// 设置生成的唯一 ID
		field.SetInt(node.Generate().Int64())
	}
}
