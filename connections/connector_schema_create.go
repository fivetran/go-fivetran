package connectors

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// ConnectorSchemaConfigUpdateService implements the Connector Management, Create a Connector Schema Config API.
// Ref. https://fivetran.com/docs/rest-api/connectors#createaconnectorschemaconfig
type ConnectorSchemaConfigCreateService struct {
	httputils.HttpService
	connectorID          *string
	schemaChangeHandling *string
	schemas              map[string]*ConnectorSchemaConfigSchema
}

func (csu *ConnectorSchemaConfigCreateService) request() *connectorSchemaConfigUpdateRequest {
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

func (csu *ConnectorSchemaConfigCreateService) ConnectorID(value string) *ConnectorSchemaConfigCreateService {
	csu.connectorID = &value
	return csu
}

func (csu *ConnectorSchemaConfigCreateService) SchemaChangeHandling(value string) *ConnectorSchemaConfigCreateService {
	csu.schemaChangeHandling = &value
	return csu
}

func (csu *ConnectorSchemaConfigCreateService) Schema(name string, schema *ConnectorSchemaConfigSchema) *ConnectorSchemaConfigCreateService {
	if csu.schemas == nil {
		csu.schemas = make(map[string]*ConnectorSchemaConfigSchema)
	}
	csu.schemas[name] = schema
	return csu
}

func (csu *ConnectorSchemaConfigCreateService) Do(ctx context.Context) (ConnectorSchemaDetailsResponse, error) {
	var response ConnectorSchemaDetailsResponse

	if csu.connectorID == nil {
		return response, fmt.Errorf("missing required ConnectorID")
	}

	url := fmt.Sprintf("/connectors/%v/schemas", *csu.connectorID)
	err := csu.HttpService.Do(ctx, "POST", url, csu.request(), nil, 200, &response)
	return response, err
}
