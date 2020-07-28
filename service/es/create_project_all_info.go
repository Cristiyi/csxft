/**
 * @Description:
 * @File: create_project_all_info
 * @Date: 2020/7/28 0028 14:27
 */

package es

import (
	"csxft/elasticsearch"
	"csxft/repo"
	"strconv"
	"context"
)

func CreateAllProjectInfo(projectId uint64) {

	projectObj, _ := repo.NewProjectRepo().GetToEsData(projectId)
	elasticsearch.GetEsCli().Index().Index("project").Id(strconv.Itoa(int(projectId))).BodyJson(projectObj).Do(context.Background())
	credObj, _ := repo.NewCredRepo().GetByProjectId(projectId)
	for _, item := range credObj {
		elasticsearch.GetEsCli().Index().Index("cred").Id(strconv.Itoa(int(item.ID))).BodyJson(item).Do(context.Background())
		houseObj, _ := repo.NewHouseRepo().GetToEsData(uint64(item.ID))
		for _, house := range houseObj {
			elasticsearch.GetEsCli().Index().Index("house").Id(strconv.Itoa(int(item.ID))).BodyJson(house).Do(context.Background())
		}
	}

}
