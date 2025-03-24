package connections

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type ConnectionSchemaConfigUpdateService struct {
    httputils.HttpService
	connectionID          *string
	schemaChangeHandling *string
	schemas              map[string]*ConnectionSchemaConfigSchema
}

func (csu *ConnectionSchemaConfigUpdateService) request() *connectionSchemaConfigUpdateRequest {
	var schemas map[string]*ConnectionSchemaConfigSchemaRequest
	if csu.schemas != nil && len(csu.schemas) != 0 {
		schemas = make(map[string]*ConnectionSchemaConfigSchemaRequest)
		for k, v := range csu.schemas {
			schemas[k] = v.Request()
		}
	}

	return &connectionSchemaConfigUpdateRequest{
		SchemaChangeHandling: csu.schemaChangeHandling,
		Schemas:              schemas,
	}
}

func (csu *ConnectionSchemaConfigUpdateService) ConnectionID(value string) *ConnectionSchemaConfigUpdateService {
	csu.connectionID = &value
	return csu
}

func (csu *ConnectionSchemaConfigUpdateService) SchemaChangeHandling(value string) *ConnectionSchemaConfigUpdateService {
	csu.schemaChangeHandling = &value
	return csu
}

func (csu *ConnectionSchemaConfigUpdateService) Schema(name string, schema *ConnectionSchemaConfigSchema) *ConnectionSchemaConfigUpdateService {
	if csu.schemas == nil {
		csu.schemas = make(map[string]*ConnectionSchemaConfigSchema)
	}
	csu.schemas[name] = schema
	return csu
}

func (csu *ConnectionSchemaConfigUpdateService) Do(ctx context.Context) (ConnectionSchemaDetailsResponse, error) {
    var response ConnectionSchemaDetailsResponse

	if csu.connectionID == nil {
		return response, fmt.Errorf("missing required ConnectionID")
	}

    url := fmt.Sprintf("/connections/%v/schemas", *csu.connectionID)
    err := csu.HttpService.Do(ctx, "PATCH", url, csu.request(), nil, 200, &response)
    return response, err
}