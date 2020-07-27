/**
 * @Description: comment model
 * @File: comment
 * @Date: 2020/7/13 0013 11:33
 */

package model

import (
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	BuildId uint64 `gorm:"not null;index:idx_build_id"`
	Pid uint64
	AvatarUrl string
	NickName string
	FromUid uint64
	ToUid uint64
	Content string
	Likes int
	Status uint32
}