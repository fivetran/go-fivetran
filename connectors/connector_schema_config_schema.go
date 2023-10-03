package connectors

type ConnectorSchemaConfigSchema struct {
	enabled *bool
	tables  map[string]*ConnectorSchemaConfigTable
}

type ConnectorSchemaConfigSchemaRequest struct {
	Enabled *bool                                         `json:"enabled,omitempty"`
	Tables  map[string]*ConnectorSchemaConfigTableRequest `json:"tables,omitempty"`
}

type ConnectorSchemaConfigSchemaResponse struct {
	NameInDestination *string                                        `json:"name_in_destination"`
	Enabled           *bool                                          `json:"enabled"`
	Tables            map[string]*ConnectorSchemaConfigTableResponse `json:"tables"`
}

func (css *ConnectorSchemaConfigSchema) Request() *ConnectorSchemaConfigSchemaRequest {
	var tables map[string]*ConnectorSchemaConfigTableRequest
	if css.tables != nil && len(css.tables) != 0 {
		tables = make(map[string]*ConnectorSchemaConfigTableRequest)
		for k, v := range css.tables {
			tables[k] = v.Request()
		}
	}

	return &ConnectorSchemaConfigSchemaRequest{
		Enabled: css.enabled,
		Tables:  tables,
	}
}

func (css *ConnectorSchemaConfigSchema) Enabled(value bool) *ConnectorSchemaConfigSchema {
	css.enabled = &value
	return css
}

func (css *ConnectorSchemaConfigSchema) Table(name string, table *ConnectorSchemaConfigTable) *ConnectorSchemaConfigSchema {
	if css.tables == nil {
		css.tables = make(map[string]*ConnectorSchemaConfigTable)
	}
	css.tables[name] = table
	return css
}
