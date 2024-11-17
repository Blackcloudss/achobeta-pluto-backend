package repo

import (
	"fmt"
	"tgwp/internal/model"
	"tgwp/internal/types"
)

// 删除架构节点
func (r StructureRepo) DeleteNode(Node types.TeamStructure) error {

	var node = model.Structure{
		CommonModel: model.CommonModel{
			ID: Node.MyselfId},
		TeamId:   Node.TeamId,
		FatherId: Node.FatherId,
		NodeName: Node.NodeName,
	}

	err := r.DB.Model(&model.Structure{}).
		Where(fmt.Sprintf("%s = ?", C_Id), Node.MyselfId).
		Delete(&node).
		Error
	return err
}
