/**
 * @Description:
 * @File: common
 * @Date: 2020/8/18 0018 10:51
 */

package es_delete

import (
	"context"
	"csxft/elasticsearch"
	"github.com/olivere/elastic/v7"
	"strconv"
)

//通用删除-根据id
func DeleteDoc(id int,index string)(*elastic.DeleteResponse, error){
	client := elasticsearch.GetEsCli()
	rsp,err := client.Delete().Index(index).Id(strconv.Itoa(id)).Do(context.Background())
	if err != nil{
		return nil,err
	}
	return rsp,nil
}


//根据开盘删除一房一价
func DeleteCredHouse(credId int)(){
	client := elasticsearch.GetEsCli()
	query := elastic.Query(elastic.NewTermQuery("CredId", credId))
	client.DeleteByQuery("house").Query(query).Do(context.Background())
}

//根据批次删除摇号结果
func DeleteBatchLotteryResult(batchId int)(){
	client := elasticsearch.GetEsCli()
	query := elastic.Query(elastic.NewTermQuery("BatchId", batchId))
	client.DeleteByQuery("lottery_result").Query(query).Do(context.Background())
}

//根据批次删除认筹结果
func DeleteBatchSolicitResult(batchId int)(){
	client := elasticsearch.GetEsCli()
	query := elastic.Query(elastic.NewTermQuery("BatchId", batchId))
	client.DeleteByQuery("solicit_result").Query(query).Do(context.Background())
}