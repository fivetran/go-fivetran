package connectors

// ConnectorAuthClientAccess builds Connector Management, Auth Client Access.
// Ref. https://fivetran.com/docs/rest-api/connectors
type ConnectorAuthClientAccess struct {
	clientID       *string
	clientSecret   *string
	userAgent      *string
	developerToken *string
}

type connectorAuthClientAccessRequest struct {
	ClientID       *string `json:"client_id,omitempty"`
	ClientSecret   *string `json:"client_secret,omitempty"`
	UserAgent      *string `json:"user_agent,omitempty"`
	DeveloperToken *string `json:"developer_token,omitempty"`
}

func (ca *ConnectorAuthClientAccess) request() *connectorAuthClientAccessRequest {
	var clientID *string
	if ca.clientID != nil {
		clientID = ca.clientID
	}

	var clientSecret *string
	if ca.clientSecret != nil {
		clientSecret = ca.clientSecret
	}

	var userAgent *string
	if ca.userAgent != nil {
		userAgent = ca.userAgent
	}

	var developerToken *string
	if ca.developerToken != nil {
		developerToken = ca.developerToken
	}

	return &connectorAuthClientAccessRequest{
		ClientID:       clientID,
		ClientSecret:   clientSecret,
		UserAgent:      userAgent,
		DeveloperToken: developerToken,
	}
}

func (ca *ConnectorAuthClientAccess) ClientID(value string) *ConnectorAuthClientAccess {
	ca.clientID = &value
	return ca
}

func (ca *ConnectorAuthClientAccess) ClientSecret(value string) *ConnectorAuthClientAccess {
	ca.clientSecret = &value
	return ca
}

func (ca *ConnectorAuthClientAccess) UserAgent(value string) *ConnectorAuthClientAccess {
	ca.userAgent = &value
	return ca
}

func (ca *ConnectorAuthClientAccess) DeveloperToken(value string) *ConnectorAuthClientAccess {
	ca.developerToken = &value
	return ca
}
