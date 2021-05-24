package fivetran

type DestinationConfig struct {
	Fhost                 string      `json:"host,omitempty"`
	Fport                 interface{} `json:"port,omitempty"` // Type should change to int when https://fivetran.height.app/T-97508 fixed.
	Fdatabase             string      `json:"database,omitempty"`
	Fauth                 string      `json:"auth,omitempty"`
	Fuser                 string      `json:"user,omitempty"`
	Fpassword             string      `json:"password,omitempty"`
	FconnectionType       string      `json:"connection_type,omitempty"`
	FtunnelHost           string      `json:"tunnel_host,omitempty"`
	FtunnelPort           string      `json:"tunnel_port,omitempty"`
	FtunnelUser           string      `json:"tunnel_user,omitempty"`
	FprojectID            string      `json:"project_id,omitempty"`
	FdataSetLocation      string      `json:"data_set_location,omitempty"`
	Fbucket               string      `json:"bucket,omitempty"`
	FserverHostName       string      `json:"server_host_name,omitempty"`
	FhttpPath             string      `json:"http_path,omitempty"`
	FpersonalAccessToken  string      `json:"personal_access_token,omitempty"`
	FcreateExternalTables *bool       `json:"create_external_tables,omitempty"`
	FexternalLocation     string      `json:"external_location,omitempty"`
	FauthType             string      `json:"auth_type,omitempty"`
	FroleArn              string      `json:"role_arn,omitempty"`
}

func NewDestinationConfig() *DestinationConfig {
	return &DestinationConfig{}
}

func (dc *DestinationConfig) Host(host string) *DestinationConfig {
	dc.Fhost = host
	return dc
}

func (dc *DestinationConfig) Port(port int) *DestinationConfig {
	dc.Fport = port
	return dc
}

func (dc *DestinationConfig) Database(database string) *DestinationConfig {
	dc.Fdatabase = database
	return dc
}

func (dc *DestinationConfig) Auth(auth string) *DestinationConfig {
	dc.Fauth = auth
	return dc
}

func (dc *DestinationConfig) User(user string) *DestinationConfig {
	dc.Fuser = user
	return dc
}

func (dc *DestinationConfig) Password(password string) *DestinationConfig {
	dc.Fpassword = password
	return dc
}

func (dc *DestinationConfig) ConnectionType(connectionType string) *DestinationConfig {
	dc.FconnectionType = connectionType
	return dc
}

func (dc *DestinationConfig) TunnelHost(tunnelHost string) *DestinationConfig {
	dc.FtunnelHost = tunnelHost
	return dc
}

func (dc *DestinationConfig) TunnelPort(tunnelPort string) *DestinationConfig {
	dc.FtunnelPort = tunnelPort
	return dc
}

func (dc *DestinationConfig) TunnelUser(tunnelUser string) *DestinationConfig {
	dc.FtunnelUser = tunnelUser
	return dc
}

func (dc *DestinationConfig) ProjectID(projectID string) *DestinationConfig {
	dc.FprojectID = projectID
	return dc
}

func (dc *DestinationConfig) DataSetLocation(dataSetLocation string) *DestinationConfig {
	dc.FdataSetLocation = dataSetLocation
	return dc
}

func (dc *DestinationConfig) Bucket(bucket string) *DestinationConfig {
	dc.Fbucket = bucket
	return dc
}

func (dc *DestinationConfig) ServerHostName(serverHostName string) *DestinationConfig {
	dc.FserverHostName = serverHostName
	return dc
}

func (dc *DestinationConfig) HttpPath(httpPath string) *DestinationConfig {
	dc.FhttpPath = httpPath
	return dc
}

func (dc *DestinationConfig) PersonalAccessToken(personalAccessToken string) *DestinationConfig {
	dc.FpersonalAccessToken = personalAccessToken
	return dc
}

func (dc *DestinationConfig) CreateExternalTables(createExternalTables bool) *DestinationConfig {
	dc.FcreateExternalTables = &createExternalTables
	return dc
}

func (dc *DestinationConfig) ExternalLocation(externalLocation string) *DestinationConfig {
	dc.FexternalLocation = externalLocation
	return dc
}

func (dc *DestinationConfig) AuthType(authType string) *DestinationConfig {
	dc.FauthType = authType
	return dc
}

func (dc *DestinationConfig) RoleArn(roleArn string) *DestinationConfig {
	dc.FroleArn = roleArn
	return dc
}
