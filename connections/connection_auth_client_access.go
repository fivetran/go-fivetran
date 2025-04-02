package connections

type ConnectionAuthClientAccess struct {
	clientID       *string
	clientSecret   *string
	userAgent      *string
	developerToken *string
}

type connectionAuthClientAccessRequest struct {
	ClientID       *string `json:"client_id,omitempty"`
	ClientSecret   *string `json:"client_secret,omitempty"`
	UserAgent      *string `json:"user_agent,omitempty"`
	DeveloperToken *string `json:"developer_token,omitempty"`
}

func (ca *ConnectionAuthClientAccess) request() *connectionAuthClientAccessRequest {
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

	return &connectionAuthClientAccessRequest{
		ClientID:       clientID,
		ClientSecret:   clientSecret,
		UserAgent:      userAgent,
		DeveloperToken: developerToken,
	}
}

func (ca *ConnectionAuthClientAccess) ClientID(value string) *ConnectionAuthClientAccess {
	ca.clientID = &value
	return ca
}

func (ca *ConnectionAuthClientAccess) ClientSecret(value string) *ConnectionAuthClientAccess {
	ca.clientSecret = &value
	return ca
}

func (ca *ConnectionAuthClientAccess) UserAgent(value string) *ConnectionAuthClientAccess {
	ca.userAgent = &value
	return ca
}

func (ca *ConnectionAuthClientAccess) DeveloperToken(value string) *ConnectionAuthClientAccess {
	ca.developerToken = &value
	return ca
}
