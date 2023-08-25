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
    GROUPID         = "group_id"
    SERVICE         = "service"
    ENABLED         = true

    WORKSPACEID     = "workspace_id"
    PRIMARYKEY      = "primary_key"
    LOGGROUPNAME    = "log_group_name"
    ROLEARN         = "role_arn"
    EXTERNALID      = "external_id"
    REGION          = "region"
    APIKEY          = "api_key"
    SUBDOMAIN       = "sub_domain"
    HOST            = "host"
    HOSTNAME        = "hostname"
    CHANNEL         = "channel"
    ENABLESSL       = "enable_ssl"
    TOKEN           = "token"
    PORT            = 443
)

func TestNewExternalLoggingCreateFullMappingMock(t *testing.T) {
    // arrange
    ftClient, mockClient := CreateTestClient()
    handler := mockClient.When(http.MethodPost, "/v1/external-logging").ThenCall(

        func(req *http.Request) (*http.Response, error) {
            body := requestBodyToJson(t, req)
            assertRequest(t, body)
            response := mock.NewResponse(req, http.StatusCreated, prepareExternalLoggingResponse())
            return response, nil
        })

    // act
    response, err := ftClient.NewExternalLoggingCreate().
        GroupID(GROUPID).
        Service(SERVICE).
        Enabled(ENABLED).
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

func prepareExternalLoggingResponse() string {
    return fmt.Sprintf(
        `{
            "code":"Created",
            "message":"External logging service has been added",
            "data":{
                "id":                           "%v",
                "service":                      "%v",
                "enabled":                      "%v"
            }
        }`,
        GROUPID,
        SERVICE,
        ENABLED,
    )
}

func prepareConfig() *fivetran.ExternalLoggingConfig {
    config := fivetran.NewExternalLoggingConfig()
    config.WorkspaceId(WORKSPACEID)
    config.PrimaryKey(PRIMARYKEY)
    config.LogGroupName(LOGGROUPNAME)
    config.RoleArn(ROLEARN)
    config.ExternalId(EXTERNALID)
    config.Region(REGION)
    config.ApiKey(APIKEY)
    config.SubDomain(SUBDOMAIN)
    config.Host(HOST)
    config.Hostname(HOSTNAME)
    config.Channel(CHANNEL)
    config.EnableSsl(ENABLESSL)
    config.Token(TOKEN)
    config.Port(PORT)
    return config
}

func assertRequest(t *testing.T, request map[string]interface{}) {
    assertKey(t, "service", request, SERVICE)
    assertKey(t, "group_id", request, GROUPID)
    assertKey(t, "enabled", request, ENABLED)

    c, ok := request["config"]
    assertEqual(t, ok, true)
    config, ok := c.(map[string]interface{})
    assertEqual(t, ok, true)

         = ""
    assertKey(t, "workspace_id", config, WORKSPACEID)
    assertKey(t, "primary_key", config, PRIMARYKEY)     
    assertKey(t, "log_group_name", config, LOGGROUPNAME)
    assertKey(t, "role_arn", config, ROLEARN)       
    assertKey(t, "external_id", config, EXTERNALID) 
    assertKey(t, "region", config, REGION)          
    assertKey(t, "api_key", config, APIKEY)          
    assertKey(t, "sub_domain", config, SUBDOMAIN)     
    assertKey(t, "host", config, HOST)
    assertKey(t, "hostname", config, HOSTNAME)
    assertKey(t, "channel", config, CHANNEL)
    assertKey(t, "enable_ssl", config, ENABLESSL)
    assertKey(t, "token", config, TOKEN)
    assertKey(t, "port", config, float64(PORT)) // json marshalling stores all numbers as float64
}

func assertResponse(t *testing.T, response fivetran.ExternalLoggingCreateResponse) {
    assertEqual(t, response.Code, "Created")
    assertNotEmpty(t, response.Message)

    assertEqual(t, response.Data.Id, GROUPID)
    assertEqual(t, response.Data.Service, SERVICE)
    assertEqual(t, response.Data.Enabled, ENABLED)
}
