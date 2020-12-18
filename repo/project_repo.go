/**
 * @Description: project repo
 * @File: project_repo
 * @Date: 2020/7/9 0009 14:39
 */

package repo

import (
	"csxft/model"
	"time"
)

type ProjectRepo interface {
	//获取插入到es的数据
	GetToEsData(id uint64) (project *model.Project, err error)
	//获取即将取证的分组
	GetPredictCredDate() (group []*model.PredictCredDate)
	//获取所有插入到es的数据
	GetAllToEsData() (projects []*model.Project, err error)
	GetOne(id uint64) (project *model.Project, err error)
}

func NewProjectRepo() ProjectRepo {
	return &projectRepo{}
}

type projectRepo struct {
	thisModel model.Project
}

func (p projectRepo) GetAllToEsData() (projects []*model.Project, err error) {
	err = model.DB.Preload("EffectImages", "type = 1").
		Preload("TempletImages", "type = 2").
		Preload("LiveImages", "type = 3").
		Preload("CircumImages", "type = 4").
		Preload("AerialImages", "type = 5").
		//Preload("HouseTypeImages", "type = 6").
		Preload("AerialMainImages", "type = 6").
		Preload("AerialUploadImages", "type = 7").
		Where("no_status = ?", 1).Find(&projects).Error
	if err == nil {
		for i, item := range projects {
			area := new(model.Area)
			if err := model.DB.Where("id = ?", item.AreaId).First(&area).Error; err == nil {
				projects[i].AreaName = area.Name
			}
			var count = 0
			model.DB.Model(model.Comment{}).Where("build_id = ? and pid = ? and status = ?", item.ID, 0, 1).Count(&count)
			projects[i].CommentCount = count
		}
	}
	return
}

func (p projectRepo) GetToEsData(id uint64) (project *model.Project, err error) {
	project = new(model.Project)
	err = model.DB.Preload("EffectImages", "type = 1").
		           Preload("TempletImages", "type = 2").
				   Preload("LiveImages", "type = 3").
				   Preload("CircumImages", "type = 4").
				   Preload("AerialImages", "type = 5").
				   //Preload("HouseTypeImages", "type = 6").
				   Preload("AerialMainImages", "type = 6").
		           Preload("AerialUploadImages", "type = 7").
				   Where("id = ?", id).First(&project).Error
	if err == nil {
		area := new(model.Area)
		if err := model.DB.Where("id = ?", project.AreaId).First(&area).Error; err == nil {
			project.AreaName = area.Name
		}
		var count = 0
		model.DB.Model(model.Comment{}).Where("build_id = ? and pid = ? and status = ?", project.ID, 0, 1).Count(&count)
		project.CommentCount = count
	}
	return
}

//获取单楼盘记录
func (p projectRepo) GetOne(id uint64) (project *model.Project, err error) {
	project = new(model.Project)
	err = model.DB.Where("id = ?", id).First(&project).Error
	return
}

func (p projectRepo) GetPredictCredDate() (group []*model.PredictCredDate) {
	var predictCredDate []*model.PredictCredDate
	var predictCredTemp []*model.PredictCredTemp
	err := model.DB.Model(p.thisModel).Select("predict_cred_date").Where("is_will_cred = ? and predict_cred_date is not null", 1).Group("predict_cred_date").Scan(&predictCredTemp).Error
	if err != nil {
		return nil
	} else {
		timeLayout := "2006年01月"
		for _, item := range predictCredTemp {
			temp := new(model.PredictCredDate)
			temp.PredictCredDate = item.PredictCredDate.Unix()
			temp.PredictCredMonth = time.Unix(item.PredictCredDate.Unix(), 0).Format(timeLayout)
			predictCredDate = append(predictCredDate, temp)
		}
		return predictCredDate
	}
}

