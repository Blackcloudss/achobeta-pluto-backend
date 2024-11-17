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
	zlog.Infof("Registering GORM hooks...")
	db.Callback().Create().Before("gorm:Create").Register("before_create_Node", BeforeCreateNode)
}

// 在操作数据库前 创建 用户 ID
func BeforeCreateNode(db *gorm.DB) {
	// 检查是否是切片
	if db.Statement.ReflectValue.Kind() == reflect.Slice {
		for i := 0; i < db.Statement.ReflectValue.Len(); i++ {
			elem := db.Statement.ReflectValue.Index(i)
			elem = reflect.Indirect(elem) // 取指针指向的值
			setIDForElem(elem)
		}
	} else {
		elem := reflect.Indirect(db.Statement.ReflectValue)
		setIDForElem(elem)
	}
}

func setIDForElem(elem reflect.Value) {
	// 确认 ID 字段存在并可设置
	if field := elem.FieldByName("ID"); field.IsValid() && field.CanSet() && field.Kind() == reflect.Int64 {
		node, err := snowflake.NewNode(global.DEFAULT_NODE_ID)
		if err != nil {
			zlog.Errorf("生成 Snowflake 节点失败: %v", err)
			return
		}
		newID := node.Generate().Int64()
		field.SetInt(newID)
		zlog.Infof("生成的 ID: %d", newID)
	}
}
