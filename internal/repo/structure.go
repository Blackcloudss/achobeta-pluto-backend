package repo

import (
	"fmt"
	"gorm.io/gorm"
)

type StructureRepo struct {
	DB *gorm.DB
}

func NewStructureRepo(db *gorm.DB) *StructureRepo {
	return &StructureRepo{DB: db}
}

type MyNode struct {
	MyselfId   int64
	StructName string
}

// 获取架构节点
func (r StructureRepo) GetNode(fatherid, teamid int64) ([]MyNode, error) {

	var mynode []MyNode

	err := r.DB.Table(StructureTableName).
		Select(MyselfId, StructureName).
		Where(fmt.Sprintf("%s = ? AND %s = ?", FatherId, TeamId), fatherid, teamid).
		Find(&mynode).
		Error
	if err != nil {
		return nil, err
	}
	return mynode, nil
}
