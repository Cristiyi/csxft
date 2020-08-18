/**
 * @Description:
 * @File: recognition
 * @Date: 2020/8/17 0017 18:39
 */

package task

import (
	"csxft/repo"
	"csxft/serializer"
	"fmt"
)

//最新取证任务
type RecognitionService struct {
}

func (service *RecognitionService) GetRecognitionTask() serializer.Response {

	data, err := repo.NewBatchRepo().GetRecognitionTask()
	if err != nil {
		return serializer.Response{
			Code: 400,
			Msg: "fail",
			Error: err.Error(),
		}
	}
	for _, item := range data {
		fmt.Println(item)
	}

	return serializer.Response{
		Code: 200,
		Msg: "success",
	}
}


