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

func NewConnectorSchemaConfigColumn() *connectors.ConnectorSchemaConfigColumn {
	return &connectors.ConnectorSchemaConfigColumn{}
}

func NewConnectorSchemaConfigSchema() *connectors.ConnectorSchemaConfigSchema {
	return &connectors.ConnectorSchemaConfigSchema{}
}

func NewConnectorSchemaConfigTable() *connectors.ConnectorSchemaConfigTable {
	return &connectors.ConnectorSchemaConfigTable{}
}
