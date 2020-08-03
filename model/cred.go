package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Cred struct {
	gorm.Model
	ProjectId    uint64 `gorm:"not null;index:idx_project_id"`
	BatchId      uint
	Cred         string `gorm:"not null;index:idx_cred"`
	BuildingNo   string
	CredDate time.Time
	Acreage string
	LandNo string
	EngineeNo string
	LandPlanNo string
}