package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/connections"
)

func main() {
	apiKey := os.Getenv("FIVETRAN_APIKEY")
	apiSecret := os.Getenv("FIVETRAN_APISECRET")
	fivetran.Debug(true)

	client := fivetran.New(apiKey, apiSecret)
	svc := client.NewConnectionCreate()

	connConfig := fivetran.NewConnectionConfig()

	connConfigReport := fivetran.NewConnectionConfigReports().ConfigType("cType1").Filter("filter1")
	connConfigReport2 := fivetran.NewConnectionConfigReports().ConfigType("cType2")
	connConfig.Reports([]*connections.ConnectionConfigReports{connConfigReport, connConfigReport2})

	pCred1 := fivetran.NewConnectionConfigProjectCredentials().APIKey("myApiKEY1").Project("myProject1")
	pCred2 := fivetran.NewConnectionConfigProjectCredentials().APIKey("myApiKEY2").Project("myProject2").SecretKey("TheSecretKEY2")
	pCred3 := fivetran.NewConnectionConfigProjectCredentials().APIKey("myApiKEY3").Project("myProject3").SecretKey("YASK3")
	connConfig.ProjectCredentials([]*connections.ConnectionConfigProjectCredentials{pCred1, pCred2, pCred3})

	cTables1 := fivetran.NewConnectionConfigCustomTables().Aggregation("aggregation1").TableName("TableNAME1")
	cTables2 := fivetran.NewConnectionConfigCustomTables().TableName("TableNAME2")
	connConfig.CustomTables([]*connections.ConnectionConfigCustomTables{cTables1, cTables2})

	adobeAnalyticsConfig1 := fivetran.NewConnectionConfigAdobeAnalyticsConfiguration().SyncMode("syncMode").Elements([]string{"elemet1", "element2"})
	connConfig.AdobeAnalyticsConfigurations([]*connections.ConnectionConfigAdobeAnalyticsConfiguration{adobeAnalyticsConfig1})

	connConfig.Schema("google_sheets5959").
		Table("table").
		SheetID("1Rmq_FN2kTNwWiT4adZKBxHBRmvfeBTIfKWi5B8ii9qk").
		NamedRange("range1")

	connAuthClientAccess := fivetran.NewConnectionAuthClientAccess().
		ClientID("thisIsClientID").
		ClientSecret("thisIsClientSecret").
		DeveloperToken("thisIsDeveloperToken").
		UserAgent("thisIsUserAgent")

	connAuth := fivetran.NewConnectionAuth()
	connAuth.ClientAccess(connAuthClientAccess).
		AccessToken("thisIsAccessToken").
		RealmID("thisIsRealmID").
		RefreshToken("thisIsRefreshToken").
		PreviousRefreshToken("thisIsPreviousRefreshToken").
		UserAccessToken("thisIsUserAccessToken").
		ConsumerSecret("thisIsConsumerSecret").
		ConsumerKey("thisIsConsumerKey").
		OauthToken("thisIsOauthToken").
		OauthTokenSecret("thisIsOauthTokenSecret").
		RoleArn("thisIsRoleArn").
		AwsAccessKey("thisIsAwsAccessKey").
		AwsSecretKey("thisIsAwsSecretKey").
		ClientId("thisIsClientId").
		KeyId("thisIsKeyId").
		TeamId("thisIsTeamId").
		ClientSecret("thisIsClientSecret")

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
