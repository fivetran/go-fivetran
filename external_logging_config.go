package fivetran

// ExternalLoggingConfig builds Log Management, Log Config.
// Ref. https://fivetran.com/docs/rest-api/log-service-management#logservicesetupconfigurations
type ExternalLoggingConfig struct {
    workspaceId         *string
    primaryKey          *string
    logGroupName        *string
    roleArn             *string
    externalId          *string
    region              *string
    apiKey              *string
    subDomain           *string
    host                *string
    hostname            *string
    channel             *string
    enableSsl           *string
    token               *string
    port                *int
}

type externalLoggingConfigRequest struct {
    WorkspaceId         *string `json:"workspace_id,omitempty"`
    PrimaryKey          *string `json:"primary_key,omitempty"`
    LogGroupName        *string `json:"log_group_name,omitempty"`
    RoleArn             *string `json:"role_arn,omitempty"`
    ExternalId          *string `json:"external_id,omitempty"`
    Region              *string `json:"region,omitempty"`
    ApiKey              *string `json:"api_key,omitempty"`
    SubDomain           *string `json:"sub_domain,omitempty"`
    Host                *string `json:"host,omitempty"`
    Hostname            *string `json:"hostname,omitempty"`
    Channel             *string `json:"channel,omitempty"`
    EnableSsl           *string `json:"enable_ssl,omitempty"`
    Token               *string `json:"token,omitempty"`
    Port                *int    `json:"port,omitempty"`
}

type ExternalLoggingConfigResponse struct {
    WorkspaceId         string `json:"workspace_id"`
    PrimaryKey          string `json:"primary_key"`
    LogGroupName        string `json:"log_group_name"`
    RoleArn             string `json:"role_arn"`
    ExternalId          string `json:"external_id"`
    Region              string `json:"region"`
    ApiKey              string `json:"api_key"`
    SubDomain           string `json:"sub_domain"`
    Host                string `json:"host"`
    Hostname            string `json:"hostname"`
    Channel             string `json:"channel"`
    EnableSsl           string `json:"enable_ssl"`
    Token               string `json:"token"`
    Port                int    `json:"port"`
}

func NewExternalLoggingConfig() *ExternalLoggingConfig {
    return &ExternalLoggingConfig{}
}

func (elc *ExternalLoggingConfig) request() *externalLoggingConfigRequest {
    return &externalLoggingConfigRequest{
        WorkspaceId:         elc.workspaceId,
        PrimaryKey:          elc.primaryKey,
        LogGroupName:        elc.logGroupName,
        RoleArn:             elc.roleArn,
        ExternalId:          elc.externalId,
        Region:              elc.region,
        ApiKey:              elc.apiKey,
        SubDomain:           elc.subDomain,
        Host:                elc.host,
        Hostname:            elc.hostname,
        Channel:             elc.channel,
        EnableSsl:           elc.enableSsl,
        Token:               elc.token,
        Port:                elc.port,
    }
}

func (elc *ExternalLoggingConfig) WorkspaceId(value string) *ExternalLoggingConfig {
    elc.workspaceId = &value
    return elc
}

func (elc *ExternalLoggingConfig) PrimaryKey(value string) *ExternalLoggingConfig {
    elc.primaryKey = &value
    return elc
}

func (elc *ExternalLoggingConfig) LogGroupName(value string) *ExternalLoggingConfig {
    elc.logGroupName = &value
    return elc
}

func (elc *ExternalLoggingConfig) RoleArn(value string) *ExternalLoggingConfig {
    elc.roleArn = &value
    return elc
}

func (elc *ExternalLoggingConfig) ExternalId(value string) *ExternalLoggingConfig {
    elc.externalId = &value
    return elc
}

func (elc *ExternalLoggingConfig) Region(value string) *ExternalLoggingConfig {
    elc.region = &value
    return elc
}

func (elc *ExternalLoggingConfig) ApiKey(value string) *ExternalLoggingConfig {
    elc.apiKey = &value
    return elc
}

func (elc *ExternalLoggingConfig) SubDomain(value string) *ExternalLoggingConfig {
    elc.subDomain = &value
    return elc
}

func (elc *ExternalLoggingConfig) Host(value string) *ExternalLoggingConfig {
    elc.host  = &value
    return elc
}

func (elc *ExternalLoggingConfig) Hostname(value string) *ExternalLoggingConfig {
    elc.hostname = &value
    return elc
}

func (elc *ExternalLoggingConfig) Channel(value string) *ExternalLoggingConfig {
    elc.channel = &value
    return elc
}

func (elc *ExternalLoggingConfig) EnableSsl(value string) *ExternalLoggingConfig {
    elc.enableSsl = &value
    return elc
}

func (elc *ExternalLoggingConfig) Token(value string) *ExternalLoggingConfig {
    elc.token = &value
    return elc
}

func (elc *ExternalLoggingConfig) Port(value int) *ExternalLoggingConfig {
    elc.port = &value
    return elc
}
