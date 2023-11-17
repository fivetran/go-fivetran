package connectors

import "github.com/fivetran/go-fivetran/utils"

// ConnectorAuth builds Connector Management, Auth.
// Ref. https://fivetran.com/docs/rest-api/connectors
type ConnectorAuth struct {
	clientAccess 				*ConnectorAuthClientAccess
	refreshToken 				*string
	accessToken  				*string
	realmID      				*string
	previousRefreshToken 		*string
	userAccessToken				*string
	consumerSecret				*string
	consumerKey					*string
	oauthToken					*string
	oauthTokenSecret			*string
	roleArn						*string
	awsAccessKey				*string
	awsSecretKey				*string
	clientId					*string
	keyId						*string
	teamId						*string
	clientSecret				*string
}

type connectorAuthRequest struct {
	ClientAccess 				*connectorAuthClientAccessRequest `json:"client_access,omitempty"`
	RefreshToken 				*string `json:"refresh_token,omitempty"`
	AccessToken  				*string `json:"access_token,omitempty"`
	RealmID      				*string `json:"realm_id,omitempty"`
	PreviousRefreshToken 		*string `json:"previous_refresh_token,omitempty"`
	UserAccessToken				*string `json:"user_access_token,omitempty"`
	ConsumerSecret				*string `json:"consumer_secret,omitempty"`
	ConsumerKey					*string `json:"consumer_key,omitempty"`
	OauthToken					*string `json:"oauth_token,omitempty"`
	OauthTokenSecret			*string `json:"oauth_token_secret,omitempty"`
	RoleArn						*string `json:"role_arn,omitempty"`
	AwsAccessKey				*string `json:"aws_access_key,omitempty"`
	AwsSecretKey				*string `json:"aws_secret_key,omitempty"`
	ClientId					*string `json:"client_id,omitempty"`
	KeyId						*string `json:"key_id,omitempty"`
	TeamId						*string `json:"team_id,omitempty"`
	ClientSecret				*string `json:"client_secret,omitempty"`
}

func (ca *ConnectorAuth) Merge(customAuth *map[string]interface{}) (*map[string]interface{}, error) {
	err := utils.MergeIntoMap(ca.Request(), customAuth)
	if err != nil {
		return nil, err
	}
	return customAuth, nil
}

func (ca *ConnectorAuth) Request() *connectorAuthRequest {
	var clientAccess *connectorAuthClientAccessRequest
	if ca.clientAccess != nil {
		clientAccess = ca.clientAccess.request()
	}

	return &connectorAuthRequest{
		ClientAccess:			clientAccess,
		RefreshToken:			ca.refreshToken,
		AccessToken:  			ca.accessToken,
		RealmID:      			ca.realmID,
		PreviousRefreshToken: 	ca.previousRefreshToken,
		UserAccessToken: 		ca.userAccessToken,
		ConsumerSecret: 		ca.consumerSecret,
		ConsumerKey:			ca.consumerKey,
		OauthToken:			 	ca.oauthToken,
		OauthTokenSecret: 		ca.oauthTokenSecret,
		RoleArn: 				ca.roleArn,
		AwsAccessKey: 			ca.awsAccessKey,
		AwsSecretKey: 			ca.awsSecretKey,
		ClientId: 				ca.clientId,
		KeyId: 					ca.keyId,
		TeamId: 				ca.teamId,
		ClientSecret: 			ca.clientSecret,
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

func (ca *ConnectorAuth) PreviousRefreshToken(value string) *ConnectorAuth {
	ca.previousRefreshToken = &value
	return ca
}

func (ca *ConnectorAuth) UserAccessToken(value string) *ConnectorAuth {
	ca.userAccessToken = &value
	return ca
}

func (ca *ConnectorAuth) ConsumerSecret(value string) *ConnectorAuth {
	ca.consumerSecret = &value
	return ca
}

func (ca *ConnectorAuth) ConsumerKey(value string) *ConnectorAuth {
	ca.consumerKey = &value
	return ca
}

func (ca *ConnectorAuth) OauthToken(value string) *ConnectorAuth {
 	ca.oauthToken = &value
	return ca
}

func (ca *ConnectorAuth) OauthTokenSecret(value string) *ConnectorAuth {
	ca.oauthTokenSecret = &value
	return ca
}

func (ca *ConnectorAuth) RoleArn(value string) *ConnectorAuth {
	ca.roleArn = &value
	return ca
}

func (ca *ConnectorAuth) AwsAccessKey(value string) *ConnectorAuth {
	ca.awsAccessKey = &value
	return ca
}

func (ca *ConnectorAuth) AwsSecretKey(value string) *ConnectorAuth {
	ca.awsSecretKey = &value
	return ca
}

func (ca *ConnectorAuth) ClientId(value string) *ConnectorAuth {
	ca.clientId = &value
	return ca
}

func (ca *ConnectorAuth) KeyId(value string) *ConnectorAuth {
	ca.keyId = &value
	return ca
}

func (ca *ConnectorAuth) TeamId(value string) *ConnectorAuth {
	ca.teamId = &value
	return ca
}

func (ca *ConnectorAuth) ClientSecret(value string) *ConnectorAuth {
	ca.clientSecret = &value
	return ca
}