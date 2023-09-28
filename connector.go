package fivetran

import "github.com/fivetran/go-fivetran/connectors"

func NewFunctionSecret() *connectors.FunctionSecret {
	return &connectors.FunctionSecret{}
}

func NewConnectorConfig() *connectors.ConnectorConfig {
	return &connectors.ConnectorConfig{}
}

func NewConnectorAuth() *connectors.ConnectorAuth {
	return &connectors.ConnectorAuth{}
}

func NewConnectorConfigReports() *connectors.ConnectorConfigReports {
	return &connectors.ConnectorConfigReports{}
}

func NewConnectorConfigProjectCredentials() *connectors.ConnectorConfigProjectCredentials {
	return &connectors.ConnectorConfigProjectCredentials{}
}

func NewConnectorConfigCustomTables() *connectors.ConnectorConfigCustomTables {
	return &connectors.ConnectorConfigCustomTables{}
}

func NewConnectorConfigAdobeAnalyticsConfiguration() *connectors.ConnectorConfigAdobeAnalyticsConfiguration {
	return &connectors.ConnectorConfigAdobeAnalyticsConfiguration{}
}

func NewConnectorAuthClientAccess() *connectors.ConnectorAuthClientAccess {
	return &connectors.ConnectorAuthClientAccess{}
}
