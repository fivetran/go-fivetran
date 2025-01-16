package fivetran

import (
	"github.com/fivetran/go-fivetran/connectors"
	"github.com/fivetran/go-fivetran/dbt"
	"github.com/fivetran/go-fivetran/destinations"
	"github.com/fivetran/go-fivetran/transformations"
	"github.com/fivetran/go-fivetran/connect_card"
	externallogging "github.com/fivetran/go-fivetran/external_logging"
	"github.com/fivetran/go-fivetran/private_link"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

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

func NewDbtProjectConfig() *dbt.DbtProjectConfig {
	return &dbt.DbtProjectConfig{}
}

func NewDbtTransformationSchedule() *dbt.DbtTransformationSchedule {
	return &dbt.DbtTransformationSchedule{}
}

func NewDestinationConfig() *destinations.DestinationConfig {
	return &destinations.DestinationConfig{}
}

func NewExternalLoggingConfig() *externallogging.ExternalLoggingConfig {
	return &externallogging.ExternalLoggingConfig{}
}

func NewTransformationConfig() *transformations.TransformationConfig {
	return &transformations.TransformationConfig{}
}

func NewTransformationSchedule() *transformations.TransformationSchedule {
	return &transformations.TransformationSchedule{}
}

func NewTransformationProjectConfig() *transformations.TransformationProjectConfig {
	return &transformations.TransformationProjectConfig{}
}

func NewConnectCardConfig() *connectcard.ConnectCardConfig {
	return &connectcard.ConnectCardConfig{}
}

func NewPrivateLinkConfig() *privatelink.PrivateLinkConfig {
	return &privatelink.PrivateLinkConfig{}
}

func Debug(value bool) {
	httputils.Debug(value)
}

func DebugAuth(value bool) {
	httputils.DebugAuth(value)
}
