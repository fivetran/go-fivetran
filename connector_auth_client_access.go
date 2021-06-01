package fivetran

type ConnectorAuthClientAccess struct {
	FClientID       *string `json:"client_id,omitempty"`
	FClientSecret   *string `json:"client_secret,omitempty"`
	FUserAgent      *string `json:"user_agent,omitempty"`
	FDeveloperToken *string `json:"developer_token,omitempty"`
}

func NewConnectorAuthClientAccess() *ConnectorAuthClientAccess {
	return &ConnectorAuthClientAccess{}
}

func (ca *ConnectorAuthClientAccess) ClientID(value string) *ConnectorAuthClientAccess {
	ca.FClientID = &value
	return ca
}

func (ca *ConnectorAuthClientAccess) ClientSecret(value string) *ConnectorAuthClientAccess {
	ca.FClientSecret = &value
	return ca
}

func (ca *ConnectorAuthClientAccess) UserAgent(value string) *ConnectorAuthClientAccess {
	ca.FUserAgent = &value
	return ca
}

func (ca *ConnectorAuthClientAccess) DeveloperToken(value string) *ConnectorAuthClientAccess {
	ca.FDeveloperToken = &value
	return ca
}
