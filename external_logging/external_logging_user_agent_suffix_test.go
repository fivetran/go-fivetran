package externallogging_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	fivetran "github.com/fivetran/go-fivetran"
	testutils "github.com/fivetran/go-fivetran/test_utils"
	"github.com/fivetran/go-fivetran/tests/mock"
)

const EXTLOG_USER_AGENT_SUFFIX = "fivetran_external_logging_v2"

func expectedExternalLoggingUserAgent() string {
	return fmt.Sprintf("%s terraform-provider-fivetran/1.9.37 %s", fivetran.DefaultUserAgent, EXTLOG_USER_AGENT_SUFFIX)
}

func prepareExternalLoggingSetupTestsResponse() string {
	return `{
		"code": "Success",
		"message": "Log service has been updated",
		"data": {
			"id": "log_id",
			"group_id": "group_id",
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

func assertSingleInteraction(t *testing.T, mockClient *mock.HttpClient, handler *mock.Handler) {
	interactions := mockClient.Interactions()
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)
}

func TestExternalLoggingServicesWithUserAgentSuffix(t *testing.T) {
	ctx := context.Background()
	config := &map[string]interface{}{"fake_field": "unmapped-value"}

	cases := []struct {
		name   string
		method string
		path   string
		status int
		body   string
		call   func(c *fivetran.Client) error
	}{
		{"CreateDo", http.MethodPost, "/v1/external-logging", http.StatusCreated, prepareExternalLoggingSetupTestsResponse(),
			func(c *fivetran.Client) error {
				_, err := c.NewExternalLoggingCreate().GroupId("group_id").Service("azure_monitor_log").DoWithUserAgentSuffix(ctx, EXTLOG_USER_AGENT_SUFFIX)
				return err
			}},
		{"CreateDoCustom", http.MethodPost, "/v1/external-logging", http.StatusCreated, prepareExternalLoggingSetupTestsResponse(),
			func(c *fivetran.Client) error {
				_, err := c.NewExternalLoggingCreate().GroupId("group_id").ConfigCustom(config).DoCustomWithUserAgentSuffix(ctx, EXTLOG_USER_AGENT_SUFFIX)
				return err
			}},
		{"CreateDoCustomMerged", http.MethodPost, "/v1/external-logging", http.StatusCreated, prepareExternalLoggingSetupTestsResponse(),
			func(c *fivetran.Client) error {
				_, err := c.NewExternalLoggingCreate().GroupId("group_id").ConfigCustom(config).DoCustomMergedWithUserAgentSuffix(ctx, EXTLOG_USER_AGENT_SUFFIX)
				return err
			}},
		{"DetailsDo", http.MethodGet, "/v1/external-logging/log_id", http.StatusOK, prepareExternalLoggingSetupTestsResponse(),
			func(c *fivetran.Client) error {
				_, err := c.NewExternalLoggingDetails().ExternalLoggingId("log_id").DoWithUserAgentSuffix(ctx, EXTLOG_USER_AGENT_SUFFIX)
				return err
			}},
		{"DetailsDoCustom", http.MethodGet, "/v1/external-logging/log_id", http.StatusOK, prepareExternalLoggingSetupTestsResponse(),
			func(c *fivetran.Client) error {
				_, err := c.NewExternalLoggingDetails().ExternalLoggingId("log_id").DoCustomWithUserAgentSuffix(ctx, EXTLOG_USER_AGENT_SUFFIX)
				return err
			}},
		{"DetailsDoCustomMerged", http.MethodGet, "/v1/external-logging/log_id", http.StatusOK, prepareExternalLoggingSetupTestsResponse(),
			func(c *fivetran.Client) error {
				_, err := c.NewExternalLoggingDetails().ExternalLoggingId("log_id").DoCustomMergedWithUserAgentSuffix(ctx, EXTLOG_USER_AGENT_SUFFIX)
				return err
			}},
		{"UpdateDo", http.MethodPatch, "/v1/external-logging/log_id", http.StatusOK, prepareExternalLoggingSetupTestsResponse(),
			func(c *fivetran.Client) error {
				_, err := c.NewExternalLoggingUpdate().ExternalLoggingId("log_id").Enabled(true).DoWithUserAgentSuffix(ctx, EXTLOG_USER_AGENT_SUFFIX)
				return err
			}},
		{"UpdateDoCustom", http.MethodPatch, "/v1/external-logging/log_id", http.StatusOK, prepareExternalLoggingSetupTestsResponse(),
			func(c *fivetran.Client) error {
				_, err := c.NewExternalLoggingUpdate().ExternalLoggingId("log_id").ConfigCustom(config).DoCustomWithUserAgentSuffix(ctx, EXTLOG_USER_AGENT_SUFFIX)
				return err
			}},
		{"UpdateDoCustomMerged", http.MethodPatch, "/v1/external-logging/log_id", http.StatusOK, prepareExternalLoggingSetupTestsResponse(),
			func(c *fivetran.Client) error {
				_, err := c.NewExternalLoggingUpdate().ExternalLoggingId("log_id").ConfigCustom(config).DoCustomMergedWithUserAgentSuffix(ctx, EXTLOG_USER_AGENT_SUFFIX)
				return err
			}},
		{"Delete", http.MethodDelete, "/v1/external-logging/log_id", http.StatusOK, `{"code": "Success", "message": "deleted"}`,
			func(c *fivetran.Client) error {
				_, err := c.NewExternalLoggingDelete().ExternalLoggingId("log_id").DoWithUserAgentSuffix(ctx, EXTLOG_USER_AGENT_SUFFIX)
				return err
			}},
		{"SetupTests", http.MethodPost, "/v1/external-logging/log_id/test", http.StatusOK, `{"code": "Success", "data": {"setup_tests": []}}`,
			func(c *fivetran.Client) error {
				_, err := c.NewExternalLoggingSetupTests().ExternalLoggingId("log_id").DoWithUserAgentSuffix(ctx, EXTLOG_USER_AGENT_SUFFIX)
				return err
			}},
		{"List", http.MethodGet, "/v1/external-logging", http.StatusOK, prepareExternalLoggingListResponse(),
			func(c *fivetran.Client) error {
				_, err := c.NewExternalLoggingList().DoWithUserAgentSuffix(ctx, EXTLOG_USER_AGENT_SUFFIX)
				return err
			}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ftClient, mockClient := testutils.CreateTestClient()
			ftClient.CustomUserAgent("terraform-provider-fivetran/1.9.37")
			handler := mockClient.When(tc.method, tc.path).ThenCall(
				func(req *http.Request) (*http.Response, error) {
					testutils.AssertEqual(t, req.Header.Get("User-Agent"), expectedExternalLoggingUserAgent())
					return mock.NewResponse(req, tc.status, tc.body), nil
				})

			if err := tc.call(ftClient); err != nil {
				t.Error(err)
			}

			assertSingleInteraction(t, mockClient, handler)
		})
	}
}

func TestExternalLoggingServicesWithoutUserAgentSuffix(t *testing.T) {
	ftClient, mockClient := testutils.CreateTestClient()
	ftClient.CustomUserAgent("terraform-provider-fivetran/1.9.37")
	handler := mockClient.When(http.MethodGet, "/v1/external-logging/log_id").ThenCall(
		func(req *http.Request) (*http.Response, error) {
			testutils.AssertEqual(t, req.Header.Get("User-Agent"), fivetran.DefaultUserAgent+" terraform-provider-fivetran/1.9.37")
			return mock.NewResponse(req, http.StatusOK, prepareExternalLoggingSetupTestsResponse()), nil
		})

	_, err := ftClient.NewExternalLoggingDetails().ExternalLoggingId("log_id").DoCustom(context.Background())
	if err != nil {
		t.Error(err)
	}

	assertSingleInteraction(t, mockClient, handler)
}

func TestExternalLoggingServicesWithUserAgentSuffixMissingId(t *testing.T) {
	ctx := context.Background()
	ftClient, mockClient := testutils.CreateTestClient()

	calls := map[string]func() error{
		"Details": func() error {
			_, err := ftClient.NewExternalLoggingDetails().DoCustomWithUserAgentSuffix(ctx, EXTLOG_USER_AGENT_SUFFIX)
			return err
		},
		"Update": func() error {
			_, err := ftClient.NewExternalLoggingUpdate().DoCustomWithUserAgentSuffix(ctx, EXTLOG_USER_AGENT_SUFFIX)
			return err
		},
		"Delete": func() error {
			_, err := ftClient.NewExternalLoggingDelete().DoWithUserAgentSuffix(ctx, EXTLOG_USER_AGENT_SUFFIX)
			return err
		},
		"SetupTests": func() error {
			_, err := ftClient.NewExternalLoggingSetupTests().DoWithUserAgentSuffix(ctx, EXTLOG_USER_AGENT_SUFFIX)
			return err
		},
	}

	for name, call := range calls {
		t.Run(name, func(t *testing.T) {
			err := call()
			testutils.AssertEqual(t, fmt.Sprint(err), "missing required externalLoggingId")
		})
	}

	testutils.AssertEqual(t, len(mockClient.Interactions()), 0)
}

func TestExternalLoggingUpdateCustomSendsRunSetupTestsAndReturnsSetupTests(t *testing.T) {
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodPatch, "/v1/external-logging/log_id").ThenCall(
		func(req *http.Request) (*http.Response, error) {
			body := testutils.RequestBodyToJson(t, req)
			testutils.AssertKey(t, "run_setup_tests", body, true)
			testutils.AssertKey(t, "enabled", body, true)
			return mock.NewResponse(req, http.StatusOK, prepareExternalLoggingSetupTestsResponse()), nil
		})

	response, err := ftClient.NewExternalLoggingUpdate().
		ExternalLoggingId("log_id").
		Enabled(true).
		RunSetupTests(true).
		ConfigCustom(&map[string]interface{}{"fake_field": "unmapped-value"}).
		DoCustomWithUserAgentSuffix(context.Background(), EXTLOG_USER_AGENT_SUFFIX)

	if err != nil {
		t.Logf("%+v\n", response)
		t.Error(err)
	}

	assertSingleInteraction(t, mockClient, handler)
	testutils.AssertHasLength(t, response.Data.SetupTests, 1)
	testutils.AssertEqual(t, response.Data.SetupTests[0].Title, "Validate Login")
	testutils.AssertEqual(t, response.Data.SetupTests[0].Status, "FAILED")
	testutils.AssertEqual(t, response.Data.SetupTests[0].Message, "Invalid login credentials.")
}

func TestExternalLoggingUpdateServiceSendsRunSetupTests(t *testing.T) {
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodPatch, "/v1/external-logging/log_id").ThenCall(
		func(req *http.Request) (*http.Response, error) {
			body := testutils.RequestBodyToJson(t, req)
			testutils.AssertKey(t, "run_setup_tests", body, true)
			return mock.NewResponse(req, http.StatusOK, prepareExternalLoggingSetupTestsResponse()), nil
		})

	response, err := ftClient.NewExternalLoggingUpdate().
		ExternalLoggingId("log_id").
		RunSetupTests(true).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", response)
		t.Error(err)
	}

	assertSingleInteraction(t, mockClient, handler)
	testutils.AssertHasLength(t, response.Data.SetupTests, 1)
}
