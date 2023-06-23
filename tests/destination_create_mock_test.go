package tests

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/tests/mock"
)

const (
	SERVICE                  = "test_service"
	ID                       = "test_id"
	REGION                   = "GCP_US_EAST4"
	TIME_ZONE                = "-5"
	SETUP_STATUS             = "connected"
	TEST_TITLE               = "Test Title"
	TEST_STATUS              = "PASSED"
	TEST_MESSAGE             = "Test message"
	HOST                     = "your.host"
	PORT                     = 443
	DATABASE                 = "fivetran"
	AUTH                     = "PASSWORD"
	USER                     = "fivetran_user"
	MASKED                   = "******"
	CONNECTION_TYPE          = "Directly"
	TUNNEL_HOST              = "tunnel.host"
	TUNNEL_PORT              = "334"
	TUNNEL_USER              = "tunnel_user"
	PROJECT_ID               = "project_id_value"
	DATA_SET_LOCATION        = "data_Set_location_value"
	LOCATION                 = "data_Set_location_value"
	BUCKET                   = "your-bucket"
	SERVER_HOST_NAME         = "server.host.name"
	HTTP_PATH                = "http.path"
	CREATE_EXTERNAL_TABLES   = true
	EXTERNAL_LOCATION        = "group_id"
	AUTH_TYPE                = "auth_type_value"
	ROLE_ARN                 = "role:arn-xxx"
	PUBLIC_KEY               = "public_key_value"
	CLUSTER_ID               = "cluster_id_value"
	CLUSTER_REGION           = "cluster_region_value"
	PASSWORD                 = "password"
	PRIVATE_KEY              = "private_key"
	SECRET_KEY               = "secret_key"
	TRUST_CERTIFICATES       = true
	TRUST_FINGERPRINTS       = true
	RUN_SETUP_TESTS          = true
	PERSONAL_ACCESS_TOKEN    = "PAT"
	ROLE                     = "role"
	PASSPHRASE               = "passphrase"
	IS_PRIVATE_KEY_ENCRYPTED = true
	CATALOG                  = "catalog"
	FIVETRAN_ROLE_ARN        = "fivetran_role_arn"
	PREFIX_PATH              = "prefix_path"
)

func TestNewDestinationCreateFullMappingMock(t *testing.T) {
	// arrange
	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(http.MethodPost, "/v1/destinations").ThenCall(

		func(req *http.Request) (*http.Response, error) {
			body := requestBodyToJson(t, req)
			assertRequest(t, body)
			response := mock.NewResponse(req, http.StatusCreated, prepareDestinationResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewDestinationCreate().
		Service(SERVICE).
		Region(REGION).
		TimeZoneOffset(TIME_ZONE).
		TrustCertificates(TRUST_CERTIFICATES).
		TrustFingerprints(TRUST_FINGERPRINTS).
		RunSetupTests(RUN_SETUP_TESTS).
		GroupID(ID).
		Config(prepareConfig()).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", response)
		t.Error(err)
	}

	// assert
	interactions := mockClient.Interactions()
	assertEqual(t, len(interactions), 1)
	assertEqual(t, interactions[0].Handler, handler)
	assertEqual(t, handler.Interactions, 1)
	assertResponse(t, response)
}

func prepareDestinationResponse() string {
	return fmt.Sprintf(
		`{
			"code":"Created",
			"message":"Destination has been created",
			"data":{
				"id":                           "%v",
				"group_id":                     "%v",
				"service":                      "%v",
				"region":                       "%v",
				"time_zone_offset":             "%v",
				"setup_status":                 "%v",
				"setup_tests":[
					{
						"title":                "%v",
						"status":               "%v",
						"message":              "%v"
					}
				],
				"config":{
					"host":                     "%v",
					"port":                     "%v",
					"database":                 "%v",
					"auth":                     "%v",
					"user":                     "%v",
					"password":                 "%v",
					"connection_type":          "%v",
					"tunnel_host":              "%v",
					"tunnel_port":              "%v",
					"tunnel_user":              "%v",
					"project_id":               "%v",
					"data_set_location":        "%v",
					"location":                 "%v",
					"bucket":                   "%v",
					"server_host_name":         "%v",
					"http_path":                "%v",
					"personal_access_token":    "%v",
					"create_external_tables":   "%v",
					"external_location":        "%v",
					"auth_type":                "%v",
					"role_arn":                 "%v",
					"secret_key":               "%v",
					"private_key":              "%v",
					"public_key":               "%v",
					"cluster_id":               "%v",
					"cluster_region":           "%v",
					"role":                     "%v",
					"is_private_key_encrypted": "%v",
					"passphrase":               "%v",
					"catalog": "%v",
					"fivetran_role_arn":"%v",
					"prefix_path":"%v",
					"region": "%v",
				}
			}
		}`,
		ID, // id
		ID, // group_id
		SERVICE,
		REGION,
		TIME_ZONE, // time_zone_offset
		SETUP_STATUS,
		TEST_TITLE,
		TEST_STATUS,
		TEST_MESSAGE,
		HOST,
		PORT,
		DATABASE,
		AUTH,
		USER,
		MASKED, // password
		CONNECTION_TYPE,
		TUNNEL_HOST,
		TUNNEL_PORT,
		TUNNEL_USER,
		PROJECT_ID,
		DATA_SET_LOCATION,
		LOCATION,
		BUCKET,
		SERVER_HOST_NAME,
		HTTP_PATH,
		MASKED, // personal_access_token
		CREATE_EXTERNAL_TABLES,
		EXTERNAL_LOCATION,
		AUTH_TYPE,
		ROLE_ARN,
		MASKED, // secret_key
		MASKED, // private_key
		PUBLIC_KEY,
		CLUSTER_ID,
		CLUSTER_REGION,
		ROLE,
		IS_PRIVATE_KEY_ENCRYPTED,
		PASSPHRASE,
		CATALOG,
		FIVETRAN_ROLE_ARN,
		PREFIX_PATH,
		REGION,
	)
}

func prepareConfig() *fivetran.DestinationConfig {
	config := fivetran.NewDestinationConfig()
	config.Host(HOST)
	config.Port(PORT)
	config.Database(DATABASE)
	config.Auth(AUTH)
	config.User(USER)
	config.Password(PASSWORD)
	config.ConnectionType(CONNECTION_TYPE)
	config.TunnelHost(TUNNEL_HOST)
	config.TunnelPort(TUNNEL_PORT)
	config.TunnelUser(TUNNEL_USER)
	config.ProjectID(PROJECT_ID)
	config.Bucket(BUCKET)
	config.ServerHostName(SERVER_HOST_NAME)
	config.HTTPPath(HTTP_PATH)
	config.PersonalAccessToken(PERSONAL_ACCESS_TOKEN)
	config.CreateExternalTables(CREATE_EXTERNAL_TABLES)
	config.ExternalLocation(EXTERNAL_LOCATION)
	config.AuthType(AUTH_TYPE)
	config.RoleArn(ROLE_ARN)
	config.SecretKey(SECRET_KEY)
	config.PrivateKey(PRIVATE_KEY)
	config.ClusterId(CLUSTER_ID)
	config.ClusterRegion(CLUSTER_REGION)
	config.Role(ROLE)
	config.IsPrivateKeyEncrypted(IS_PRIVATE_KEY_ENCRYPTED)
	config.Passphrase(PASSPHRASE)
	config.Catalog(CATALOG)
	config.FivetranRoleArn(FIVETRAN_ROLE_ARN)
	config.PrefixPath(PREFIX_PATH)
	config.Region(REGION)
	return config
}

func assertRequest(t *testing.T, request map[string]interface{}) {
	assertKey(t, "service", request, SERVICE)
	assertKey(t, "region", request, REGION)
	assertKey(t, "time_zone_offset", request, TIME_ZONE)
	assertKey(t, "trust_certificates", request, TRUST_CERTIFICATES)
	assertKey(t, "trust_fingerprints", request, TRUST_FINGERPRINTS)
	assertKey(t, "run_setup_tests", request, RUN_SETUP_TESTS)
	assertKey(t, "group_id", request, ID)

	c, ok := request["config"]
	assertEqual(t, ok, true)
	config, ok := c.(map[string]interface{})
	assertEqual(t, ok, true)

	assertKey(t, "host", config, HOST)
	assertKey(t, "port", config, float64(PORT)) // json marshalling stores all numbers as float64
	assertKey(t, "database", config, DATABASE)
	assertKey(t, "auth", config, AUTH)
	assertKey(t, "user", config, USER)
	assertKey(t, "password", config, PASSWORD)
	assertKey(t, "connection_type", config, CONNECTION_TYPE)
	assertKey(t, "tunnel_host", config, TUNNEL_HOST)
	assertKey(t, "tunnel_port", config, TUNNEL_PORT)
	assertKey(t, "tunnel_user", config, TUNNEL_USER)
	assertKey(t, "project_id", config, PROJECT_ID)
	assertKey(t, "bucket", config, BUCKET)
	assertKey(t, "server_host_name", config, SERVER_HOST_NAME)
	assertKey(t, "http_path", config, HTTP_PATH)
	assertKey(t, "personal_access_token", config, PERSONAL_ACCESS_TOKEN)
	assertKey(t, "create_external_tables", config, CREATE_EXTERNAL_TABLES)
	assertKey(t, "external_location", config, EXTERNAL_LOCATION)
	assertKey(t, "auth_type", config, AUTH_TYPE)
	assertKey(t, "role_arn", config, ROLE_ARN)
	assertKey(t, "secret_key", config, SECRET_KEY)
	assertKey(t, "private_key", config, PRIVATE_KEY)
	assertKey(t, "cluster_id", config, CLUSTER_ID)
	assertKey(t, "cluster_region", config, CLUSTER_REGION)
	assertKey(t, "role", config, ROLE)
	assertKey(t, "is_private_key_encrypted", config, IS_PRIVATE_KEY_ENCRYPTED)
	assertKey(t, "passphrase", config, PASSPHRASE)
	assertKey(t, "catalog", config, CATALOG)
	assertKey(t, "fivetran_role_arn", config, FIVETRAN_ROLE_ARN)
	assertKey(t, "prefix_path", config, PREFIX_PATH)
	assertKey(t, "region", config, REGION)
}

func assertResponse(t *testing.T, response fivetran.DestinationCreateResponse) {

	assertEqual(t, response.Code, "Created")
	assertNotEmpty(t, response.Message)

	assertEqual(t, response.Data.GroupID, ID)
	assertEqual(t, response.Data.ID, ID)
	assertEqual(t, response.Data.Region, REGION)
	assertEqual(t, response.Data.Service, SERVICE)
	assertEqual(t, response.Data.SetupStatus, SETUP_STATUS)
	assertEqual(t, response.Data.TimeZoneOffset, TIME_ZONE)

	assertEqual(t, len(response.Data.SetupTests), 1)
	assertEqual(t, response.Data.SetupTests[0].Message, TEST_MESSAGE)
	assertEqual(t, response.Data.SetupTests[0].Status, TEST_STATUS)
	assertEqual(t, response.Data.SetupTests[0].Title, TEST_TITLE)

	assertEqual(t, response.Data.Config.Auth, AUTH)
	assertEqual(t, response.Data.Config.AuthType, AUTH_TYPE)
	assertEqual(t, response.Data.Config.Bucket, BUCKET)
	assertEqual(t, response.Data.Config.ClusterId, CLUSTER_ID)
	assertEqual(t, response.Data.Config.ClusterRegion, CLUSTER_REGION)
	assertEqual(t, response.Data.Config.ConnectionType, CONNECTION_TYPE)
	assertEqual(t, response.Data.Config.CreateExternalTables, boolToStr(CREATE_EXTERNAL_TABLES)) // Inconsistent response
	assertEqual(t, response.Data.Config.DataSetLocation, DATA_SET_LOCATION)
	assertEqual(t, response.Data.Config.Database, DATABASE)
	assertEqual(t, response.Data.Config.ExternalLocation, EXTERNAL_LOCATION)
	assertEqual(t, response.Data.Config.HTTPPath, HTTP_PATH)
	assertEqual(t, response.Data.Config.Host, HOST)
	assertEqual(t, response.Data.Config.Password, MASKED)
	assertEqual(t, response.Data.Config.PersonalAccessToken, MASKED)
	assertEqual(t, response.Data.Config.Port, strconv.Itoa(PORT))
	assertEqual(t, response.Data.Config.PrivateKey, MASKED)
	assertEqual(t, response.Data.Config.ProjectID, PROJECT_ID)
	assertEqual(t, response.Data.Config.PublicKey, PUBLIC_KEY)
	assertEqual(t, response.Data.Config.RoleArn, ROLE_ARN)
	assertEqual(t, response.Data.Config.SecretKey, MASKED)
	assertEqual(t, response.Data.Config.ServerHostName, SERVER_HOST_NAME)
	assertEqual(t, response.Data.Config.TunnelHost, TUNNEL_HOST)
	assertEqual(t, response.Data.Config.TunnelPort, TUNNEL_PORT)
	assertEqual(t, response.Data.Config.TunnelUser, TUNNEL_USER)
	assertEqual(t, response.Data.Config.User, USER)
	assertEqual(t, response.Data.Config.Location, LOCATION)
	assertEqual(t, response.Data.Config.Role, ROLE)
	assertEqual(t, response.Data.Config.IsPrivateKeyEncrypted, boolToStr(IS_PRIVATE_KEY_ENCRYPTED))
	assertEqual(t, response.Data.Config.Passphrase, PASSPHRASE)
	assertEqual(t, response.Data.Config.Catalog, CATALOG)
	assertEqual(t, response.Data.Config.FivetranRoleArn, FIVETRAN_ROLE_ARN)
	assertEqual(t, response.Data.Config.PrefixPath, PREFIX_PATH)
	assertEqual(t, response.Data.Config.Region, REGION)
}
