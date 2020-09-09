package es

import (
	"csxft/repo"
	"csxft/serializer"
)

// InitBaseService 初始化基础数据服务
type CreateService struct {
	Type        uint32 `form:"type" json:"type" binding:"required"`
	TargetId    uint64 `form:"target_id" json:"targetId" binding:"required"`
}

//初始化楼盘相关所有信息
type ProjectAllInfoService struct {
	ProjectId    uint64 `form:"project_id" json:"project_id" binding:"required"`
}

//初始化楼盘相关所有信息
type InitProjectAllService struct {
}

//初始化批次相关所有信息
type InitBatchAllService struct {
}

//初始化一房一价相关所有信息
type InitHouseAllService struct {
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

// 初始化
func (service *ProjectAllInfoService) Create() serializer.Response {

	CreateAllProjectInfo(service.ProjectId)

	return serializer.Response{
		Code: 200,
		Msg: "success",
	}

}

// 初始化所有楼盘到es
func (service *InitProjectAllService) InitProjectAll() serializer.Response {

	projects, err := repo.NewProjectRepo().GetAllToEsData()
	if err == nil {
		for _, item := range projects {
			strategy := NewCreateContext(1)
			strategy.Create(uint64(item.ID))
		}
	}

	return serializer.Response{
		Code: 200,
		Msg: "success",
	}

}

// 初始化所有批次到es
func (service *InitBatchAllService) InitBatchAll() serializer.Response {

	batches, err := repo.NewBatchRepo().GetAllToEsData()
	if err == nil {
		for _, item := range batches {
			strategy := NewCreateContext(11)
			strategy.Create(uint64(item.ID))
		}
	}

	return serializer.Response{
		Code: 200,
		Msg: "success",
	}

}

// 初始化所有一房一价到es
func (service *InitHouseAllService) InitHouseAll() serializer.Response {

	batches, err := repo.NewBatchRepo().GetAllToEsData()
	if err == nil {
		for _, item := range batches {
			for _, cred := range item.Creds {
				if err == nil {
					strategy := NewCreateContext(3)
					strategy.Create(uint64(cred.ID))
				}
			}
		}
	}

	return serializer.Response{
		Code: 200,
		Msg: "success",
	}

}








