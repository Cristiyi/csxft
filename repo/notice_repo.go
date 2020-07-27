/**
 * @Description:
 * @File: notice_repo
 * @Date: 2020/7/24 0024 18:00
 */

package repo

import "CMD-XuanFangTong-Server/model"

type NoticeRepo interface {
	//获取插入到es的数据
	GetToEsData(id uint64) (notice *model.Notice, err error)
}

func NewNoticeRepo() NoticeRepo {
	return &noticeRepo{}
}

type noticeRepo struct {
	thisModel model.Notice
}

func (c noticeRepo) GetToEsData(id uint64) (notice *model.Notice, err error) {
	notice = new(model.Notice)
	err = model.DB.Model(c.thisModel).Where("id = ?", id).First(&notice).Error
	return
}

