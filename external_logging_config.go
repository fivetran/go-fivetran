package fivetran

// ExternalLoggingConfig builds Log Management, Log Config.
// Ref. https://fivetran.com/docs/rest-api/log-service-management#logservicesetupconfigurations
type ExternalLoggingConfig struct {
	workspaceId         *string
	privateKey          *string
	logGroupName        *string
	roleArn             *string
	externalId          *string
	region              *string
	apiKey              *string
	subDomain           *string
	host       			*string
	hostname            *string
	channel             *string
	enableSsl           *string
	token	 			*string
	port                *int
}

type externalLoggingConfigRequest struct {
	WorkspaceId         *string `json:"workspace_id,omitempty"`
	PrivateKey          *string `json:"primary_key,omitempty"`
	LogGroupName        *string `json:"log_group_name,omitempty"`
	RoleArn             *string `json:"role_arn,omitempty"`
	ExternalId          *string `json:"external_id,omitempty"`
	Region              *string `json:"region,omitempty"`
	ApiKey              *string `json:"api_key,omitempty"`
	SubDomain           *string `json:"sub_domain,omitempty"`
	Host       			*string `json:"host,omitempty"`
	Hostname            *string `json:"hostname,omitempty"`
	Channel             *string `json:"channel,omitempty"`
	EnableSsl           *string `json:"enable_ssl,omitempty"`
	Token 				*string `json:"token,omitempty"`
	Port                *int 	`json:"port,omitempty"`
}

type ExternalLoggingConfigResponse struct {
	WorkspaceId         string `json:"workspace_id"`
	PrivateKey          string `json:"primary_key"`
	LogGroupName        string `json:"log_group_name"`
	RoleArn             string `json:"role_arn"`
	ExternalId          string `json:"external_id"`
	Region              string `json:"region"`
	ApiKey              string `json:"api_key"`
	SubDomain           string `json:"sub_domain"`
	Host       			string `json:"host"`
	Hostname            string `json:"hostname"`
	Channel             string `json:"channel"`
	EnableSsl           string `json:"enable_ssl"`
	Token 				string `json:"token"`
	Port                int 	`json:"port"`
}

func NewExternalLoggingConfig() *ExternalLoggingConfig {
	return &ExternalLoggingConfig{}
}

func (dc *ExternalLoggingConfig) request() *externalLoggingConfigRequest {
	return &externalLoggingConfigRequest{
		WorkspaceId         string `json:"workspace_id"`
	PrivateKey          string `json:"primary_key"`
	LogGroupName        string `json:"log_group_name"`
	RoleArn             string `json:"role_arn"`
	ExternalId          string `json:"external_id"`
	Region              string `json:"region"`
	ApiKey              string `json:"api_key"`
	SubDomain           string `json:"sub_domain"`
	Host       			string `json:"host"`
	Hostname            string `json:"hostname"`
	Channel             string `json:"channel"`
	EnableSsl           string `json:"enable_ssl"`
	Token 				string `json:"token"`
	Port                int 	`json:"port"`

		Host:                  dc.host,
		Port:                  dc.port,
		Database:              dc.database,
		Auth:                  dc.auth,
		User:                  dc.user,
		Password:              dc.password,
		ConnectionType:        dc.connectionType,
		TunnelHost:            dc.tunnelHost,
		TunnelPort:            dc.tunnelPort,
		TunnelUser:            dc.tunnelUser,
		ProjectID:             dc.projectID,
		DataSetLocation:       dc.dataSetLocation,
		Bucket:                dc.bucket,
		ServerHostName:        dc.serverHostName,
		HTTPPath:              dc.httpPath,
		PersonalAccessToken:   dc.personalAccessToken,
		CreateExternalTables:  dc.createExternalTables,
		ExternalLocation:      dc.externalLocation,
		AuthType:              dc.authType,
		RoleArn:               dc.roleArn,
		SecretKey:             dc.secretKey,
		PrivateKey:            dc.privateKey,
		ClusterId:             dc.clusterId,
		ClusterRegion:         dc.clusterRegion,
		Role:                  dc.role,
		IsPrivateKeyEncrypted: dc.isPrivateKeyEncrypted,
		Passphrase:            dc.passphrase,
		Catalog:               dc.catalog,
		FivetranRoleArn:       dc.fivetranRoleArn,
		PrefixPath:            dc.prefixPath,
		Region:                dc.region,
	}
}

func (dc *DestinationConfig) Host(value string) *DestinationConfig {
	dc.host = &value
	return dc
}

func (dc *DestinationConfig) Port(value int) *DestinationConfig {
	dc.port = &value
	return dc
}

func (dc *DestinationConfig) Database(value string) *DestinationConfig {
	dc.database = &value
	return dc
}

func (dc *DestinationConfig) Auth(value string) *DestinationConfig {
	dc.auth = &value
	return dc
}

func (dc *DestinationConfig) User(value string) *DestinationConfig {
	dc.user = &value
	return dc
}

func (dc *DestinationConfig) Password(value string) *DestinationConfig {
	dc.password = &value
	return dc
}

func (dc *DestinationConfig) ConnectionType(value string) *DestinationConfig {
	dc.connectionType = &value
	return dc
}

func (dc *DestinationConfig) TunnelHost(value string) *DestinationConfig {
	dc.tunnelHost = &value
	return dc
}

func (dc *DestinationConfig) TunnelPort(value string) *DestinationConfig {
	dc.tunnelPort = &value
	return dc
}

func (dc *DestinationConfig) TunnelUser(value string) *DestinationConfig {
	dc.tunnelUser = &value
	return dc
}

func (dc *DestinationConfig) ProjectID(value string) *DestinationConfig {
	dc.projectID = &value
	return dc
}

func (dc *DestinationConfig) DataSetLocation(value string) *DestinationConfig {
	dc.dataSetLocation = &value
	return dc
}

func (dc *DestinationConfig) Bucket(value string) *DestinationConfig {
	dc.bucket = &value
	return dc
}

func (dc *DestinationConfig) ServerHostName(value string) *DestinationConfig {
	dc.serverHostName = &value
	return dc
}

func (dc *DestinationConfig) HTTPPath(value string) *DestinationConfig {
	dc.httpPath = &value
	return dc
}

func (dc *DestinationConfig) PersonalAccessToken(value string) *DestinationConfig {
	dc.personalAccessToken = &value
	return dc
}

func (dc *DestinationConfig) CreateExternalTables(value bool) *DestinationConfig {
	dc.createExternalTables = &value
	return dc
}

func (dc *DestinationConfig) ExternalLocation(value string) *DestinationConfig {
	dc.externalLocation = &value
	return dc
}

func (dc *DestinationConfig) AuthType(value string) *DestinationConfig {
	dc.authType = &value
	return dc
}

func (dc *DestinationConfig) RoleArn(value string) *DestinationConfig {
	dc.roleArn = &value
	return dc
}

func (dc *DestinationConfig) SecretKey(value string) *DestinationConfig {
	dc.secretKey = &value
	return dc
}

func (dc *DestinationConfig) PrivateKey(value string) *DestinationConfig {
	dc.privateKey = &value
	return dc
}

func (dc *DestinationConfig) ClusterId(value string) *DestinationConfig {
	dc.clusterId = &value
	return dc
}

func (dc *DestinationConfig) ClusterRegion(value string) *DestinationConfig {
	dc.clusterRegion = &value
	return dc
}

func (dc *DestinationConfig) Role(value string) *DestinationConfig {
	dc.role = &value
	return dc
}

func (dc *DestinationConfig) IsPrivateKeyEncrypted(value bool) *DestinationConfig {
	dc.isPrivateKeyEncrypted = &value
	return dc
}

func (dc *DestinationConfig) Passphrase(value string) *DestinationConfig {
	dc.passphrase = &value
	return dc
}

func (dc *DestinationConfig) Catalog(value string) *DestinationConfig {
	dc.catalog = &value
	return dc
}

func (dc *DestinationConfig) FivetranRoleArn(value string) *DestinationConfig {
	dc.fivetranRoleArn = &value
	return dc
}

func (dc *DestinationConfig) PrefixPath(value string) *DestinationConfig {
	dc.prefixPath = &value
	return dc
}

func (dc *DestinationConfig) Region(value string) *DestinationConfig {
	dc.region = &value
	return dc
}
