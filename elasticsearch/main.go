package elasticsearch

import (
	"context"
	"github.com/olivere/elastic/v7"
	"log"
	"os"
)

var esCli *elastic.Client

func initEngine() {
	var err error
	elastic.SetSniff(false)
	esCli, err = elastic.NewClient(elastic.SetURL(os.Getenv("ES_URL")))
	if err != nil {
		log.Fatal(err)
	}
	_,_,err = esCli.Ping(os.Getenv("ES_URL")).Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
func GetEsCli() *elastic.Client {
	if esCli == nil {
		initEngine()
	}
	return esCli
}
