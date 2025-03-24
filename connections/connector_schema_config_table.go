package connectors

type ConnectorSchemaConfigTable struct {
	enabled  *bool
	syncMode *string
	columns  map[string]*ConnectorSchemaConfigColumn
}

type ConnectorSchemaConfigTableRequest struct {
	Enabled  *bool                                          `json:"enabled,omitempty"`
	SyncMode *string                                        `json:"sync_mode,omitempty"`
	Columns  map[string]*ConnectorSchemaConfigColumnRequest `json:"columns,omitempty"`
}

type ConnectorSchemaConfigTableResponse struct {
	NameInDestination     *string                                         `json:"name_in_destination"`
	Enabled               *bool                                           `json:"enabled"`
	SyncMode              *string                                         `json:"sync_mode"`
	Columns               map[string]*ConnectorSchemaConfigColumnResponse `json:"columns"`
	SupportsColumnsConfig *bool                                           `json:"supports_columns_config"`
	EnabledPatchSettings  struct {
		Allowed    *bool   `json:"allowed"`
		ReasonCode *string `json:"reason_code"`
		Reason     *string `json:"reason"`
	} `json:"enabled_patch_settings"`
}

func (cst *ConnectorSchemaConfigTable) Request() *ConnectorSchemaConfigTableRequest {
	var columns map[string]*ConnectorSchemaConfigColumnRequest
	if cst.columns != nil && len(cst.columns) != 0 {
		columns = make(map[string]*ConnectorSchemaConfigColumnRequest)
		for k, v := range cst.columns {
			columns[k] = v.Request()
		}
	}

	return &ConnectorSchemaConfigTableRequest{
		Enabled:  cst.enabled,
		SyncMode: cst.syncMode,
		Columns:  columns,
	}
}

func (cst *ConnectorSchemaConfigTable) Enabled(value bool) *ConnectorSchemaConfigTable {
	cst.enabled = &value
	return cst
}

func (cst *ConnectorSchemaConfigTable) SyncMode(value string) *ConnectorSchemaConfigTable {
	cst.syncMode = &value
	return cst
}

func (cst *ConnectorSchemaConfigTable) Column(name string, value *ConnectorSchemaConfigColumn) *ConnectorSchemaConfigTable {
	if cst.columns == nil {
		cst.columns = make(map[string]*ConnectorSchemaConfigColumn)
	}
	cst.columns[name] = value
	return cst
}
