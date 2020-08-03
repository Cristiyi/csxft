/**
 * @Description:
 * @File: init_fdc_service.go
 * @Date: 2020/7/30 0030 18:25
 */

package init

import (
	"context"
	"csxft/model"
	"csxft/repo"
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

var fdcResult chan bson.M = make(chan bson.M, 2000)

// InitBaseService 初始化基础数据服务
type InitFdcService struct {
	Date        string `form:"date" json:"date" binding:"required"`
}

// 初始化
func (service *InitFdcService) Init() serializer.Response {

	collection = client.Database(os.Getenv("MONGO_DATABASE")).Collection("FdccomBaseItem")
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
	//遍历游标
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	for _, result := range results {
		fdcResult <- result
		wg.Add(1)
		go checkFdcMySQL()
	}

	wg.Wait()

	return serializer.Response{
		Code: 200,
	}

}

//mysql查重
func checkFdcMySQL() {
	tempResult :=  <- fdcResult
	count := 0
	model.DB.Model(&model.Fdc{}).Where("project_name = ? and area = ?", util.InConvertString(tempResult["projectName"]), util.InConvertString(tempResult["cred"])).Count(&count)
	if count == 0 {
		insertFdcToMysql(tempResult)
	}
	defer wg.Done()
}

func insertFdcToMysql(tempResult bson.M) {
	if tempResult != nil {
		//入库
		fdc := model.Fdc{Cred: util.InConvertString(tempResult["cred"]), ProjectName: util.InConvertString(tempResult["projectName"]), SaleStatus: util.InConvertString(tempResult["saleStatus"]),
			HouseType: util.InConvertString(tempResult["houseType"]), ReferencePrice: util.InConvertString(tempResult["referencePrice"]), PrimaryHouseStructure: util.InConvertString(tempResult["primaryHouseStructure"]),
			Area: util.InConvertString(tempResult["area"]), DecorationStatus: util.InConvertString(tempResult["decorationStatus"]), PlannedAllHome: util.InConvertString(tempResult["plannedAllHome"]),
			CompletedTime: util.InConvertString(tempResult["completedTime"]), StallCount: util.InConvertString(tempResult["stallCount"]), OpenTime: util.InConvertString(tempResult["openTime"]),
			AllAcreage: util.InConvertString(tempResult["allAcreage"]), AllFloorAcreage: util.InConvertString(tempResult["allFloorAcreage"]), VolumeRatio: util.InConvertString(tempResult["volumeRatio"]),
			GreenRate: util.InConvertString(tempResult["greenRate"]), DesignCompany: util.InConvertString(tempResult["designCompany"]), PropertyCost: util.InConvertString(tempResult["propertyCost"]),
			Property: util.InConvertString(tempResult["property"]), ConstructionCompany: util.InConvertString(tempResult["constructionCompany"]), DevelopCompany: util.InConvertString(tempResult["developCompnay"]),
			Address: util.InConvertString(tempResult["address"]), BusLine: util.InConvertString(tempResult["busLine"]), SurroundingFacility: util.InConvertString(tempResult["surroundingFacility"]),
		}
		//入库失败 记录到mongo
		if err := model.DB.Create(&fdc).Error; err != nil {
			fmt.Println("失败原因")
			fmt.Println(err)
			collection = client.Database(os.Getenv("MONGO_DATABASE")).Collection("FdcError")
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
			fmt.Println("fdc插入失败，入库mongo，自增ID", id.Hex())
		}

		var area model.Area
		if err := model.DB.Where("pid = ? and name = ?", "1482", util.InConvertString(tempResult["area"])).First(&area).Error; err == nil {
			var creds []*model.Cred
			creds = repo.NewCredRepo().GetByCred(util.InConvertString(tempResult["cred"]))
			if creds != nil {
				for _, item := range creds {
					project := &model.Project{}
					err := model.DB.Model(model.Project{}).Where("id = ? and area_id = ?", item.ProjectId, area.ID).First(project).Error
					fmt.Println("查找project失败")
					fmt.Println(err)
					if err == nil {
						err := model.DB.Model(project).Update("promotion_first_name", util.InConvertString(tempResult["projectName"])).Error
						if err != nil {
							fmt.Println("修改失败")
							fmt.Println(err)
						}
					}
				}
			}
		}

	}
}
