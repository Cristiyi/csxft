/**
 * @Description:
 * @File: solicit_result_repo
 * @Date: 2020/7/23 0023 18:06
 */

package repo

import (
	"CMD-XuanFangTong-Server/model"
	"CMD-XuanFangTong-Server/util"
)

type SolicitResultRepo interface {
	//获取插入到es的数据
	GetToEsData(id uint64) (solicitResult *model.SolicitResult, err error)
}

func NewSolicitResultRepo() SolicitResultRepo {
	return &solicitResultRepo{}
}

type solicitResultRepo struct {
	thisModel model.SolicitResult
}

func (c solicitResultRepo) GetToEsData(id uint64) (solicitResult *model.SolicitResult, err error) {
	solicitResult = new(model.SolicitResult)
	err = model.DB.Model(c.thisModel).Where("id = ?", id).First(&solicitResult).Error

	if solicitResult.IdCard != "" {
		solicitResult.IdCardBack = solicitResult.IdCard
		solicitResult.IdCard = util.HideIdCard(solicitResult.IdCard)
	}
	return
}
