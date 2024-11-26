package connectors

type ConnectorSchemaConfigColumn struct {
	enabled      *bool
	hashed       *bool
	isPrimaryKey *bool
}

type ConnectorSchemaConfigColumnRequest struct {
	Enabled      *bool `json:"enabled,omitempty"`
	Hashed       *bool `json:"hashed,omitempty"`
	IsPrimaryKey *bool `json:"is_primary_key"`
}

type ConnectorSchemaConfigColumnResponse struct {
	NameInDestination    *string `json:"name_in_destination"`
	Enabled              *bool   `json:"enabled"`
	Hashed               *bool   `json:"hashed"`
	IsPrimaryKey         *bool   `json:"is_primary_key"`
	EnabledPatchSettings struct {
		Allowed    *bool   `json:"allowed"`
		ReasonCode *string `json:"reason_code"`
		Reason     *string `json:"reason"`
	} `json:"enabled_patch_settings"`
}

func (csc *ConnectorSchemaConfigColumn) Request() *ConnectorSchemaConfigColumnRequest {
	return &ConnectorSchemaConfigColumnRequest{
		Enabled:      csc.enabled,
		Hashed:       csc.hashed,
		IsPrimaryKey: csc.isPrimaryKey,
	}
}

func (csc *ConnectorSchemaConfigColumn) Enabled(value bool) *ConnectorSchemaConfigColumn {
	csc.enabled = &value
	return csc
}

func (csc *ConnectorSchemaConfigColumn) Hashed(value bool) *ConnectorSchemaConfigColumn {
	csc.hashed = &value
	return csc
}

func (csc *ConnectorSchemaConfigColumn) IsPrimaryKey(value bool) *ConnectorSchemaConfigColumn {
	csc.isPrimaryKey = &value
	return csc
}
