package repo

import (
	"tgwp/global"
	"tgwp/internal/model"
)

type TestRepo struct {
}

func NewTestRepo() *TestRepo {
	return &TestRepo{}
}
func (r TestRepo) GetUserById(userId string) (testUser model.Test, err error) {
	err = global.DB.Model(&model.Test{}).Where(model.Test{UserID: userId}).First(&testUser).Error
	return
}
