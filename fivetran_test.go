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

var apiKey string
var apiSecret string
var CertificateHash string
var EncodedCertificate string
var PredefinedGroupId string = "climbed_consulted"
var PredefinedUserId string = "cherry_spoilt"

func init() {
	apiKey = os.Getenv("FIVETRAN_APIKEY")
	apiSecret = os.Getenv("FIVETRAN_APISECRET")
	CertificateHash = os.Getenv("FIVETRAN_TEST_CERTIFICATE_HASH")
	EncodedCertificate = os.Getenv("FIVETRAN_TEST_CERTIFICATE")
	Client = fivetran.New(apiKey, apiSecret)
	Client.BaseURL("https://api.fivetran.com/v1")
	cleanupAccount()
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
		GroupID("climbed_consulted").
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

func CreateConnector(t *testing.T) string {
	t.Helper()
	created, err := Client.NewConnectorCreate().
		GroupID("climbed_consulted").
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
	cleanupGroups()
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