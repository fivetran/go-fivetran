package connections

type ConnectionConfigReports struct {
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

type connectionConfigReportsRequest struct {
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

type ConnectionConfigReportsResponse struct {
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

func (r *ConnectionConfigReports) request() *connectionConfigReportsRequest {
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

	return &connectionConfigReportsRequest{
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

func (r *ConnectionConfigReports) Table(value string) *ConnectionConfigReports {
	r.table = &value
	return r
}

func (r *ConnectionConfigReports) ConfigType(value string) *ConnectionConfigReports {
	r.configType = &value
	return r
}

func (r *ConnectionConfigReports) PrebuiltReport(value string) *ConnectionConfigReports {
	r.prebuiltReport = &value
	return r
}

func (r *ConnectionConfigReports) ReportType(value string) *ConnectionConfigReports {
	r.reportType = &value
	return r
}

func (r *ConnectionConfigReports) Fields(value []string) *ConnectionConfigReports {
	r.fields = value
	return r
}

func (r *ConnectionConfigReports) Dimensions(value []string) *ConnectionConfigReports {
	r.dimensions = value
	return r
}

func (r *ConnectionConfigReports) Metrics(value []string) *ConnectionConfigReports {
	r.metrics = value
	return r
}

func (r *ConnectionConfigReports) Segments(value []string) *ConnectionConfigReports {
	r.segments = value
	return r
}

func (r *ConnectionConfigReports) Filter(value string) *ConnectionConfigReports {
	r.filter = &value
	return r
}
