package testutils

import (
	"context"
	"fmt"
	"net/http"
	"io"
	"bytes"
	"log"
	"os"
	"reflect"
	"encoding/json"
	"testing"
	"time"
	"github.com/fivetran/go-fivetran/tests/mock"
	"github.com/fivetran/go-fivetran"
)

var Client *fivetran.Client

var CertificateHash string
var EncodedCertificate string

// Tests should be re-written to not use a pre-defined user and group
const (
	PredefinedGroupName      = "GoSdkTesting"
	PredefinedUserEmail      = "dev-markov+go-fivetran-sdk@fivetran.com"
	PredefinedUserGivenName  = "Go"
	PredefinedUserFamilyName = "5Tran"
	PredefinedUserPhone      = "+1234567890"
	BqProjectId              = "dulcet-yew-246109"

	// ! WARNING !
	// ! Do not change these values on production ones !
	// ! When running e2e tests locally endure you're using BLANK ACCOUNT credentials !
	PredefinedGroupId = "sepulchre_settlement"
	PredefinedUserId  = "recoup_befell"
)

var (
	cleanup = false
	TEST_KEY    = "test_key"
	TEST_SECRET = "test_secret"

	TEST_CONNECTOR_ID = "test_connector_id"
	TEST_HASH         = "test_hash"
	TEST_PUBLIC_KEY   = "test_public_key"
)

func InitE2E() {
	var apiUrl string
	var apiKey string
	var apiSecret string

	valuesToLoad := map[string]*string{
		"FIVETRAN_API_URL":               &apiUrl,
		"FIVETRAN_APIKEY":                &apiKey,
		"FIVETRAN_APISECRET":             &apiSecret,
		"FIVETRAN_TEST_CERTIFICATE_HASH": &CertificateHash,
		"FIVETRAN_TEST_CERTIFICATE":      &EncodedCertificate,
	}

	for name, value := range valuesToLoad {
		*value = os.Getenv(name)
		if *value == "" {
			log.Fatalf("Environment variable %s is not set!\n", name)
		}
	}

	Client = fivetran.New(apiKey, apiSecret)
	Client.BaseURL(apiUrl)
	if IsPredefinedUserExist() && IsPredefinedGroupExist() {
		if !cleanup {
			CleanupAccount()
			cleanup = true
		}
	} else {
		log.Fatalln("The predefined user doesn't belong to the Testing account. Make sure that credentials are using in the tests belong to the Testing account.")
	}
}

func CreateTestClient() (*fivetran.Client, *mock.HttpClient) {
	ftClient := fivetran.New(TEST_KEY, TEST_SECRET)
	mockClient := mock.NewHttpClient()
	ftClient.SetHttpClient(mockClient)
	return ftClient, mockClient
}

func RequestBodyToJson(t *testing.T, req *http.Request) map[string]interface{} {
	t.Helper()

	bodyBytes, err := io.ReadAll(req.Body)
	if err != nil {
		t.Errorf("requestBodyToJson, cannot read request body: %s", err)
	}
	req.Body.Close()
	req.Body = io.NopCloser(bytes.NewReader(bodyBytes))

	result := map[string]interface{}{}
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		t.Errorf("requestBodyToJson, cannot parse request body: %s", err)
	}

	return result
}

func CreateDbtDestination(t *testing.T) {
	t.Helper()
	destination, err := Client.NewDestinationCreate().
		GroupID(PredefinedGroupId).
		Service("big_query").
		Region("GCP_US_EAST4").
		RunSetupTests(true).
		TimeZoneOffset("-5").
		Config(
			fivetran.NewDestinationConfig().
				ProjectID(BqProjectId).
				DataSetLocation("US")).
		Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", destination)
		t.Error(err)
	}
}

func DeleteDbtDestination() {
	resp, err := Client.NewDestinationDelete().DestinationID(PredefinedGroupId).Do(context.Background())
	if err != nil {
		log.Fatal(resp, err)
	}
}

func CreateUser(t *testing.T) string {
	t.Helper()
	user, err := Client.NewUserInvite().
		Email("william_addison.@fivetran.com").
		GivenName("William").
		FamilyName("Addison").
		Phone("+19876543210").
		Role("Account Reviewer").
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", user)
		t.Error(err)
	}
	return user.Data.ID
}

func CreateTempUser(t *testing.T) string {
	t.Helper()
	userId := CreateUser(t)
	t.Cleanup(func() { DeleteUser(t, userId) })
	return userId
}

func DeleteUser(t *testing.T, id string) {
	t.Helper()
	user, err := Client.NewUserDelete().UserID(id).Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", user)
		t.Error(err)
	}
}

func CreateGroup(t *testing.T) string {
	t.Helper()
	created, err := Client.NewGroupCreate().Name("test").Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", created)
		t.Error(err)
	}
	return created.Data.ID
}

func CreateDbtProject(t *testing.T) string {
	t.Helper()
	CreateDbtDestination(t)
	created, err := Client.NewDbtProjectCreate().
		GroupID(PredefinedGroupId).
		DbtVersion("1.3.1").
		ProjectConfig(fivetran.NewDbtProjectConfig().
			GitRemoteUrl("https://github.com/fivetran/dbt_demo").
			FolderPath("").
			GitBranch("main")).
		DefaultSchema("test_schema").
		TargetName("").
		Threads(4).
		Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", created)
		t.Error(err)
	}
	return created.Data.ID
}

func CleanupDbtProjects() {
	projects, err := Client.NewDbtProjectsList().Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	for _, project := range projects.Data.Items {
		CleanupDbtTransformations(project.ID, "")
		_, err := Client.NewDbtProjectDelete().DbtProjectID(project.ID).Do(context.Background())
		if err != nil && err.Error() != "status code: 404; expected: 200" {
			log.Fatal(err)
		}
	}
	if projects.Data.NextCursor != "" {
		CleanupDbtProjects()
	}
}

func CleanupDbtTransformations(projectId, nextCursor string) {
	svc := Client.NewDbtModelsList().ProjectId(projectId)

	if nextCursor != "" {
		svc.Cursor(nextCursor)
	}

	models, err := svc.Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	for _, model := range models.Data.Items {
		if model.Scheduled {
			_, err := Client.NewDbtTransformationDeleteService().TransformationId(model.ID).Do(context.Background())
			if err != nil && err.Error() != "status code: 404; expected: 200" {
				log.Fatal(err)
			}
		}
	}

	if models.Data.NextCursor != "" {
		CleanupDbtTransformations(projectId, models.Data.NextCursor)
	}
}

func CreateTempGroup(t *testing.T) string {
	t.Helper()
	groupId := CreateGroup(t)
	t.Cleanup(func() { DeleteGroup(t, groupId) })
	return groupId
}

func DeleteGroup(t *testing.T, id string) {
	t.Helper()
	deleted, err := Client.NewGroupDelete().GroupID(id).Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", deleted)
		t.Error(err)
	}
}

func AddUserToGroup(t *testing.T, groupId string, email string) {
	t.Helper()
	created, err := Client.NewGroupAddUser().GroupID(groupId).Email(email).Role("Destination Administrator").Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", created)
		t.Error(err)
	}
}

func RemoveUserFromGroup(t *testing.T, groupId string, userId string) {
	t.Helper()
	deleted, err := Client.NewGroupRemoveUser().GroupID(groupId).UserID(userId).Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", deleted)
		t.Error(err)
	}
}

func DeleteDestination(t *testing.T, id string) {
	t.Helper()
	deleted, err := Client.NewDestinationDelete().DestinationID(id).Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", deleted)
		t.Error(err)
	}
}

func CreateDestination(t *testing.T) string {
	t.Helper()
	created, err := Client.NewDestinationCreate().
		GroupID(PredefinedGroupId).
		Service("snowflake").
		TimeZoneOffset("+10").
		RunSetupTests(false).
		Config(fivetran.NewDestinationConfig().
			Host("your-account.snowflakecomputing.com").
			Port(443).
			Database("fivetran").
			Auth("PASSWORD").
			User("fivetran_user").
			Password("123456")).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", created)
		t.Error(err)
	}
	return created.Data.ID
}

func CreateTempDestination(t *testing.T) string {
	t.Helper()
	destinationId := CreateDestination(t)
	t.Cleanup(func() { DeleteDestination(t, destinationId) })
	return destinationId
}

func CreateDbtTransformation(t *testing.T) string {
	t.Helper()
	created, err := Client.NewDbtTransformationCreateService().
		DbtModelId("").
		Schedule(fivetran.NewDbtTransformationSchedule().
			ScheduleType("INTEGRATED").
			DaysOfWeek([]string{}).
			Interval(0).
			TimeOfDay("")).
		RunTests(true).
		Paused(true).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", created)
		t.Error(err)
	}
	return created.Data.ID
}

func CreateTempDbtTransformation(t *testing.T) string {
	t.Helper()
	dbtTransformationId := CreateDbtTransformation(t)
	t.Cleanup(func() { DeleteDbtTransformation(t, dbtTransformationId) })
	return dbtTransformationId
}

func DeleteDbtTransformation(t *testing.T, id string) {
	t.Helper()
	deleted, err := Client.NewDbtTransformationDeleteService().
		TransformationId(id).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", deleted)
		t.Error(err)
	}
}

func CreateConnector(t *testing.T) string {
	t.Helper()
	created, err := Client.NewConnectorCreate().
		GroupID(PredefinedGroupId).
		Service("itunes_connect").
		RunSetupTests(false).
		Config(fivetran.NewConnectorConfig().
			Schema("itunes_e2e_connect").
			Username("fivetran").
			Password("fivetran-api-e2e")).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", created)
		t.Error(err)
	}

	return created.Data.ID
}

func CreateTempConnector(t *testing.T) string {
	t.Helper()
	connectorId := CreateConnector(t)
	t.Cleanup(func() { DeleteConnector(t, connectorId) })
	return connectorId
}

func DeleteConnector(t *testing.T, id string) {
	t.Helper()
	deleted, err := Client.NewConnectorDelete().ConnectorID(id).Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", deleted)
		t.Error(err)
	}
}

func printError(t *testing.T, actual interface{}, expected interface{}) {
	t.Helper()
	t.Errorf("Expected: %s"+
		"\n     but: <%s>\n",
		fmt.Sprintf("value equal to <%v>", expected),
		fmt.Sprintf("%v", actual),
	)
}

func AssertHasLength(t *testing.T, actual interface{}, expected int) {
	t.Helper()

	if actual == nil {
		printError(t, actual, fmt.Sprintf("value with length %v", expected))
	} else {
		lenOfActual := reflect.ValueOf(actual).Len()
		if lenOfActual != expected {
			printError(t, fmt.Sprintf("length was %d", lenOfActual), fmt.Sprintf("value with length %v", expected))
		}
	}
}

func AssertEqual(t *testing.T, actual interface{}, expected interface{}) {
	t.Helper()

	if !reflect.DeepEqual(expected, actual) {
		printError(t, actual, expected)
	}
}

func AssertTimeEqual(t *testing.T, actualTime time.Time, expectedTime string) {
	ex, _ := time.Parse(time.RFC3339, expectedTime)
	AssertEqual(t, ex, actualTime)
}

func AssertEmpty(t *testing.T, actual interface{}) {
	t.Helper()

	var isEmpty bool = isEmpty(actual)

	if !isEmpty {
		printError(t, actual, "empty value")
	}
}

func AssertNotEmpty(t *testing.T, actual interface{}) {
	t.Helper()

	var isEmpty bool = isEmpty(actual)

	if isEmpty {
		printError(t, actual, "none-empty value")
	}
}

func AssertKey(t *testing.T, key string, requestPart map[string]interface{}, expectedValue interface{}) {
	v, ok := requestPart[key]
	AssertEqual(t, ok, true)
	AssertEqual(t, v, expectedValue)
}

func AssertHasKey(t *testing.T, source map[string]interface{}, key string) {
	t.Helper()
	_, ok := source[key]
	if !ok {
		t.Errorf("Expected Key not found in map: %s", key)
	}
}

func AssertHasNoKey(t *testing.T, source map[string]interface{}, key string) {
	t.Helper()
	_, ok := source[key]
	if ok {
		t.Errorf("Unexpected Key found in map: %s", key)
	}
}

func BoolToStr(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

func AssertIsNotNil(t *testing.T, value interface{}) {
	t.Helper()

	if value == nil {
		printError(t, value, "non-nil value")
	}
}

func AssertKeyValue(t *testing.T, source map[string]interface{}, key string, expected interface{}) {
	t.Helper()
	AssertHasKey(t, source, key)
	actual := source[key]
	if !reflect.DeepEqual(actual, expected) {
		printError(t, actual, expected)
	}
}

func isEmpty(actual interface{}) bool {
	var isEmpty bool = false

	if actual == nil {
		isEmpty = true
	} else if actualValue, ok := actual.(string); ok {
		isEmpty = actualValue == ""
	} else if reflect.ValueOf(actual).Len() == 0 {
		isEmpty = true
	}
	return isEmpty
}

func CleanupAccount() {
	CleanupUsers()
	CleanupDestinations()
	CleanupDbtProjects()
	CleanupGroups()
	CleanupExternalLogging()
	CleanupPrivateLinks()
	CleanupWebhooks()
	CleanupTeams()
}

func IsPredefinedUserExist() bool {
	user, err := Client.NewUserDetails().UserID(PredefinedUserId).Do(context.Background())
	if err != nil {
		return false
	}
	return user.Data.GivenName == PredefinedUserGivenName
}

func IsPredefinedGroupExist() bool {
	group, err := Client.NewGroupDetails().GroupID(PredefinedGroupId).Do(context.Background())
	if err != nil {
		return false
	}
	return group.Data.Name == PredefinedGroupName
}

func CleanupUsers() {
	users, err := Client.NewUsersList().Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	for _, user := range users.Data.Items {
		if user.ID != PredefinedUserId {
			_, err := Client.NewUserDelete().UserID(user.ID).Do(context.Background())
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func CleanupDestinations() {
	groups, err := Client.NewGroupsList().Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	for _, group := range groups.Data.Items {
		_, err := Client.NewDestinationDelete().DestinationID(group.ID).Do(context.Background())
		if err != nil && err.Error() != "status code: 404; expected: 200" {
			log.Fatal(err)
		}
	}
}

func CleanupGroups() {
	groups, err := Client.NewGroupsList().Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	for _, group := range groups.Data.Items {
		CleanupConnectors(group.ID)
		if group.ID != PredefinedGroupId {
			_, err := Client.NewGroupDelete().GroupID(group.ID).Do(context.Background())
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func CleanupConnectors(groupId string) {
	connectors, err := Client.NewGroupListConnectors().GroupID(groupId).Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	for _, connector := range connectors.Data.Items {
		_, err := Client.NewConnectorDelete().ConnectorID(connector.ID).Do(context.Background())
		if err != nil {
			log.Fatal(err)
		}
	}
}

func CleanupExternalLogging() {
	groups, err := Client.NewGroupsList().Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	for _, group := range groups.Data.Items {
		_, err := Client.NewExternalLoggingDelete().ExternalLoggingId(group.ID).Do(context.Background())
		if err != nil && err.Error() != "status code: 404; expected: 200" {
			log.Fatal(err)
		}
	}
}

func CleanupPrivateLinks() {
	groups, err := Client.NewGroupsList().Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	for _, group := range groups.Data.Items {
		privateLinks, err := Client.NewGroupListPrivateLinks().GroupID(group.ID).Do(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		for _, privateLinks := range privateLinks.Data.Items {
			_, err := Client.NewPrivateLinksDelete().PrivateLinkId(privateLinks.Id).Do(context.Background())
			if err != nil && err.Error() != "status code: 404; expected: 200" {
				log.Fatal(err)
			}
		}
	}
}

func CleanupWebhooks() {
	list, err := Client.NewWebhookList().Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	for _, webhook := range list.Data.Items {
		_, err := Client.NewWebhookDelete().WebhookId(webhook.Id).Do(context.Background())
		if err != nil && err.Error() != "status code: 404; expected: 200" {
			log.Fatal(err)
		}
	}
}

func CleanupTeams() {
	list, err := Client.NewTeamsList().Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	for _, team := range list.Data.Items {
		_, err := Client.NewTeamsDelete().TeamId(team.Id).Do(context.Background())
		if err != nil && err.Error() != "status code: 404; expected: 200" {
			log.Fatal(err)
		}
	}

	if list.Data.NextCursor != "" {
		CleanupTeams()
	}
}

func CreateTempExternalLogging(t *testing.T) string {
	t.Helper()
	externalLoggingId := CreateExternalLogging(t)
	t.Cleanup(func() { DeleteExternalLogging(t, externalLoggingId) })
	return externalLoggingId
}

func DeleteExternalLogging(t *testing.T, id string) {
	t.Helper()
	deleted, err := Client.NewExternalLoggingDelete().ExternalLoggingId(id).Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", deleted)
		t.Error(err)
	}
}

func CreateExternalLogging(t *testing.T) string {
	t.Helper()
	created, err := Client.NewExternalLoggingCreate().
		GroupId(PredefinedGroupId).
		Service("azure_monitor_log").
		Enabled(true).
		Config(fivetran.NewExternalLoggingConfig().
			WorkspaceId("workspace_id").
			PrimaryKey("PASSWORD")).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", created)
		t.Error(err)
	}
	return created.Data.Id
}

/* Private Links */
func CreatePrivateLinkGroup(t *testing.T) string {
	t.Helper()
	created, err := Client.NewGroupCreate().Name("private_link_test").Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", created)
		t.Error(err)
	}
	return created.Data.ID
}

func CreatePrivateLinkDestination(t *testing.T, id string) string {
	t.Helper()
	destination, err := Client.NewDestinationCreate().
		GroupID(id).
		Service("big_query").
		Region("AWS_US_EAST_1").
		RunSetupTests(false).
		TimeZoneOffset("-5").
		Config(
			fivetran.NewDestinationConfig().
				ProjectID(BqProjectId).
				DataSetLocation("US")).
		Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", destination)
		t.Error(err)
	}
	return destination.Data.ID
}

func CreatePrivateLink(t *testing.T) (string, string, string) {
	t.Helper()
	plGroupId := CreatePrivateLinkGroup(t)
	t.Logf("groupId %+v\n", plGroupId)
	plDestinationId := CreatePrivateLinkDestination(t,plGroupId)
	t.Logf("plDestinationId %+v\n", plDestinationId)

	created, err := Client.NewPrivateLinksCreate().
		Name("test").
		GroupId(plGroupId).
		Service("redshift").
		Config(fivetran.NewPrivateLinksConfig().
			AwsAccountId("account_id.cloud_region_name.privatelink.snowflakecomputing.com").
    		ClusterIdentifier("account_id.cloud_region_name.privatelink.snowflakecomputing.com")).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", created)
		t.Error(err)
	}
	return created.Data.Id, plDestinationId, plGroupId
}

func DeletePrivateLink(t *testing.T, id string) {
	t.Helper()
	deleted, err := Client.NewPrivateLinksDelete().PrivateLinkId(id).Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", deleted)
		//t.Error(err)
	}
}

func CreateTempPrivateLink(t *testing.T) (string, string) {
	t.Helper()
	privateLinkId, plDestinationId, plGroupId := CreatePrivateLink(t)

	t.Cleanup(func() {
		//DeletePrivateLink(t, privateLinkId) 
		DeleteDestination(t, plDestinationId)
		DeleteGroup(t, plGroupId)
	})
	return privateLinkId, plGroupId
}

/* Private Links */

func CreateTempWebhook(t *testing.T) string {
	t.Helper()
	webhookId := CreateWebhookAccount(t)
	t.Cleanup(func() { DeleteWebhook(t, webhookId) })
	return webhookId
}

func DeleteWebhook(t *testing.T, id string) {
	t.Helper()
	deleted, err := Client.NewWebhookDelete().WebhookId(id).Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", deleted)
		t.Error(err)
	}
}

func CreateWebhookAccount(t *testing.T) string {
	t.Helper()
	created, err := Client.NewWebhookAccountCreate().
		Url("https://localhost:12345").
		Secret("my_secret").
		Active(false).
		Events([]string{"sync_start", "sync_end"}).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", created)
		t.Error(err)
	}
	return created.Data.Id
}


/* Begin Team Management */
func CreateTeam(t *testing.T) string {
	t.Helper()
	created, err := Client.NewTeamsCreate().
		Name("test_team").
		Description("test_description").
		Role("Account Reviewer").
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", created)
		t.Error(err)
	}
	return created.Data.Id
}

func DeleteTeamConnector(t *testing.T, teamId string, connectorId string) {
	t.Helper()
	deleted, err := Client.NewTeamConnectorMembershipDelete().TeamId(teamId).ConnectorId(connectorId).Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", deleted)
		t.Error(err)
	}
}

func CreateTeamConnector(t *testing.T, teamId string, connectorId string) {
	t.Helper()
	created, err := Client.NewTeamConnectorMembershipCreate().
		TeamId(teamId).
		ConnectorId(connectorId).
		Role("Connector Administrator").
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", created)
		t.Error(err)
	}
}

func DeleteTeamUser(t *testing.T, teamId string, userId string) {
	t.Helper()
	deleted, err := Client.NewTeamUserMembershipDelete().
		TeamId(teamId).
		UserId(userId).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", deleted)
		t.Error(err)
	}
}

func CreateTeamUser(t *testing.T, teamId string, userId string) {
	t.Helper()
	created, err := Client.NewTeamUserMembershipCreate().
		TeamId(teamId).
		UserId(userId).
		Role("Team Member").
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", created)
		t.Error(err)
	}
}

func DeleteTeamGroup(t *testing.T, teamId string, groupId string) {
	t.Helper()
	deleted, err := Client.NewTeamGroupMembershipDelete().
		TeamId(teamId).
		GroupId(groupId).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", deleted)
		t.Error(err)
	}
}

func CreateTeamGroup(t *testing.T, teamId string, groupId string) {
	t.Helper()
	created, err := Client.NewTeamGroupMembershipCreate().
		TeamId(teamId).
		GroupId(groupId).
		Role("Destination Analyst").
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", created)
		t.Error(err)
	}
}

func DeleteTeam(t *testing.T, id string) {
	t.Helper()
	deleted, err := Client.NewTeamsDelete().TeamId(id).Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", deleted)
		t.Error(err)
	}
}

/* End Team Management */

func DeleteUserConnector(t *testing.T, userId string, connectorId string) {
	t.Helper()
	deleted, err := Client.NewUserConnectorMembershipDelete().UserId(userId).ConnectorId(connectorId).Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", deleted)
		t.Error(err)
	}
}

func CreateUserConnector(t *testing.T, userId string, connectorId string) {
	t.Helper()
	created, err := Client.NewUserConnectorMembershipCreate().
		UserId(userId).
		ConnectorId(connectorId).
		Role("Connector Administrator").
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", created)
		t.Error(err)
	}
}

func DeleteUserGroup(t *testing.T, userId string, groupId string) {
	t.Helper()
	deleted, err := Client.NewUserGroupMembershipDelete().
		UserId(userId).
		GroupId(groupId).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", deleted)
		t.Error(err)
	}
}

func CreateUserGroup(t *testing.T, userId string, groupId string) {
	t.Helper()
	created, err := Client.NewUserGroupMembershipCreate().
		UserId(userId).
		GroupId(groupId).
		Role("Destination Analyst").
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", created)
		t.Error(err)
	}
}