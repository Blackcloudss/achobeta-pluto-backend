package repo

import (
	"gorm.io/gorm"
	"tgwp/internal/model"
)

type StructureRepo struct {
	DB *gorm.DB
}

func NewStructureRepo(db *gorm.DB) *StructureRepo {
	return &StructureRepo{DB: db}
}

type MyNode struct {
	MyselfId int64
	NodeName string
}

// 获取架构节点
func (r StructureRepo) GetNode(fatherid, teamid int64) ([]MyNode, error) {

	var mynode []MyNode

	err := r.DB.Model(&model.Structure{}).
		Select(C_Id, C_NodeName).
		Where(&model.Structure{
			FatherId: fatherid,
			TeamId:   teamid,
		}).
		Find(&mynode).
		Error
	if err != nil {
		return nil, err
	}
	return mynode, nil
}
