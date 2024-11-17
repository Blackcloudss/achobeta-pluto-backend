package repo

import (
	"tgwp/internal/model"
	"tgwp/internal/types"
	"time"
)

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

	err := r.DB.Model(&model.Structure{}).
		Create(&node).
		Error
	return err
}
