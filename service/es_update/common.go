/**
 * @Description:
 * @File: common
 * @Date: 2020/8/18 0018 10:12
 */

package es_update

import (
	"context"
	"csxft/elasticsearch"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"strconv"
)

//修改封装
func Update(updateField *map[string]interface{},id int,index string) (*elastic.UpdateResponse,error) {

	client := elasticsearch.GetEsCli()
	if !IsDocExists(id,index){
		return nil,fmt.Errorf("id不存在")
	}
	rsp, err := client.Update().Index(index).Id(strconv.Itoa(id)).Doc(updateField).Do(context.Background())
	if err != nil{
		fmt.Println(err)
		return nil,err
	}
	return rsp,nil

}

//判断文档是否存在
func IsDocExists(id int,index string) bool {
	client := elasticsearch.GetEsCli()
	exist,_ := client.Exists().Index(index).Id(strconv.Itoa(id)).Do(context.Background())
	if !exist{
		log.Println("ID may be incorrect! ",id)
		return false
	}
	return true
}
