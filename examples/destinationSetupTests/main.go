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

	svc := client.NewDestinationSetupTests()

	svc.TrustCertificates(true)
	svc.TrustFingerprints(false)

	value, err := svc.DestinationID("replying_ministry").Do(context.Background())
	if err != nil {
		fmt.Printf("%+v\n", value)
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", value)
}
