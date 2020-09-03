package model

import (
	"github.com/jinzhu/gorm"
)

type House struct {
	gorm.Model
	CredId uint64 `gorm:"not null;index:idx_cred_id"`
	HouseNo string
	FloorNo int
	PurposeId uint
	//房屋目的 文字（仅用于展示）
	PurposeString string  `gorm:"-"`
	TypeId uint
	DecorationId uint
	//装修 文字（仅用于展示）
	DecorationString string  `gorm:"-"`
	//类型 文字（仅用于展示）
	TypeString string  `gorm:"-"`
	HouseAcreage float64  `gorm:"type:decimal(10,2);default:null"`
	UseAcreage float64  `gorm:"type:decimal(10,2);default:null"`
	ShareAcreage float64  `gorm:"type:decimal(10,2);default:null"`
	UnitPrice float64  `gorm:"type:decimal(10,2);default:null"`
	//RoughcastUnitPrice float64  `gorm:"type:decimal(10,2);default:null"`
	DownPaymentThird float64  `gorm:"type:decimal(10,2);default:null"`
	DownPaymentSixth float64  `gorm:"type:decimal(10,2);default:null"`
	TotalPrice float64  `gorm:"type:decimal(10,2);default:null"`
	SaleStatus int32 `gorm:"default:1"`
	//楼栋号
	BuildNo string  `gorm:"-"`
}
