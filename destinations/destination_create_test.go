package destinations_test

import (
    "context"
    "fmt"
    "net/http"
    "strconv"
    "testing"

    "github.com/fivetran/go-fivetran"
    "github.com/fivetran/go-fivetran/destinations"

    "github.com/fivetran/go-fivetran/tests/mock"

    testutils "github.com/fivetran/go-fivetran/test_utils"
)

const (
    SERVICE                   = "test_service"
    ID                        = "test_id"
    REGION                    = "GCP_US_EAST4"
    TIME_ZONE                 = "-5"
    SETUP_STATUS              = "connected"
    TEST_TITLE                = "Test Title"
    TEST_STATUS               = "PASSED"
    TEST_MESSAGE              = "Test message"
    HOST                      = "your.host"
    PORT                      = 443
    DATABASE                  = "fivetran"
    AUTH                      = "PASSWORD"
    USER                      = "fivetran_user"
    MASKED                    = "******"
    CONNECTION_TYPE           = "Directly"
    TUNNEL_HOST               = "tunnel.host"
    TUNNEL_PORT               = "334"
    TUNNEL_USER               = "tunnel_user"
    PROJECT_ID                = "project_id_value"
    DATA_SET_LOCATION         = "data_Set_location_value"
    LOCATION                  = "data_Set_location_value"
    BUCKET                    = "your-bucket"
    SERVER_HOST_NAME          = "server.host.name"
    HTTP_PATH                 = "http.path"
    CREATE_EXTERNAL_TABLES    = true
    EXTERNAL_LOCATION         = "group_id"
    AUTH_TYPE                 = "auth_type_value"
    ROLE_ARN                  = "role:arn-xxx"
    PUBLIC_KEY                = "public_key_value"
    CLUSTER_ID                = "cluster_id_value"
    CLUSTER_REGION            = "cluster_region_value"
    PASSWORD                  = "password"
    PRIVATE_KEY               = "private_key"
    SECRET_KEY                = "secret_key"
    TRUST_CERTIFICATES        = true
    TRUST_FINGERPRINTS        = true
    RUN_SETUP_TESTS           = true
    DAYLIGHTSAVINGTIMEENABLED = true
    PERSONAL_ACCESS_TOKEN     = "PAT"
    ROLE                      = "role"
    PASSPHRASE                = "passphrase"
    IS_PRIVATE_KEY_ENCRYPTED  = true
    CATALOG                   = "catalog"
    FIVETRAN_ROLE_ARN         = "fivetran_role_arn"
    PREFIX_PATH               = "prefix_path"
    STORAGE_ACCOUNT_NAME      = "storage_account_name"
    CONTAINER_NAME            = "container_name"
    TENANT_ID                 = "tenant_id"
    CLIENT_ID                 = "client_id"
    SECRET_VALUE              = "secret_value"
    WORKSPACE_NAME            = "workspace_name"
    LAKEHOUSE_NAME            = "lakehouse_name"
    HYBRIDDEPLOYMENTAGENTID   = "hybrid_deployment_agent_id"
    PRIVATELINKID             = "private_link_id"
    NETWORKINGMETHOD          = "Direct"
)

func TestNewDestinationCreateFullMappingMock(t *testing.T) {
    // arrange
    ftClient, mockClient := testutils.CreateTestClient()
    handler := mockClient.When(http.MethodPost, "/v1/destinations").ThenCall(
        func(req *http.Request) (*http.Response, error) {
            body := testutils.RequestBodyToJson(t, req)
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
        DaylightSavingTimeEnabled(DAYLIGHTSAVINGTIMEENABLED).
        HybridDeploymentAgentId(HYBRIDDEPLOYMENTAGENTID).
        PrivateLinkId(PRIVATELINKID).
        NetworkingMethod(NETWORKINGMETHOD).
        GroupID(ID).
        Config(prepareConfig()).
        Do(context.Background())

    if err != nil {
        t.Logf("%+v\n", response)
        t.Error(err)
    }

    // testutils.Assert
    interactions := mockClient.Interactions()
    testutils.AssertEqual(t, len(interactions), 1)
    testutils.AssertEqual(t, interactions[0].Handler, handler)
    testutils.AssertEqual(t, handler.Interactions, 1)
    assertResponse(t, response)
}

func TestNewDestinationCreateCustomFullMappingMock(t *testing.T) {
    // arrange
    ftClient, mockClient := testutils.CreateTestClient()
    handler := mockClient.When(http.MethodPost, "/v1/destinations").ThenCall(
        func(req *http.Request) (*http.Response, error) {
            body := testutils.RequestBodyToJson(t, req)
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
        DaylightSavingTimeEnabled(DAYLIGHTSAVINGTIMEENABLED).
        HybridDeploymentAgentId(HYBRIDDEPLOYMENTAGENTID).
        PrivateLinkId(PRIVATELINKID).
        NetworkingMethod(NETWORKINGMETHOD).
        GroupID(ID).
        ConfigCustom(prepareConfigCustom()).
        DoCustom(context.Background())

    if err != nil {
        t.Logf("%+v\n", response)
        t.Error(err)
    }

    // testutils.Assert
    interactions := mockClient.Interactions()
    testutils.AssertEqual(t, len(interactions), 1)
    testutils.AssertEqual(t, interactions[0].Handler, handler)
    testutils.AssertEqual(t, handler.Interactions, 1)
    asserResponseCustom(t, response)
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
                "daylight_saving_time_enabled": %v,
                "hybrid_deployment_agent_id":   "%v",
                "private_link_id":              "%v",
                "networking_method":            "%v",
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
                    "is_private_key_encrypted": %v,
                    "passphrase":               "%v",
                    "catalog":                  "%v",
                    "fivetran_role_arn":        "%v",
                    "prefix_path":              "%v",
                    "region":                   "%v",
                    "storage_account_name":     "%v",
                    "container_name":           "%v",
                    "tenant_id":                "%v",
                    "client_id":                "%v",
                    "secret_value":             "%v",
                    "workspace_name":           "%v",
                    "lakehouse_name":           "%v"
                }
            }
        }`,
        ID, // id
        ID, // group_id
        SERVICE,
        REGION,
        TIME_ZONE, // time_zone_offset
        DAYLIGHTSAVINGTIMEENABLED,
        HYBRIDDEPLOYMENTAGENTID,
        PRIVATELINKID,
        NETWORKINGMETHOD,
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
        STORAGE_ACCOUNT_NAME,
        CONTAINER_NAME,
        TENANT_ID,
        CLIENT_ID,
        SECRET_VALUE,
        WORKSPACE_NAME,
        LAKEHOUSE_NAME,
    )
}

func prepareConfigCustom() *map[string]interface{} {
    return &(map[string]interface{}{
        "host":                     HOST,
        "port":                     PORT,
        "database":                 DATABASE,
        "auth":                     AUTH,
        "user":                     USER,
        "password":                 PASSWORD,
        "connection_type":          CONNECTION_TYPE,
        "tunnel_host":              TUNNEL_HOST,
        "tunnel_port":              TUNNEL_PORT,
        "tunnel_user":              TUNNEL_USER,
        "project_id":               PROJECT_ID,
        "bucket":                   BUCKET,
        "server_host_name":         SERVER_HOST_NAME,
        "http_path":                HTTP_PATH,
        "personal_access_token":    PERSONAL_ACCESS_TOKEN,
        "create_external_tables":   CREATE_EXTERNAL_TABLES,
        "external_location":        EXTERNAL_LOCATION,
        "auth_type":                AUTH_TYPE,
        "role_arn":                 ROLE_ARN,
        "secret_key":               SECRET_KEY,
        "private_key":              PRIVATE_KEY,
        "cluster_id":               CLUSTER_ID,
        "cluster_region":           CLUSTER_REGION,
        "role":                     ROLE,
        "is_private_key_encrypted": IS_PRIVATE_KEY_ENCRYPTED,
        "passphrase":               PASSPHRASE,
        "catalog":                  CATALOG,
        "fivetran_role_arn":        FIVETRAN_ROLE_ARN,
        "region":                   REGION,
        "prefix_path":              PREFIX_PATH,
        "storage_account_name":     STORAGE_ACCOUNT_NAME,
        "container_name":           CONTAINER_NAME,
        "tenant_id":                TENANT_ID,
        "client_id":                CLIENT_ID,
        "secret_value":             SECRET_VALUE,
        "workspace_name":           WORKSPACE_NAME,
        "lakehouse_name":           LAKEHOUSE_NAME,
    })
}

func prepareConfig() *destinations.DestinationConfig {
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
    config.StorageAccountName(STORAGE_ACCOUNT_NAME)
    config.ContainerName(CONTAINER_NAME)
    config.TenantId(TENANT_ID)
    config.ClientId(CLIENT_ID)
    config.SecretValue(SECRET_VALUE)
    config.WorkspaceName(WORKSPACE_NAME)
    config.LakehouseName(LAKEHOUSE_NAME)

    return config
}

func assertRequest(t *testing.T, request map[string]interface{}) {
    testutils.AssertKey(t, "service", request, SERVICE)
    testutils.AssertKey(t, "region", request, REGION)
    testutils.AssertKey(t, "time_zone_offset", request, TIME_ZONE)
    testutils.AssertKey(t, "trust_certificates", request, TRUST_CERTIFICATES)
    testutils.AssertKey(t, "trust_fingerprints", request, TRUST_FINGERPRINTS)
    testutils.AssertKey(t, "run_setup_tests", request, RUN_SETUP_TESTS)
    testutils.AssertKey(t, "daylight_saving_time_enabled", request, DAYLIGHTSAVINGTIMEENABLED)
    testutils.AssertKey(t, "hybrid_deployment_agent_id", request, HYBRIDDEPLOYMENTAGENTID)
    testutils.AssertKey(t, "private_link_id", request, PRIVATELINKID)
    testutils.AssertKey(t, "networking_method", request, NETWORKINGMETHOD)
    testutils.AssertKey(t, "group_id", request, ID)

    c, ok := request["config"]
    testutils.AssertEqual(t, ok, true)
    config, ok := c.(map[string]interface{})
    testutils.AssertEqual(t, ok, true)

    testutils.AssertKey(t, "host", config, HOST)
    testutils.AssertKey(t, "port", config, float64(PORT)) // json marshalling stores all numbers as float64
    testutils.AssertKey(t, "database", config, DATABASE)
    testutils.AssertKey(t, "auth", config, AUTH)
    testutils.AssertKey(t, "user", config, USER)
    testutils.AssertKey(t, "password", config, PASSWORD)
    testutils.AssertKey(t, "connection_type", config, CONNECTION_TYPE)
    testutils.AssertKey(t, "tunnel_host", config, TUNNEL_HOST)
    testutils.AssertKey(t, "tunnel_port", config, TUNNEL_PORT)
    testutils.AssertKey(t, "tunnel_user", config, TUNNEL_USER)
    testutils.AssertKey(t, "project_id", config, PROJECT_ID)
    testutils.AssertKey(t, "bucket", config, BUCKET)
    testutils.AssertKey(t, "server_host_name", config, SERVER_HOST_NAME)
    testutils.AssertKey(t, "http_path", config, HTTP_PATH)
    testutils.AssertKey(t, "personal_access_token", config, PERSONAL_ACCESS_TOKEN)
    testutils.AssertKey(t, "create_external_tables", config, CREATE_EXTERNAL_TABLES)
    testutils.AssertKey(t, "external_location", config, EXTERNAL_LOCATION)
    testutils.AssertKey(t, "auth_type", config, AUTH_TYPE)
    testutils.AssertKey(t, "role_arn", config, ROLE_ARN)
    testutils.AssertKey(t, "secret_key", config, SECRET_KEY)
    testutils.AssertKey(t, "private_key", config, PRIVATE_KEY)
    testutils.AssertKey(t, "cluster_id", config, CLUSTER_ID)
    testutils.AssertKey(t, "cluster_region", config, CLUSTER_REGION)
    testutils.AssertKey(t, "role", config, ROLE)
    testutils.AssertKey(t, "is_private_key_encrypted", config, IS_PRIVATE_KEY_ENCRYPTED)
    testutils.AssertKey(t, "passphrase", config, PASSPHRASE)
    testutils.AssertKey(t, "catalog", config, CATALOG)
    testutils.AssertKey(t, "fivetran_role_arn", config, FIVETRAN_ROLE_ARN)
    testutils.AssertKey(t, "prefix_path", config, PREFIX_PATH)
    testutils.AssertKey(t, "region", config, REGION)
    testutils.AssertKey(t, "storage_account_name", config, STORAGE_ACCOUNT_NAME)
    testutils.AssertKey(t, "container_name", config, CONTAINER_NAME)
    testutils.AssertKey(t, "tenant_id", config, TENANT_ID)
    testutils.AssertKey(t, "client_id", config, CLIENT_ID)
    testutils.AssertKey(t, "secret_value", config, SECRET_VALUE)
    testutils.AssertKey(t, "workspace_name", config, WORKSPACE_NAME)
    testutils.AssertKey(t, "lakehouse_name", config, LAKEHOUSE_NAME)
}

func asserResponseCustom(t *testing.T, response destinations.DestinationDetailsWithSetupTestsCustomResponse) {
    testutils.AssertEqual(t, response.Code, "Created")
    testutils.AssertNotEmpty(t, response.Message)

    testutils.AssertEqual(t, response.Data.GroupID, ID)
    testutils.AssertEqual(t, response.Data.ID, ID)
    testutils.AssertEqual(t, response.Data.Region, REGION)
    testutils.AssertEqual(t, response.Data.Service, SERVICE)
    testutils.AssertEqual(t, response.Data.DaylightSavingTimeEnabled, DAYLIGHTSAVINGTIMEENABLED)
    testutils.AssertEqual(t, response.Data.HybridDeploymentAgentId, HYBRIDDEPLOYMENTAGENTID)
    testutils.AssertEqual(t, response.Data.PrivateLinkId, PRIVATELINKID)
    testutils.AssertEqual(t, response.Data.NetworkingMethod, NETWORKINGMETHOD)
    testutils.AssertEqual(t, response.Data.SetupStatus, SETUP_STATUS)
    testutils.AssertEqual(t, response.Data.TimeZoneOffset, TIME_ZONE)

    testutils.AssertEqual(t, len(response.Data.SetupTests), 1)
    testutils.AssertEqual(t, response.Data.SetupTests[0].Message, TEST_MESSAGE)
    testutils.AssertEqual(t, response.Data.SetupTests[0].Status, TEST_STATUS)
    testutils.AssertEqual(t, response.Data.SetupTests[0].Title, TEST_TITLE)

    testutils.AssertEqual(t, response.Data.Config["auth"], AUTH)
    testutils.AssertEqual(t, response.Data.Config["auth_type"], AUTH_TYPE)
    testutils.AssertEqual(t, response.Data.Config["bucket"], BUCKET)
    testutils.AssertEqual(t, response.Data.Config["cluster_id"], CLUSTER_ID)
    testutils.AssertEqual(t, response.Data.Config["cluster_region"], CLUSTER_REGION)
    testutils.AssertEqual(t, response.Data.Config["connection_type"], CONNECTION_TYPE)
    testutils.AssertEqual(t, response.Data.Config["create_external_tables"], testutils.BoolToStr(CREATE_EXTERNAL_TABLES)) // Inconsistent response
    testutils.AssertEqual(t, response.Data.Config["data_set_location"], DATA_SET_LOCATION)
    testutils.AssertEqual(t, response.Data.Config["database"], DATABASE)
    testutils.AssertEqual(t, response.Data.Config["external_location"], EXTERNAL_LOCATION)
    testutils.AssertEqual(t, response.Data.Config["http_path"], HTTP_PATH)
    testutils.AssertEqual(t, response.Data.Config["host"], HOST)
    testutils.AssertEqual(t, response.Data.Config["password"], MASKED)
    testutils.AssertEqual(t, response.Data.Config["personal_access_token"], MASKED)
    testutils.AssertEqual(t, response.Data.Config["port"], strconv.Itoa(PORT))
    testutils.AssertEqual(t, response.Data.Config["private_key"], MASKED)
    testutils.AssertEqual(t, response.Data.Config["project_id"], PROJECT_ID)
    testutils.AssertEqual(t, response.Data.Config["public_key"], PUBLIC_KEY)
    testutils.AssertEqual(t, response.Data.Config["role_arn"], ROLE_ARN)
    testutils.AssertEqual(t, response.Data.Config["secret_key"], MASKED)
    testutils.AssertEqual(t, response.Data.Config["server_host_name"], SERVER_HOST_NAME)
    testutils.AssertEqual(t, response.Data.Config["tunnel_host"], TUNNEL_HOST)
    testutils.AssertEqual(t, response.Data.Config["tunnel_port"], TUNNEL_PORT)
    testutils.AssertEqual(t, response.Data.Config["tunnel_user"], TUNNEL_USER)
    testutils.AssertEqual(t, response.Data.Config["user"], USER)
    testutils.AssertEqual(t, response.Data.Config["location"], LOCATION)
    testutils.AssertEqual(t, response.Data.Config["role"], ROLE)
    testutils.AssertEqual(t, response.Data.Config["is_private_key_encrypted"], IS_PRIVATE_KEY_ENCRYPTED)
    testutils.AssertEqual(t, response.Data.Config["passphrase"], PASSPHRASE)
    testutils.AssertEqual(t, response.Data.Config["catalog"], CATALOG)
    testutils.AssertEqual(t, response.Data.Config["fivetran_role_arn"], FIVETRAN_ROLE_ARN)
    testutils.AssertEqual(t, response.Data.Config["prefix_path"], PREFIX_PATH)
    testutils.AssertEqual(t, response.Data.Config["region"], REGION)
    testutils.AssertEqual(t, response.Data.Config["storage_account_name"], STORAGE_ACCOUNT_NAME)
    testutils.AssertEqual(t, response.Data.Config["container_name"], CONTAINER_NAME)
    testutils.AssertEqual(t, response.Data.Config["tenant_id"], TENANT_ID)
    testutils.AssertEqual(t, response.Data.Config["client_id"], CLIENT_ID)
    testutils.AssertEqual(t, response.Data.Config["secret_value"], SECRET_VALUE)
    testutils.AssertEqual(t, response.Data.Config["workspace_name"], WORKSPACE_NAME)
    testutils.AssertEqual(t, response.Data.Config["lakehouse_name"], LAKEHOUSE_NAME)
}

func assertResponse(t *testing.T, response destinations.DestinationDetailsWithSetupTestsResponse) {

    testutils.AssertEqual(t, response.Code, "Created")
    testutils.AssertNotEmpty(t, response.Message)

    testutils.AssertEqual(t, response.Data.GroupID, ID)
    testutils.AssertEqual(t, response.Data.ID, ID)
    testutils.AssertEqual(t, response.Data.Region, REGION)
    testutils.AssertEqual(t, response.Data.Service, SERVICE)
    testutils.AssertEqual(t, response.Data.DaylightSavingTimeEnabled, DAYLIGHTSAVINGTIMEENABLED)
    testutils.AssertEqual(t, response.Data.HybridDeploymentAgentId, HYBRIDDEPLOYMENTAGENTID)
    testutils.AssertEqual(t, response.Data.PrivateLinkId, PRIVATELINKID)
    testutils.AssertEqual(t, response.Data.NetworkingMethod, NETWORKINGMETHOD)
    testutils.AssertEqual(t, response.Data.SetupStatus, SETUP_STATUS)
    testutils.AssertEqual(t, response.Data.TimeZoneOffset, TIME_ZONE)

    testutils.AssertEqual(t, len(response.Data.SetupTests), 1)
    testutils.AssertEqual(t, response.Data.SetupTests[0].Message, TEST_MESSAGE)
    testutils.AssertEqual(t, response.Data.SetupTests[0].Status, TEST_STATUS)
    testutils.AssertEqual(t, response.Data.SetupTests[0].Title, TEST_TITLE)

    testutils.AssertEqual(t, response.Data.Config.Auth, AUTH)
    testutils.AssertEqual(t, response.Data.Config.AuthType, AUTH_TYPE)
    testutils.AssertEqual(t, response.Data.Config.Bucket, BUCKET)
    testutils.AssertEqual(t, response.Data.Config.ClusterId, CLUSTER_ID)
    testutils.AssertEqual(t, response.Data.Config.ClusterRegion, CLUSTER_REGION)
    testutils.AssertEqual(t, response.Data.Config.ConnectionType, CONNECTION_TYPE)
    testutils.AssertEqual(t, response.Data.Config.CreateExternalTables, testutils.BoolToStr(CREATE_EXTERNAL_TABLES)) // Inconsistent response
    testutils.AssertEqual(t, response.Data.Config.DataSetLocation, DATA_SET_LOCATION)
    testutils.AssertEqual(t, response.Data.Config.Database, DATABASE)
    testutils.AssertEqual(t, response.Data.Config.ExternalLocation, EXTERNAL_LOCATION)
    testutils.AssertEqual(t, response.Data.Config.HTTPPath, HTTP_PATH)
    testutils.AssertEqual(t, response.Data.Config.Host, HOST)
    testutils.AssertEqual(t, response.Data.Config.Password, MASKED)
    testutils.AssertEqual(t, response.Data.Config.PersonalAccessToken, MASKED)
    testutils.AssertEqual(t, response.Data.Config.Port, strconv.Itoa(PORT))
    testutils.AssertEqual(t, response.Data.Config.PrivateKey, MASKED)
    testutils.AssertEqual(t, response.Data.Config.ProjectID, PROJECT_ID)
    testutils.AssertEqual(t, response.Data.Config.PublicKey, PUBLIC_KEY)
    testutils.AssertEqual(t, response.Data.Config.RoleArn, ROLE_ARN)
    testutils.AssertEqual(t, response.Data.Config.SecretKey, MASKED)
    testutils.AssertEqual(t, response.Data.Config.ServerHostName, SERVER_HOST_NAME)
    testutils.AssertEqual(t, response.Data.Config.TunnelHost, TUNNEL_HOST)
    testutils.AssertEqual(t, response.Data.Config.TunnelPort, TUNNEL_PORT)
    testutils.AssertEqual(t, response.Data.Config.TunnelUser, TUNNEL_USER)
    testutils.AssertEqual(t, response.Data.Config.User, USER)
    testutils.AssertEqual(t, response.Data.Config.Location, LOCATION)
    testutils.AssertEqual(t, response.Data.Config.Role, ROLE)
    testutils.AssertEqual(t, response.Data.Config.IsPrivateKeyEncrypted, IS_PRIVATE_KEY_ENCRYPTED)
    testutils.AssertEqual(t, response.Data.Config.Passphrase, PASSPHRASE)
    testutils.AssertEqual(t, response.Data.Config.Catalog, CATALOG)
    testutils.AssertEqual(t, response.Data.Config.FivetranRoleArn, FIVETRAN_ROLE_ARN)
    testutils.AssertEqual(t, response.Data.Config.PrefixPath, PREFIX_PATH)
    testutils.AssertEqual(t, response.Data.Config.Region, REGION)
    testutils.AssertEqual(t, response.Data.Config.StorageAccountName, STORAGE_ACCOUNT_NAME)
    testutils.AssertEqual(t, response.Data.Config.ContainerName, CONTAINER_NAME)
    testutils.AssertEqual(t, response.Data.Config.TenantId, TENANT_ID)
    testutils.AssertEqual(t, response.Data.Config.ClientId, CLIENT_ID)
    testutils.AssertEqual(t, response.Data.Config.SecretValue, SECRET_VALUE)
    testutils.AssertEqual(t, response.Data.Config.WorkspaceName, WORKSPACE_NAME)
    testutils.AssertEqual(t, response.Data.Config.LakehouseName, LAKEHOUSE_NAME)
}
