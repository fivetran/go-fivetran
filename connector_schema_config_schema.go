package fivetran

type ConnectorSchemaConfigSchema struct {
	enabled *bool
	tables  map[string]*ConnectorSchemaConfigTable
}

type connectorSchemaConfigSchemaRequest struct {
	Enabled *bool                                         `json:"enabled,omitempty"`
	Tables  map[string]*connectorSchemaConfigTableRequest `json:"tables,omitempty"`
}

type ConnectorSchemaConfigSchemaResponse struct {
	Enabled *bool                                          `json:"enabled"`
	Tables  map[string]*ConnectorSchemaConfigTableResponse `json:"tables"`
}

func NewConnectorSchemaConfigSchema() *ConnectorSchemaConfigSchema {
	return &ConnectorSchemaConfigSchema{}
}

func (css *ConnectorSchemaConfigSchema) request() *connectorSchemaConfigSchemaRequest {
	var tables map[string]*connectorSchemaConfigTableRequest
	if css.tables != nil && len(css.tables) != 0 {
		tables = make(map[string]*connectorSchemaConfigTableRequest)
		for k, v := range css.tables {
			tables[k] = v.request()
		}
	}

	return &connectorSchemaConfigSchemaRequest{
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
