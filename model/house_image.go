/**
 * @Description:
 * @File: house_image
 * @Date: 2020/12/18 0018 15:08
 */

package model

import "github.com/jinzhu/gorm"

type HouseImage struct {
	gorm.Model
	ProjectId uint `gorm:"not null;index:idx_project_id"`
	BatchId uint
	Images string
	Name string
}
