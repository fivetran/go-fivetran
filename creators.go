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
		HttpService: c.NewHttpService(),
	}
}

func (c *Client) NewCertificateConnectorCertificateApprove() *certificates.ConnectorCertificateApproveService {
	return &certificates.ConnectorCertificateApproveService{
		HttpService: c.NewHttpService(),
	}
}

func (c *Client) NewCertificateDestinationCertificateApprove() *certificates.DestinationCertificateApproveService {
	return &certificates.DestinationCertificateApproveService{
		HttpService: c.NewHttpService(),
	}
}

func (c *Client) NewCertificateConnectorFingerprintApprove() *fingerprints.ConnectorFingerprintApproveService {
	return &fingerprints.ConnectorFingerprintApproveService{
		HttpService: c.NewHttpService(),
	}
}

func (c *Client) NewCertificateDestinationFingerprintApprove() *fingerprints.DestinationFingerprintApproveService {
	return &fingerprints.DestinationFingerprintApproveService{
		HttpService: c.NewHttpService(),
	}
}

func (c *Client) NewConnectorCertificateRevoke() *certificates.ConnectorCertificateRevokeService {
	return &certificates.ConnectorCertificateRevokeService{
		HttpService: c.NewHttpService(),
	}
}

func (c *Client) NewDestinationCertificateRevoke() *certificates.DestinationCertificateRevokeService {
	return &certificates.DestinationCertificateRevokeService{
		HttpService: c.NewHttpService(),
	}
}

func (c *Client) NewConnectorCertificatesList() *certificates.ConnectorCertificatesListService {
	return &certificates.ConnectorCertificatesListService{
		HttpService: c.NewHttpService(),
	}
}

func (c *Client) NewDestinationCertificatesList() *certificates.DestinationCertificatesListService {
	return &certificates.DestinationCertificatesListService{
		HttpService: c.NewHttpService(),
	}
}

func (c *Client) NewConnectorCertificateDetails() *certificates.ConnectorCertificateDetailsService {
	return &certificates.ConnectorCertificateDetailsService{
		HttpService: c.NewHttpService(),
	}
}

func (c *Client) NewDestinationCertificateDetails() *certificates.DestinationCertificateDetailsService {
	return &certificates.DestinationCertificateDetailsService{
		HttpService: c.NewHttpService(),
	}
}

func (c *Client) NewConnectorFingerprintRevoke() *fingerprints.ConnectorFingerprintRevokeService {
	return &fingerprints.ConnectorFingerprintRevokeService{
		HttpService: c.NewHttpService(),
	}
}

func (c *Client) NewDestinationFingerprintRevoke() *fingerprints.DestinationFingerprintRevokeService {
	return &fingerprints.DestinationFingerprintRevokeService{
		HttpService: c.NewHttpService(),
	}
}

func (c *Client) NewConnectorFingerprintsList() *fingerprints.ConnectorFingerprintsListService {
	return &fingerprints.ConnectorFingerprintsListService{
		HttpService: c.NewHttpService(),
	}
}

func (c *Client) NewDestinationFingerprintsList() *fingerprints.DestinationFingerprintsListService {
	return &fingerprints.DestinationFingerprintsListService{
		HttpService: c.NewHttpService(),
	}
}

func (c *Client) NewConnectorFingerprintDetails() *fingerprints.ConnectorFingerprintDetailsService {
	return &fingerprints.ConnectorFingerprintDetailsService{
		HttpService: c.NewHttpService(),
	}
}

func (c *Client) NewDestinationFingerprintDetails() *fingerprints.DestinationFingerprintDetailsService {
	return &fingerprints.DestinationFingerprintDetailsService{
		HttpService: c.NewHttpService(),
	}
}
