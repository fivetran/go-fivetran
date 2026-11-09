package externallogging_test

import (
	"context"
	"net/http"
	"testing"

	fivetran "github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/common"
	externallogging "github.com/fivetran/go-fivetran/external_logging"
	testutils "github.com/fivetran/go-fivetran/test_utils"
	"github.com/fivetran/go-fivetran/tests/mock"
)

const (
	accountExternalLoggingPath     = "/v1/external-logging/account"
	accountExternalLoggingUASuffix = "fivetran_account_external_logging_v2"
)

func prepareAccountExternalLoggingResponse() string {
	return `{
		"code": "Success",
		"message": "Account external logging service has been added",
		"data": {
			"id": "account_log_id",
			"service": "azure_monitor_log",
			"enabled": true,
			"setup_tests": [
				{
					"title": "Validate Login",
					"status": "FAILED",
					"message": "Invalid login credentials.",
					"details": "Check the client secret"
				}
			],
			"config": {
				"workspace_id": "workspace_id",
				"primary_key": "******",
				"fake_field": "unmapped-value"
			}
		}
	}`
}

func prepareAccountExternalLoggingConfigCustom() *map[string]interface{} {
	return &map[string]interface{}{
		"workspace_id": "workspace_id",
		"primary_key":  "primary_key",
		"fake_field":   "unmapped-value",
	}
}

func assertAccountExternalLoggingConfigRequest(t *testing.T, config map[string]interface{}) {
	testutils.AssertKey(t, "workspace_id", config, "workspace_id")
	testutils.AssertKey(t, "primary_key", config, "primary_key")
	testutils.AssertKey(t, "fake_field", config, "unmapped-value")
}

func assertAccountExternalLoggingCustomResponse(t *testing.T, response externallogging.ExternalLoggingCustomResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, response.Data.Id, "account_log_id")
	testutils.AssertEqual(t, response.Data.Service, "azure_monitor_log")
	testutils.AssertEqual(t, response.Data.Enabled, true)

	testutils.AssertHasLength(t, response.Data.SetupTests, 1)
	testutils.AssertEqual(t, response.Data.SetupTests[0].Title, "Validate Login")
	testutils.AssertEqual(t, response.Data.SetupTests[0].Status, "FAILED")
	testutils.AssertEqual(t, response.Data.SetupTests[0].Message, "Invalid login credentials.")

	testutils.AssertKey(t, "workspace_id", response.Data.Config, "workspace_id")
	testutils.AssertKey(t, "primary_key", response.Data.Config, "******")
	testutils.AssertKey(t, "fake_field", response.Data.Config, "unmapped-value")
}

// withUserAgent wraps a handler with a User-Agent check: the suffix is expected
// at the end when it is set, and the client's own User-Agent otherwise.
func withUserAgent(t *testing.T, suffix *string, next mock.HandlerFunc) mock.HandlerFunc {
	return func(req *http.Request) (*http.Response, error) {
		expected := fivetran.DefaultUserAgent + " terraform-provider-fivetran/1.9.37"
		if suffix != nil {
			expected += " " + *suffix
		}
		testutils.AssertEqual(t, req.Header.Get("User-Agent"), expected)
		return next(req)
	}
}

func newAccountTestClient() (*fivetran.Client, *mock.HttpClient) {
	ftClient, mockClient := testutils.CreateTestClient()
	ftClient.CustomUserAgent("terraform-provider-fivetran/1.9.37")
	return ftClient, mockClient
}

func suffixCases() map[string]*string {
	suffix := accountExternalLoggingUASuffix
	return map[string]*string{"WithoutUserAgentSuffix": nil, "WithUserAgentSuffix": &suffix}
}

func TestExternalLoggingAccountCreateService(t *testing.T) {
	for name, suffix := range suffixCases() {
		t.Run(name, func(t *testing.T) {
			ftClient, mockClient := newAccountTestClient()
			handler := mockClient.When(http.MethodPost, accountExternalLoggingPath).ThenCall(withUserAgent(t, suffix,
				func(req *http.Request) (*http.Response, error) {
					body := testutils.RequestBodyToJson(t, req)
					testutils.AssertHasNoKey(t, body, "group_id")
					testutils.AssertKey(t, "service", body, "azure_monitor_log")
					testutils.AssertKey(t, "enabled", body, true)
					testutils.AssertKey(t, "run_setup_tests", body, true)
					assertAccountExternalLoggingConfigRequest(t, body["config"].(map[string]interface{}))
					return mock.NewResponse(req, http.StatusCreated, prepareAccountExternalLoggingResponse()), nil
				}))

			svc := ftClient.NewExternalLoggingAccountCreate().
				Service("azure_monitor_log").
				Enabled(true).
				RunSetupTests(true).
				ConfigCustom(prepareAccountExternalLoggingConfigCustom())

			var response externallogging.ExternalLoggingCustomResponse
			var err error
			if suffix == nil {
				response, err = svc.DoCustom(context.Background())
			} else {
				response, err = svc.DoCustomWithUserAgentSuffix(context.Background(), *suffix)
			}

			if err != nil {
				t.Logf("%+v\n", response)
				t.Error(err)
			}

			assertSingleInteraction(t, mockClient, handler)
			assertAccountExternalLoggingCustomResponse(t, response)
		})
	}
}

func TestExternalLoggingAccountDetailsService(t *testing.T) {
	for name, suffix := range suffixCases() {
		t.Run(name, func(t *testing.T) {
			ftClient, mockClient := newAccountTestClient()
			handler := mockClient.When(http.MethodGet, accountExternalLoggingPath).ThenCall(withUserAgent(t, suffix,
				func(req *http.Request) (*http.Response, error) {
					return mock.NewResponse(req, http.StatusOK, prepareAccountExternalLoggingResponse()), nil
				}))

			svc := ftClient.NewExternalLoggingAccountDetails()

			var response externallogging.ExternalLoggingCustomResponse
			var err error
			if suffix == nil {
				response, err = svc.DoCustom(context.Background())
			} else {
				response, err = svc.DoCustomWithUserAgentSuffix(context.Background(), *suffix)
			}

			if err != nil {
				t.Logf("%+v\n", response)
				t.Error(err)
			}

			assertSingleInteraction(t, mockClient, handler)
			assertAccountExternalLoggingCustomResponse(t, response)
		})
	}
}

func TestExternalLoggingAccountUpdateService(t *testing.T) {
	for name, suffix := range suffixCases() {
		t.Run(name, func(t *testing.T) {
			ftClient, mockClient := newAccountTestClient()
			handler := mockClient.When(http.MethodPatch, accountExternalLoggingPath).ThenCall(withUserAgent(t, suffix,
				func(req *http.Request) (*http.Response, error) {
					body := testutils.RequestBodyToJson(t, req)
					testutils.AssertKey(t, "enabled", body, true)
					testutils.AssertKey(t, "run_setup_tests", body, true)
					testutils.AssertHasNoKey(t, body, "service")
					assertAccountExternalLoggingConfigRequest(t, body["config"].(map[string]interface{}))
					return mock.NewResponse(req, http.StatusOK, prepareAccountExternalLoggingResponse()), nil
				}))

			svc := ftClient.NewExternalLoggingAccountUpdate().
				Enabled(true).
				RunSetupTests(true).
				ConfigCustom(prepareAccountExternalLoggingConfigCustom())

			var response externallogging.ExternalLoggingCustomResponse
			var err error
			if suffix == nil {
				response, err = svc.DoCustom(context.Background())
			} else {
				response, err = svc.DoCustomWithUserAgentSuffix(context.Background(), *suffix)
			}

			if err != nil {
				t.Logf("%+v\n", response)
				t.Error(err)
			}

			assertSingleInteraction(t, mockClient, handler)
			assertAccountExternalLoggingCustomResponse(t, response)
		})
	}
}

func TestExternalLoggingAccountDeleteService(t *testing.T) {
	for name, suffix := range suffixCases() {
		t.Run(name, func(t *testing.T) {
			ftClient, mockClient := newAccountTestClient()
			handler := mockClient.When(http.MethodDelete, accountExternalLoggingPath).ThenCall(withUserAgent(t, suffix,
				func(req *http.Request) (*http.Response, error) {
					return mock.NewResponse(req, http.StatusOK, `{
						"code": "Success",
						"message": "Account external logging service has been deleted"
					}`), nil
				}))

			svc := ftClient.NewExternalLoggingAccountDelete()

			var response common.CommonResponse
			var err error
			if suffix == nil {
				response, err = svc.Do(context.Background())
			} else {
				response, err = svc.DoWithUserAgentSuffix(context.Background(), *suffix)
			}

			if err != nil {
				t.Error(err)
			}

			assertSingleInteraction(t, mockClient, handler)
			testutils.AssertEqual(t, response.Code, "Success")
			testutils.AssertEqual(t, response.Message, "Account external logging service has been deleted")
		})
	}
}

func TestExternalLoggingAccountSetupTestsService(t *testing.T) {
	for name, suffix := range suffixCases() {
		t.Run(name, func(t *testing.T) {
			ftClient, mockClient := newAccountTestClient()
			handler := mockClient.When(http.MethodPost, accountExternalLoggingPath+"/test").ThenCall(withUserAgent(t, suffix,
				func(req *http.Request) (*http.Response, error) {
					return mock.NewResponse(req, http.StatusOK, `{
						"code": "Success",
						"message": "Account external logging service setup tests have been completed",
						"data": {
							"setup_tests": [
								{"title": "Validate Login", "status": "PASSED", "message": ""}
							]
						}
					}`), nil
				}))

			svc := ftClient.NewExternalLoggingAccountSetupTests()

			var response externallogging.ExternalLoggingSetupTestsResponse
			var err error
			if suffix == nil {
				response, err = svc.Do(context.Background())
			} else {
				response, err = svc.DoWithUserAgentSuffix(context.Background(), *suffix)
			}

			if err != nil {
				t.Logf("%+v\n", response)
				t.Error(err)
			}

			assertSingleInteraction(t, mockClient, handler)
			testutils.AssertEqual(t, response.Code, "Success")
			testutils.AssertHasLength(t, response.Data.SetupTests, 1)
			testutils.AssertEqual(t, response.Data.SetupTests[0].Title, "Validate Login")
			testutils.AssertEqual(t, response.Data.SetupTests[0].Status, "PASSED")
		})
	}
}
