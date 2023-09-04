package tests

import (
    "context"
    "net/http"
    "testing"

    "github.com/fivetran/go-fivetran"
    "github.com/fivetran/go-fivetran/tests/mock"
)

func TestMetadataColumnsListServiceDo(t *testing.T) {
    ftClient, mockClient := CreateTestClient()

    responseData := `{
  "code": "Success",
  "data": {
    "items": [
      {
        "id": "NTY4ODgzNDI",
        "parent_id": "NjUwMTU",
        "name_in_source": "id",
        "name_in_destination": "id",
        "type_in_source": "Integer",
        "type_in_destination": "Integer",
        "is_primary_key": true,
        "is_foreign_key": false
      },
      {
        "id": "NTY4ODgzNDM",
        "parent_id": "NjUwMTU",
        "name_in_source": "FirstName",
        "name_in_destination": "first_name",
        "type_in_source": "String",
        "type_in_destination": "Text",
        "is_primary_key": false,
        "is_foreign_key": false
      },
      {
        "id": "NTY4ODgzNDQ",
        "parent_id": "NjUwMTU",
        "name_in_source": "LastName",
        "name_in_destination": "last_name",
        "type_in_source": "String",
        "type_in_destination": "Text",
        "is_primary_key": false,
        "is_foreign_key": false
      }
    ],
    "next_cursor": "YUWEudlwIjoxkK"
  }
}`

    handler := mockClient.When(http.MethodGet, "/v1/metadata/connectors/test_connector/columns").ThenCall(
        func(req *http.Request) (*http.Response, error) {
            response := mock.NewResponse(req, http.StatusOK, responseData)
            return response, nil
        },
    )

    service := ftClient.NewMetadataColumnsList().ConnectorId("test_connector")

    response, err := service.Do(context.Background())
    if err != nil {
        t.Error(err)
    }

    flag_true := true
    flag_false := false
    expectedResponse := fivetran.MetadataColumnsListResponse{
        Code:    "Success",
        Data:    struct {
            Items []struct {
                Id                      string    `json:"id"`
                ParentId                string    `json:"parent_id"`
                NameInSource            string    `json:"name_in_source"`
                NameInDestination       string    `json:"name_in_destination"`
                TypeInSource            string    `json:"type_in_source"`
                TypeInDestination       string    `json:"type_in_destination"`
                IsPrimaryKey            *bool      `json:"is_primary_key"`
                IsForeignKey            *bool      `json:"is_foreign_key"`
            } `json:"items"`
            NextCursor string `json:"next_cursor"`
        }{
            Items: []struct {
                Id                      string    `json:"id"`
                ParentId                string    `json:"parent_id"`
                NameInSource            string    `json:"name_in_source"`
                NameInDestination       string    `json:"name_in_destination"`
                TypeInSource            string    `json:"type_in_source"`
                TypeInDestination       string    `json:"type_in_destination"`
                IsPrimaryKey            *bool      `json:"is_primary_key"`
                IsForeignKey            *bool      `json:"is_foreign_key"`
            }{
                {
                    Id:                 "NTY4ODgzNDI",
                    ParentId:           "NjUwMTU",
                    NameInSource:       "id",
                    NameInDestination:  "id",
                    TypeInSource:       "Integer",
                    TypeInDestination:  "Integer",
                    IsPrimaryKey:       &flag_true,
                    IsForeignKey:       &flag_false,
                },
                {
                    Id:                 "NTY4ODgzNDM",
                    ParentId:           "NjUwMTU",
                    NameInSource:       "FirstName",
                    NameInDestination:  "first_name",
                    TypeInSource:       "String",
                    TypeInDestination:  "Text",
                    IsPrimaryKey:       &flag_false,
                    IsForeignKey:       &flag_false,
                },
                {
                    Id:                 "NTY4ODgzNDQ",
                    ParentId:           "NjUwMTU",
                    NameInSource:       "LastName",
                    NameInDestination:  "last_name",
                    TypeInSource:       "String",
                    TypeInDestination:  "Text",
                    IsPrimaryKey:       &flag_false,
                    IsForeignKey:       &flag_false,
                },
            },
            NextCursor: "YUWEudlwIjoxkK",
        },
    }

    assertMetadataColumnsListResponse(t, response, expectedResponse)

    interactions := mockClient.Interactions()
    assertEqual(t, len(interactions), 1)
    assertEqual(t, interactions[0].Handler, handler)
    assertEqual(t, handler.Interactions, 1)

}

func assertMetadataColumnsListResponse(t *testing.T, response fivetran.MetadataColumnsListResponse, expected fivetran.MetadataColumnsListResponse) {
    assertEqual(t, response.Code, expected.Code)
    assertEqual(t, response.Data.NextCursor, expected.Data.NextCursor)

    // Assert items
    assertEqual(t, len(response.Data.Items), len(expected.Data.Items))
    for i, item := range response.Data.Items {
        assertEqual(t, item.Id, expected.Data.Items[i].Id)
        assertEqual(t, item.ParentId, expected.Data.Items[i].ParentId)
        assertEqual(t, item.NameInSource, expected.Data.Items[i].NameInSource)
        assertEqual(t, item.NameInDestination, expected.Data.Items[i].NameInDestination)
        assertEqual(t, item.TypeInSource, expected.Data.Items[i].TypeInSource)
        assertEqual(t, item.TypeInDestination, expected.Data.Items[i].TypeInDestination)
        assertEqual(t, *item.IsPrimaryKey, *expected.Data.Items[i].IsPrimaryKey)
        assertEqual(t, *item.IsForeignKey, *expected.Data.Items[i].IsForeignKey)
    }
}
