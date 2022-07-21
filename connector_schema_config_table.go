package fivetran

type ConnectorSchemaConfigTable struct {
	enabled *bool
	columns map[string]*ConnectorSchemaConfigColumn
}

type connectorSchemaConfigTableRequest struct {
	Enabled *bool                                          `json:"enabled,omitempty"`
	Columns map[string]*connectorSchemaConfigColumnRequest `json:"columns,omitempty"`
}

type ConnectorSchemaConfigTableResponse struct {
	NameInDestination    *string                                         `json:"name_in_destination"`
	Enabled              *bool                                           `json:"enabled"`
	Columns              map[string]*ConnectorSchemaConfigColumnResponse `json:"columns"`
	EnabledPatchSettings struct {
		Allowed    *bool   `json:"allowed"`
		ReasonCode *string `json:"reason_code"`
		Reason     *string `json:"reason"`
	} `json:"enabled_patch_settings"`
}

func NewConnectorSchemaConfigTable() *ConnectorSchemaConfigTable {
	return &ConnectorSchemaConfigTable{}
}

func (cst *ConnectorSchemaConfigTable) request() *connectorSchemaConfigTableRequest {
	var columns map[string]*connectorSchemaConfigColumnRequest
	if cst.columns != nil && len(cst.columns) != 0 {
		columns = make(map[string]*connectorSchemaConfigColumnRequest)
		for k, v := range cst.columns {
			columns[k] = v.request()
		}
	}

	return &connectorSchemaConfigTableRequest{
		Enabled: cst.enabled,
		Columns: columns,
	}
}

func (cst *ConnectorSchemaConfigTable) Enabled(value bool) *ConnectorSchemaConfigTable {
	cst.enabled = &value
	return cst
}

func (cst *ConnectorSchemaConfigTable) Column(name string, value *ConnectorSchemaConfigColumn) *ConnectorSchemaConfigTable {
	if cst.columns == nil {
		cst.columns = make(map[string]*ConnectorSchemaConfigColumn)
	}
	cst.columns[name] = value
	return cst
}
