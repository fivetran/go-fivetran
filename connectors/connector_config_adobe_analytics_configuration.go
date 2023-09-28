package connectors

type ConnectorConfigAdobeAnalyticsConfiguration struct {
	syncMode          *string
	reportSuites      []string
	elements          []string
	metrics           []string
	calculatedMetrics []string
	segments          []string
}

type connectorConfigAdobeAnalyticsConfigurationRequest struct {
	SyncMode          *string  `json:"sync_mode,omitempty"`
	ReportSuites      []string `json:"report_suites,omitempty"`
	Elements          []string `json:"elements,omitempty"`
	Metrics           []string `json:"metrics,omitempty"`
	CalculatedMetrics []string `json:"calculated_metrics,omitempty"`
	Segments          []string `json:"segments,omitempty"`
}

type ConnectorConfigAdobeAnalyticsConfigurationResponse struct {
	SyncMode          string   `json:"sync_mode"`
	ReportSuites      []string `json:"report_suites"`
	Elements          []string `json:"elements"`
	Metrics           []string `json:"metrics"`
	CalculatedMetrics []string `json:"calculated_metrics"`
	Segments          []string `json:"segments"`
}

func (c *ConnectorConfigAdobeAnalyticsConfiguration) request() *connectorConfigAdobeAnalyticsConfigurationRequest {
	var syncMode *string
	if c.syncMode != nil {
		syncMode = c.syncMode
	}

	var reportSuites []string
	if c.reportSuites != nil {
		reportSuites = c.reportSuites
	}

	var elements []string
	if c.elements != nil {
		elements = c.elements
	}

	var metrics []string
	if c.metrics != nil {
		metrics = c.metrics
	}

	var calculatedMetrics []string
	if c.calculatedMetrics != nil {
		calculatedMetrics = c.calculatedMetrics
	}

	var segments []string
	if c.segments != nil {
		segments = c.segments
	}

	return &connectorConfigAdobeAnalyticsConfigurationRequest{
		SyncMode:          syncMode,
		ReportSuites:      reportSuites,
		Elements:          elements,
		Metrics:           metrics,
		CalculatedMetrics: calculatedMetrics,
		Segments:          segments,
	}
}

func (c *ConnectorConfigAdobeAnalyticsConfiguration) SyncMode(value string) *ConnectorConfigAdobeAnalyticsConfiguration {
	c.syncMode = &value
	return c
}

func (c *ConnectorConfigAdobeAnalyticsConfiguration) ReportSuites(value []string) *ConnectorConfigAdobeAnalyticsConfiguration {
	c.reportSuites = value
	return c
}

func (c *ConnectorConfigAdobeAnalyticsConfiguration) Elements(value []string) *ConnectorConfigAdobeAnalyticsConfiguration {
	c.elements = value
	return c
}

func (c *ConnectorConfigAdobeAnalyticsConfiguration) Metrics(value []string) *ConnectorConfigAdobeAnalyticsConfiguration {
	c.metrics = value
	return c
}

func (c *ConnectorConfigAdobeAnalyticsConfiguration) CalculatedMetrics(value []string) *ConnectorConfigAdobeAnalyticsConfiguration {
	c.calculatedMetrics = value
	return c
}

func (c *ConnectorConfigAdobeAnalyticsConfiguration) Segments(value []string) *ConnectorConfigAdobeAnalyticsConfiguration {
	c.segments = value
	return c
}
