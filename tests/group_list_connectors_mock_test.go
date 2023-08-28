package tests

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/tests/mock"
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
	return `{
		"code": "Success",
		"data": {
			"items": [
				{
					"id": "iodize_impressive",
					"group_id": "projected_sickle",
					"service": "salesforce",
					"service_version": 1,
					"schema": "salesforce",
					"connected_by": "concerning_batch",
					"created_at": "2018-07-21T22:55:21.724201Z",
					"succeeded_at": "2018-12-26T17:58:18.245Z",
					"failed_at": "2018-08-24T15:24:58.872491Z",
					"sync_frequency": 60,
					"status": {
						"setup_state": "connected",
						"sync_state": "paused",
						"update_state": "delayed",
						"is_historical_sync": false,
						"tasks": [],
						"warnings": []
					}
				}
			],
			"next_cursor": "eyJza2lwIjoxfQ"
		}
	}`
}

func assertGroupListConnectorsResponse(t *testing.T, response fivetran.GroupListConnectorsResponse) {
	isHistoricalSyncFlag := false
	createdAt, _ := time.Parse(time.RFC3339, "2018-07-21T22:55:21.724201Z")
	succeededAt, _ := time.Parse(time.RFC3339, "2018-12-26T17:58:18.245Z")
	failedAt, _ := time.Parse(time.RFC3339, "2018-08-24T15:24:58.872491Z")

	assertEqual(t, response.Code, "Success")
	assertEqual(t, len(response.Data.Items), 1)
	item := response.Data.Items[0]
	connectorStatus := fivetran.ConnectorsStatus{
		SetupState:       "connected",
		SyncState:        "paused",
		UpdateState:      "delayed",
		IsHistoricalSync: &(isHistoricalSyncFlag),
		Tasks:            []fivetran.ConnectorTasks{},
		Warnings:         []fivetran.ConnectorWarning{},
	}

	assertEqual(t, item.ID, "iodize_impressive")
	assertEqual(t, item.GroupID, "projected_sickle")
	assertEqual(t, item.Service, "salesforce")
	assertEqual(t, *item.ServiceVersion, 1)
	assertEqual(t, item.Schema, "salesforce")
	assertEqual(t, item.ConnectedBy, "concerning_batch")
	assertEqual(t, item.CreatedAt, createdAt)
	assertEqual(t, item.SucceededAt, succeededAt)
	assertEqual(t, item.FailedAt, failedAt)
	assertEqual(t, *item.SyncFrequency, 60)
	assertEqual(t, item.Status.IsHistoricalSync, connectorStatus.IsHistoricalSync)
	assertEqual(t, item.Status.SetupState, connectorStatus.SetupState)
	assertEqual(t, item.Status.SyncState, connectorStatus.SyncState)
	assertEqual(t, item.Status.Tasks, connectorStatus.Tasks)
	assertEqual(t, item.Status.UpdateState, connectorStatus.UpdateState)
	assertEqual(t, item.Status.Warnings, connectorStatus.Warnings)
}
