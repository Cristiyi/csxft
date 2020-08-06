/**
 * @Description:
 * @File: Notice
 * @Date: 2020/7/24 0024 17:55
 */

package model

import "github.com/jinzhu/gorm"

type Notice struct {
	gorm.Model
	ProjectId uint `gorm:"not null;index:idx_project_id"`
	BatchId uint `gorm:"not null;index:idx_batch_id"`
	Type int32
	NoticeType int32
	Content string
	Url string
	Status int32
}
