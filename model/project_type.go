package model

import "github.com/jinzhu/gorm"

type ProjectType struct {
	gorm.Model
	Type int32
}