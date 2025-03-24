package connections

type ConnectionSchemaConfigColumn struct {
	enabled      *bool
	hashed       *bool
	isPrimaryKey *bool
}

type ConnectionSchemaConfigColumnRequest struct {
	Enabled      *bool `json:"enabled,omitempty"`
	Hashed       *bool `json:"hashed,omitempty"`
	IsPrimaryKey *bool `json:"is_primary_key"`
}

type ConnectionSchemaConfigColumnResponse struct {
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

func (csc *ConnectionSchemaConfigColumn) Request() *ConnectionSchemaConfigColumnRequest {
	return &ConnectionSchemaConfigColumnRequest{
		Enabled:      csc.enabled,
		Hashed:       csc.hashed,
		IsPrimaryKey: csc.isPrimaryKey,
	}
}

func (csc *ConnectionSchemaConfigColumn) Enabled(value bool) *ConnectionSchemaConfigColumn {
	csc.enabled = &value
	return csc
}

func (csc *ConnectionSchemaConfigColumn) Hashed(value bool) *ConnectionSchemaConfigColumn {
	csc.hashed = &value
	return csc
}

func (csc *ConnectionSchemaConfigColumn) IsPrimaryKey(value bool) *ConnectionSchemaConfigColumn {
	csc.isPrimaryKey = &value
	return csc
}
