package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/connectors"
)

func main() {
	apiKey := os.Getenv("FIVETRAN_APIKEY")
	apiSecret := os.Getenv("FIVETRAN_APISECRET")
	fivetran.Debug(true)

	client := fivetran.New(apiKey, apiSecret)

	svc := client.NewConnectorModify()

	connConfig := fivetran.NewConnectorConfig().
		NamedRange("range1")

	cTables1 := fivetran.NewConnectorConfigCustomTables().Aggregation("aggregation1").TableName("theName")
	cTables2 := fivetran.NewConnectorConfigCustomTables().TableName("theName").Fields([]string{"FIELD ONE", "FIELD TWO"})

	connConfig.CustomTables([]*connectors.ConnectorConfigCustomTables{cTables1, cTables2})

	adobeAnalyticsConfig1 := fivetran.NewConnectorConfigAdobeAnalyticsConfiguration().SyncMode("syncMode").Elements([]string{"elemet1", "element2"})
	connConfig.AdobeAnalyticsConfigurations([]*connectors.ConnectorConfigAdobeAnalyticsConfiguration{adobeAnalyticsConfig1})

	svc.ConnectorID("grateful_vertices")
	svc.Paused(true)
	svc.SyncFrequency("5")
	svc.Config(connConfig)

	value, err := svc.Do(context.Background())
	if err != nil {
		fmt.Printf("%+v\n", value)
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", value)
}
