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
)

type CredIdResult struct {
	ID uint64
}

type PurposeIdResult struct {
	ID uint
}

type TypeIdResult struct {
	ID uint
}

type DecorationResult struct {
	ID uint
}

// InitCredService 初始化开盘房屋数据服务
type InitHouseService struct {
	Date string `form:"date" json:"date" binding:"required"`
}

// 初始化
func (service *InitHouseService) Init() serializer.Response {

	collection = client.Database(os.Getenv("MONGO_DATABASE")).Collection("HouseItem")
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
		checkHouseMySQL(result)
	}

	return serializer.Response{
		Code: 200,
	}

}

//mysql查重
func checkHouseMySQL(tempResult bson.M) {
	var credIdResult CredIdResult
	if err := model.DB.Table("xft_creds").Select("id").Where("cred = ?", util.InConvertString(tempResult["cred"])).Scan(&credIdResult).Error; err == nil {
		count := 0
		model.DB.Model(&model.House{}).Where("cred_id = ? and house_no = ?", credIdResult.ID, util.InConvertString(tempResult["houseNo"])).Count(&count)
		if count == 0 {
			insertHouseToMysql(tempResult, credIdResult.ID)
		}
	}
}

func insertHouseToMysql(tempResult bson.M, credId uint64) {
	if tempResult != nil {
		//房屋用途
		var purposeIdResult PurposeIdResult
		if err := model.DB.Table("xft_project_purposes").Select("id").Where("name = ?", util.InConvertString(tempResult["housePurpose"])).Scan(&purposeIdResult).Error; err != nil {
			purposeIdResult.ID = 0
		}
		//房屋类型
		var typeIdResult TypeIdResult
		if err := model.DB.Table("xft_project_types").Select("id").Where("name = ?", util.InConvertString(tempResult["hourseType"])).Scan(&typeIdResult).Error; err != nil {
			typeIdResult.ID = 0
		}
		//房屋装修
		var decorationResult DecorationResult
		if err := model.DB.Table("xft_decorations").Select("id").Where("name = ?", util.InConvertString(tempResult["decorationStatus"])).Scan(&decorationResult).Error; err != nil {
			decorationResult.ID = 0
		}
		//房屋入库
		house := model.House{CredId: credId, HouseNo: util.InConvertString(tempResult["houseNo"]), FloorNo: util.InConvertInt(tempResult["floorNo"]),
			PurposeId: purposeIdResult.ID, TypeId: typeIdResult.ID, DecorationId: decorationResult.ID, HouseAcreage: util.InConvertFloat64(tempResult["houseAcreage"]),
			UseAcreage: util.InConvertFloat64(tempResult["useAcreage"]), ShareAcreage: util.InConvertFloat64(tempResult["shareAcreage"]),
		}
		//入库失败 记录到mongo
		if err := model.DB.Create(&house).Error; err != nil {
			fmt.Println("失败原因")
			fmt.Println(err)
			collection = client.Database(os.Getenv("MONGO_DATABASE")).Collection("HouseError")
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
			fmt.Println("house插入失败，入库mongo，自增ID", id.Hex())
		}
	}
}



