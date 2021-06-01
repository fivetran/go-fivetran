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

	connConfigReport := fivetran.NewConnectorConfigReports().ConfigType("cType1").Filter("filter1")
	connConfigReport2 := fivetran.NewConnectorConfigReports().ConfigType("cType2")

	connConfig := fivetran.NewConnectorConfig()
	connConfig.Reports([]fivetran.ConnectorConfigReports{*connConfigReport, *connConfigReport2})

	connConfig.Schema("google_sheets2").
		Table("table").
		SheetID("1Rmq_FN2kTNwWiT4adZKBxHBRmvfeBTIfKWi5B8ii9qk").
		NamedRange("range1")

	connAuthClientAccess := fivetran.NewConnectorAuthClientAccess().
		ClientID("thisIsClientID").
		ClientSecret("thisIsClientSecret").
		DeveloperToken("thisIsDeveloperToken").
		UserAgent("thisIsUserAgent")

	connAuth := fivetran.NewConnectorAuth()
	connAuth.ClientAccess(*connAuthClientAccess).
		AccessToken("thisIsAccessToken").
		RealmID("thisIsRealmID").
		RefreshToken("thisIsRefreshToken")

	svc.GroupID("replying_ministry")
	svc.Service("google_sheets")
	svc.Config(connConfig)
	// svc.Auth(connAuth)
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
