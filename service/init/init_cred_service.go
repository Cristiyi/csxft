package init

import (
	"context"
	"csxft/model"
	"csxft/serializer"
	"csxft/util"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"sync"
)

type ProjectIdResult struct {
	ID uint64
}
var mongoResult chan bson.M = make(chan bson.M, 1000)
var wg sync.WaitGroup

// InitCredService 初始化预售证号数据服务
type InitCredService struct {
	Date string `form:"date" json:"date" binding:"required"`
}

// 初始化
func (service *InitCredService) Init() serializer.Response {

	collection = client.Database(os.Getenv("MONGO_DATABASE")).Collection("CredItem")
	filter := bson.M{"created_at":service.Date}
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

	for _, result := range results {
		mongoResult <- result
		wg.Add(1)
		go checkCredMySQL()
	}

	wg.Wait()

	return serializer.Response{
		Code: 200,
	}

}

//mysql查重
func checkCredMySQL() {
	tempResult :=  <- mongoResult
	var projectIdResult ProjectIdResult
	if err := model.DB.Table("xft_projects").Select("id").Where("project_name = ? and area_origin = ?", util.InConvertString(tempResult["name"]), util.InConvertString(tempResult["area"])).Scan(&projectIdResult).Error; err == nil {
		count := 0
		model.DB.Model(&model.Cred{}).Where("cred = ? and project_id = ?", util.InConvertString(tempResult["cred"]), projectIdResult.ID).Count(&count)
		if count == 0 {
			insertCredToMysql(tempResult, projectIdResult.ID)
		}
	}
	defer wg.Done()
}

func insertCredToMysql(tempResult bson.M, projectId uint64) {
	if tempResult != nil {
		//入库
		cred := model.Cred{ProjectId: projectId, Cred: util.InConvertString(tempResult["cred"]), BuildingNo: util.InConvertString(tempResult["houseNo"]),
			CredDate: util.StrToTime(util.InConvertString(tempResult["credDate"])), Acreage: util.InConvertString(tempResult["acreage"]), LandNo: util.InConvertString(tempResult["landNo"] ),
			EngineeNo: util.InConvertString(tempResult["engineeNo"]), LandPlanNo: util.InConvertString(tempResult["landPlanNo"]),
		}
		//入库失败 记录到mongo
		if err := model.DB.Create(&cred).Error; err != nil {
			fmt.Println("失败原因")
			fmt.Println(err)
			collection = client.Database(os.Getenv("MONGO_DATABASE")).Collection("CredError")
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
			fmt.Println("cred插入失败，入库mongo，自增ID", id.Hex())
		}
	}
}



