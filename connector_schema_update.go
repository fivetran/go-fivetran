package fivetran

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fivetran/go-fivetran/connectors"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// ConnectorSchemaConfigUpdateService implements the Connector Management, Modify a Connector Schema Config API.
// Ref. https://fivetran.com/docs/rest-api/connectors#modifyaconnectorschemaconfig
type ConnectorSchemaConfigUpdateService struct {
	c                    *Client
	connectorID          *string
	schemaChangeHandling *string
	schemas              map[string]*connectors.ConnectorSchemaConfigSchema
}

type connectorSchemaConfigUpdateRequest struct {
	SchemaChangeHandling *string                                                   `json:"schema_change_handling,omitempty"`
	Schemas              map[string]*connectors.ConnectorSchemaConfigSchemaRequest `json:"schemas,omitempty"`
}

func (c *Client) NewConnectorSchemaUpdateService() *ConnectorSchemaConfigUpdateService {
	return &ConnectorSchemaConfigUpdateService{c: c}
}

func (csu *ConnectorSchemaConfigUpdateService) request() *connectorSchemaConfigUpdateRequest {
	var schemas map[string]*connectors.ConnectorSchemaConfigSchemaRequest
	if csu.schemas != nil && len(csu.schemas) != 0 {
		schemas = make(map[string]*connectors.ConnectorSchemaConfigSchemaRequest)
		for k, v := range csu.schemas {
			schemas[k] = v.Request()
		}
	}

	return &connectorSchemaConfigUpdateRequest{
		SchemaChangeHandling: csu.schemaChangeHandling,
		Schemas:              schemas,
	}
}

func (csu *ConnectorSchemaConfigUpdateService) ConnectorID(value string) *ConnectorSchemaConfigUpdateService {
	csu.connectorID = &value
	return csu
}

func (csu *ConnectorSchemaConfigUpdateService) SchemaChangeHandling(value string) *ConnectorSchemaConfigUpdateService {
	csu.schemaChangeHandling = &value
	return csu
}

func (csu *ConnectorSchemaConfigUpdateService) Schema(name string, schema *connectors.ConnectorSchemaConfigSchema) *ConnectorSchemaConfigUpdateService {
	if csu.schemas == nil {
		csu.schemas = make(map[string]*connectors.ConnectorSchemaConfigSchema)
	}
	csu.schemas[name] = schema
	return csu
}

func (csu *ConnectorSchemaConfigUpdateService) Do(ctx context.Context) (connectors.ConnectorSchemaDetailsResponse, error) {
	var response connectors.ConnectorSchemaDetailsResponse

	if csu.connectorID == nil {
		return response, fmt.Errorf("missing required ConnectorID")
	}

	reqBody, err := json.Marshal(csu.request())
	if err != nil {
		return response, err
	}

	url := fmt.Sprintf("%v/connectors/%v/schemas/", csu.c.baseURL, *csu.connectorID)
	expectedStatus := 200

	headers := csu.c.commonHeaders()
	headers["Content-Type"] = "application/json"
	headers["Accept"] = restAPIv2

	r := httputils.Request{
		Method:           "PATCH",
		Url:              url,
		Body:             reqBody,
		Queries:          nil,
		Headers:          headers,
		Client:           csu.c.httpClient,
		HandleRateLimits: csu.c.handleRateLimits,
		MaxRetryAttempts: csu.c.maxRetryAttempts,
	}

	respBody, respStatus, err := r.Do(ctx)
	if err != nil {
		return response, err
	}

	if err := json.Unmarshal(respBody, &response); err != nil {
		return response, err
	}

	if respStatus != expectedStatus {
		err := fmt.Errorf("status code: %v; expected: %v", respStatus, expectedStatus)
		return response, err
	}

	return response, nil
}
