package fivetran

type ConnectorSchemaConfigColumn struct {
	enabled *bool
	hashed  *bool
}

type connectorSchemaConfigColumnRequest struct {
	Enabled *bool `json:"enabled,omitempty"`
	Hashed  *bool `json:"hashed,omitempty"`
}

type ConnectorSchemaConfigColumnResponse struct {
	NameInDestination    *string `json:"name_in_destination"`
	Enabled              *bool   `json:"enabled"`
	Hashed               *bool   `json:"hashed"`
	EnabledPatchSettings struct {
		Allowed    *bool   `json:"allowed"`
		ReasonCode *string `json:"reason_code"`
		Reason     *string `json:"reason"`
	} `json:"enabled_patch_settings"`
}

func NewConnectorSchemaConfigColumn() *ConnectorSchemaConfigColumn {
	return &ConnectorSchemaConfigColumn{}
}

func (csc *ConnectorSchemaConfigColumn) request() *connectorSchemaConfigColumnRequest {
	return &connectorSchemaConfigColumnRequest{
		Enabled: csc.enabled,
		Hashed:  csc.hashed,
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
