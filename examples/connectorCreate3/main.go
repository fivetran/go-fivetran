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

	connConfig := fivetran.NewConnectorConfig().
		SchemaPrefix("test_postgres_terraform_20210601v2").
		Host("terraform-pgsql-connector-test.cp0rdhwjbsae.us-east-1.rds.amazonaws.com").
		Port(5432).
		Database("fivetran").
		User("postgres").
		Port(5432).
		Password("thisIsMyNewFiveTranP4ssw0rd123").
		UpdateMethod("XMIN")

	svc.GroupID("replying_ministry")
	svc.Service("postgres_rds")
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
