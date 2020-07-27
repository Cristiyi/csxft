package model

import "github.com/jinzhu/gorm"

type Area struct {
	gorm.Model
	Name string
	Pid int64
}
