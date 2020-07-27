/**
 * @Description: cred repo
 * @File: cred_repo
 * @Date: 2020/7/9 0009 14:39
 */

package repo

import "CMD-XuanFangTong-Server/model"

type CredRepo interface {
	//获取插入到es的数据
	GetToEsData(id uint64) (cred *model.Cred, err error)
}

func NewCredRepo() CredRepo {
	return &credRepo{}
}

type credRepo struct {
	thisModel model.Cred
}

func (c credRepo) GetToEsData(id uint64) (cred *model.Cred, err error) {
	cred = new(model.Cred)
	err = model.DB.Model(c.thisModel).Where("id = ?", id).First(&cred).Error
	if cred.Renovation != 0 {
		if cred.Renovation == 1 {
			cred.RenovationString = "精装"
		} else {
			cred.RenovationString = "毛坯"
		}
	}
	switch cred.Status {
	case 1:
		cred.StatusName = "即将取证"
		break
	case 2:
		cred.StatusName = "最新取证"
		break
	case 3:
		cred.StatusName = "正在认筹"
		break
	case 4:
		cred.StatusName = "最新摇号"
		break
	case 5:
		cred.StatusName = "在售楼盘"
		break
	}
	return
}