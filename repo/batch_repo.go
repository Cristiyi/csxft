/**
 * @Description:
 * @File: batch_repo
 * @Date: 2020/7/29 0029 11:43
 */

package repo

import "csxft/model"

type BatchRepo interface {
	//获取插入到es的数据
	GetToEsData(id uint64) (batch *model.Batch, err error)
}

func NewBatchRepo() BatchRepo {
	return &batchRepo{}
}

type batchRepo struct {
	thisModel model.Iottery
}

func (c batchRepo) GetToEsData(id uint64) (batch *model.Batch, err error) {
	batch = new(model.Batch)
	err = model.DB.Model(c.thisModel).Preload("Creds").Where("id = ?", id).First(&batch).Error
	if batch.Renovation != 0 {
		if batch.Renovation == 1 {
			batch.RenovationString = "精装"
		} else {
			batch.RenovationString = "毛坯"
		}
	}
	switch batch.Status {
	case 1:
		batch.StatusName = "即将取证"
		break
	case 2:
		batch.StatusName = "最新取证"
		break
	case 3:
		batch.StatusName = "正在认筹"
		break
	case 4:
		batch.StatusName = "最新摇号"
		break
	case 5:
		batch.StatusName = "在售楼盘"
		break
	}
	return
}
