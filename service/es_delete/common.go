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

func DeleteDoc(id int,index string)(*elastic.DeleteResponse, error){
	client := elasticsearch.GetEsCli()
	rsp,err := client.Delete().Index(index).Id(strconv.Itoa(id)).Do(context.Background())
	if err != nil{
		return nil,err
	}
	return rsp,nil
}
