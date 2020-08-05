/**
 * @Description: es创建通用策略
 * @File: create_common_strategy
 * @Date: 2020/7/9 0009 14:39
 */

package es

import (
	"context"
	"csxft/elasticsearch"
	"csxft/repo"
	"fmt"
	"strconv"
)

type createHandler interface {
	Create(id uint64) (code int, msg string)
}

//基础数据
type baseHandler struct {
}

func newBaseHandler() baseHandler {
	instance := new(baseHandler)
	return *instance
}

func (b baseHandler) Create(id uint64) (code int, msg string) {

	projectObj, err := repo.NewProjectRepo().GetToEsData(id)
	if err != nil {
		code = 400
		msg = "未找到数据"
	} else {
		code = 200
		msg = "success"
	}

	_, err = elasticsearch.GetEsCli().Index().Index("project").Id(strconv.Itoa(int(id))).BodyJson(projectObj).Do(context.Background())
	if err != nil {
		code = 400
		msg = "存储失败"
	}

	return

}

//预售证数据
type credHandler struct {

}

func newCredHandler() credHandler {
	instance := new(credHandler)
	return *instance
}

func (c credHandler) Create(id uint64) (code int, msg string) {

	credObj, err := repo.NewCredRepo().GetToEsData(id)
	if err != nil {
		fmt.Println(err)
		code = 400
		msg = "未找到数据"
	} else {
		code = 200
		msg = "success"
	}

	_, err = elasticsearch.GetEsCli().Index().Index("cred").Id(strconv.Itoa(int(id))).BodyJson(credObj).Do(context.Background())
	if err != nil {
		code = 400
		msg = "存储失败"
	}
	//for _, item := range credObj {
	//	_, err = elasticsearch.GetEsCli().Index().Index("cred").Id(strconv.Itoa(int(item.ID))).BodyJson(item).Do(context.Background())
	//}

	return

}

//房屋数据
type houseHandler struct {

}

//房屋handler
func newHouseHandler() houseHandler {
	instance := new(houseHandler)
	return *instance
}

func (h houseHandler) Create(id uint64) (code int, msg string) {

	houseObj, err := repo.NewHouseRepo().GetToEsData(id)
	if err != nil {
		code = 400
		msg = "未找到数据"
	} else {
		code = 200
		msg = "success"
	}

	for _, item := range houseObj {
		_, err = elasticsearch.GetEsCli().Index().Index("house").Id(strconv.Itoa(int(item.ID))).BodyJson(item).Do(context.Background())
	}

	return

}

//摇号数据
type iotteryHandler struct {

}

//摇号handler
func newIotteryHandler() iotteryHandler {
	instance := new(iotteryHandler)
	return *instance
}

func (h iotteryHandler) Create(id uint64) (code int, msg string) {

	iotteryObj, err := repo.NewIotteryRepo().GetToEsData(id)
	if err != nil {
		code = 400
		msg = "未找到数据"
	} else {
		code = 200
		msg = "success"
	}

	_, err = elasticsearch.GetEsCli().Index().Index("iottery").Id(strconv.Itoa(int(id))).BodyJson(iotteryObj).Do(context.Background())
	if err != nil {
		code = 400
		msg = "存储失败"
	}

	return

}

//楼盘动态数据
type dynamicHandler struct {

}

func newDynamicHandler() dynamicHandler {
	instance := new(dynamicHandler)
	return *instance
}

func (c dynamicHandler) Create(id uint64) (code int, msg string) {

	dynamicObj, err := repo.NewDynamicRepo().GetToEsData(id)
	if err != nil {
		fmt.Println(err)
		code = 400
		msg = "未找到数据"
	} else {
		code = 200
		msg = "success"
	}

	_, err = elasticsearch.GetEsCli().Index().Index("dynamic").Id(strconv.Itoa(int(id))).BodyJson(dynamicObj).Do(context.Background())
	if err != nil {
		code = 400
		msg = "存储失败"
	}

	return

}

//摇号结果handler
type lotteryResultHandler struct {

}

func newLotteryResultHandler() lotteryResultHandler {
	instance := new(lotteryResultHandler)
	return *instance
}

func (c lotteryResultHandler) Create(id uint64) (code int, msg string) {

	lotteryResultObj, err := repo.NewLotteryResultRepo().GetToEsData(id)
	if err != nil {
		fmt.Println(err)
		code = 400
		msg = "未找到数据"
	} else {
		code = 200
		msg = "success"
	}

	_, err = elasticsearch.GetEsCli().Index().Index("lottery_result").Id(strconv.Itoa(int(id))).BodyJson(lotteryResultObj).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		code = 400
		msg = "存储失败"
	}

	return

}

//认筹结果handler
type solicitResultHandler struct {

}

func newSolicitResultHandler() solicitResultHandler {
	instance := new(solicitResultHandler)
	return *instance
}

func (c solicitResultHandler) Create(id uint64) (code int, msg string) {

	solicitResultObj, err := repo.NewSolicitResultRepo().GetToEsData(id)
	if err != nil {
		fmt.Println(err)
		code = 400
		msg = "未找到数据"
	} else {
		code = 200
		msg = "success"
	}

	_, err = elasticsearch.GetEsCli().Index().Index("solicit_result").Id(strconv.Itoa(int(id))).BodyJson(solicitResultObj).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		code = 400
		msg = "存储失败"
	}

	return

}

//地区handler
type areaHandler struct {

}

func newAreaHandler() areaHandler {
	instance := new(areaHandler)
	return *instance
}

func (c areaHandler) Create(id uint64) (code int, msg string) {

	areaObj, err := repo.NewAreaRepo().GetToEsData(id)
	if err != nil {
		fmt.Println(err)
		code = 400
		msg = "未找到数据"
	} else {
		code = 200
		msg = "success"
	}

	for _, item := range areaObj {
		_, err = elasticsearch.GetEsCli().Index().Index("cs_area").Id(strconv.Itoa(int(item.ID))).BodyJson(item).Do(context.Background())
	}

	return

}

//户型图handler
type houseTypeImageHandler struct {

}

func newHouseTypeImageHandler() houseTypeImageHandler {
	instance := new(houseTypeImageHandler)
	return *instance
}

func (c houseTypeImageHandler) Create(id uint64) (code int, msg string) {

	houseTypeImage, err := repo.NewHouseTypeImageRepo().GetToEsData(id)
	if err != nil {
		fmt.Println(err)
		code = 400
		msg = "未找到数据"
	} else {
		code = 200
		msg = "success"
	}

	_, err = elasticsearch.GetEsCli().Index().Index("house_type_image").Id(strconv.Itoa(int(id))).BodyJson(houseTypeImage).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		code = 400
		msg = "存储失败"
	}

	return

}

//公告handler
type noticeHandler struct {

}

func newNoticeHandler() noticeHandler {
	instance := new(noticeHandler)
	return *instance
}

func (c noticeHandler) Create(id uint64) (code int, msg string) {

	noticeObj, err := repo.NewNoticeRepo().GetToEsData(id)
	if err != nil {
		fmt.Println(err)
		code = 400
		msg = "未找到数据"
	} else {
		code = 200
		msg = "success"
	}

	_, err = elasticsearch.GetEsCli().Index().Index("notice").Id(strconv.Itoa(int(id))).BodyJson(noticeObj).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		code = 400
		msg = "存储失败"
	}

	return

}

//批次handler
type batchHandler struct {

}

func newBatchHandler() batchHandler {
	instance := new(batchHandler)
	return *instance
}

func (c batchHandler) Create(id uint64) (code int, msg string) {

	batchObj, err := repo.NewBatchRepo().GetToEsData(id)
	if err != nil {
		fmt.Println(err)
		code = 400
		msg = "未找到数据"
	} else {
		code = 200
		msg = "success"
	}

	_, err = elasticsearch.GetEsCli().Index().Index("batch").Id(strconv.Itoa(int(id))).BodyJson(batchObj).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		code = 400
		msg = "存储失败"
	}

	return

}


type CreateContext struct {
	Strategy createHandler
}

//生产策略
func NewCreateContext(insertType uint32) CreateContext {
	c := new(CreateContext)
	switch insertType {
	case 1:
		c.Strategy = newBaseHandler()  //基础数据策略
	case 2:
		c.Strategy = newCredHandler()  //预售数据策略
	case 3:
		c.Strategy = newHouseHandler()  //房屋数据策略
	case 4:
		c.Strategy = newIotteryHandler()  //摇号数据策略
	case 5:
		c.Strategy = newDynamicHandler()  //楼盘动态数据策略
	case 6:
		c.Strategy = newLotteryResultHandler()  //摇号结果数据策略
	case 7:
		c.Strategy = newSolicitResultHandler()  //认筹结果数据策略
	case 8:
		c.Strategy = newAreaHandler()  //地区
	case 9:
		c.Strategy = newHouseTypeImageHandler()  //户型图
	case 10:
		c.Strategy = newNoticeHandler()  //公告
	case 11:
		c.Strategy = newBatchHandler()  //批次
	default:
		c.Strategy = newBaseHandler()
	}
	return *c
}

//在策略生产成功后，直接调用策略的函数。
func (c CreateContext) Create(id uint64) (code int, msg string) {
	return c.Strategy.Create(id)
}

