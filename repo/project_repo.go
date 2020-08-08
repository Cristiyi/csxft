/**
 * @Description: project repo
 * @File: project_repo
 * @Date: 2020/7/9 0009 14:39
 */

package repo

import (
	"csxft/model"
	"fmt"
	"time"
)

type ProjectRepo interface {
	//获取插入到es的数据
	GetToEsData(id uint64) (project *model.Project, err error)
	//获取即将取证的分组
	GetPredictCredDate() (group []*model.PredictCredDate)
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

func (p projectRepo) GetPredictCredDate() (group []*model.PredictCredDate) {
	var predictCredDate []*model.PredictCredDate
	var predictCredTemp []*model.PredictCredTemp
	err := model.DB.Model(p.thisModel).Select("predict_cred_date").Where("is_will_cred = ?", 1).Group("predict_cred_date").Scan(&predictCredTemp).Error
	if err != nil {
		return nil
	} else {
		fmt.Println(predictCredTemp)
		fmt.Println(predictCredTemp[0])
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

