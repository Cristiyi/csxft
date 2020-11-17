/**
 * @Description:
 * @File: SolicitResult
 * @Date: 2020/7/23 0023 18:03
 */

package model

import "github.com/jinzhu/gorm"

type SolicitResult struct {
	gorm.Model
	ProjectId uint `gorm:"not null;index:idx_project_id"`
	BatchId int32
	SolicitNo string
	Username string
	IdCard string
	//Title string
	//身份证备份
	IdCardBack string  `gorm:"-"`
	Type int32
	ContentType int32
	Url string
}
