package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/fivetran/go-fivetran"
)

func main() {
	apiKey := os.Getenv("FIVETRAN_APIKEY")
	apiSecret := os.Getenv("FIVETRAN_APISECRET")
	fivetran.Debug(true)

	client := fivetran.NewClient(apiKey, apiSecret)

	svc := client.NewConnectorModify()

	connConfig := fivetran.NewConnectorConfig().
		NamedRange("range1")

	cTables1 := fivetran.NewConnectorConfigCustomTables().Aggregation("aggregation1").TableName("theName")
	cTables2 := fivetran.NewConnectorConfigCustomTables().TableName("theName")
	connConfig.CustomTables(&[]*fivetran.ConnectorConfigCustomTables{cTables1, cTables2})

	svc.ConnectorID("grateful_vertices")
	svc.Paused(true)
	svc.SyncFrequency(5)
	svc.Config(connConfig)

	value, err := svc.Do(context.Background())
	checkErr(err, value)

	fmt.Printf("%+v\n", value)
}

func checkErr(err error, value interface{}) {
	if err != nil {
		fmt.Printf("%+v\n", value)
		log.Fatal(err)
	}
}
