package connections

type ConnectionSchemaConfigSchema struct {
	enabled *bool
	tables  map[string]*ConnectionSchemaConfigTable
}

type ConnectionSchemaConfigSchemaRequest struct {
	Enabled *bool                                         `json:"enabled,omitempty"`
	Tables  map[string]*ConnectionSchemaConfigTableRequest `json:"tables,omitempty"`
}

type ConnectionSchemaConfigSchemaResponse struct {
	NameInDestination *string                                        `json:"name_in_destination"`
	Enabled           *bool                                          `json:"enabled"`
	Tables            map[string]*ConnectionSchemaConfigTableResponse `json:"tables"`
}

func (css *ConnectionSchemaConfigSchema) Request() *ConnectionSchemaConfigSchemaRequest {
	var tables map[string]*ConnectionSchemaConfigTableRequest
	if css.tables != nil && len(css.tables) != 0 {
		tables = make(map[string]*ConnectionSchemaConfigTableRequest)
		for k, v := range css.tables {
			tables[k] = v.Request()
		}
	}

	return &ConnectionSchemaConfigSchemaRequest{
		Enabled: css.enabled,
		Tables:  tables,
	}
}

func (css *ConnectionSchemaConfigSchema) Enabled(value bool) *ConnectionSchemaConfigSchema {
	css.enabled = &value
	return css
}

func (css *ConnectionSchemaConfigSchema) Table(name string, table *ConnectionSchemaConfigTable) *ConnectionSchemaConfigSchema {
	if css.tables == nil {
		css.tables = make(map[string]*ConnectionSchemaConfigTable)
	}
	css.tables[name] = table
	return css
}
