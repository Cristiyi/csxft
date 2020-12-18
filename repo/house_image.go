/**
 * @Description:
 * @File: house_image
 * @Date: 2020/12/18 0018 15:10
 */

package repo

import "csxft/model"

type HouseImageRepo interface {
	//获取插入到es的数据
	GetToEsData(id uint64) (houseImage *model.HouseImage, err error)
}

func NewHouseImageRepo() HouseImageRepo {
	return &houseImageRepo{}
}

type houseImageRepo struct {
	thisModel model.HouseImage
}

func (c houseImageRepo) GetToEsData(id uint64) (houseImage *model.HouseImage, err error) {
	houseImage = new(model.HouseImage)
	err = model.DB.Model(c.thisModel).Where("id = ?", id).First(&houseImage).Error
	return
}

