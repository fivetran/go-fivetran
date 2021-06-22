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
	svc := client.NewConnectorCreate()

	connConfig := fivetran.NewConnectorConfig()

	// connConfigReport := fivetran.NewConnectorConfigReports().ConfigType("cType1").Filter("filter1")
	// connConfigReport2 := fivetran.NewConnectorConfigReports().ConfigType("cType2")
	// connConfig.Reports(&[]*fivetran.ConnectorConfigReports{connConfigReport, connConfigReport2})

	// pCred1 := fivetran.NewConnectorConfigProjectCredentials().APIKey("myApiKEY").Project("myProject")
	// pCred2 := fivetran.NewConnectorConfigProjectCredentials().APIKey("myApiKEY").Project("myProject").SecretKey("TheSecretKEY")
	// pCred3 := fivetran.NewConnectorConfigProjectCredentials().APIKey("myApiKEY").Project("myProject").SecretKey("YASK")
	// connConfig.ProjectCredentials(&[]*fivetran.ConnectorConfigProjectCredentials{pCred1, pCred2, pCred3})

	// cTables1 := fivetran.NewConnectorConfigCustomTables().Aggregation("aggregation1").TableName("theName")
	// cTables2 := fivetran.NewConnectorConfigCustomTables().TableName("theName")
	// connConfig.CustomTables(&[]*fivetran.ConnectorConfigCustomTables{cTables1, cTables2})

	connConfig.Schema("google_sheets7").
		Table("table").
		SheetID("1Rmq_FN2kTNwWiT4adZKBxHBRmvfeBTIfKWi5B8ii9qk").
		NamedRange("range1")

	connAuthClientAccess := fivetran.NewConnectorAuthClientAccess().
		ClientID("thisIsClientID").
		ClientSecret("thisIsClientSecret").
		DeveloperToken("thisIsDeveloperToken").
		UserAgent("thisIsUserAgent")

	connAuth := fivetran.NewConnectorAuth()
	connAuth.ClientAccess(connAuthClientAccess).
		AccessToken("thisIsAccessToken").
		RealmID("thisIsRealmID").
		RefreshToken("thisIsRefreshToken")

	svc.GroupID("replying_ministry")
	svc.Service("google_sheets")
	svc.Config(connConfig)
	svc.Auth(connAuth)
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
