package tests

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/common"
	"github.com/fivetran/go-fivetran/tests/mock"
)

const (
	LIST_CONNECTORS_ID                 = "iodize_impressive"
	LIST_CONNECTORS_GROUP_ID           = "projected_sickle"
	LIST_CONNECTORS_SERVICE            = "salesforce"
	LIST_CONNECTORS_SERVICE_VERSION    = 1
	LIST_CONNECTORS_SCHEMA             = "salesforce"
	LIST_CONNECTORS_CONNECTED_BY       = "concerning_batch"
	LIST_CONNECTORS_CREATED_AT         = "2018-07-21T22:55:21.724201Z"
	LIST_CONNECTORS_SUCCEEDED_AT       = "2018-12-26T17:58:18.245Z"
	LIST_CONNECTORS_FAILED_AT          = "2018-08-24T15:24:58.872491Z"
	LIST_CONNECTORS_SYNC_FREQUENCY     = 60
	LIST_CONNECTORS_SETUP_STATE        = "connected"
	LIST_CONNECTORS_SYNC_STATE         = "paused"
	LIST_CONNECTORS_UPDATE_STATE       = "delayed"
	LIST_CONNECTORS_IS_HISTORICAL_SYNC = false
)

func TestGroupListConnectorsServiceDo(t *testing.T) {
	// arrange
	groupID := "projected_sickle"
	limit := 10
	cursor := "eyJza2lwIjoxfQ"
	schema := "salesforce"

	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(http.MethodGet, fmt.Sprintf("/v1/groups/%s/connectors", groupID)).
		ThenCall(func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, prepareGroupListConnectorsResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewGroupListConnectors().
		GroupID(groupID).
		Limit(limit).
		Cursor(cursor).
		Schema(schema).
		Do(context.Background())

	// assert
	if err != nil {
		t.Error(err)
	}

	interactions := mockClient.Interactions()
	assertEqual(t, len(interactions), 1)
	assertEqual(t, interactions[0].Handler, handler)
	assertEqual(t, handler.Interactions, 1)
	assertGroupListConnectorsResponse(t, response)
}

func prepareGroupListConnectorsResponse() string {
	value := fmt.Sprintf(`{
		"code": "Success",
		"data": {
			"items": [
				{
					"id": "%s",
					"group_id": "%s",
					"service": "%s",
					"service_version": %d,
					"schema": "%s",
					"connected_by": "%s",
					"created_at": "%s",
					"succeeded_at": "%s",
					"failed_at": "%s",
					"sync_frequency": %d,
					"status": {
						"setup_state": "%s",
						"sync_state": "%s",
						"update_state": "%s",
						"is_historical_sync": %t,
						"tasks": [],
						"warnings": []
					}
				}
				],
				"next_cursor": "eyJza2lwIjoxfQ"
			}		
			}`,
		LIST_CONNECTORS_ID,
		LIST_CONNECTORS_GROUP_ID,
		LIST_CONNECTORS_SERVICE,
		LIST_CONNECTORS_SERVICE_VERSION,
		LIST_CONNECTORS_SCHEMA,
		LIST_CONNECTORS_CONNECTED_BY,
		LIST_CONNECTORS_CREATED_AT,
		LIST_CONNECTORS_SUCCEEDED_AT,
		LIST_CONNECTORS_FAILED_AT,
		LIST_CONNECTORS_SYNC_FREQUENCY,
		LIST_CONNECTORS_SETUP_STATE,
		LIST_CONNECTORS_SYNC_STATE,
		LIST_CONNECTORS_UPDATE_STATE,
		LIST_CONNECTORS_IS_HISTORICAL_SYNC,
	)
	return value
}

func assertGroupListConnectorsResponse(t *testing.T, response fivetran.GroupListConnectorsResponse) {

	assertEqual(t, response.Code, "Success")
	assertEqual(t, len(response.Data.Items), 1)
	item := response.Data.Items[0]

	assertEqual(t, item.ID, LIST_CONNECTORS_ID)
	assertEqual(t, item.GroupID, LIST_CONNECTORS_GROUP_ID)
	assertEqual(t, item.Service, LIST_CONNECTORS_SERVICE)
	assertEqual(t, *item.ServiceVersion, LIST_CONNECTORS_SERVICE_VERSION)
	assertEqual(t, item.Schema, LIST_CONNECTORS_SCHEMA)
	assertEqual(t, item.ConnectedBy, LIST_CONNECTORS_CONNECTED_BY)
	assertTimeEqual(t, item.CreatedAt, LIST_CONNECTORS_CREATED_AT)
	assertTimeEqual(t, item.SucceededAt, LIST_CONNECTORS_SUCCEEDED_AT)
	assertTimeEqual(t, item.FailedAt, LIST_CONNECTORS_FAILED_AT)
	assertEqual(t, *item.SyncFrequency, LIST_CONNECTORS_SYNC_FREQUENCY)
	assertEqual(t, *item.Status.IsHistoricalSync, LIST_CONNECTORS_IS_HISTORICAL_SYNC)
	assertEqual(t, item.Status.SetupState, LIST_CONNECTORS_SETUP_STATE)
	assertEqual(t, item.Status.SyncState, LIST_CONNECTORS_SYNC_STATE)
	assertEqual(t, item.Status.Tasks, []common.CommonResponse{})
	assertEqual(t, item.Status.UpdateState, LIST_CONNECTORS_UPDATE_STATE)
	assertEqual(t, item.Status.Warnings, []common.CommonResponse{})
}
