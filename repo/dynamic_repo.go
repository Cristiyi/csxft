/**
 * @Description:
 * @File: dynamic_repo
 * @Date: 2020/7/22 0022 20:39
 */

package repo

import "CMD-XuanFangTong-Server/model"

type DynamicRepo interface {
	//获取插入到es的数据
	GetToEsData(id uint64) (dynamic *model.Dynamic, err error)
}

func NewDynamicRepo() DynamicRepo {
	return &dynamicRepo{}
}

type dynamicRepo struct {
	thisModel model.Dynamic
}

func (c dynamicRepo) GetToEsData(id uint64) (dynamic *model.Dynamic, err error) {
	dynamic = new(model.Dynamic)
	err = model.DB.Model(c.thisModel).Where("id = ?", id).First(&dynamic).Error
	return
}
