package repo

import (
	"fmt"
	"gorm.io/gorm"
	"tgwp/internal/model"
	"tgwp/internal/types"
	"time"
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

	err := r.DB.Table(StructureTableName).
		Select(C_Id, C_NodeName).
		Where(fmt.Sprintf("%s = ? AND %s = ?", C_FatherId, C_TeamId), fatherid, teamid).
		Find(&mynode).
		Error
	if err != nil {
		return nil, err
	}
	return mynode, nil
}

// 新增架构节点
func (r StructureRepo) InsertNode(Node types.TeamStructure) error {

	var node = model.Structure{
		CommonModel: model.CommonModel{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		TeamId:   Node.TeamId,
		FatherId: Node.FatherId,
		NodeName: Node.NodeName,
	}

	err := r.DB.Table(StructureTableName).
		Create(&node).
		Error
	return err
}

// 删除架构节点
func (r StructureRepo) DeleteNode(Node types.TeamStructure) error {

	var node = model.Structure{
		CommonModel: model.CommonModel{
			ID: Node.MyselfId},
		TeamId:   Node.TeamId,
		FatherId: Node.FatherId,
		NodeName: Node.NodeName,
	}

	err := r.DB.Table(StructureTableName).
		Where(fmt.Sprintf("%s = ? AND %s =? AND %s = ?", C_Id, C_FatherId, C_TeamId), Node.MyselfId, Node.FatherId, Node.TeamId).
		Delete(&node).
		Error
	return err
}
