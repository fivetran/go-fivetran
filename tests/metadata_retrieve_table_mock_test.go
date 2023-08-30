package tests

import (
    "context"
    "net/http"
    "testing"

    "github.com/fivetran/go-fivetran"
    "github.com/fivetran/go-fivetran/tests/mock"
)

func TestMetadataTableListServiceDo(t *testing.T) {
    ftClient, mockClient := CreateTestClient()

    responseData := `{
                        "code": "Success",
                        "data": {
                        "items": [
                            {
                                "id": "NjUwMTU",
                                "parent_id": "bWFpbl9wdWJsaWM",
                                "name_in_source": "User Accounts",
                                "name_in_destination": "user_accounts"
                            },
                            {
                                "id": "NjUwMTY",
                                "parent_id": "bWFpbl9wdWJsaWM",
                                "name_in_source": "User Subscriptions",
                                "name_in_destination": "user_subscriptions"
                            },
                            {
                                "id": "NjUwMTW",
                                "parent_id": "bWFpbl9wdWJsaWM",
                                "name_in_source": "Account Details",
                                "name_in_destination": "account_details"
                            }
                            ],
                            "next_cursor": "YUWEudlwIjoxkK"
                        }
                    }`

    handler := mockClient.When(http.MethodGet, "/v1/metadata/connectors/test_connector/tables").ThenCall(
        func(req *http.Request) (*http.Response, error) {
            response := mock.NewResponse(req, http.StatusOK, responseData)
            return response, nil
        },
    )

    service := ftClient.NewMetadataTableList().ConnectorId("test_connector")

    response, err := service.Do(context.Background())
    if err != nil {
        t.Error(err)
    }

    expectedResponse := fivetran.MetadataTableListResponse{
        Code:    "Success",
        Data:    struct {
            Items []struct {
                Id                     string    `json:"id"`
                ParentId               string    `json:"parent_id"`
                NameInSource           string    `json:"name_in_source"`
                NameInDestination      string    `json:"name_in_destination"`
            } `json:"items"`
            NextCursor string `json:"next_cursor"`
        }{
            Items: []struct {
                Id                     string    `json:"id"`
                ParentId               string    `json:"parent_id"`
                NameInSource           string    `json:"name_in_source"`
                NameInDestination      string    `json:"name_in_destination"`
            }{
                {
                    Id:                 "NjUwMTU",
                    ParentId:           "bWFpbl9wdWJsaWM",
                    NameInSource:       "User Accounts",
                    NameInDestination:  "user_accounts",
                },
                {
                    Id:                 "NjUwMTY",
                    ParentId:           "bWFpbl9wdWJsaWM",
                    NameInSource:       "User Subscriptions",
                    NameInDestination:  "user_subscriptions",
                },
                {
                    Id:                 "NjUwMTW",
                    ParentId:           "bWFpbl9wdWJsaWM",
                    NameInSource:       "Account Details",
                    NameInDestination:  "account_details",
                },
            },
            NextCursor: "YUWEudlwIjoxkK",
        },
    }

    assertMetadataTableListResponse(t, response, expectedResponse)

    interactions := mockClient.Interactions()
    assertEqual(t, len(interactions), 1)
    assertEqual(t, interactions[0].Handler, handler)
    assertEqual(t, handler.Interactions, 1)

}

func assertMetadataTableListResponse(t *testing.T, response fivetran.MetadataTableListResponse, expected fivetran.MetadataTableListResponse) {
    assertEqual(t, response.Code, expected.Code)
    assertEqual(t, response.Data.NextCursor, expected.Data.NextCursor)

    // Assert items
    assertEqual(t, len(response.Data.Items), len(expected.Data.Items))
    for i, item := range response.Data.Items {
        assertEqual(t, item.Id, expected.Data.Items[i].Id)
        assertEqual(t, item.ParentId, expected.Data.Items[i].ParentId)
        assertEqual(t, item.NameInSource, expected.Data.Items[i].NameInSource)
        assertEqual(t, item.NameInDestination, expected.Data.Items[i].NameInDestination)
    }
}
