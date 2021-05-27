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

	svc := client.NewConnectorCreateService()

	connConfig := fivetran.NewConnectorConfig()
	connConfig.Schema("google_sheets")
	connConfig.Table("table")
	connConfig.SheetID("1Rmq_FN2kTNwWiT4adZKBxHBRmvfeBTIfKWi5B8ii9qk")
	connConfig.NamedRange("range1")

	svc.GroupID("replying_ministry")
	svc.Service("google_sheets")
	svc.Config(connConfig)
	svc.Paused(false)
	svc.TrustCertificates(true)
	svc.TrustFingerprints(true)
	svc.RunSetupTests(true)

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
