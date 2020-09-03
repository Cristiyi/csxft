package model

import "github.com/jinzhu/gorm"

type ProjectType struct {
	gorm.Model
	Type int32
}

type TypeNameResult struct {
	Name string
}

func GetTypeNameById(id uint) (nameResult TypeNameResult) {
	DB.Model(ProjectType{}).Where("id = ?", id).Select("name").Scan(&nameResult)
	return
}