package externallogging

import "github.com/fivetran/go-fivetran/common"

type ExternalLoggingConfig struct {
	workspaceId  *string
	primaryKey   *string
	logGroupName *string
	roleArn      *string
	externalId   *string
	region       *string
	apiKey       *string
	subDomain    *string
	host         *string
	hostname     *string
	channel      *string
	enableSsl    *bool
	token        *string
	port         *int
	projectId    *string
}

/* Requests */
type externalLoggingConfigRequest struct {
	WorkspaceId  *string `json:"workspace_id,omitempty"`
	PrimaryKey   *string `json:"primary_key,omitempty"`
	LogGroupName *string `json:"log_group_name,omitempty"`
	RoleArn      *string `json:"role_arn,omitempty"`
	ExternalId   *string `json:"external_id,omitempty"`
	Region       *string `json:"region,omitempty"`
	ApiKey       *string `json:"api_key,omitempty"`
	SubDomain    *string `json:"sub_domain,omitempty"`
	Host         *string `json:"host,omitempty"`
	Hostname     *string `json:"hostname,omitempty"`
	Channel      *string `json:"channel,omitempty"`
	EnableSsl    *bool   `json:"enable_ssl,omitempty"`
	Token        *string `json:"token,omitempty"`
	Port         *int    `json:"port,omitempty"`
	ProjectId    *string `json:"project_id,omitempty"`
}

type externalLoggingCreateRequest struct {
	externalLoggingCreateRequestBase
	Config any `json:"config,omitempty"`
}

type externalLoggingCustomCreateRequest struct {
	externalLoggingCreateRequestBase
	Config *map[string]interface{} `json:"config,omitempty"`
}

type externalLoggingModifyRequestBase struct {
	Enabled       *bool `json:"enabled,omitempty"`
	RunSetupTests *bool `json:"run_setup_tests,omitempty"`
}

type externalLoggingModifyRequest struct {
	externalLoggingModifyRequestBase
	Config any `json:"config,omitempty"`
}

type externalLoggingCustomModifyRequest struct {
	externalLoggingModifyRequestBase
	Config *map[string]interface{} `json:"config,omitempty"`
}

/* Responses */
type ExternalLoggingConfigResponse struct {
	WorkspaceId  string `json:"workspace_id"`
	PrimaryKey   string `json:"primary_key"`
	LogGroupName string `json:"log_group_name"`
	RoleArn      string `json:"role_arn"`
	ExternalId   string `json:"external_id"`
	Region       string `json:"region"`
	ApiKey       string `json:"api_key"`
	SubDomain    string `json:"sub_domain"`
	Host         string `json:"host"`
	Hostname     string `json:"hostname"`
	Channel      string `json:"channel"`
	EnableSsl    bool   `json:"enable_ssl"`
	Token        string `json:"token"`
	Port         int    `json:"port"`
	ProjectId    string `json:"project_id"`
}

type ExternalLoggingResponseBase struct {
	Id      string `json:"id"`
	Service string `json:"service"`
	Enabled bool   `json:"enabled"`
}

type ExternalLoggingResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ExternalLoggingResponseBase
		Config ExternalLoggingConfigResponse `json:"config"`
	} `json:"data"`
}

type ExternalLoggingCustomResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ExternalLoggingResponseBase
		Config map[string]interface{} `json:"config"`
	} `json:"data"`
}

type ExternalLoggingCustomMergedResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ExternalLoggingResponseBase
		CustomConfig map[string]interface{}        `json:"config"`
		Config       ExternalLoggingConfigResponse // no mapping here
	} `json:"data"`
}

type externalLoggingCreateRequestBase struct {
	Id      *string `json:"id,omitempty"`
	GroupId *string `json:"group_id,omitempty"`
	Service *string `json:"service,omitempty"`
	Enabled *bool   `json:"enabled,omitempty"`
}

type ExternalLoggingModifyResponseDataBased struct {
	Id      string `json:"id"`
	Service string `json:"service"`
	Enabled bool   `json:"enabled"`
}

type ExternalLoggingModifyResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ExternalLoggingModifyResponseDataBased
		Config ExternalLoggingConfigResponse `json:"config"`
	} `json:"data"`
}

type ExternalLoggingModifyCustomResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ExternalLoggingModifyResponseDataBased
		Config map[string]interface{} `json:"config"`
	} `json:"data"`
}

type ExternalLoggingModifyCustomMergedResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ExternalLoggingModifyResponseDataBased
		CustomConfig map[string]interface{}                        `json:"config"`
		Config       ExternalLoggingConfigResponse // no mapping here
	} `json:"data"`
}

type ExternalLoggingSetupTestsResponse struct {
	common.CommonResponse
	Data struct {
		SetupTests []common.SetupTestResponse `json:"setup_tests"`
	} `json:"data"`
}

type ExternalLoggingListResponse struct {
    common.CommonResponse
    Data struct {
        Items      []ExternalLoggingResponseBase `json:"items"`
        NextCursor string                   	 `json:"next_cursor"`
    } `json:"data"`
}