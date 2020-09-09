package server

import (
	"csxft/api"
	"csxft/middleware"
	"github.com/gin-gonic/gin"
	"os"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.LogerMiddleware())
	//r.Use(middleware.CurrentUser())

	//初始化路由
	initData := r.Group("/api/init")
	{
		initData.GET("base", api.InitBase)
		initData.GET("cred", api.InitCred)
		initData.GET("house", api.InitHouse)
		initData.GET("fdc", api.InitFdc)
	}

	//es
	esData := r.Group("/api/es")
	{
		esData.POST("create", api.DataToEs)
		esData.GET("search_project", api.SearchProject)
		esData.GET("index", api.Index)
		esData.GET("get_project", api.GetProject)
		esData.GET("get_project_cred", api.GetProjectCred)
		//esData.GET("get_cred_house", api.GetCredHouse)
		esData.GET("get_project_detail", api.ProjectDetail)
		esData.GET("get_hot_project", api.GetHotProject)
		esData.GET("get_history_iottery", api.GetHistoryIottery)
		esData.GET("get_house", api.GetHouse)
		esData.GET("get_build_no", api.GetBuildNo)
		esData.GET("get_time_line", api.GetTimeLine)
		esData.GET("get_dynamic_count", api.GetDynamicCount)
		esData.GET("get_dynamic", api.GetDynamic)
		esData.GET("get_lottery_result", api.GetLotteryResult)
		esData.GET("get_solicit_result", api.GetSolicitResult)
		esData.GET("get_cs_area", api.GetCsArea)
		esData.GET("house_type_image_group", api.GetHouseTypeImageGroup)
		esData.GET("house_type_image", api.GetHouseTypeImage)
		esData.GET("project_detail_check", api.ProjectDetailCheck)
		esData.GET("get_notice", api.GetNotice)
		esData.POST("get_subway_house_count", api.GetSubwayProjectCount)
		esData.POST("get_subway_house", api.GetSubwayProject)
		esData.GET("get_predict_cred_date", api.GetPredictCredDate)
	}

	taskData := r.Group("/api/task")
	{
		taskData.GET("new_cred_task", api.NewCredTask)
		taskData.GET("not_new_cred_task", api.NotNewCredTask)
		taskData.GET("recognition", api.NewRecognitionTask)
		taskData.GET("not_recognition", api.NotNewRecognitionTask)
	}

	deleteData := r.Group("/api/delete")
	{
		deleteData.POST("delete_project", api.DeleteProject)
		deleteData.POST("delete_batch", api.DeleteBatch)
		deleteData.POST("delete_notice", api.DeleteNotice)
		deleteData.POST("delete_house", api.DeleteHouse)
		deleteData.POST("delete_cred_house", api.DeleteCredHouse)
		deleteData.POST("delete_lottery_result", api.DeleteLotteryResult)
		deleteData.POST("delete_batch_lottery_result", api.DeleteBatchLotteryResult)
		deleteData.POST("delete_dynamics", api.DeleteDynamics)
	}

	initAllEs := r.Group("/api/init_all")
	{
		initAllEs.GET("project", api.InitAllProject)
		initAllEs.GET("batch", api.InitAllBatch)
		initAllEs.GET("house", api.InitAllHouse)
		initAllEs.GET("dynamic", api.InitAllDynamic)
		initAllEs.GET("lottery_result", api.InitAllLotteryResult)
		initAllEs.GET("notice", api.InitAllNoticeResult)
	}

	return r
}
