package model

// Test 一个实体类，对应数据库一张表
type Test struct {
	CommonModel
	UserID string
	Name   string
	Age    int
}

func (t *Test) TableName() {

}
