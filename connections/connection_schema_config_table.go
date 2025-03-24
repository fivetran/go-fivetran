package connections

type ConnectionSchemaConfigTable struct {
	enabled  *bool
	syncMode *string
	columns  map[string]*ConnectionSchemaConfigColumn
}

type ConnectionSchemaConfigTableRequest struct {
	Enabled  *bool                                          `json:"enabled,omitempty"`
	SyncMode *string                                        `json:"sync_mode,omitempty"`
	Columns  map[string]*ConnectionSchemaConfigColumnRequest `json:"columns,omitempty"`
}

type ConnectionSchemaConfigTableResponse struct {
	NameInDestination     *string                                         `json:"name_in_destination"`
	Enabled               *bool                                           `json:"enabled"`
	SyncMode              *string                                         `json:"sync_mode"`
	Columns               map[string]*ConnectionSchemaConfigColumnResponse `json:"columns"`
	SupportsColumnsConfig *bool                                           `json:"supports_columns_config"`
	EnabledPatchSettings  struct {
		Allowed    *bool   `json:"allowed"`
		ReasonCode *string `json:"reason_code"`
		Reason     *string `json:"reason"`
	} `json:"enabled_patch_settings"`
}

func (cst *ConnectionSchemaConfigTable) Request() *ConnectionSchemaConfigTableRequest {
	var columns map[string]*ConnectionSchemaConfigColumnRequest
	if cst.columns != nil && len(cst.columns) != 0 {
		columns = make(map[string]*ConnectionSchemaConfigColumnRequest)
		for k, v := range cst.columns {
			columns[k] = v.Request()
		}
	}

	return &ConnectionSchemaConfigTableRequest{
		Enabled:  cst.enabled,
		SyncMode: cst.syncMode,
		Columns:  columns,
	}
}

func (cst *ConnectionSchemaConfigTable) Enabled(value bool) *ConnectionSchemaConfigTable {
	cst.enabled = &value
	return cst
}

func (cst *ConnectionSchemaConfigTable) SyncMode(value string) *ConnectionSchemaConfigTable {
	cst.syncMode = &value
	return cst
}

func (cst *ConnectionSchemaConfigTable) Column(name string, value *ConnectionSchemaConfigColumn) *ConnectionSchemaConfigTable {
	if cst.columns == nil {
		cst.columns = make(map[string]*ConnectionSchemaConfigColumn)
	}
	cst.columns[name] = value
	return cst
}
