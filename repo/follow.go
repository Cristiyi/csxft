/**
 * @Description:
 * @File: follow
 * @Date: 2020/7/27 0027 12:14
 */

package repo

import "csxft/model"

type FollowRepo interface {
	//获取插入到es的数据
	Get(userId uint64, projectId uint64) (follow *model.Follow, err error)
	GetCount(projectId uint64) (count int)
}

func NewFollowRepo() FollowRepo {
	return &followRepo{}
}

type followRepo struct {
	thisModel model.Follow
}

func (c followRepo) Get(userId uint64, projectId uint64) (follow *model.Follow, err error) {
	follow = new(model.Follow)
	err = model.DB.Model(c.thisModel).Where("user_id = ? and build_id = ? and type = ?", userId, projectId, 1).First(&follow).Error
	return
}

func (c followRepo) GetCount(projectId uint64) (count int) {
	model.DB.Model(c.thisModel).Where("build_id = ? and type = ? and status = ?", projectId, 1, 1).Count(&count)
	return
}
