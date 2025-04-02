package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/connections"
)

func main() {
	apiKey := os.Getenv("FIVETRAN_APIKEY")
	apiSecret := os.Getenv("FIVETRAN_APISECRET")
	syncFrequency := 5;
	fivetran.Debug(true)

	client := fivetran.New(apiKey, apiSecret)

	svc := client.NewConnectionUpdate()

	connConfig := fivetran.NewConnectionConfig().
		NamedRange("range1")

	cTables1 := fivetran.NewConnectionConfigCustomTables().Aggregation("aggregation1").TableName("theName")
	cTables2 := fivetran.NewConnectionConfigCustomTables().TableName("theName").Fields([]string{"FIELD ONE", "FIELD TWO"})

	connConfig.CustomTables([]*connections.ConnectionConfigCustomTables{cTables1, cTables2})

	adobeAnalyticsConfig1 := fivetran.NewConnectionConfigAdobeAnalyticsConfiguration().SyncMode("syncMode").Elements([]string{"elemet1", "element2"})
	connConfig.AdobeAnalyticsConfigurations([]*connections.ConnectionConfigAdobeAnalyticsConfiguration{adobeAnalyticsConfig1})

	svc.ConnectionID("grateful_vertices")
	svc.Paused(true)
	svc.SyncFrequency(&syncFrequency)
	svc.Config(connConfig)

	value, err := svc.Do(context.Background())
	if err != nil {
		fmt.Printf("%+v\n", value)
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", value)
}
