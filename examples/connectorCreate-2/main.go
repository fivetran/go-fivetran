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
	svc := client.NewConnectorCreate()

	connConfig := fivetran.NewConnectorConfig()

	connConfigReport := fivetran.NewConnectorConfigReports().ConfigType("cType1").Filter("filter1")
	connConfigReport2 := fivetran.NewConnectorConfigReports().ConfigType("cType2")
	connConfig.Reports([]*fivetran.ConnectorConfigReports{connConfigReport, connConfigReport2})

	pCred1 := fivetran.NewConnectorConfigProjectCredentials().APIKey("myApiKEY1").Project("myProject1")
	pCred2 := fivetran.NewConnectorConfigProjectCredentials().APIKey("myApiKEY2").Project("myProject2").SecretKey("TheSecretKEY2")
	pCred3 := fivetran.NewConnectorConfigProjectCredentials().APIKey("myApiKEY3").Project("myProject3").SecretKey("YASK3")
	connConfig.ProjectCredentials([]*fivetran.ConnectorConfigProjectCredentials{pCred1, pCred2, pCred3})

	cTables1 := fivetran.NewConnectorConfigCustomTables().Aggregation("aggregation1").TableName("TableNAME1")
	cTables2 := fivetran.NewConnectorConfigCustomTables().TableName("TableNAME2")
	connConfig.CustomTables([]*fivetran.ConnectorConfigCustomTables{cTables1, cTables2})

	adobeAnalyticsConfig1 := fivetran.NewConnectorConfigAdobeAnalyticsConfiguration().SyncMode("syncMode").Elements([]string{"elemet1", "element2"})
	connConfig.AdobeAnalyticsConfigurations([]*fivetran.ConnectorConfigAdobeAnalyticsConfiguration{adobeAnalyticsConfig1})

	connConfig.Schema("google_sheets5959").
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
	if err != nil {
		fmt.Printf("%+v\n", value)
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", value)
}
