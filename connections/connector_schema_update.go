package connectors

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// ConnectorSchemaConfigUpdateService implements the Connector Management, Modify a Connector Schema Config API.
// Ref. https://fivetran.com/docs/rest-api/connectors#modifyaconnectorschemaconfig
type ConnectorSchemaConfigUpdateService struct {
    httputils.HttpService
	connectorID          *string
	schemaChangeHandling *string
	schemas              map[string]*ConnectorSchemaConfigSchema
}

func (csu *ConnectorSchemaConfigUpdateService) request() *connectorSchemaConfigUpdateRequest {
	var schemas map[string]*ConnectorSchemaConfigSchemaRequest
	if csu.schemas != nil && len(csu.schemas) != 0 {
		schemas = make(map[string]*ConnectorSchemaConfigSchemaRequest)
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

    url := fmt.Sprintf("/connectors/%v/schemas", *csu.connectorID)
    err := csu.HttpService.Do(ctx, "PATCH", url, csu.request(), nil, 200, &response)
    return response, err
}