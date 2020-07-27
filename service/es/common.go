package es

import (
	"context"
	"CMD-XuanFangTong-Server/elasticsearch"
	"fmt"
	"github.com/olivere/elastic/v7"
	"strings"
)

//搜索楼盘
func QueryProject(start,size int, commonParams map[string]string) *elastic.SearchResult {
	sortType := true
	if commonParams["sortType"] == "desc" {
		sortType = false
	}
	searchService := elasticsearch.GetEsCli().Search("project")

	//搜索条件构建
	queryService := elastic.NewBoolQuery()
	if commonParams["name"] != "" {
		queryService.Must(elastic.NewMatchQuery("ProjectName", commonParams["name"]))
	}
	if commonParams["IsWillCred"] != "" {
		queryService.Must(elastic.NewQueryStringQuery("IsWillCred:"+"1"))
	}
	if commonParams["IsNewCred"] != "" {
		queryService.Must(elastic.NewQueryStringQuery("IsNewCred:"+"1"))
	}
	if commonParams["IsRecognition"] != "" {
		queryService.Must(elastic.NewQueryStringQuery("IsRecognition:"+"1"))
	}
	if commonParams["IsIottery"] != "" {
		queryService.Must(elastic.NewQueryStringQuery("IsIottery:"+"1"))
	}
	if commonParams["IsSell"] != "" {
		queryService.Must(elastic.NewQueryStringQuery("IsSell:"+"1"))
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
		searchService = searchService.Query(elastic.NewQueryStringQuery("IsWillCred:"+"1"))
		break
	case "2":
		searchService = searchService.Query(elastic.NewQueryStringQuery("IsNewCred:"+"1"))
	case "3":
		searchService = searchService.Query(elastic.NewQueryStringQuery("IsRecognition:"+"1"))
	case "4":
		searchService = searchService.Query(elastic.NewQueryStringQuery("IsIottery:"+"1"))
	case "5":
		searchService = searchService.Query(elastic.NewQueryStringQuery("IsSell:"+"1"))
	default:
		searchService = searchService.Query(elastic.NewQueryStringQuery("IsSell:"+"1"))
	}
	count, err := searchService.Do(context.Background())
	if err != nil {
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

//根据开盘信息获取楼盘信息
func GetCredHouse (start,size int, commonParams map[string]string, credIds []string) *elastic.SearchResult {

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
func GetProjectIottery (start,size int, commonParams map[string]string) *elastic.SearchResult {
	sortType := true
	if commonParams["sortType"] == "desc" {
		sortType = false
	}
	searchService := elasticsearch.GetEsCli().Search("iottery")
	searchService = searchService.Query(elastic.NewTermQuery("ProjectId", commonParams["ProjectId"]))
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
	if commonParams["IsNew"] != "" {
		queryService.Must(elastic.NewQueryStringQuery("IsNew:"+"1"))
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
	queryService.Must(elastic.NewTermQuery("ProjectId",commonParams["ProjectId"]))
	if commonParams["IsNew"] != "" {
		queryService.Must(elastic.NewQueryStringQuery("IsNew:"+"1"))
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
	if commonParams["NoticeType"] != "" {
		fmt.Println(commonParams)
		queryService.Must(elastic.NewQueryStringQuery("NoticeType:" +commonParams["NoticeType"]))
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

