package fivetran

import (
	"github.com/fivetran/go-fivetran/connections"
	"github.com/fivetran/go-fivetran/destinations"
	"github.com/fivetran/go-fivetran/transformations"
	"github.com/fivetran/go-fivetran/connect_card"
	externallogging "github.com/fivetran/go-fivetran/external_logging"
	"github.com/fivetran/go-fivetran/private_link"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

func NewFunctionSecret() *connections.FunctionSecret {
	return &connections.FunctionSecret{}
}

func NewConnectionConfig() *connections.ConnectionConfig {
	return &connections.ConnectionConfig{}
}

func NewConnectionAuth() *connections.ConnectionAuth {
	return &connections.ConnectionAuth{}
}

func NewConnectionConfigReports() *connections.ConnectionConfigReports {
	return &connections.ConnectionConfigReports{}
}

func NewConnectionConfigProjectCredentials() *connections.ConnectionConfigProjectCredentials {
	return &connections.ConnectionConfigProjectCredentials{}
}

func NewConnectionConfigCustomTables() *connections.ConnectionConfigCustomTables {
	return &connections.ConnectionConfigCustomTables{}
}

func NewConnectionConfigAdobeAnalyticsConfiguration() *connections.ConnectionConfigAdobeAnalyticsConfiguration {
	return &connections.ConnectionConfigAdobeAnalyticsConfiguration{}
}

func NewConnectionAuthClientAccess() *connections.ConnectionAuthClientAccess {
	return &connections.ConnectionAuthClientAccess{}
}

func NewConnectionSchemaConfigColumn() *connections.ConnectionSchemaConfigColumn {
	return &connections.ConnectionSchemaConfigColumn{}
}

func NewConnectionSchemaConfigSchema() *connections.ConnectionSchemaConfigSchema {
	return &connections.ConnectionSchemaConfigSchema{}
}

func NewConnectionSchemaConfigTable() *connections.ConnectionSchemaConfigTable {
	return &connections.ConnectionSchemaConfigTable{}
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
