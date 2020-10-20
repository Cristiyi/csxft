/**
 * @Description:
 * @File: solicit_result_repo
 * @Date: 2020/7/23 0023 18:06
 */

package repo

import (
	"csxft/model"
)

type SolicitResultRepo interface {
	//获取插入到es的数据
	GetToEsData(id uint64) (solicitResult []*model.SolicitResult, err error)
}

func NewSolicitResultRepo() SolicitResultRepo {
	return &solicitResultRepo{}
}

type solicitResultRepo struct {
	thisModel model.SolicitResult
}

func (c solicitResultRepo) GetToEsData(id uint64) (solicitResult []*model.SolicitResult, err error) {
	err = model.DB.Model(c.thisModel).Where("batch_id = ?", id).Find(&solicitResult).Error
	for i, _ := range solicitResult {
		solicitResult[i].IdCardBack = solicitResult[i].IdCard
		//solicitResult[i].IdCard = util.HideIdCard(solicitResult[i].IdCard)
	}
	return
}
