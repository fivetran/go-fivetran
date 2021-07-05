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

	svc := client.NewDestinationCreate()

	destConfig := fivetran.NewDestinationConfig().
		Host("10.20.30.40").
		Port(443).
		Database("myDatabase").
		User("myUsername").
		Password("myPassword").
		Auth("PASSWORD")

	svc.GroupID("anyplace_silvery")
	svc.Service("snowflake")
	svc.Region("US")
	svc.TimeZoneOffset("-5")
	svc.Config(destConfig)

	value, err := svc.Do(context.Background())
	if err != nil {
		fmt.Printf("%+v\n", value)
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", value)
}
