package fivetran

// DestinationConfig builds Destination Management, Destination Config.
// Ref. https://fivetran.com/docs/rest-api/destinations/config
type DestinationConfig struct {
	host                 *string
	port                 *int
	database             *string
	auth                 *string
	user                 *string
	password             *string
	connectionType       *string
	tunnelHost           *string
	tunnelPort           *string
	tunnelUser           *string
	projectID            *string
	dataSetLocation      *string
	bucket               *string
	serverHostName       *string
	httpPath             *string
	personalAccessToken  *string
	createExternalTables *bool
	externalLocation     *string
	authType             *string
	roleArn              *string
	secretKey            *string
}

type destinationConfigRequest struct {
	Host                 *string `json:"host,omitempty"`
	Port                 *int    `json:"port,omitempty"`
	Database             *string `json:"database,omitempty"`
	Auth                 *string `json:"auth,omitempty"`
	User                 *string `json:"user,omitempty"`
	Password             *string `json:"password,omitempty"`
	ConnectionType       *string `json:"connection_type,omitempty"`
	TunnelHost           *string `json:"tunnel_host,omitempty"`
	TunnelPort           *string `json:"tunnel_port,omitempty"`
	TunnelUser           *string `json:"tunnel_user,omitempty"`
	ProjectID            *string `json:"project_id,omitempty"`
	DataSetLocation      *string `json:"data_set_location,omitempty"`
	Bucket               *string `json:"bucket,omitempty"`
	ServerHostName       *string `json:"server_host_name,omitempty"`
	HTTPPath             *string `json:"http_path,omitempty"`
	PersonalAccessToken  *string `json:"personal_access_token,omitempty"`
	CreateExternalTables *bool   `json:"create_external_tables,omitempty"`
	ExternalLocation     *string `json:"external_location,omitempty"`
	AuthType             *string `json:"auth_type,omitempty"`
	RoleArn              *string `json:"role_arn,omitempty"`
	SecretKey            *string `json:"secret_key,omitempty"`
}

type DestinationConfigResponse struct {
	Host                 string `json:"host"`
	Port                 string `json:"port"` // Port is sent as `string` but returned as `int`. T-97508
	Database             string `json:"database"`
	Auth                 string `json:"auth"`
	User                 string `json:"user"`
	Password             string `json:"password"`
	ConnectionMethod     string `json:"connection_method"` // ConnectionMethod is the REST API's response of ConnectionType. T-111758
	TunnelHost           string `json:"tunnel_host"`
	TunnelPort           string `json:"tunnel_port"`
	TunnelUser           string `json:"tunnel_user"`
	ProjectID            string `json:"project_id"`
	DataSetLocation      string `json:"data_set_location"`
	Bucket               string `json:"bucket"`
	ServerHostName       string `json:"server_host_name"`
	HTTPPath             string `json:"http_path"`
	PersonalAccessToken  string `json:"personal_access_token"`
	CreateExternalTables *bool  `json:"create_external_tables"`
	ExternalLocation     string `json:"external_location"`
	AuthType             string `json:"auth_type"`
	RoleArn              string `json:"role_arn"`
	SecretKey            string `json:"secret_key"`
}

func NewDestinationConfig() *DestinationConfig {
	return &DestinationConfig{}
}

func (dc *DestinationConfig) request() *destinationConfigRequest {
	return &destinationConfigRequest{
		Host:                 dc.host,
		Port:                 dc.port,
		Database:             dc.database,
		Auth:                 dc.auth,
		User:                 dc.user,
		Password:             dc.password,
		ConnectionType:       dc.connectionType,
		TunnelHost:           dc.tunnelHost,
		TunnelPort:           dc.tunnelPort,
		TunnelUser:           dc.tunnelUser,
		ProjectID:            dc.projectID,
		DataSetLocation:      dc.dataSetLocation,
		Bucket:               dc.bucket,
		ServerHostName:       dc.serverHostName,
		HTTPPath:             dc.httpPath,
		PersonalAccessToken:  dc.personalAccessToken,
		CreateExternalTables: dc.createExternalTables,
		ExternalLocation:     dc.externalLocation,
		AuthType:             dc.authType,
		RoleArn:              dc.roleArn,
		SecretKey:            dc.secretKey,
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
