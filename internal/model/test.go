package model

type Test struct {
	CommonModel
	UserID string
	Name   string
	Age    int
}

func (t *Test) TableName() {

}
