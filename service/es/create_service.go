package es

import (
	"csxft/serializer"
)

// InitBaseService 初始化基础数据服务
type CreateService struct {
	Type        uint32 `form:"type" json:"type" binding:"required"`
	TargetId    uint64 `form:"target_id" json:"targetId" binding:"required"`
}

// 初始化
func (service *CreateService) CommonCreate() serializer.Response {

	strategy := NewCreateContext(service.Type)
	code, msg := strategy.Create(service.TargetId)

	return serializer.Response{
		Code: code,
		Msg: msg,
	}

}




