package fivetran

type connectorAuth struct {
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

func NewConnectorAuth() *connectorAuth {
	return &connectorAuth{}
}

func (ca *connectorAuth) request() *connectorAuthRequest {
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

func (ca *connectorAuth) ClientAccess(value *ConnectorAuthClientAccess) *connectorAuth {
	ca.clientAccess = value
	return ca
}

func (ca *connectorAuth) RefreshToken(value string) *connectorAuth {
	ca.refreshToken = &value
	return ca
}

func (ca *connectorAuth) AccessToken(value string) *connectorAuth {
	ca.accessToken = &value
	return ca
}

func (ca *connectorAuth) RealmID(value string) *connectorAuth {
	ca.realmID = &value
	return ca
}
