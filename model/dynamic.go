/**
 * @Description:
 * @File: dynamic
 * @Date: 2020/7/22 0022 20:20
 */

package model

import "github.com/jinzhu/gorm"

type Dynamic struct {
	gorm.Model
	ProjectId uint `gorm:"not null;index:idx_project_id"`
	Title string
	Type int32
	Content string
}

