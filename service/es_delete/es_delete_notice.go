/**
 * @Description:
 * @File: es_delete_notice
 * @Date: 2020/8/28 0028 11:40
 */

package es_delete

import "csxft/serializer"

//删除楼盘服务
type DeleteNoticeService struct {
	NoticeId  int `form:"notice_id" json:"notice_id" binding:"required"`
}

func (service *DeleteNoticeService) DeleteNotice() serializer.Response {

	DeleteDoc(service.NoticeId, "notice")

	return serializer.Response{
		Code: 200,
		Msg: "success",
	}

}
