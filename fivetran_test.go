package fivetran_test

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/fivetran/go-fivetran"
)

var Clients map[string]*fivetran.Client

var CertificateHash string
var EncodedCertificate string
var PredefinedGroupId string = "climbed_consulted"
var PredefinedUserId string = "cherry_spoilt"

func init() {
	Clients = getClients()
}

func CreateUser(t *testing.T) string {
	t.Helper()
	user, err := Clients["v1"].NewUserInvite().
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
	user, err := Clients["v1"].NewUserDelete().UserID(id).Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", user)
		t.Error(err)
	}
}

func CreateGroup(t *testing.T) string {
	t.Helper()
	created, err := Clients["v1"].NewGroupCreate().Name("test").Do(context.Background())
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
	deleted, err := Clients["v1"].NewGroupDelete().GroupID(id).Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", deleted)
		t.Error(err)
	}
}

func AddUserToGroup(t *testing.T, groupId string, email string) {
	t.Helper()
	created, err := Clients["v1"].NewGroupAddUser().GroupID(groupId).Email(email).Role("Destination Administrator").Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", created)
		t.Error(err)
	}
}

func RemoveUserFromGroup(t *testing.T, groupId string, userId string) {
	t.Helper()
	deleted, err := Clients["v1"].NewGroupRemoveUser().GroupID(groupId).UserID(userId).Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", deleted)
		t.Error(err)
	}
}

func DeleteDestination(t *testing.T, id string) {
	t.Helper()
	deleted, err := Clients["v1"].NewDestinationDelete().DestinationID(id).Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", deleted)
		t.Error(err)
	}
}

func CreateDestination(t *testing.T) string {
	t.Helper()
	created, err := Clients["v1"].NewDestinationCreate().
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
	destinationId := CreateDestination(t)
	t.Cleanup(func() { DeleteDestination(t, destinationId) })
	return destinationId
}

func CreateConnector(t *testing.T) string {
	created, err := Clients["v1"].NewConnectorCreate().
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
	deleted, err := Clients["v1"].NewConnectorDelete().ConnectorID(id).Do(context.Background())

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

func getClients() map[string]*fivetran.Client {
	apiKey := os.Getenv("FIVETRAN_API_KEY")
	apiSecret := os.Getenv("FIVETRAN_API_SECRET")
	CertificateHash = os.Getenv("FIVETRAN_TEST_CERTIFICATE_HASH")
	EncodedCertificate = os.Getenv("FIVETRAN_TEST_CERTIFICATE")

	clients := make(map[string]*fivetran.Client)

	versions := [...]string{"v1", "v2"}
	for _, version := range versions {
		client := fivetran.New(apiKey, apiSecret)
		client.BaseURL("https://api.fivetran.com/" + version)
		clients[version] = client
	}
	return clients
}
