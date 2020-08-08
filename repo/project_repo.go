/**
 * @Description: project repo
 * @File: project_repo
 * @Date: 2020/7/9 0009 14:39
 */

package repo

import (
	"csxft/model"
)

type ProjectRepo interface {
	//获取插入到es的数据
	GetToEsData(id uint64) (project *model.Project, err error)
}

func NewProjectRepo() ProjectRepo {
	return &projectRepo{}
}

type projectRepo struct {
	thisModel model.Project
}


func (p projectRepo) GetToEsData(id uint64) (project *model.Project, err error) {
	project = new(model.Project)
	err = model.DB.Preload("EffectImages", "type = 1").
		           Preload("TempletImages", "type = 2").
				   Preload("LiveImages", "type = 3").
				   Preload("CircumImages", "type = 4").
				   Preload("AerialImages", "type = 5").
				   Preload("HouseTypeImages", "type = 6").
				   Where("id = ?", id).First(&project).Error
	if err == nil {
		area := new(model.Area)
		if err := model.DB.Where("id = ?", project.AreaId).First(&area).Error; err == nil {
			project.AreaName = area.Name
		}
		//var count = 0
		//model.DB.Model(model.Comment{}).Where("build_id = ? and pid = ?", project.ID, 0).Count(&count)
		//project.CommentCount = count
	}
	return
}

