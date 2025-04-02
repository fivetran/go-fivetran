package connections

import "github.com/fivetran/go-fivetran/utils"

type ConnectionAuth struct {
	clientAccess 				*ConnectionAuthClientAccess
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

type connectionAuthRequest struct {
	ClientAccess 				*connectionAuthClientAccessRequest `json:"client_access,omitempty"`
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

func (ca *ConnectionAuth) Merge(customAuth *map[string]interface{}) (*map[string]interface{}, error) {
	err := utils.MergeIntoMap(ca.Request(), customAuth)
	if err != nil {
		return nil, err
	}
	return customAuth, nil
}

func (ca *ConnectionAuth) Request() *connectionAuthRequest {
	var clientAccess *connectionAuthClientAccessRequest
	if ca.clientAccess != nil {
		clientAccess = ca.clientAccess.request()
	}

	return &connectionAuthRequest{
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

func (ca *ConnectionAuth) ClientAccess(value *ConnectionAuthClientAccess) *ConnectionAuth {
	ca.clientAccess = value
	return ca
}

func (ca *ConnectionAuth) RefreshToken(value string) *ConnectionAuth {
	ca.refreshToken = &value
	return ca
}

func (ca *ConnectionAuth) AccessToken(value string) *ConnectionAuth {
	ca.accessToken = &value
	return ca
}

func (ca *ConnectionAuth) RealmID(value string) *ConnectionAuth {
	ca.realmID = &value
	return ca
}

func (ca *ConnectionAuth) PreviousRefreshToken(value string) *ConnectionAuth {
	ca.previousRefreshToken = &value
	return ca
}

func (ca *ConnectionAuth) UserAccessToken(value string) *ConnectionAuth {
	ca.userAccessToken = &value
	return ca
}

func (ca *ConnectionAuth) ConsumerSecret(value string) *ConnectionAuth {
	ca.consumerSecret = &value
	return ca
}

func (ca *ConnectionAuth) ConsumerKey(value string) *ConnectionAuth {
	ca.consumerKey = &value
	return ca
}

func (ca *ConnectionAuth) OauthToken(value string) *ConnectionAuth {
 	ca.oauthToken = &value
	return ca
}

func (ca *ConnectionAuth) OauthTokenSecret(value string) *ConnectionAuth {
	ca.oauthTokenSecret = &value
	return ca
}

func (ca *ConnectionAuth) RoleArn(value string) *ConnectionAuth {
	ca.roleArn = &value
	return ca
}

func (ca *ConnectionAuth) AwsAccessKey(value string) *ConnectionAuth {
	ca.awsAccessKey = &value
	return ca
}

func (ca *ConnectionAuth) AwsSecretKey(value string) *ConnectionAuth {
	ca.awsSecretKey = &value
	return ca
}

func (ca *ConnectionAuth) ClientId(value string) *ConnectionAuth {
	ca.clientId = &value
	return ca
}

func (ca *ConnectionAuth) KeyId(value string) *ConnectionAuth {
	ca.keyId = &value
	return ca
}

func (ca *ConnectionAuth) TeamId(value string) *ConnectionAuth {
	ca.teamId = &value
	return ca
}

func (ca *ConnectionAuth) ClientSecret(value string) *ConnectionAuth {
	ca.clientSecret = &value
	return ca
}