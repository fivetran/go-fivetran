package externallogging

import "github.com/fivetran/go-fivetran/utils"

func (elc *ExternalLoggingConfig) Request() *externalLoggingConfigRequest {
	return &externalLoggingConfigRequest{
		WorkspaceId:  elc.workspaceId,
		PrimaryKey:   elc.primaryKey,
		LogGroupName: elc.logGroupName,
		RoleArn:      elc.roleArn,
		ExternalId:   elc.externalId,
		Region:       elc.region,
		ApiKey:       elc.apiKey,
		SubDomain:    elc.subDomain,
		Host:         elc.host,
		Hostname:     elc.hostname,
		Channel:      elc.channel,
		EnableSsl:    elc.enableSsl,
		Token:        elc.token,
		Port:         elc.port,
		ProjectId:    elc.projectId,
	}
}

func (cc *ExternalLoggingConfig) Merge(customConfig *map[string]interface{}) (*map[string]interface{}, error) {
	err := utils.MergeIntoMap(cc.Request(), customConfig)
	if err != nil {
		return nil, err
	}
	return customConfig, nil
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
	elc.host = &value
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

func (elc *ExternalLoggingConfig) EnableSsl(value bool) *ExternalLoggingConfig {
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

func (elc *ExternalLoggingConfig) ProjectId(value string) *ExternalLoggingConfig {
	elc.projectId = &value
	return elc
}
