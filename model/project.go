package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Project struct {
	gorm.Model
	ConstructionNo string `gorm:"not null;index:idx_construction_no"`
	ProjectName string `gorm:"not null"`
	AreaId uint
	Approval string
	DevelopCompany string
	CountBuilding string
	ProjectAddress string
	MinPrice string
	SaleAddress string
	SalePhone string
	AllHome string
	BusLine string
	School string
	ShoppingMall string
	Hospital string
	AllAcreage string
	DesignCompany string
	AllArchitectureAcreage string
	SaleAgency string
	VolumeRatio string
	Property string
	GreenRate string
	ConstructionCompany string
	PropertyCost string
	CompleteTime time.Time
	Introduction string
	ProjectPurpose string
	ProjectType string
	ProjectDecoration string
	CheckStatus uint32
	ViewCount uint64 `gorm:"default:0"`
	AveragePrice float64 `gorm:"type:decimal(10,2);default:null"`
	SaleStatus int32
	IsWillCred int32 `gorm:"default:0"`
	IsNewCred int32  `gorm:"default:1"`
	IsRecognition int32  `gorm:"default:0"`
	IsIottery int32  `gorm:"default:0"`
	IsSell int32  `gorm:"default:0"`
	//效果图
	EffectImages []Image `gorm:"ForeignKey:ProjectId"`
	//样板间图
	TempletImages []Image `gorm:"ForeignKey:ProjectId"`
	//实景图
	LiveImages []Image `gorm:"ForeignKey:ProjectId"`
	//周边图
	CircumImages []Image `gorm:"ForeignKey:ProjectId"`
	//鸟瞰图
	AerialImages []Image `gorm:"ForeignKey:ProjectId"`
	//户型图
	HouseTypeImages []Image `gorm:"ForeignKey:ProjectId"`
	//评论数量 仅用于查询后的展示
	CommentCount int  `gorm:"-"`
	//地区名称 仅用于查询后的展示
	AreaName string `gorm:"-"`
}

