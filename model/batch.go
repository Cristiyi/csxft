/**
 * @Description:
 * @File: batch
 * @Date: 2020/7/28 0028 20:26
 */

package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Batch struct {
	gorm.Model
	ProjectId uint64 `gorm:"not null;index:idx_project_id"`
	BatchNo int64
	Name string
	PreSellTime time.Time `gorm:"default:null"`
	SolicitBegin time.Time `gorm:"default:null"`
	SolicitEnd time.Time  `gorm:"default:null"`
	SolicitTime time.Time  `gorm:"default:null"`
	LotteryTime time.Time  `gorm:"default:null"`
	LotteryBegin time.Time  `gorm:"default:null"`
	LotteryEnd time.Time  `gorm:"default:null"`
	ChooseHouseBegin time.Time  `gorm:"default:null"`
	ChooseHouseEnd time.Time  `gorm:"default:null"`
	Renovation uint32
	//装修情况 文字（仅用于展示）
	RenovationString string  `gorm:"-"`
	RenovationNo string
	RoughcastNo string
	MinArea float64  `gorm:"type:decimal(10,2);default:null"`
	MaxArea float64  `gorm:"type:decimal(10,2);default:null"`
	AllNo int
	LotteryNo int
	RigidNo int
	OrdinaryNo int
	LeftNo int
	MinPrice float64  `gorm:"type:decimal(10,2);default:null"`
	MaxPrice float64  `gorm:"type:decimal(10,2);default:null"`
	CustomPrice string
	ProjectTypeId int
	CredDate time.Time
	Acreage float64  `gorm:"type:decimal(10,2);default:null"`
	VerifyMoney string
	Status int32 `gorm:"default:2"`
	IsWillCred int32 `gorm:"default:0"`
	IsNewCred int32  `gorm:"default:0"`
	IsRecognition int32  `gorm:"default:0"`
	IsIottery int32  `gorm:"default:0"`
	IsSell int32  `gorm:"default:0"`
	//状态名称（仅用于展示）
	StatusName string  `gorm:"-"`
	Creds []Cred `gorm:"ForeignKey:BatchID;"`
}
