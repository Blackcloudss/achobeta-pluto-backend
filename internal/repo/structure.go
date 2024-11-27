package repo

import (
	"fmt"
	"gorm.io/gorm"
	"tgwp/internal/model"
	"tgwp/internal/types"
	"tgwp/log/zlog"
)

const c_NodeName = "node_name"

type StructureRepo struct {
	DB *gorm.DB
}

func NewStructureRepo(db *gorm.DB) *StructureRepo {
	return &StructureRepo{DB: db}
}

// GetNode
//
//	@Description:  获取架构节点
//	@receiver r
//	@param fatherid
//	@param teamid
//	@return []MyNode
//	@return error
type MyNode struct {
	MyselfId int64  `gorm:"column:id"`
	NodeName string `gorm:"column:node_name"`
}

func (r StructureRepo) GetNode(fatherid, teamid int64) ([]MyNode, error) {

	var mynode []MyNode
	err := r.DB.Model(&model.Structure{}).
		Select(C_Id, c_NodeName).
		Where("structure.deleted_at IS NULL AND father_id = ? AND team_id = ?", fatherid, teamid).
		Find(&mynode).
		Error
	if err != nil {
		zlog.Errorf("获取当前节点失败：%v", err)
		return nil, err
	}
	return mynode, nil
}

// CreateNode
//
//	@Description: 新增架构节点
//	@receiver r
//	@param Node
//	@return error
func (r StructureRepo) CreateNode(Node types.TeamStructure) error {

	var node = model.Structure{
		TeamId:   Node.TeamId,
		FatherId: Node.FatherId,
		NodeName: Node.NodeName,
	}

	err := r.DB.Model(&model.Structure{}).
		Create(&node).
		Error
	return err
}

// DeleteNode
//
//	@Description: 删除架构节点
//	@receiver r
//	@param Node
//	@return error
const C_FATHER = "father_id"

func (r StructureRepo) DeleteNode(Node types.TeamStructure) error {

	//将 被删除的节点 的 子节点 的 父节点 变为 被删除的节点 的 父节点
	err := r.DB.Model(&model.Structure{}).
		Where("father_id = ? AND team_id = ?", Node.MyselfId, Node.TeamId).
		Update("father_id", Node.FatherId).
		Error
	if err != nil {
		zlog.Errorf("被删除节点的子节点的父节点更改失败%v", err)
		return err
	}

	err = r.DB.Model(&model.Structure{}).
		Where(fmt.Sprintf("%s = ? AND %s = ?", C_Id, C_FATHER), Node.MyselfId, Node.FatherId).
		Delete(&model.Structure{}).
		Error
	if err != nil {
		zlog.Errorf("指定节点删除失败%v", err)
		return err
	}
	return err
}
