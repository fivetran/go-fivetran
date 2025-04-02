package destinations

type DestinationConfig struct {
    host                  *string
    port                  *int
    database              *string
    auth                  *string
    user                  *string
    password              *string
    connectionType        *string
    tunnelHost            *string
    tunnelPort            *string
    tunnelUser            *string
    projectID             *string
    dataSetLocation       *string
    bucket                *string
    serverHostName        *string
    httpPath              *string
    personalAccessToken   *string
    createExternalTables  *bool
    externalLocation      *string
    authType              *string
    roleArn               *string
    secretKey             *string
    privateKey            *string
    clusterId             *string
    clusterRegion         *string
    role                  *string
    isPrivateKeyEncrypted *bool
    passphrase            *string
    catalog               *string
    fivetranRoleArn       *string
    prefixPath            *string
    region                *string
    storageAccountName    *string
    containerName         *string
    tenantId              *string
    clientId              *string
    secretValue           *string
    workspaceName         *string
    lakehouseName         *string
}

type destinationConfigRequest struct {
    Host                  *string `json:"host,omitempty"`
    Port                  *int    `json:"port,omitempty"`
    Database              *string `json:"database,omitempty"`
    Auth                  *string `json:"auth,omitempty"`
    User                  *string `json:"user,omitempty"`
    Password              *string `json:"password,omitempty"`
    ConnectionType        *string `json:"connection_type,omitempty"`
    TunnelHost            *string `json:"tunnel_host,omitempty"`
    TunnelPort            *string `json:"tunnel_port,omitempty"`
    TunnelUser            *string `json:"tunnel_user,omitempty"`
    ProjectID             *string `json:"project_id,omitempty"`
    DataSetLocation       *string `json:"data_set_location,omitempty"`
    Bucket                *string `json:"bucket,omitempty"`
    ServerHostName        *string `json:"server_host_name,omitempty"`
    HTTPPath              *string `json:"http_path,omitempty"`
    PersonalAccessToken   *string `json:"personal_access_token,omitempty"`
    CreateExternalTables  *bool   `json:"create_external_tables,omitempty"`
    ExternalLocation      *string `json:"external_location,omitempty"`
    AuthType              *string `json:"auth_type,omitempty"`
    RoleArn               *string `json:"role_arn,omitempty"`
    SecretKey             *string `json:"secret_key,omitempty"`
    PrivateKey            *string `json:"private_key,omitempty"`
    ClusterId             *string `json:"cluster_id,omitempty"`
    ClusterRegion         *string `json:"cluster_region,omitempty"`
    Role                  *string `json:"role,omitempty"`
    IsPrivateKeyEncrypted *bool   `json:"is_private_key_encrypted,omitempty"`
    Passphrase            *string `json:"passphrase,omitempty"`
    Catalog               *string `json:"catalog,omitempty"`
    FivetranRoleArn       *string `json:"fivetran_role_arn,omitempty"`
    PrefixPath            *string `json:"prefix_path,omitempty"`
    Region                *string `json:"region,omitempty"`
    StorageAccountName    *string `json:"storage_account_name,omitempty"`
    ContainerName         *string `json:"container_name,omitempty"`
    TenantId              *string `json:"tenant_id,omitempty"`
    ClientId              *string `json:"client_id,omitempty"`
    SecretValue           *string `json:"secret_value,omitempty"`
    WorkspaceName         *string `json:"workspace_name,omitempty"`
    LakehouseName         *string `json:"lakehouse_name,omitempty"`
}

type DestinationConfigResponse struct {
    Host                  string `json:"host"`
    Port                  string `json:"port"` // Port is sent as `string` but returned as `int`. T-97508
    Database              string `json:"database"`
    Auth                  string `json:"auth"`
    User                  string `json:"user"`
    Password              string `json:"password"`
    ConnectionType        string `json:"connection_type"` // ConnectionMethod is the REST API's response of ConnectionType. T-111758
    TunnelHost            string `json:"tunnel_host"`
    TunnelPort            string `json:"tunnel_port"`
    TunnelUser            string `json:"tunnel_user"`
    ProjectID             string `json:"project_id"`
    DataSetLocation       string `json:"data_set_location"`
    Location              string `json:"location"` // Big Query returns `data_set_location` as `location` in response (will be fixed with migration to API V2)
    Bucket                string `json:"bucket"`
    ServerHostName        string `json:"server_host_name"`
    HTTPPath              string `json:"http_path"`
    PersonalAccessToken   string `json:"personal_access_token"`
    CreateExternalTables  string `json:"create_external_tables"`
    ExternalLocation      string `json:"external_location"`
    AuthType              string `json:"auth_type"`
    RoleArn               string `json:"role_arn"`
    SecretKey             string `json:"secret_key"`
    PublicKey             string `json:"public_key"` // Readonly field
    PrivateKey            string `json:"private_key"`
    ClusterId             string `json:"cluster_id"`
    ClusterRegion         string `json:"cluster_region"`
    Role                  string `json:"role"`
    IsPrivateKeyEncrypted bool   `json:"is_private_key_encrypted"`
    Passphrase            string `json:"passphrase"`
    Catalog               string `json:"catalog"`
    FivetranRoleArn       string `json:"fivetran_role_arn"`
    PrefixPath            string `json:"prefix_path"`
    Region                string `json:"region"`
    StorageAccountName    string `json:"storage_account_name"`
    ContainerName         string `json:"container_name"`
    TenantId              string `json:"tenant_id"`
    ClientId              string `json:"client_id"`
    SecretValue           string `json:"secret_value"`
    WorkspaceName         string `json:"workspace_name"`
    LakehouseName         string `json:"lakehouse_name"`
}

func (dc *DestinationConfig) Request() *destinationConfigRequest {
    return &destinationConfigRequest{
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
        StorageAccountName:    dc.storageAccountName,
        ContainerName:         dc.containerName,
        TenantId:              dc.tenantId,
        ClientId:              dc.clientId,
        SecretValue:           dc.secretValue,
        WorkspaceName:         dc.workspaceName,
        LakehouseName:         dc.lakehouseName,
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


func (dc *DestinationConfig) StorageAccountName(value string) *DestinationConfig {
    dc.storageAccountName = &value
    return dc
}


func (dc *DestinationConfig) ContainerName(value string) *DestinationConfig {
    dc.containerName = &value
    return dc
}


func (dc *DestinationConfig) TenantId(value string) *DestinationConfig {
    dc.tenantId = &value
    return dc
}


func (dc *DestinationConfig) ClientId(value string) *DestinationConfig {
    dc.clientId = &value
    return dc
}


func (dc *DestinationConfig) SecretValue(value string) *DestinationConfig {
    dc.secretValue = &value
    return dc
}

func (dc *DestinationConfig) WorkspaceName(value string) *DestinationConfig {
    dc.workspaceName = &value
    return dc
}
 
func (dc *DestinationConfig) LakehouseName(value string) *DestinationConfig {
    dc.lakehouseName = &value
    return dc
}