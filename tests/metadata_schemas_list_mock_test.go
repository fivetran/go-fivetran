package tests

import (
    "context"
    "net/http"
    "testing"

    "github.com/fivetran/go-fivetran"
    "github.com/fivetran/go-fivetran/tests/mock"
)

func TestMetadataSchemasListServiceDo(t *testing.T) {
    ftClient, mockClient := CreateTestClient()

    responseData := `{
                        "code": "Success",
                        "data": {
                        "items": [
                            {
                                "id": "bWFpbl9wdWJsaWM",
                                "name_in_source": "Main Public",
                                "name_in_destination": "main_public"
                            },
                            {
                                "id": "dXNlcl9zZXR0aW5ncw",
                                "name_in_source": "Local backup",
                                "name_in_destination": "local_backup"
                            }],
                        "next_cursor": "eyJza2lwIjoxfQ"
                        }
                    }`

    handler := mockClient.When(http.MethodGet, "/v1/metadata/connectors/test_connector/schemas").ThenCall(
        func(req *http.Request) (*http.Response, error) {
            response := mock.NewResponse(req, http.StatusOK, responseData)
            return response, nil
        },
    )

    service := ftClient.NewMetadataSchemasList().ConnectorId("test_connector")

    response, err := service.Do(context.Background())
    if err != nil {
        t.Error(err)
    }

    expectedResponse := fivetran.MetadataSchemasListResponse{
        Code:    "Success",
        Data:    struct {
            Items []struct {
                Id                     string    `json:"id"`
                NameInSource           string    `json:"name_in_source"`
                NameInDestination      string    `json:"name_in_destination"`
            } `json:"items"`
            NextCursor string `json:"next_cursor"`
        }{
            Items: []struct {
                Id                     string    `json:"id"`
                NameInSource           string    `json:"name_in_source"`
                NameInDestination      string    `json:"name_in_destination"`
            }{
                {
                    Id:                 "bWFpbl9wdWJsaWM",
                    NameInSource:       "Main Public",
                    NameInDestination:  "main_public",
                },
                {
                    Id:                 "dXNlcl9zZXR0aW5ncw",
                    NameInSource:       "Local backup",
                    NameInDestination:  "local_backup",
                },
            },
            NextCursor: "eyJza2lwIjoxfQ",
        },
    }

    assertMetadataSchemasListResponse(t, response, expectedResponse)

    interactions := mockClient.Interactions()
    assertEqual(t, len(interactions), 1)
    assertEqual(t, interactions[0].Handler, handler)
    assertEqual(t, handler.Interactions, 1)

}

func assertMetadataSchemasListResponse(t *testing.T, response fivetran.MetadataSchemasListResponse, expected fivetran.MetadataSchemasListResponse) {
    assertEqual(t, response.Code, expected.Code)
    assertEqual(t, response.Data.NextCursor, expected.Data.NextCursor)

    // Assert items
    assertEqual(t, len(response.Data.Items), len(expected.Data.Items))
    for i, item := range response.Data.Items {
        assertEqual(t, item.Id, expected.Data.Items[i].Id)
        assertEqual(t, item.NameInSource, expected.Data.Items[i].NameInSource)
        assertEqual(t, item.NameInDestination, expected.Data.Items[i].NameInDestination)
    }
}
