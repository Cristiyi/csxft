/**
 * @Description:
 * @File: fdc
 * @Date: 2020/7/30 0030 20:32
 */

package model

import "github.com/jinzhu/gorm"

type Fdc struct {
	gorm.Model
	Cred string
	ProjectName string
	SaleStatus string
	HouseType string
	ReferencePrice string
	PrimaryHouseStructure string
	Area string
	DecorationStatus string
	PlannedAllHome string
	CompletedTime string
	StallCount string
	OpenTime string
	AllAcreage string
	AllFloorAcreage string
	VolumeRatio string
	GreenRate string
	DesignCompany string
	PropertyCost string
	Property string
	ConstructionCompany string
	DevelopCompany string
	Address string
	BusLine string
	SurroundingFacility string
}
