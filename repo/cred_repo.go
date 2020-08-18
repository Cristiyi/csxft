/**
 * @Description: cred repo
 * @File: cred_repo
 * @Date: 2020/7/9 0009 14:39
 */

package repo

import (
	"csxft/model"
	"os"
)

type CredRepo interface {
	//获取插入到es的数据
	GetToEsData(id uint64) (cred *model.Cred, err error)
	GetByProjectId(projectId uint64) (creds []*model.Cred, err error)
	GetByCred(cred string) (creds []*model.Cred)
	//获取最新取证自动化任务数据
	GetNewCredTask() (creds []model.Cred, err error)
	//获取最新取证到期自动化任务数据
	GetNotNewCredTask() (creds []model.Cred, err error)
}

func NewCredRepo() CredRepo {
	return &credRepo{}
}

type credRepo struct {
	thisModel model.Cred
}

//func (c credRepo) GetToEsData(batchId uint64) (creds []*model.Cred, err error) {
//	err = model.DB.Model(c.thisModel).Where("batch_id = ?", batchId).Find(&creds).Error
//	for i, cred := range creds {
//		if cred.Renovation != 0 {
//			if cred.Renovation == 1 {
//				creds[i].RenovationString = "精装"
//			} else {
//				creds[i].RenovationString = "毛坯"
//			}
//		}
//		switch cred.Status {
//		case 1:
//			creds[i].StatusName = "即将取证"
//			break
//		case 2:
//			creds[i].StatusName = "最新取证"
//			break
//		case 3:
//			creds[i].StatusName = "正在认筹"
//			break
//		case 4:
//			creds[i].StatusName = "最新摇号"
//			break
//		case 5:
//			creds[i].StatusName = "在售楼盘"
//			break
//		}
//	}
//	return
//}

func (c credRepo) GetToEsData(id uint64) (cred *model.Cred, err error) {
	cred = new(model.Cred)
	err = model.DB.Model(c.thisModel).Where("id = ?", id).First(&cred).Error
	//if cred.Renovation != 0 {
	//	if cred.Renovation == 1 {
	//		cred.RenovationString = "精装"
	//	} else {
	//		cred.RenovationString = "毛坯"
	//	}
	//}
	//switch cred.Status {
	//case 1:
	//	cred.StatusName = "即将取证"
	//	break
	//case 2:
	//	cred.StatusName = "最新取证"
	//	break
	//case 3:
	//	cred.StatusName = "正在认筹"
	//	break
	//case 4:
	//	cred.StatusName = "最新摇号"
	//	break
	//case 5:
	//	cred.StatusName = "在售楼盘"
	//	break
	//}
	return
}

func (c credRepo) GetByProjectId(projectId uint64) (creds []*model.Cred, err error) {
	err = model.DB.Model(c.thisModel).Where("project_id = ?", projectId).Find(&creds).Error
	//for i, cred := range creds {
	//	if cred.Renovation != 0 {
	//		if cred.Renovation == 1 {
	//			creds[i].RenovationString = "精装"
	//		} else {
	//			creds[i].RenovationString = "毛坯"
	//		}
	//	}
	//	switch cred.Status {
	//	case 1:
	//		creds[i].StatusName = "即将取证"
	//		break
	//	case 2:
	//		creds[i].StatusName = "最新取证"
	//		break
	//	case 3:
	//		creds[i].StatusName = "正在认筹"
	//		break
	//	case 4:
	//		creds[i].StatusName = "最新摇号"
	//		break
	//	case 5:
	//		creds[i].StatusName = "在售楼盘"
	//		break
	//	}
	//}
	return
}

func (c credRepo) GetByCred(cred string) (creds []*model.Cred) {
	err := model.DB.Model(c.thisModel).Where("cred = ?", cred).Find(&creds).Error
	if err != nil {
		return nil
	}
	return creds
}

//获取最新取证自动化任务数据
func (c credRepo) GetNewCredTask() (creds []model.Cred, err error) {
	//err = model.DB.Raw("SELECT * FROM xft_creds WHERE TO_DAYS(DATE_ADD(cred_date, INTERVAL" + os.Getenv("NEW_CRED_TIME_BASE") + "DAY))>=TO_DAYS(now())").Find(&creds).Error
	err = model.DB.Model(c.thisModel).Where("TO_DAYS(DATE_ADD(cred_date, INTERVAL " + os.Getenv("NEW_CRED_TIME_BASE") + " DAY)) >= TO_DAYS(now())").Find(&creds).Error
	return
}

//获取最新取证到期自动化任务数据
func (c credRepo) GetNotNewCredTask() (creds []model.Cred, err error) {
	err = model.DB.Table("xft_creds").
		           Select("xft_creds.*").
		           Joins("left join xft_batches as batch on xft_creds.batch_id = batch.id").
		           Where("TO_DAYS(DATE_ADD(xft_creds.cred_date, INTERVAL " + os.Getenv("NEW_CRED_TIME_BASE") + " DAY)) < TO_DAYS(now()) and batch.is_new_cred = 1").
		           Find(&creds).
		           Error
	return

}