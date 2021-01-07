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
	searchService := elasticsearch.GetEsCli().Search("project").SearchType("dfs_query_then_fetch")

	//搜索条件构建
	queryService := elastic.NewBoolQuery()
	if commonParams["name"] != "" {
		//queryService.Must(elastic.NewQueryStringQuery("*Name:" + commonParams["name"]))
		nameQueryService := elastic.NewBoolQuery()
		nameQueryService.Should(elastic.NewQueryStringQuery("ProjectName:"+commonParams["name"]))
		nameQueryService.Should(elastic.NewQueryStringQuery("PromotionFirstName:"+commonParams["name"]))
		nameQueryService.Should(elastic.NewQueryStringQuery("PromotionSecondName:"+commonParams["name"]))
		//nameQueryService.Should(elastic.NewMatchPhrasePrefixQuery("ProjectName", commonParams["name"]).MaxExpansions(50).Slop(0).Boost(0.1))
		//nameQueryService.Should(elastic.NewMatchPhrasePrefixQuery("PromotionFirstName", commonParams["name"]).MaxExpansions(50).Slop(0).Boost(0.1))
		//nameQueryService.Should(elastic.NewMatchPhrasePrefixQuery("PromotionSecondName", commonParams["name"]).MaxExpansions(50).Slop(0).Boost(0.1))

		//nameQueryService.Should(elastic.NewMatchPhrasePrefixQuery("ProjectName", commonParams["name"]).Boost(0.3))
		//nameQueryService.Should(elastic.NewMatchPhrasePrefixQuery("PromotionFirstName", commonParams["name"]).Boost(0.3))
		//nameQueryService.Should(elastic.NewMatchPhrasePrefixQuery("PromotionSecondName", commonParams["name"]).Boost(0.3))

		//nameQueryService.Should(elastic.NewMatchQuery("ProjectName", commonParams["name"]).Analyzer("ngramSearchAnalyzer").Boost(5))
		//nameQueryService.Should(elastic.NewMatchQuery("PromotionFirstName", commonParams["name"]).Analyzer("ngramSearchAnalyzer").Boost(5))
		//nameQueryService.Should(elastic.NewMatchQuery("PromotionSecondName", commonParams["name"]).Analyzer("ngramSearchAnalyzer").Boost(5))

		queryService.Must(nameQueryService)
	}

	//全部楼盘筛选
	if commonParams["IsAll"] != "" {
		isAllService := elastic.NewBoolQuery()
		isAllService.Should(elastic.NewTermQuery("NoStatus", 1))
		isAllService.Should(elastic.NewTermQuery("NoStatus", 4))
		isAllService.Should(elastic.NewTermQuery("NoStatus", 10))
		queryService.Must(isAllService)
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
	if commonParams["HasAerialUpload"] != "" {
		HasAerialUploadService := elastic.NewBoolQuery()
		HasAerialUploadService.Must(elastic.NewExistsQuery("AerialMainImages"))
		queryService.Must(HasAerialUploadService)
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
		minUnitPriceRangeQuery := elastic.NewRangeQuery("AveragePrice")
		minUnitPriceRangeQuery.Lte(calParams["MaxUnitPrice"])
		queryService.Must(minUnitPriceRangeQuery)
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
	//if commonParams["IsAll"] != "" {
	//	searchService.Sort("NoStatus", false)
	//}
	//searchResult, err := searchService.
	//	Sort("_score", false).
	//	Sort(commonParams["sort"], sortType).
	//	From(start).Size(size).
	//	Pretty(true).
	//	Do(context.Background())
	if commonParams["IsAll"] != "" {
		searchService.Sort("NoStatus", true)
	}
	searchService.Sort("_score", false).Sort(commonParams["sort"], sortType)
	searchResult, err := searchService.
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
		isAllService := elastic.NewBoolQuery()
		isAllService.Should(elastic.NewTermQuery("NoStatus", 1))
		isAllService.Should(elastic.NewTermQuery("NoStatus", 4))
		isAllService.Should(elastic.NewTermQuery("NoStatus", 10))
		searchService.Query(isAllService)
		break
	}
	if commonParams["areaId"] != "" {
		areaId, err := strconv.Atoi(commonParams["areaId"])
		fmt.Println(err)
		searchService = searchService.Query(elastic.NewTermQuery("AreaId", areaId))
	}
	count, err := searchService.Do(context.Background())
	if err != nil {
		fmt.Println(err)
		count = 0
	}
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

//根据开盘信息获取一房一价
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
		//queryService.Must(elastic.NewMatchQuery("BuildNo", commonParams["BuildNo"]), credQueryService)
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
	if commonParams["Type"] != "" {
		queryService.Must(elastic.NewTermQuery("Type", commonParams["Type"]))
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
	if commonParams["Type"] != "" {
		queryService.Must(elastic.NewTermQuery("Type", commonParams["Type"]))
	}
	if commonParams["ContentType"] != "" {
		queryService.Must(elastic.NewTermQuery("ContentType", commonParams["ContentType"]))
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
	//queryService.Must(elastic.NewTermQuery("ProjectId", commonParams["ProjectId"]))
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

//获取批次
func GetBatchAll(projectId string, status int32) *elastic.SearchResult {
	searchService := elasticsearch.GetEsCli().Search("batch")
	queryService := elastic.NewBoolQuery()
	queryService.Must(elastic.NewTermQuery("ProjectId", projectId))
	if status != 0 {
		queryService.Must(elastic.NewTermQuery("Status", status))
	}
	searchService = searchService.Query(queryService)
	searchResult, err := searchService.
		Sort("BatchNo", false).
		From(0).Size(100).
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


//根据坐标获取周边楼盘（地铁用）
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

//猜你喜欢
func GetRecommendProject(projectId string, pointRange util.PointRange) *elastic.SearchResult {

	searchService := elasticsearch.GetEsCli().Search("project")
	queryService := elastic.NewBoolQuery()

	queryService.MustNot(elastic.NewTermQuery("ID", projectId))

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
		From(0).Size(3).
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

//获取一房一价图
func GetHouseImage(commonParams map[string]string) *elastic.SearchResult {
	searchService := elasticsearch.GetEsCli().Search("house_image")
	queryService := elastic.NewBoolQuery()
	//queryService.Must(elastic.NewTermQuery("ProjectId", commonParams["ProjectId"]))
	queryService.Must(elastic.NewTermQuery("BatchId", commonParams["BatchId"]))
	searchService = searchService.Query(queryService)
	searchResult, err := searchService.
		Sort("UpdatedAt", true).
		From(0).Size(20).
		Pretty(true).
		Do(context.Background())
	if err != nil {
		return nil
	}
	return searchResult
}

//根据批次信息搜索楼盘
func QueryBatchProject(start,size int, commonParams map[string]string, calParams map[string]float64) *elastic.SearchResult {
	sortType := true
	if commonParams["sortType"] == "desc" {
		sortType = false
	}
	searchService := elasticsearch.GetEsCli().Search("batch_project").SearchType("dfs_query_then_fetch")

	//搜索条件构建
	queryService := elastic.NewBoolQuery()
	queryService.Must(elastic.NewTermQuery("SaleStatus", 1))
	if commonParams["name"] != "" {
		nameQueryService := elastic.NewBoolQuery()
		nameQueryService.Should(elastic.NewQueryStringQuery("Project.ProjectName:"+commonParams["name"]))
		nameQueryService.Should(elastic.NewQueryStringQuery("Project.PromotionFirstName:"+commonParams["name"]))
		nameQueryService.Should(elastic.NewQueryStringQuery("Project.PromotionSecondName:"+commonParams["name"]))
		queryService.Must(nameQueryService)
	}

	if commonParams["IsWillCred"] != "" {
		queryService.Must(elastic.NewTermQuery("Project.IsWillCred", 1))
	}
	if commonParams["IsNewCred"] != "" {
		queryService.Must(elastic.NewTermQuery("Project.IsNewCred", 1))
	}
	if commonParams["IsRecognition"] != "" {
		queryService.Must(elastic.NewTermQuery("Project.IsRecognition", 1))
	}
	if commonParams["IsIottery"] != "" {
		queryService.Must(elastic.NewTermQuery("Project.IsIottery", 1))
	}
	if commonParams["IsSell"] != "" {
		queryService.Must(elastic.NewTermQuery("Project.IsSell", 1))
	}
	if commonParams["AreaId"] != "" {
		areaArr := strings.Split(commonParams["AreaId"], ",")
		if len(areaArr) > 0 {
			areaQueryService := elastic.NewBoolQuery()
			for _, item := range areaArr {
				areaQueryService.Should(elastic.NewTermQuery("Project.AreaId", item))
			}
			queryService.Must(areaQueryService)
		}
	}
	if commonParams["Renovation"] != "" {
		queryService.Must(elastic.NewTermQuery("Renovation", commonParams["Renovation"]))
	}

	if commonParams["HasAerialUpload"] != "" {
		HasAerialUploadService := elastic.NewBoolQuery()
		HasAerialUploadService.Must(elastic.NewExistsQuery("Project.AerialMainImages"))
		queryService.Must(HasAerialUploadService)
	}

	//if calParams["MaxAcreage"] != 0 {
	//	maxAcreageRangeQuery := elastic.NewRangeQuery("MaxAcreage")
	//	maxAcreageRangeQuery.Lte(calParams["MaxArea"])
	//	queryService.Must(maxAcreageRangeQuery)
	//}
	//if calParams["MinAcreage"] != 0 {
	//	minAcreageRangeQuery := elastic.NewRangeQuery("MinArea")
	//	minAcreageRangeQuery.Gte(calParams["MinArea"])
	//	queryService.Must(minAcreageRangeQuery)
	//}
	//if calParams["MaxTotalPrice"] != 0 {
	//	maxTotalPriceRangeQuery := elastic.NewRangeQuery("MaxTotalPrice")
	//	maxTotalPriceRangeQuery.Lte(calParams["MaxTotalPrice"])
	//	queryService.Must(maxTotalPriceRangeQuery)
	//}
	//if calParams["MinTotalPrice"] != 0 {
	//	minTotalPriceRangeQuery := elastic.NewRangeQuery("MinTotalPrice")
	//	minTotalPriceRangeQuery.Gte(calParams["MinTotalPrice"])
	//	queryService.Must(minTotalPriceRangeQuery)
	//}
	//if calParams["MaxPrice"] != 0 {
	//	maxUnitPriceRangeQuery := elastic.NewRangeQuery("MaxPrice")
	//	maxUnitPriceRangeQuery.Lte(calParams["MaxPrice"])
	//	queryService.Must(maxUnitPriceRangeQuery)
	//}
	//if calParams["MinPrice"] != 0 {
	//	minUnitPriceRangeQuery := elastic.NewRangeQuery("MinPrice")
	//	minUnitPriceRangeQuery.Gte(calParams["MinPrice"])
	//	queryService.Must(minUnitPriceRangeQuery)
	//}
	//总价筛选
	if calParams["MaxTotalPrice"] != 0 {

		queryService.MustNot(elastic.NewTermQuery("MinTotalPrice", 0))
		queryService.MustNot(elastic.NewTermQuery("MaxTotalPrice", 0))

		//筛选区间： A(最小总价) B(最大总价)
		//楼盘实际价格（为丙）：甲（最小总价） 乙（最大总价）
		//条件： A<=甲&&B>=甲 || A>甲&&A<=乙

		maxTotalPriceQuery := elastic.NewBoolQuery()
		maxTotalPriceQueryOne := elastic.NewBoolQuery()
		maxTotalPriceQueryTwo := elastic.NewBoolQuery()

		//A<=甲&&B>=甲
		maxTotalPriceRangeOne := elastic.NewRangeQuery("MinTotalPrice")
		maxTotalPriceRangeOne.Gte(calParams["MinTotalPrice"])
		maxTotalPriceRangeTwo := elastic.NewRangeQuery("MinTotalPrice")
		maxTotalPriceRangeTwo.Lte(calParams["MaxTotalPrice"])
		maxTotalPriceQueryOne.Must(maxTotalPriceRangeOne)
		maxTotalPriceQueryOne.Must(maxTotalPriceRangeTwo)

		//A>甲&&A<=乙
		maxTotalPriceRangeThree := elastic.NewRangeQuery("MinTotalPrice")
		maxTotalPriceRangeThree.Lt(calParams["MinTotalPrice"])
		maxTotalPriceRangeFour := elastic.NewRangeQuery("MaxTotalPrice")
		maxTotalPriceRangeFour.Gte(calParams["MinTotalPrice"])
		maxTotalPriceQueryTwo.Must(maxTotalPriceRangeThree)
		maxTotalPriceQueryTwo.Must(maxTotalPriceRangeFour)

		maxTotalPriceQuery.Should(maxTotalPriceQueryOne)
		maxTotalPriceQuery.Should(maxTotalPriceQueryTwo)

		queryService.Must(maxTotalPriceQuery)
	}

	//单价筛选
	if calParams["MaxPrice"] != 0 {
		//筛选区间： A(最小单价) B(最大单价)
		//楼盘实际价格（为丙）：甲（最小单价） 乙（最大单价）
		//条件： A<=甲&&B>=甲 || A>甲&&A<=乙

		maxPriceQuery := elastic.NewBoolQuery()
		maxPriceQueryOne := elastic.NewBoolQuery()
		maxPriceQueryTwo := elastic.NewBoolQuery()

		//A<=甲&&B>=甲
		maxPriceRangeOne := elastic.NewRangeQuery("MinPrice")
		maxPriceRangeOne.Gte(calParams["MinPrice"])
		maxPriceRangeTwo := elastic.NewRangeQuery("MinPrice")
		maxPriceRangeTwo.Lte(calParams["MaxPrice"])
		maxPriceQueryOne.Must(maxPriceRangeOne)
		maxPriceQueryOne.Must(maxPriceRangeTwo)

		//A>甲&&A<=乙
		maxPriceRangeThree := elastic.NewRangeQuery("MinPrice")
		maxPriceRangeThree.Lt(calParams["MinPrice"])
		maxPriceRangeFour := elastic.NewRangeQuery("MaxPrice")
		maxPriceRangeFour.Gte(calParams["MinPrice"])
		maxPriceQueryTwo.Must(maxPriceRangeThree)
		maxPriceQueryTwo.Must(maxPriceRangeFour)

		maxPriceQuery.Should(maxPriceQueryOne)
		maxPriceQuery.Should(maxPriceQueryTwo)

		queryService.MustNot(elastic.NewTermQuery("MinPrice", 0))
		queryService.MustNot(elastic.NewTermQuery("MaxPrice", 0))

		queryService.Must(maxPriceQuery)
	}

	//面积筛选
	if calParams["MaxAcreage"] != 0 {
		//筛选区间： A(最小面积) B(最大面积)
		//楼盘实际面积（为丙）：甲（最小面积） 乙（最大面积）
		//条件： A<=甲&&B>=甲 || A>甲&&A<=乙

		maxAcreageQuery := elastic.NewBoolQuery()
		maxAcreageQueryOne := elastic.NewBoolQuery()
		maxAcreageQueryTwo := elastic.NewBoolQuery()

		//A<=甲&&B>=甲
		maxAcreageRangeOne := elastic.NewRangeQuery("MinArea")
		maxAcreageRangeOne.Gte(calParams["MinAcreage"])
		maxAcreageRangeTwo := elastic.NewRangeQuery("MinArea")
		maxAcreageRangeTwo.Lte(calParams["MaxAcreage"])
		maxAcreageQueryOne.Must(maxAcreageRangeOne)
		maxAcreageQueryOne.Must(maxAcreageRangeTwo)

		//A>甲&&A<=乙
		maxAcreageRangeThree := elastic.NewRangeQuery("MinArea")
		maxAcreageRangeThree.Lt(calParams["MinAcreage"])
		maxAcreageRangeFour := elastic.NewRangeQuery("MaxArea")
		maxAcreageRangeFour.Gte(calParams["MinAcreage"])
		maxAcreageQueryTwo.Must(maxAcreageRangeThree)
		maxAcreageQueryTwo.Must(maxAcreageRangeFour)

		maxAcreageQuery.Should(maxAcreageQueryOne)
		maxAcreageQuery.Should(maxAcreageQueryTwo)

		queryService.MustNot(elastic.NewTermQuery("MinPrice", 0))
		queryService.MustNot(elastic.NewTermQuery("MaxPrice", 0))

		queryService.Must(maxAcreageQuery)
	}

	if calParams["PredictCredDate"] != 0 {
		queryService.Must(elastic.NewTermQuery("PredictCredDate", calParams["PredictCredDate"]))
	}

	if calParams["IsNearLineOne"] != 0 {
		queryService.Must(elastic.NewTermQuery("Project.IsNearLineOne", 1))
	}
	if calParams["IsNearLineTwo"] != 0 {
		queryService.Must(elastic.NewTermQuery("Project.IsNearLineTwo", 1))
	}
	if calParams["IsNearLineThird"] != 0 {
		queryService.Must(elastic.NewTermQuery("Project.IsNearLineThird", 1))
	}
	if calParams["IsNearLineFouth"] != 0 {
		queryService.Must(elastic.NewTermQuery("Project.IsNearLineFouth", 1))
	}
	if calParams["IsNearLineFifth"] != 0 {
		queryService.Must(elastic.NewTermQuery("Project.IsNearLineFifth", 1))
	}
	if calParams["IsNearLineSixth"] != 0 {
		queryService.Must(elastic.NewTermQuery("Project.IsNearLineSixth", 1))
	}


	searchService = searchService.Query(queryService)

	//分页构建
	if start == 1 || start == 0 {
		start = 0
	} else {
		start = (start-1)*size
	}
	if commonParams["IsAll"] != "" {
		searchService.Sort("NoStatus", true)
	}
	searchService.Sort("BatchNo", false)
	searchService.Sort("_score", false).Sort(commonParams["sort"], sortType)
	searchResult, err := searchService.
		From(start).Size(size).
		Pretty(true).
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return searchResult
}