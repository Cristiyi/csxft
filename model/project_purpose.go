package model

import "github.com/jinzhu/gorm"

type ProjectPurpose struct {
	gorm.Model
	Purpose int32
	Name string
}

type PurposeNameResult struct {
	Name string
}

func GetPurposeNameById(id uint) (nameResult PurposeNameResult) {
	DB.Model(ProjectPurpose{}).Where("id = ?", id).Select("name").Scan(&nameResult)
	return
}