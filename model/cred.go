package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Cred struct {
	gorm.Model
	ProjectId uint64 `gorm:"not null;index:idx_project_id"`
	batch int
	name string
	SaleScope string
	Cred string `gorm:"not null;index:idx_cred"`
	BuildingNo string
	PreSellTime time.Time `gorm:"default:null"`
	SolicitBegin time.Time `gorm:"default:null"`
	SolicitEnd time.Time  `gorm:"default:null"`
	SolicitTime time.Time  `gorm:"default:null"`
	LotteryTime time.Time  `gorm:"default:null"`
	LotteryBegin time.Time  `gorm:"default:null"`
	LotteryEnd time.Time  `gorm:"default:null"`
	ChooseHouseBegin time.Time  `gorm:"default:null"`
	ChooseHouseEnd time.Time  `gorm:"default:null"`
	LotteryStatus int
	Renovation uint32
	//装修情况 文字（仅用于展示）
	RenovationString string  `gorm:"-"`
	RenovationNo string
	RoughcastNo string
	MinArea float64  `gorm:"type:decimal(10,2);default:null"`
	MaxArea float64  `gorm:"type:decimal(10,2);default:null"`
	AllNo int
	ShedNo int
	RigidNo int
	OrdinaryNo int
	LeftNo int
	MinPrice float64  `gorm:"type:decimal(10,2);default:null"`
	MaxPrice float64  `gorm:"type:decimal(10,2);default:null"`
	CustomPrice string
	TypeId uint
	CredDate time.Time
	Acreage string
	LandNo string
	EngineeNo string
	LandPlanNo string
	VerifyMoney string
	IsNew uint32 `gorm:"default:0"`
	Status int32 `gorm:"default:2"`
	//状态名称（仅用于展示）
	StatusName string  `gorm:"-"`
}