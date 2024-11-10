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

// GetUserById 数据库操作层，数据库操作层应该与逻辑层解耦 所有数据库和redis的操作应该放在repo包下内执行
func (r TestRepo) GetUserById(userId string) (testUser model.Test, err error) {
	err = global.DB.Model(&model.Test{}).Where(model.Test{UserID: userId}).First(&testUser).Error
	return
}
