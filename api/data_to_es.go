package api

import (
	service "csxft/service/es"
	"fmt"
	"github.com/gin-gonic/gin"
)

//初始化基础数据
func DataToEs(c *gin.Context) {
	var service service.CreateService
	if err := c.ShouldBind(&service); err == nil {
		res := service.CommonCreate()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

//搜索楼盘
func SearchProject(c *gin.Context) {
	var service service.SearchProjectService
	if err := c.ShouldBind(&service); err == nil {
		res := service.SearchProject()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

//首页数据
func Index(c *gin.Context) {
	var service service.IndexService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Index()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

//根据id获取楼盘
func GetProject(c *gin.Context) {
	var service service.GetProjectService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetProject()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

//根据楼盘id获取开盘信息
func GetProjectCred(c *gin.Context) {
	var service service.GetProjectCredService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetProjectCred()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

//根据楼盘id获取开盘信息
//func GetCredHouse(c *gin.Context) {
//	var service service.GetCredHouseService
//	if err := c.ShouldBind(&service); err == nil {
//		res := service.GetCredHouse()
//		c.JSON(200, res)
//	} else {
//		c.JSON(400, ErrorResponse(err))
//	}
//}

//根据楼盘id获取最新开盘信息
func ProjectDetail(c *gin.Context) {
	var service service.ProjectDetailService
	if err := c.ShouldBind(&service); err == nil {
		res := service.ProjectDetail()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

//获取热门搜索
func GetHotProject(c *gin.Context) {
	var service service.GetHotService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetHot()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

//获取历史摇号
func GetHistoryIottery(c *gin.Context) {
	var service service.HistoryIotteryService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetHistoryIottery()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

//获取一房一价
func GetHouse(c *gin.Context) {
	var service service.NewCredHouseService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetNewCredHouse()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

//获取楼栋
func GetBuildNo(c *gin.Context) {
	var service service.AllBuildNoService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetAllBuildNo()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

//获取时间轴
func GetTimeLine(c *gin.Context) {
	var service service.TimeLineService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetTimeLine()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

//获取楼盘动态数量
func GetDynamicCount(c *gin.Context) {
	var service service.DynamicCountService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetDynamicCount()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

//获取楼盘动态
func GetDynamic(c *gin.Context) {
	var service service.DynamicService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetDynamic()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

//获取摇号结果
func GetLotteryResult(c *gin.Context) {
	var service service.LotteryResultService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetLotteryResult()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

//获取认筹结果
func GetSolicitResult(c *gin.Context) {
	var service service.SolicitResultService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetSolicitResult()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

//获取长沙地区
func GetCsArea(c *gin.Context) {
	var service service.CsAreaService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetCsArea()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

//获取户型图分组
func GetHouseTypeImageGroup(c *gin.Context) {
	var service service.GetHouseTypeImageNumService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetGroup()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

//获取户型图分组
func GetHouseTypeImage(c *gin.Context) {
	var service service.GetHouseTypeImageService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetHouseTypeImage()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

//详情页数据检测
func ProjectDetailCheck(c *gin.Context) {
	var service service.DetailCheckService
	if err := c.ShouldBind(&service); err == nil {
		res := service.DetailCheck()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

//获取公告
func GetNotice(c *gin.Context) {
	var service service.GetNoticeService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetNotice()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

//获取公告
func GetPredictCredDate(c *gin.Context) {
	var service service.GetPredictCredDate
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetPredictCredDate()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

////获取公告
//func GetNotice(c *gin.Context) {
//	var service service.GetNoticeService
//	if err := c.ShouldBind(&service); err == nil {
//		res := service.GetNotice()
//		c.JSON(200, res)
//	} else {
//		c.JSON(400, ErrorResponse(err))
//	}
//}

//猜你喜欢
func GetRecommendProject(c *gin.Context) {
	var service service.RecommendProjectService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetRecommendProject()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

//楼盘所有批次
func GetProjectBatch(c *gin.Context) {
	var service service.ProjectBatchService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetProjectBatch()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

//获取楼盘下所有的的一房一价
func GetProjectHouse(c *gin.Context) {
	var service service.ProjectHouseService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetProjectHouse()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

//生成一房一价图
func GenHouseImage(c *gin.Context) {
	var service service.GenHouseImageService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GenHouseImage()
		fmt.Println(res)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}


