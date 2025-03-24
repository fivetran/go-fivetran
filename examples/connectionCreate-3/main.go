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

	client := fivetran.New(apiKey, apiSecret)
	svc := client.NewConnectionCreate()

	connConfig := fivetran.NewConnectionConfig().
		SchemaPrefix("test_postgres_go_1").
		Host("terraform-pgsql-connection-test.cp0rdhwjbsae.us-east-1.rds.amazonaws.com").
		Port(5432).
		Database("fivetran").
		User("postgres").
		Password("mYP4ssw0rd").
		UpdateMethod("XMIN")

	svc.GroupID("replying_ministry")
	svc.Service("postgres_rds")
	svc.Config(connConfig)
	svc.Paused(false)
	svc.TrustCertificates(true)
	svc.TrustFingerprints(true)
	svc.RunSetupTests(true)

	value, err := svc.Do(context.Background())
	if err != nil {
		fmt.Printf("%+v\n", value)
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", value)
}
