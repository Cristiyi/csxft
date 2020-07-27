/**
 * @Description:
 * @File: Iottery
 * @Date: 2020/7/17 0017 11:38
 */

package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Iottery struct {
	gorm.Model
	ProjectId uint `gorm:"not null;index:idx_project_id"`
	IotteryName string
	IotteryPrice string
	PreSellDate time.Time
	AllHome uint
	ProbabilityHasHome string
	ProbabilityNoHome string
}
