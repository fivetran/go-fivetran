package fivetran

type ConnectorAuth struct {
	FClientAccess *ConnectorAuthClientAccess `json:"client_access,omitempty"`
	FRefreshToken *string                    `json:"refresh_token,omitempty"`
	FAccessToken  *string                    `json:"access_token,omitempty"`
	FRealmID      *string                    `json:"realm_id,omitempty"`
}

func NewConnectorAuth() *ConnectorAuth {
	return &ConnectorAuth{}
}

func (a *ConnectorAuth) ClientAccess(value ConnectorAuthClientAccess) *ConnectorAuth {
	a.FClientAccess = &value
	return a
}

func (a *ConnectorAuth) RefreshToken(value string) *ConnectorAuth {
	a.FRefreshToken = &value
	return a
}

func (a *ConnectorAuth) AccessToken(value string) *ConnectorAuth {
	a.FAccessToken = &value
	return a
}

func (a *ConnectorAuth) RealmID(value string) *ConnectorAuth {
	a.FRealmID = &value
	return a
}
