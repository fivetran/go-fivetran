package connections

type ConnectionConfigAdobeAnalyticsConfiguration struct {
	syncMode          *string
	reportSuites      []string
	elements          []string
	metrics           []string
	calculatedMetrics []string
	segments          []string
}

type connectionConfigAdobeAnalyticsConfigurationRequest struct {
	SyncMode          *string  `json:"sync_mode,omitempty"`
	ReportSuites      []string `json:"report_suites,omitempty"`
	Elements          []string `json:"elements,omitempty"`
	Metrics           []string `json:"metrics,omitempty"`
	CalculatedMetrics []string `json:"calculated_metrics,omitempty"`
	Segments          []string `json:"segments,omitempty"`
}

type ConnectionConfigAdobeAnalyticsConfigurationResponse struct {
	SyncMode          string   `json:"sync_mode"`
	ReportSuites      []string `json:"report_suites"`
	Elements          []string `json:"elements"`
	Metrics           []string `json:"metrics"`
	CalculatedMetrics []string `json:"calculated_metrics"`
	Segments          []string `json:"segments"`
}

func (c *ConnectionConfigAdobeAnalyticsConfiguration) request() *connectionConfigAdobeAnalyticsConfigurationRequest {
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

	return &connectionConfigAdobeAnalyticsConfigurationRequest{
		SyncMode:          syncMode,
		ReportSuites:      reportSuites,
		Elements:          elements,
		Metrics:           metrics,
		CalculatedMetrics: calculatedMetrics,
		Segments:          segments,
	}
}

func (c *ConnectionConfigAdobeAnalyticsConfiguration) SyncMode(value string) *ConnectionConfigAdobeAnalyticsConfiguration {
	c.syncMode = &value
	return c
}

func (c *ConnectionConfigAdobeAnalyticsConfiguration) ReportSuites(value []string) *ConnectionConfigAdobeAnalyticsConfiguration {
	c.reportSuites = value
	return c
}

func (c *ConnectionConfigAdobeAnalyticsConfiguration) Elements(value []string) *ConnectionConfigAdobeAnalyticsConfiguration {
	c.elements = value
	return c
}

func (c *ConnectionConfigAdobeAnalyticsConfiguration) Metrics(value []string) *ConnectionConfigAdobeAnalyticsConfiguration {
	c.metrics = value
	return c
}

func (c *ConnectionConfigAdobeAnalyticsConfiguration) CalculatedMetrics(value []string) *ConnectionConfigAdobeAnalyticsConfiguration {
	c.calculatedMetrics = value
	return c
}

func (c *ConnectionConfigAdobeAnalyticsConfiguration) Segments(value []string) *ConnectionConfigAdobeAnalyticsConfiguration {
	c.segments = value
	return c
}
