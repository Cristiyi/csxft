/**
 * @Description:
 * @File: batch_repo
 * @Date: 2020/7/29 0029 11:43
 */

package repo

import (
	"csxft/model"
	"os"
)

type BatchRepo interface {
	//获取插入到es的数据
	GetToEsData(id uint64) (batch *model.Batch, err error)
	//获取正在认筹任务数据
	GetRecognitionTask() (batches []*model.Batch, err error)
	//获取正在认筹到期任务数据
	GetNotRecognitionTask() (batches []*model.Batch, err error)
	//获取正在摇号任务数据
	GetLotteryTask() (batches []*model.Batch, err error)
	//获取正在摇号到期任务数据
	GetNotLotteryTask() (batches []*model.Batch, err error)
	//获取所有插入到es的数据
	GetAllToEsData() (batches []*model.Batch, err error)
}

func NewBatchRepo() BatchRepo {
	return &batchRepo{}
}

type batchRepo struct {
	thisModel model.Batch
}

func (c batchRepo) GetAllToEsData() (batches []*model.Batch, err error) {
	err = model.DB.Model(c.thisModel).Preload("Creds").Find(&batches).Error
	if err != nil {
		for i, batch := range batches {
			if batch.Renovation != 0 {
				if batch.Renovation == 1 {
					batches[i].RenovationString = "精装"
				} else {
					batches[i].RenovationString = "毛坯"
				}
			}
			switch batch.Status {
			case 1:
				batches[i].StatusName = "即将取证"
				break
			case 2:
				batches[i].StatusName = "最新取证"
				break
			case 3:
				batches[i].StatusName = "正在认筹"
				break
			case 4:
				batches[i].StatusName = "最新摇号"
				break
			case 5:
				batches[i].StatusName = "在售楼盘"
				break
			}
		}
	}
	return
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

//获取正在认筹的任务
func (c batchRepo) GetRecognitionTask() (batches []*model.Batch, err error) {
	err = model.DB.Model(c.thisModel).Where("TO_DAYS(solicit_begin) <= TO_DAYS(now()) and TO_DAYS(solicit_end) >= TO_DAYS(now()) and status != 3").Find(&batches).Error
	return
}

//获取正在认筹的到期任务
func (c batchRepo) GetNotRecognitionTask() (batches []*model.Batch, err error) {
	err = model.DB.Model(c.thisModel).Where("TO_DAYS(solicit_end) < TO_DAYS(now()) and status = 3").Find(&batches).Error
	return
}

//获取正在摇号的任务
func (c batchRepo) GetLotteryTask() (batches []*model.Batch, err error) {
	err = model.DB.Model(c.thisModel).Where("TO_DAYS(DATE_ADD(lottery_time, INTERVAL " + os.Getenv("NEW_LOTTERY_TIME_BASE") + " DAY)) >= TO_DAYS(now())").Find(&batches).Error
	return
}

//获取正在要好的到期任务
func (c batchRepo) GetNotLotteryTask() (batches []*model.Batch, err error) {
	err = model.DB.Model(c.thisModel).Where("TO_DAYS(DATE_ADD(lottery_time, INTERVAL " + os.Getenv("NEW_LOTTERY_TIME_BASE") + " DAY)) < TO_DAYS(now()) and status = 4").Find(&batches).Error
	return
}

