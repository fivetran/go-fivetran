package fivetran_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/fivetran/go-fivetran"
)

var Client *fivetran.Client

var CertificateHash string
var EncodedCertificate string

var PredefinedGroupId string
var PredefinedUserId string

// Tests should be re-written to not use a pre-defined user and group
const (
	PredefinedGroupName      string = "GoSdkTesting"
	PredefinedUserEmail      string = "dev-markov+go-fivetran-sdk@fivetran.com"
	PredefinedUserGivenName  string = "Go"
	PredefinedUserFamilyName string = "5Tran"
	PredefinedUserPhone      string = "+1234567890"
)

func init() {
	var apiUrl string
	var apiKey string
	var apiSecret string

	valuesToLoad := map[string]*string{
		"FIVETRAN_API_URL":               &apiUrl,
		"FIVETRAN_APIKEY":                &apiKey,
		"FIVETRAN_APISECRET":             &apiSecret,
		"FIVETRAN_TEST_CERTIFICATE_HASH": &CertificateHash,
		"FIVETRAN_TEST_CERTIFICATE":      &EncodedCertificate,
		"FIVETRAN_GROUP_ID":              &PredefinedGroupId,
		"FIVETRAN_USER_ID":               &PredefinedUserId,
	}

	for name, value := range valuesToLoad {
		*value = os.Getenv(name)
		if *value == "" {
			log.Fatalf("Environment variable %s is not set!\n", name)
		}
	}

	Client = fivetran.New(apiKey, apiSecret)
	Client.BaseURL(apiUrl)
	if isPredefinedUserExist() {
		cleanupAccount()
	} else {
		log.Fatalln("The predefined user doesn't belong to the Testing account. Make sure that credentials are using in the tests belong to the Testing account.")
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
	created, err := Client.NewDbtProjectCreate().
		GroupID(PredefinedGroupId).
		DbtVersion("1.3.1").
		ProjectConfig(fivetran.NewDbtProjectConfig().
			GitRemoteUrl("https://github.com/fivetran/dbt_demo").
			FolderPath("").
			GitBranch("main")).
		DefaultSchema("").
		TargetName("").
		Threads(4).
		Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", created)
		t.Error(err)
	}
	return created.Data.ID
}

func cleanupDbtProjects() {
	projects, err := Client.NewDbtProjectsList().Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	for _, project := range projects.Data.Items {
		cleanupDbtTransformations(project.ID, "")
		_, err := Client.NewDbtProjectDelete().DbtProjectID(project.ID).Do(context.Background())
		if err != nil && err.Error() != "status code: 404; expected: 200" {
			log.Fatal(err)
		}
	}
	if projects.Data.NextCursor != "" {
		cleanupDbtProjects()
	}
}

func cleanupDbtTransformations(projectId, nextCursor string) {
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
		cleanupDbtTransformations(projectId, models.Data.NextCursor)
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

func cleanupAccount() {
	cleanupUsers()
	cleanupDestinations()
	cleanupDbtProjects()
	cleanupGroups()
	cleanupExternalLogging()
	cleanupWebhooks()
}

func isPredefinedUserExist() bool {
	user, err := Client.NewUserDetails().UserID(PredefinedUserId).Do(context.Background())
	if err != nil {
		return false
	}
	return user.Data.ID == PredefinedUserId
}

func cleanupUsers() {
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

func cleanupDestinations() {
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

func cleanupGroups() {
	groups, err := Client.NewGroupsList().Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	for _, group := range groups.Data.Items {
		cleanupConnectors(group.ID)
		if group.ID != PredefinedGroupId {
			_, err := Client.NewGroupDelete().GroupID(group.ID).Do(context.Background())
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func cleanupConnectors(groupId string) {
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

func cleanupExternalLogging() {
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

func cleanupWebhooks() {
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
