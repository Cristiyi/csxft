/**
 * @Description:
 * @File: follow
 * @Date: 2020/7/27 0027 12:11
 */

package model

import "github.com/jinzhu/gorm"

type Follow struct {
	gorm.Model
	UserId uint64
	BuildId uint64
	Type int32
	Status int32
}
