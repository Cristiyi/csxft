/**
 * @Description:
 * @File: lotteryResult
 * @Date: 2020/7/23 0023 14:26
 */

package model

import "github.com/jinzhu/gorm"

type LotteryResult struct {
	gorm.Model
	ProjectId uint `gorm:"not null;index:idx_project_id"`
	//UserId uint `gorm:"not null;index:idx_user_id"`
	Type int32
	BatchId int64
	No int64
	Username string
	IdCard string
	SolicitNo string
	Title string
	TotalHouse int `gorm:"-"`
	TotalPerson int `gorm:"-"`
	//类型 文字（仅用于展示）
	TypeString string  `gorm:"-"`
	//身份证备份
	//IdCardBack string  `gorm:"-"`
}