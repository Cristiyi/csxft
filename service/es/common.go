package es

import (
	"context"
	"csxft/elasticsearch"
	"csxft/util"
	"fmt"
	"github.com/olivere/elastic/v7"
	"strconv"
	"strings"
)

//搜索楼盘
func QueryProject(start,size int, commonParams map[string]string, calParams map[string]float64) *elastic.SearchResult {
	sortType := true
	if commonParams["sortType"] == "desc" {
		sortType = false
	}
	searchService := elasticsearch.GetEsCli().Search("project")

	//搜索条件构建
	queryService := elastic.NewBoolQuery()
	if commonParams["name"] != "" {
		queryService.Must(elastic.NewQueryStringQuery("*Name$:" + commonParams["name"]))
	}
	if commonParams["IsWillCred"] != "" {
		queryService.Must(elastic.NewTermQuery("IsWillCred", 1))
	}
	if commonParams["IsNewCred"] != "" {
		queryService.Must(elastic.NewTermQuery("IsNewCred", 1))
	}
	if commonParams["IsRecognition"] != "" {
		queryService.Must(elastic.NewTermQuery("IsRecognition", 1))
	}
	if commonParams["IsIottery"] != "" {
		queryService.Must(elastic.NewTermQuery("IsIottery", 1))
	}
	if commonParams["IsSell"] != "" {
		queryService.Must(elastic.NewTermQuery("IsSell", 1))
	}
	if commonParams["AreaId"] != "" {
		areaArr := strings.Split(commonParams["AreaId"], ",")
		if len(areaArr) > 0 {
			areaQueryService := elastic.NewBoolQuery()
			for _, item := range areaArr {
				areaQueryService.Should(elastic.NewTermQuery("AreaId", item))
			}
			queryService.Must(areaQueryService)
		}
	}

	if calParams["MaxAcreage"] != 0 {
		maxAcreageRangeQuery := elastic.NewRangeQuery("AverageAcreage")
		maxAcreageRangeQuery.Lte(calParams["MaxAcreage"])
		queryService.Must(maxAcreageRangeQuery)
	}
	if calParams["MinAcreage"] != 0 {
		minAcreageRangeQuery := elastic.NewRangeQuery("AverageAcreage")
		minAcreageRangeQuery.Gte(calParams["MinAcreage"])
		queryService.Must(minAcreageRangeQuery)
	}
	if calParams["MaxTotalPrice"] != 0 {
		maxTotalPriceRangeQuery := elastic.NewRangeQuery("AverageTotalPrice")
		maxTotalPriceRangeQuery.Lte(calParams["MaxTotalPrice"])
		queryService.Must(maxTotalPriceRangeQuery)
	}
	if calParams["MinTotalPrice"] != 0 {
		minTotalPriceRangeQuery := elastic.NewRangeQuery("AverageTotalPrice")
		minTotalPriceRangeQuery.Gte(calParams["MinTotalPrice"])
		queryService.Must(minTotalPriceRangeQuery)
	}
	if calParams["MaxUnitPrice"] != 0 {
		maxUnitPriceRangeQuery := elastic.NewRangeQuery("AveragePrice")
		maxUnitPriceRangeQuery.Lte(calParams["MaxUnitPrice"])
		queryService.Must(maxUnitPriceRangeQuery)
	}
	if calParams["MinUnitPrice"] != 0 {
		maxUnitPriceRangeQuery := elastic.NewRangeQuery("AveragePrice")
		maxUnitPriceRangeQuery.Lte(calParams["MaxUnitPrice"])
		queryService.Must(maxUnitPriceRangeQuery)
	}
	if calParams["IsDecoration"] != 0 {
		queryService.Must(elastic.NewTermQuery("IsDecoration", 1))
	}
	if calParams["IsDecoration"]!= 0 {
		queryService.Must(elastic.NewTermQuery("IsNotDecoration", 1))
	}
	if calParams["IsNearLineOne"] != 0 {
		queryService.Must(elastic.NewTermQuery("IsNearLineOne", 1))
	}
	if calParams["IsNearLineTwo"] != 0 {
		queryService.Must(elastic.NewTermQuery("IsNearLineTwo", 1))
	}
	if calParams["IsNearLineThird"] != 0 {
		queryService.Must(elastic.NewTermQuery("IsNearLineThird", 1))
	}
	if calParams["IsNearLineFouth"] != 0 {
		queryService.Must(elastic.NewTermQuery("IsNearLineFouth", 1))
	}
	if calParams["IsNearLineFifth"] != 0 {
		queryService.Must(elastic.NewTermQuery("IsNearLineFifth", 1))
	}
	if calParams["IsNearLineSixth"] != 0 {
		queryService.Must(elastic.NewTermQuery("IsNearLineSixth", 1))
	}
	if calParams["PredictCredDate"] != 0 {
		queryService.Must(elastic.NewTermQuery("PredictCredDate", calParams["PredictCredDate"]))
	}

	searchService = searchService.Query(queryService)

	//分页构建
	if start == 1 || start == 0 {
		start = 0
	} else {
		start = (start-1)*size
	}
	searchResult, err := searchService.
		Sort(commonParams["sort"], sortType).
		From(start).Size(size).
		Pretty(true).
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return searchResult
}

//查询楼盘数量
func QueryProjectCount(commonParams map[string]string) (count int64) {
	searchService := elasticsearch.GetEsCli().Count("project")
	switch commonParams["type"] {
	case "1":
		searchService = searchService.Query(elastic.NewTermQuery("IsWillCred", 1))
		break
	case "2":
		searchService = searchService.Query(elastic.NewTermQuery("IsNewCred", 1))
		break
	case "3":
		searchService = searchService.Query(elastic.NewTermQuery("IsRecognition", 1))
		break
	case "4":
		searchService = searchService.Query(elastic.NewTermQuery("IsIottery", 1))
		break
	case "5":
		searchService = searchService.Query(elastic.NewTermQuery("IsSell", 1))
		break
	default:
		searchService = searchService.Query(elastic.NewTermQuery("IsSell", 1))
		break
	}
	if commonParams["areaId"] != "" {
		areaId, err := strconv.Atoi(commonParams["areaId"])
		fmt.Println(err)
		fmt.Println(areaId)
		searchService = searchService.Query(elastic.NewTermQuery("AreaId", areaId))
	}
	count, err := searchService.Do(context.Background())
	if err != nil {
		fmt.Println(err)
		count = 0
	}
	fmt.Println(count)
	return
}

//查询楼盘数量
func QueryProjectAreaCount(areaId uint) (count int64) {
	searchService := elasticsearch.GetEsCli().Count("project")
	searchService = searchService.Query(elastic.NewTermQuery("AreaId", areaId))
	count, err := searchService.Do(context.Background())
	if err != nil {
		fmt.Println(err)
		count = 0
	}
	fmt.Println(count)
	return
}

//获取楼盘
func GetProject(id string) *elastic.GetResult {
	res, err := elasticsearch.GetEsCli().Get().Index("project").Id(id).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return res
}

//获取开盘
func GetCred(id string) *elastic.GetResult {
	res, err := elasticsearch.GetEsCli().Get().Index("cred").Id(id).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return res
}

//根据楼盘获取开盘信息
func GetProjectCred (start,size int, commonParams map[string]string) *elastic.SearchResult {

	sortType := true
	if commonParams["sortType"] == "desc" {
		sortType = false
	}

	//搜索条件构建
	searchService := elasticsearch.GetEsCli().Search("cred")
	queryService := elastic.NewBoolQuery()
	queryService.Must(elastic.NewTermQuery("ProjectId", commonParams["ProjectId"]))
	if commonParams["IsNew"] != "" {
		queryService.Must(elastic.NewQueryStringQuery("IsNew:"+"1"))
	}
	searchService = searchService.Query(queryService)

	//分页构建
	if start == 1 || start == 0 {
		start = 0
	} else {
		start = (start-1)*size
	}
	searchResult, err := searchService.
		Sort(commonParams["sort"], sortType).
		From(start).Size(size).
		Pretty(true).
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return searchResult
}

//根据开盘信息获取楼盘信息
func GetCredHouse (start,size int, commonParams map[string]string, credIds []int) *elastic.SearchResult {

	var sortType bool
	if commonParams["sortType"] == "desc" {
		sortType = false
	} else {
		sortType = true
	}
	searchService := elasticsearch.GetEsCli().Search("house")

	//预售证id or查询
	queryService := elastic.NewBoolQuery()
	if commonParams["BuildNo"] != "" {
		credQueryService := elastic.NewBoolQuery()
		for _, item := range credIds {
			credQueryService.Should(elastic.NewTermQuery("CredId", item))
		}
		queryService.Must(elastic.NewQueryStringQuery("BuildNo.keyword:"+commonParams["BuildNo"]), credQueryService)
	} else {
		for _, item := range credIds {
			queryService.Should(elastic.NewTermQuery("CredId", item))
		}
	}
	searchService = searchService.Query(queryService)
	if start == 1 || start == 0 {
		start = 0
	} else {
		start = (start-1)*size
	}
	searchResult, err := searchService.
		Sort(commonParams["sort"], sortType).
		TypedKeys(true).
		From(start).Size(size).
		Pretty(true).
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return searchResult
}

//根据楼盘获取摇号
//func GetProjectIottery (start,size int, commonParams map[string]string) *elastic.SearchResult {
//	sortType := true
//	if commonParams["sortType"] == "desc" {
//		sortType = false
//	}
//	searchService := elasticsearch.GetEsCli().Search("iottery")
//	searchService = searchService.Query(elastic.NewTermQuery("ProjectId", commonParams["ProjectId"]))
//	if start == 1 || start == 0 {
//		start = 0
//	} else {
//		start = (start-1)*size
//	}
//	searchResult, err := searchService.
//		Sort(commonParams["sort"], sortType).
//		From(start).Size(size).
//		Pretty(true).
//		Do(context.Background())
//	if err != nil {
//		fmt.Println(err)
//		return nil
//	}
//	return searchResult
//}

//查询楼盘动态数量
func QueryDynamicCount(commonParams map[string]string) (count int64) {
	searchService := elasticsearch.GetEsCli().Count("dynamic")
	queryService := elastic.NewBoolQuery()
	queryService.Must(elastic.NewTermQuery("ProjectId",commonParams["ProjectId"]))
	switch commonParams["type"] {
	case "1":
		queryService.Must(elastic.NewQueryStringQuery("Type:"+"1"))
		break
	case "2":
		queryService.Must(elastic.NewQueryStringQuery("Type:"+"2"))
		break
	case "3":
		queryService.Must(elastic.NewQueryStringQuery("Type:"+"3"))
		break
	default:
		break
	}
	searchService = searchService.Query(queryService)
	count, err := searchService.Do(context.Background())
	if err != nil {
		fmt.Println(err)
		count = 0
	}
	return
}

//获取楼盘动态
func QueryDynamic(start,size int, commonParams map[string]string) *elastic.SearchResult {
	sortType := true
	if commonParams["sortType"] == "desc" {
		sortType = false
	}
	searchService := elasticsearch.GetEsCli().Search("dynamic")

	queryService := elastic.NewBoolQuery()
	queryService.Must(elastic.NewTermQuery("ProjectId",commonParams["ProjectId"]))
	switch commonParams["type"] {
	case "1":
		queryService.Must(elastic.NewQueryStringQuery("Type:"+"1"))
		break
	case "2":
		queryService.Must(elastic.NewQueryStringQuery("Type:"+"2"))
		break
	case "3":
		queryService.Must(elastic.NewQueryStringQuery("Type:"+"3"))
		break
	default:
		break
	}
	searchService = searchService.Query(queryService)

	//分页构建
	if start == 1 || start == 0 {
		start = 0
	} else {
		start = (start-1)*size
	}
	searchResult, err := searchService.
		Sort(commonParams["sort"], sortType).
		From(start).Size(size).
		Pretty(true).
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return searchResult
}

//获取摇号结果
func QueryLotteryResult(start,size int, commonParams map[string]string) *elastic.SearchResult {
	sortType := true
	if commonParams["sortType"] == "desc" {
		sortType = false
	}
	searchService := elasticsearch.GetEsCli().Search("lottery_result")

	queryService := elastic.NewBoolQuery()
	queryService.Must(elastic.NewTermQuery("ProjectId",commonParams["ProjectId"]))
	if commonParams["BatchId"] != "" {
		queryService.Must(elastic.NewTermQuery("BatchId", commonParams["BatchId"]))
	}
	if commonParams["Search"] != "" {
		searchQuery := elastic.NewBoolQuery()
		searchQuery.Should(elastic.NewTermQuery("No",commonParams["Search"]), elastic.NewTermQuery("IdCardBack",commonParams["Search"]), elastic.NewTermQuery("SolicitNo",commonParams["Search"]))
		queryService.Must(searchQuery)
	}
	searchService = searchService.Query(queryService)

	//分页构建
	if start == 1 || start == 0 {
		start = 0
	} else {
		start = (start-1)*size
	}
	searchResult, err := searchService.
		Sort(commonParams["sort"], sortType).
		From(start).Size(size).
		Pretty(true).
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return searchResult
}

//获取认筹结果
func QuerySolicitResult(start,size int, commonParams map[string]string) *elastic.SearchResult {
	sortType := true
	if commonParams["sortType"] == "desc" {
		sortType = false
	}
	searchService := elasticsearch.GetEsCli().Search("solicit_result")

	queryService := elastic.NewBoolQuery()
	queryService.Must(elastic.NewTermQuery("ProjectId", commonParams["ProjectId"]))
	if commonParams["BatchId"] != "" {
		queryService.Must(elastic.NewTermQuery("BatchId", commonParams["BatchId"]))
	}
	if commonParams["Search"] != "" {
		searchQuery := elastic.NewBoolQuery()
		searchQuery.Should(elastic.NewTermQuery("IdCardBack",commonParams["Search"]), elastic.NewTermQuery("SolicitNo",commonParams["Search"]))
		queryService.Must(searchQuery)
	}
	searchService = searchService.Query(queryService)

	//分页构建
	if start == 1 || start == 0 {
		start = 0
	} else {
		start = (start-1)*size
	}
	searchResult, err := searchService.
		Sort(commonParams["sort"], sortType).
		From(start).Size(size).
		Pretty(true).
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return searchResult
}

//获取长沙地区
func GetCsArea() *elastic.SearchResult {
	searchService := elasticsearch.GetEsCli().Search("cs_area")
	searchResult, err := searchService.
		Sort("ID", true).
		Pretty(true).
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return searchResult
}

//获取户型图
func QueryHouseTypeImage(start,size int, commonParams map[string]string) *elastic.SearchResult {
	sortType := true
	if commonParams["sortType"] == "desc" {
		sortType = false
	}
	searchService := elasticsearch.GetEsCli().Search("house_type_image")

	queryService := elastic.NewBoolQuery()
	queryService.Must(elastic.NewTermQuery("ProjectId", commonParams["ProjectId"]))
	if commonParams["BatchId"] != "" {
		queryService.Must(elastic.NewTermQuery("BatchId", commonParams["BatchId"]))
	}
	if commonParams["HomeNum"] != "" {
		queryService.Must(elastic.NewTermQuery("HomeNum.keyword", commonParams["HomeNum"]))
	}
	searchService = searchService.Query(queryService)

	//分页构建
	if start == 1 || start == 0 {
		start = 0
	} else {
		start = (start-1)*size
	}
	searchResult, err := searchService.
		Sort(commonParams["sort"], sortType).
		From(start).Size(size).
		Pretty(true).
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return searchResult
}

//查询户型图数量
func HouseTypeImageCount(commonParams map[string]string) (count int64) {
	searchService := elasticsearch.GetEsCli().Count("house_type_image")
	queryService := elastic.NewBoolQuery()
	queryService.Must(elastic.NewTermQuery("ProjectId", commonParams["ProjectId"]))
	if commonParams["BatchId"] != "" {
		queryService.Must(elastic.NewTermQuery("BatchId", commonParams["BatchId"]))
	}
	if commonParams["HomeNum"] != "" {
		queryService.Must(elastic.NewTermQuery("HomeNum.keyword", commonParams["HomeNum"]))
	}
	searchService = searchService.Query(queryService)
	count, err := searchService.Do(context.Background())
	if err != nil {
		fmt.Println(err)
		count = 0
	}
	return
}

//获取公告
func GetNotice(commonParams map[string]string) *elastic.SearchResult {
	searchService := elasticsearch.GetEsCli().Search("notice")
	queryService := elastic.NewBoolQuery()
	queryService.Must(elastic.NewTermQuery("ProjectId", commonParams["ProjectId"]))
	queryService.Must(elastic.NewTermQuery("Status", 1))
	if commonParams["BatchId"] != "" {
		queryService.Must(elastic.NewTermQuery("BatchId", commonParams["BatchId"]))
	}
	if commonParams["NoticeType"] != "" {
		queryService.Must(elastic.NewTermQuery("NoticeType", commonParams["NoticeType"]))
	}
	searchService = searchService.Query(queryService)
	searchResult, err := searchService.
		Sort("UpdatedAt", true).
		From(0).Size(5).
		Pretty(true).
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return searchResult
}

//获取批次
func GetBatch(projectId string, status int32) *elastic.SearchResult {
	searchService := elasticsearch.GetEsCli().Search("batch")
	queryService := elastic.NewBoolQuery()
	queryService.Must(elastic.NewTermQuery("ProjectId", projectId))
	if status != 0 {
		queryService.Must(elastic.NewTermQuery("Status", status))
	}
	searchService = searchService.Query(queryService)
	searchResult, err := searchService.
		Sort("BatchNo", false).
		From(0).Size(1).
		Pretty(true).
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return searchResult
}

//根据经纬度范围获取楼盘数量
func GetProjectCountByPoint(pointRange util.PointRange) (count int64) {
	searchService := elasticsearch.GetEsCli().Count("project")
	queryService := elastic.NewBoolQuery()
	longRangeQuery := elastic.NewRangeQuery("Longitude")
	longRangeQuery.Gte(pointRange.MinLng)
	longRangeQuery.Lte(pointRange.MaxLng)
	latRangeQuery := elastic.NewRangeQuery("Latitude")
	latRangeQuery.Gte(pointRange.MinLat)
	latRangeQuery.Lte(pointRange.MaxLat)
	queryService.Must(longRangeQuery, latRangeQuery)
	searchService = searchService.Query(queryService)
	count, err := searchService.Do(context.Background())
	if err != nil {
		fmt.Println(err)
		count = 0
	}
	return
}


//根据经纬度范围获取楼盘数量
func GetProjectByPoint(pointRange util.PointRange) *elastic.SearchResult {

	searchService := elasticsearch.GetEsCli().Search("project")
	queryService := elastic.NewBoolQuery()
	longRangeQuery := elastic.NewRangeQuery("Longitude")
	longRangeQuery.Gte(pointRange.MinLng)
	longRangeQuery.Lte(pointRange.MaxLng)
	latRangeQuery := elastic.NewRangeQuery("Latitude")
	latRangeQuery.Gte(pointRange.MinLat)
	latRangeQuery.Lte(pointRange.MaxLat)
	queryService.Must(longRangeQuery, latRangeQuery)

	searchService = searchService.Query(queryService)

	searchResult, err := searchService.
		Sort("CreatedAt", false).
		Pretty(true).
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return searchResult
}

//根据楼盘获取摇号
func GetProjectIottery(start, size, batchNo, projectId int, commonParams map[string]string) *elastic.SearchResult {
	sortType := true
	if commonParams["sortType"] == "desc" {
		sortType = false
	}
	//批次号小于传入批次号的定义为历史摇号
	searchService := elasticsearch.GetEsCli().Search("batch")
	queryService := elastic.NewBoolQuery()
	batchNoQuery := elastic.NewRangeQuery("BatchNo")
	batchNoQuery.Lt(batchNo)
	projectIdQuery := elastic.NewTermQuery("ProjectId", projectId)
	queryService.Must(batchNoQuery, projectIdQuery)

	searchService = searchService.Query(queryService)
	if start == 1 || start == 0 {
		start = 0
	} else {
		start = (start-1)*size
	}
	searchResult, err := searchService.
		Sort(commonParams["sort"], sortType).
		From(start).Size(size).
		Pretty(true).
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return searchResult
}

//根据id获取批次
func GetBatchById(batchId int) *elastic.SearchResult {
	searchService := elasticsearch.GetEsCli().Search("batch")
	queryService := elastic.NewBoolQuery()
	queryService.Must(elastic.NewTermQuery("ID", batchId))
	searchService = searchService.Query(queryService)
	searchResult, err := searchService.
		Sort("BatchNo", false).
		From(0).Size(1).
		Pretty(true).
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return searchResult
}