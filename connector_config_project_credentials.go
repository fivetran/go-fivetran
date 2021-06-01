package fivetran

type ConnectorConfigProjectCredentials struct {
	FProject   *string `json:"project,omitempty"`
	FAPIKey    *string `json:"api_key,omitempty"`
	FSecretKey *string `json:"secret_key,omitempty"`
}

func NewConnectorConfigProjectCredentials() *ConnectorConfigProjectCredentials {
	return &ConnectorConfigProjectCredentials{}
}

func (pc *ConnectorConfigProjectCredentials) Project(value string) *ConnectorConfigProjectCredentials {
	pc.FProject = &value
	return pc
}

func (pc *ConnectorConfigProjectCredentials) APIKey(value string) *ConnectorConfigProjectCredentials {
	pc.FAPIKey = &value
	return pc
}

func (pc *ConnectorConfigProjectCredentials) SecretKey(value string) *ConnectorConfigProjectCredentials {
	pc.FSecretKey = &value
	return pc
}
