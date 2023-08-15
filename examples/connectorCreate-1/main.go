package main

import (
	"context"
	"fmt"
	"log"

	"github.com/fivetran/go-fivetran"
)

func main() {
	apiKey := "_moonbeam_acc_accountworthy_api_key_rbac"  //os.Getenv("FIVETRAN_APIKEY")
	apiSecret := "_moonbeam_acc_accountworthy_api_secret" //os.Getenv("FIVETRAN_APISECRET")
	fivetran.Debug(true)

	client := fivetran.New(apiKey, apiSecret)
	client.BaseURL("http://localhost:8001/v1")

	connConfig := fivetran.NewConnectorConfig().
		Schema("google_sheets2").
		Table("table").
		SheetID("1Rmq_FN2kTNwWiT4adZKBxHBRmvfeBTIfKWi5B8ii9qk").
		NamedRange("range1")

	svc := client.NewConnectorCreate().
		GroupID("_moonbeam").
		Service("google_sheets").
		Config(connConfig).
		Paused(false).
		TrustCertificates(true).
		TrustFingerprints(true).
		RunSetupTests(true)

	value, err := svc.Do(context.Background())
	if err != nil {
		fmt.Printf("%+v\n", value)
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", value)
}
