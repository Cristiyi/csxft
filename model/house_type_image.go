/**
 * @Description:
 * @File: house_type_image
 * @Date: 2020/7/24 0024 11:51
 */

package model

import "github.com/jinzhu/gorm"

type HouseTypeImage struct {
	gorm.Model
	ProjectId uint `gorm:"not null;index:idx_project_id"`
	HomeNum string
	HouseStruct string
	Acreage string
	Tag string
	ImageUrl string
}
