package model

import (
	"github.com/jinzhu/gorm"
)

type Decoration struct {
	gorm.Model
	Status uint32
}

type DecorationNameResult struct {
	Name string
}

func GetDecorationNameById(id uint) (nameResult DecorationNameResult) {
	DB.Model(Decoration{}).Where("id = ?", id).Select("name").Scan(&nameResult)
	return
}