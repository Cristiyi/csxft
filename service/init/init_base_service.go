package init

import (
	"context"
	"csxft/model"
	"csxft/mongodb"
	"csxft/serializer"
	"csxft/util"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var (
	client     = mongodb.GetMgoCli()
	collection *mongo.Collection
	err        error
	cursor     *mongo.Cursor
)

////分组查询出的工程施工许可证号
//var constructionNo chan string = make(chan string,1000)
////查重筛选出的工程施工许可证号
//var constructionNoResult chan string = make(chan string,1000)
//var mongoResult chan bson.M = make(chan bson.M, 1000)
//var wg sync.WaitGroup

// InitBaseService 初始化基础数据服务
type InitBaseService struct {
	Date        string `form:"date" json:"date" binding:"required"`
}

// 初始化
func (service *InitBaseService) Init() serializer.Response {

	collection = client.Database(os.Getenv("MONGO_DATABASE")).Collection("CredItem")
	groupStage := mongo.Pipeline{
		{{"$group", bson.D{{"_id", "$constructionNo"}}}},
		//{{"$match", bson.D{{"created_at", service.Date}}}},
	}

	if cursor, err = collection.Aggregate(context.TODO(), groupStage, ); err != nil {
		log.Fatal(err)
	}
	//延迟关闭游标
	defer func() {
		if err = cursor.Close(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()
	//遍历游标
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	for _, result := range results {
		getFromMongo(service.Date, result["_id"].(string))
	}

	//wg.Wait()

	return serializer.Response{
		Code: 200,
	}

}

//mysql查重
func checkMySQL(results []bson.M) {
	for _, result := range results {
		tempNo := util.DeleteTailBlank(util.InConvertString(result["constructionNo"]))
		if tempNo != "" {
			count := 0
			model.DB.Model(&model.Project{}).Where("construction_no = ?", tempNo).Count(&count)
			if count == 0 {
				fmt.Println(count)
				insertToMysql(result)
			} else {
				continue
			}
		}
	}
}

//从mongo中提取要插入的数据
func getFromMongo(date string, tempNo string) {
	if tempNo != "" {
		collection = client.Database(os.Getenv("MONGO_DATABASE")).Collection("CredItem")
		filter := bson.M{"created_at":date, "constructionNo":tempNo}
		if cursor, err = collection.Find(context.TODO(), filter, options.Find()); err != nil {
			log.Fatal(err)
		}
		//延迟关闭游标
		defer func() {
			if err = cursor.Close(context.TODO()); err != nil {
				log.Fatal(err)
			}
		}()
		var results []bson.M
		if err = cursor.All(context.TODO(), &results); err != nil {
			log.Fatal(err)
		}
		checkMySQL(results)
	}
}

func insertToMysql(tempResult bson.M) {
	if tempResult != nil {
		var area model.Area
		if err := model.DB.Where("pid = ? and name = ?", "1482", util.InConvertString(tempResult["area"])).First(&area).Error; err != nil {
			area.ID = 0
		}
		project := model.Project{ConstructionNo: util.InConvertString(tempResult["constructionNo"]), ProjectName: util.InConvertString(tempResult["name"]), AreaId: area.ID,
			Approval: util.InConvertString(tempResult["approval"]), DevelopCompany: util.InConvertString(tempResult["developCompany"]), CountBuilding: util.InConvertString(tempResult["countBuilding"] ),
			ProjectAddress: util.InConvertString(tempResult["address"]), MinPrice: util.InConvertString(tempResult["minPrice"]), SaleAddress: util.InConvertString(tempResult["saleAddress"]),
			SalePhone: util.InConvertString(tempResult["salePhone"]), AllHome: util.InConvertString(tempResult["allHome"]), BusLine: util.InConvertString(tempResult["busLine"]),
			AllAcreage: util.InConvertString(tempResult["allAcreage"]), DesignCompany: util.InConvertString(tempResult["designCompany"]), AllArchitectureAcreage: util.InConvertString(tempResult["allArchitectureAcreage"]),
			SaleAgency: util.InConvertString(tempResult["saleAgency"]), VolumeRatio: util.InConvertString(tempResult["volumeRatio"]), Property: util.InConvertString(tempResult["property"]),
			GreenRate: util.InConvertString(tempResult["greenRate"]), ConstructionCompany: util.InConvertString(tempResult["constructionCompany"]), PropertyCost: util.InConvertString(tempResult["propertyCost"]),
			CompleteTime: util.StrToTime(util.InConvertString(tempResult["completeTime"])), Introduction: util.InConvertString(tempResult["introduction"]),
		}
		if err := model.DB.Create(&project).Error; err != nil {
			fmt.Println("失败原因")
			fmt.Println(err)
			collection = client.Database(os.Getenv("MONGO_DATABASE")).Collection("BaseError")
			var (
				iResult    *mongo.InsertOneResult
				id         primitive.ObjectID
			)
			//插入某一条数据
			if iResult, err = collection.InsertOne(context.TODO(), tempResult); err != nil {
				fmt.Print(err)
				return
			}
			//_id:默认生成一个全局唯一ID
			id = iResult.InsertedID.(primitive.ObjectID)
			fmt.Println("base插入失败，入库mongo，自增ID", id.Hex())
		}
	}
}



