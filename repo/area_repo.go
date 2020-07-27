/**
 * @Description:
 * @File: area_repo
 * @Date: 2020/7/23 0023 19:16
 */

package repo

import "csxft/model"

type AreaRepo interface {
	//获取插入到es的数据
	GetToEsData(id uint64) (area []model.Area, err error)
}

func NewAreaRepo() AreaRepo {
	return &areaRepo{}
}

type areaRepo struct {
	thisModel model.Area
}

func (c areaRepo) GetToEsData(id uint64) (area []model.Area, err error) {
	err = model.DB.Model(c.thisModel).Where("pid = ?", id).Find(&area).Error
	return
}

