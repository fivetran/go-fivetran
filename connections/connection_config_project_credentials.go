package connections

type ConnectionConfigProjectCredentials struct {
	project   *string
	apiKey    *string
	secretKey *string
}

type connectionConfigProjectCredentialsRequest struct {
	Project   *string `json:"project,omitempty"`
	APIKey    *string `json:"api_key,omitempty"`
	SecretKey *string `json:"secret_key,omitempty"`
}

type ConnectionConfigProjectCredentialsResponse struct {
	Project   string `json:"project"`
	APIKey    string `json:"api_key"`
	SecretKey string `json:"secret_key"`
}

func (pc *ConnectionConfigProjectCredentials) request() *connectionConfigProjectCredentialsRequest {
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

	return &connectionConfigProjectCredentialsRequest{
		Project:   project,
		APIKey:    apiKey,
		SecretKey: secretKey,
	}
}

func (pc *ConnectionConfigProjectCredentials) Project(value string) *ConnectionConfigProjectCredentials {
	pc.project = &value
	return pc
}

func (pc *ConnectionConfigProjectCredentials) APIKey(value string) *ConnectionConfigProjectCredentials {
	pc.apiKey = &value
	return pc
}

func (pc *ConnectionConfigProjectCredentials) SecretKey(value string) *ConnectionConfigProjectCredentials {
	pc.secretKey = &value
	return pc
}
