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

	svc := client.NewDestinationModify()

	destConfig := fivetran.NewDestinationConfig().
		Host("10.99.99.99").
		Port(443).
		Database("myNewDatabase").
		User("myNewUser").
		Password("MyNewPass").
		Auth("PASSWORD")

	svc.DestinationID("replying_ministry")
	svc.Region("EU")
	svc.TimeZoneOffset("0")
	svc.Config(destConfig)

	value, err := svc.Do(context.Background())
	if err != nil {
		fmt.Printf("%+v\n", value)
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", value)
}
