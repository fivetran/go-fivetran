package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
)

// ConnectorSchemaConfigUpdateService implements the Connector Management, Modify a Connector Schema Config API.
// Ref. https://fivetran.com/docs/rest-api/connectors#modifyaconnectorschemaconfig
type ConnectorSchemaConfigUpdateService struct {
	c                    *Client
	connectorID          *string
	schemaChangeHandling *string
	schemas              map[string]*ConnectorSchemaConfigSchema
}

type connectorSchemaConfigUpdateRequest struct {
	SchemaChangeHandling *string                                        `json:"schema_change_handling,omitempty"`
	Schemas              map[string]*connectorSchemaConfigSchemaRequest `json:"schemas,omitempty"`
}

func (c *Client) NewConnectorSchemaUpdateService() *ConnectorSchemaConfigUpdateService {
	return &ConnectorSchemaConfigUpdateService{c: c}
}

func (csu *ConnectorSchemaConfigUpdateService) request() *connectorSchemaConfigUpdateRequest {
	var schemas map[string]*connectorSchemaConfigSchemaRequest
	if csu.schemas != nil && len(csu.schemas) != 0 {
		schemas = make(map[string]*connectorSchemaConfigSchemaRequest)
		for k, v := range csu.schemas {
			schemas[k] = v.request()
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

func (csu *ConnectorSchemaConfigUpdateService) Schema(name string, schema *ConnectorSchemaConfigSchema) *ConnectorSchemaConfigUpdateService {
	if csu.schemas == nil {
		csu.schemas = make(map[string]*ConnectorSchemaConfigSchema)
	}
	csu.schemas[name] = schema
	return csu
}

func (csu *ConnectorSchemaConfigUpdateService) Do(ctx context.Context) (ConnectorSchemaDetailsResponse, error) {
	var response ConnectorSchemaDetailsResponse

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

	r := request{
		method:  "PATCH",
		url:     url,
		body:    reqBody,
		queries: nil,
		headers: headers,
		client:  csu.c.httpClient,
	}

	respBody, respStatus, err := r.httpRequest(ctx)
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
