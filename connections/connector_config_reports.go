package connectors

// ConnectorConfigReports builds Connector Management, Connector Config Reports.
// Ref. https://fivetran.com/docs/rest-api/connectors/config
type ConnectorConfigReports struct {
	table          *string
	configType     *string
	prebuiltReport *string
	reportType     *string
	fields         []string
	dimensions     []string
	metrics        []string
	segments       []string
	filter         *string
}

type connectorConfigReportsRequest struct {
	Table          *string  `json:"table,omitempty"`
	ConfigType     *string  `json:"config_type,omitempty"`
	PrebuiltReport *string  `json:"prebuilt_report,omitempty"`
	ReportType     *string  `json:"report_type,omitempty"`
	Fields         []string `json:"fields,omitempty"`
	Dimensions     []string `json:"dimensions,omitempty"`
	Metrics        []string `json:"metrics,omitempty"`
	Segments       []string `json:"segments,omitempty"`
	Filter         *string  `json:"filter,omitempty"`
}

type ConnectorConfigReportsResponse struct {
	Table          string   `json:"table"`
	ConfigType     string   `json:"config_type"`
	PrebuiltReport string   `json:"prebuilt_report"`
	ReportType     string   `json:"report_type"`
	Fields         []string `json:"fields"`
	Dimensions     []string `json:"dimensions"`
	Metrics        []string `json:"metrics"`
	Segments       []string `json:"segments"`
	Filter         string   `json:"filter"`
}

func (r *ConnectorConfigReports) request() *connectorConfigReportsRequest {
	var table *string
	if r.table != nil {
		table = r.table
	}

	var configType *string
	if r.configType != nil {
		configType = r.configType
	}

	var prebuiltReport *string
	if r.prebuiltReport != nil {
		prebuiltReport = r.prebuiltReport
	}

	var reportType *string
	if r.reportType != nil {
		reportType = r.reportType
	}

	var filter *string
	if r.filter != nil {
		filter = r.filter
	}

	return &connectorConfigReportsRequest{
		Table:          table,
		ConfigType:     configType,
		PrebuiltReport: prebuiltReport,
		ReportType:     reportType,
		Fields:         r.fields,
		Dimensions:     r.dimensions,
		Metrics:        r.metrics,
		Segments:       r.segments,
		Filter:         filter,
	}
}

func (r *ConnectorConfigReports) Table(value string) *ConnectorConfigReports {
	r.table = &value
	return r
}

func (r *ConnectorConfigReports) ConfigType(value string) *ConnectorConfigReports {
	r.configType = &value
	return r
}

func (r *ConnectorConfigReports) PrebuiltReport(value string) *ConnectorConfigReports {
	r.prebuiltReport = &value
	return r
}

func (r *ConnectorConfigReports) ReportType(value string) *ConnectorConfigReports {
	r.reportType = &value
	return r
}

func (r *ConnectorConfigReports) Fields(value []string) *ConnectorConfigReports {
	r.fields = value
	return r
}

func (r *ConnectorConfigReports) Dimensions(value []string) *ConnectorConfigReports {
	r.dimensions = value
	return r
}

func (r *ConnectorConfigReports) Metrics(value []string) *ConnectorConfigReports {
	r.metrics = value
	return r
}

func (r *ConnectorConfigReports) Segments(value []string) *ConnectorConfigReports {
	r.segments = value
	return r
}

func (r *ConnectorConfigReports) Filter(value string) *ConnectorConfigReports {
	r.filter = &value
	return r
}
