package fivetran

type ConnectorAuth struct {
	clientAccess *ConnectorAuthClientAccess
	refreshToken *string
	accessToken  *string
	realmID      *string
}

type connectorAuthRequest struct {
	ClientAccess *connectorAuthClientAccessRequest `json:"client_access,omitempty"`
	RefreshToken *string                           `json:"refresh_token,omitempty"`
	AccessToken  *string                           `json:"access_token,omitempty"`
	RealmID      *string                           `json:"realm_id,omitempty"`
}

func NewConnectorAuth() *ConnectorAuth {
	return &ConnectorAuth{}
}

func (ca *ConnectorAuth) request() *connectorAuthRequest {
	var clientAccess *connectorAuthClientAccessRequest
	if ca.clientAccess != nil {
		clientAccess = ca.clientAccess.request()
	}

	return &connectorAuthRequest{
		ClientAccess: clientAccess,
		RefreshToken: ca.refreshToken,
		AccessToken:  ca.accessToken,
		RealmID:      ca.realmID,
	}
}

func (ca *ConnectorAuth) ClientAccess(value *ConnectorAuthClientAccess) *ConnectorAuth {
	ca.clientAccess = value
	return ca
}

func (ca *ConnectorAuth) RefreshToken(value string) *ConnectorAuth {
	ca.refreshToken = &value
	return ca
}

func (ca *ConnectorAuth) AccessToken(value string) *ConnectorAuth {
	ca.accessToken = &value
	return ca
}

func (ca *ConnectorAuth) RealmID(value string) *ConnectorAuth {
	ca.realmID = &value
	return ca
}
