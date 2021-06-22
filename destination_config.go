package fivetran

type destinationConfig struct {
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
}

type DestinationConfigResponse struct {
	Host                 string `json:"host"`
	Port                 string `json:"port"`
	Database             string `json:"database"`
	Auth                 string `json:"auth"`
	User                 string `json:"user"`
	Password             string `json:"password"`
	ConnectionType       string `json:"connection_type"`
	TunnelHost           string `json:"tunnel_host"`
	TunnelPort           string `json:"tunnel_port"`
	TunnelUser           string `json:"tunnel_user"`
	ProjectID            string `json:"project_id"`
	DataSetLocation      string `json:"data_set_location"`
	Bucket               string `json:"bucket"`
	ServerHostName       string `json:"server_host_name"`
	HTTPPath             string `json:"http_path"`
	PersonalAccessToken  string `json:"personal_access_token"`
	CreateExternalTables bool   `json:"create_external_tables"`
	ExternalLocation     string `json:"external_location"`
	AuthType             string `json:"auth_type"`
	RoleArn              string `json:"role_arn"`
}

func NewDestinationConfig() *destinationConfig {
	return &destinationConfig{}
}

func (dc *destinationConfig) request() *destinationConfigRequest {
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
	}
}

func (dc *destinationConfig) Host(value string) *destinationConfig {
	dc.host = &value
	return dc
}

func (dc *destinationConfig) Port(value int) *destinationConfig {
	dc.port = &value
	return dc
}

func (dc *destinationConfig) Database(value string) *destinationConfig {
	dc.database = &value
	return dc
}

func (dc *destinationConfig) Auth(value string) *destinationConfig {
	dc.auth = &value
	return dc
}

func (dc *destinationConfig) User(value string) *destinationConfig {
	dc.user = &value
	return dc
}

func (dc *destinationConfig) Password(value string) *destinationConfig {
	dc.password = &value
	return dc
}

func (dc *destinationConfig) ConnectionType(value string) *destinationConfig {
	dc.connectionType = &value
	return dc
}

func (dc *destinationConfig) TunnelHost(value string) *destinationConfig {
	dc.tunnelHost = &value
	return dc
}

func (dc *destinationConfig) TunnelPort(value string) *destinationConfig {
	dc.tunnelPort = &value
	return dc
}

func (dc *destinationConfig) TunnelUser(value string) *destinationConfig {
	dc.tunnelUser = &value
	return dc
}

func (dc *destinationConfig) ProjectID(value string) *destinationConfig {
	dc.projectID = &value
	return dc
}

func (dc *destinationConfig) DataSetLocation(value string) *destinationConfig {
	dc.dataSetLocation = &value
	return dc
}

func (dc *destinationConfig) Bucket(value string) *destinationConfig {
	dc.bucket = &value
	return dc
}

func (dc *destinationConfig) ServerHostName(value string) *destinationConfig {
	dc.serverHostName = &value
	return dc
}

func (dc *destinationConfig) HTTPPath(value string) *destinationConfig {
	dc.httpPath = &value
	return dc
}

func (dc *destinationConfig) PersonalAccessToken(value string) *destinationConfig {
	dc.personalAccessToken = &value
	return dc
}

func (dc *destinationConfig) CreateExternalTables(value bool) *destinationConfig {
	dc.createExternalTables = &value
	return dc
}

func (dc *destinationConfig) ExternalLocation(value string) *destinationConfig {
	dc.externalLocation = &value
	return dc
}

func (dc *destinationConfig) AuthType(value string) *destinationConfig {
	dc.authType = &value
	return dc
}

func (dc *destinationConfig) RoleArn(value string) *destinationConfig {
	dc.roleArn = &value
	return dc
}
