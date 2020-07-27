/**
 * @Description:
 * @File: house_type_image
 * @Date: 2020/7/24 0024 11:58
 */

package repo

import (
	"CMD-XuanFangTong-Server/model"
)

type HouseTypeImageRepo interface {
	//获取插入到es的数据
	GetToEsData(id uint64) (houseTypeImage *model.HouseTypeImage, err error)
	//获取户型图分组
	GetHouseImageGroup(projectId uint64) (houseTypeImage []HouseImageGroup)
}

func NewHouseTypeImageRepo() HouseTypeImageRepo {
	return &houseTypeImageRepo{}
}

type houseTypeImageRepo struct {
	thisModel model.LotteryResult
}

type HouseImageGroup struct {
	HomeNum string
}

//获取插入到es的数据
func (c houseTypeImageRepo) GetToEsData(id uint64) (houseTypeImage *model.HouseTypeImage, err error) {
	houseTypeImage = new(model.HouseTypeImage)
	err = model.DB.Model(c.thisModel).Where("id = ?", id).First(&houseTypeImage).Error
	return
}

//获取户型图分组
func (c houseTypeImageRepo) GetHouseImageGroup(projectId uint64) (houseTypeImage []HouseImageGroup) {
	err := model.DB.Table("xft_house_type_images").Select("home_num").Where("project_id = ?", projectId).Group("home_num").Scan(&houseTypeImage).Error
	if err != nil {
		return nil
	}
	return
}
