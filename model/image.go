/**
 * @Description: 图片模型
 * @File: Image
 * @Date: 2020/7/10 0010 16:19
 */

package model

import "github.com/jinzhu/gorm"

type Image struct {
	gorm.Model
	ProjectId uint
	Type uint32
	ImageUrl string
}
