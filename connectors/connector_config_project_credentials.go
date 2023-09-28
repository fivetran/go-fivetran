package connectors

// ConnectorConfigProjectCredentials builds Connector Management, Connector Config Project Credentials.
// Ref. https://fivetran.com/docs/rest-api/connectors/config
type ConnectorConfigProjectCredentials struct {
	project   *string
	apiKey    *string
	secretKey *string
}

type connectorConfigProjectCredentialsRequest struct {
	Project   *string `json:"project,omitempty"`
	APIKey    *string `json:"api_key,omitempty"`
	SecretKey *string `json:"secret_key,omitempty"`
}

type ConnectorConfigProjectCredentialsResponse struct {
	Project   string `json:"project"`
	APIKey    string `json:"api_key"`
	SecretKey string `json:"secret_key"`
}

func (pc *ConnectorConfigProjectCredentials) request() *connectorConfigProjectCredentialsRequest {
	var project *string
	if pc.project != nil {
		project = pc.project
	}

	var apiKey *string
	if pc.apiKey != nil {
		apiKey = pc.apiKey
	}

	var secretKey *string
	if pc.secretKey != nil {
		secretKey = pc.secretKey
	}

	return &connectorConfigProjectCredentialsRequest{
		Project:   project,
		APIKey:    apiKey,
		SecretKey: secretKey,
	}
}

func (pc *ConnectorConfigProjectCredentials) Project(value string) *ConnectorConfigProjectCredentials {
	pc.project = &value
	return pc
}

func (pc *ConnectorConfigProjectCredentials) APIKey(value string) *ConnectorConfigProjectCredentials {
	pc.apiKey = &value
	return pc
}

func (pc *ConnectorConfigProjectCredentials) SecretKey(value string) *ConnectorConfigProjectCredentials {
	pc.secretKey = &value
	return pc
}
