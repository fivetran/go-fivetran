package fivetran

import (
	"github.com/fivetran/go-fivetran/certificates"
	"github.com/fivetran/go-fivetran/connectors"
	"github.com/fivetran/go-fivetran/dbt"
	"github.com/fivetran/go-fivetran/destinations"
	externallogging "github.com/fivetran/go-fivetran/external_logging"
	"github.com/fivetran/go-fivetran/fingerprints"
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

func (c *Client) NewConnectorSync() *ConnectorSyncService {
	return &ConnectorSyncService{
		HttpService: c.NewHttpService(NewConnectorSyncRequestParams()),
	}
}

func (c *Client) NewCertificateConnectorCertificateApprove() *certificates.ConnectorCertificateApproveService {
	return &certificates.ConnectorCertificateApproveService{
		HttpService: c.NewHttpService(certificates.NewApproveConnectorCertificateRequestParams()),
	}
}

func (c *Client) NewCertificateDestinationCertificateApprove() *certificates.DestinationCertificateApproveService {
	return &certificates.DestinationCertificateApproveService{
		HttpService: c.NewHttpService(certificates.NewApproveDestinationCertificateRequestParams()),
	}
}

func (c *Client) NewCertificateConnectorFingerprintApprove() *fingerprints.ConnectorFingerprintApproveService {
	return &fingerprints.ConnectorFingerprintApproveService{
		HttpService: c.NewHttpService(fingerprints.NewApproveConnectorFingerprintsRequestParams()),
	}
}

func (c *Client) NewCertificateDestinationFingerprintApprove() *fingerprints.DestinationFingerprintApproveService {
	return &fingerprints.DestinationFingerprintApproveService{
		HttpService: c.NewHttpService(fingerprints.NewApproveDestinationFingerprintsRequestParams()),
	}
}
