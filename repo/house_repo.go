/**
 * @Description: house repo
 * @File: house_repo
 * @Date: 2020/7/10 0010 18:24
 */

package repo

import (
	"csxft/model"
)

type HouseRepo interface {
	//获取插入到es的数据
	GetToEsData(id uint64) (houses []model.House, err error)
	//获取单条预售证下的一房一价
	GetOneCredHouseData(id uint64) (houses []model.House, err error)
}

func NewHouseRepo() HouseRepo {
	return &houseRepo{}
}

type houseRepo struct {
	thisModel model.House
}

//获取插入到es的数据
func (c houseRepo) GetToEsData(id uint64) (houses []model.House, err error) {
	err = model.DB.Model(c.thisModel).Where("cred_id = ?", id).Find(&houses).Error
	cred := new(model.Cred)
	model.DB.Model(model.Cred{}).Where("id = ?", id).First(&cred)
	for i, item := range houses {
		if item.DecorationId >= 0 {
			houses[i].DecorationString = model.GetDecorationNameById(item.DecorationId).Name
		}
		if item.PurposeId >= 0 {
			houses[i].PurposeString = model.GetPurposeNameById(item.PurposeId).Name
		}
		if item.TypeId >= 0 {
			houses[i].TypeString = model.GetTypeNameById(item.TypeId).Name
		}
		houses[i].BuildNo = cred.BuildingNo
	}
	return
}


//获取单条预售证下的一房一价
func (c houseRepo) GetOneCredHouseData(id uint64) (houses []model.House, err error) {
	err = model.DB.Model(c.thisModel).Where("cred_id = ?", id).Find(&houses).Error
	return
}