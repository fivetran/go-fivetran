package fivetran

type ConnectorConfigReports struct {
	FTable          *string   `json:"table,omitempty"`
	FConfigType     *string   `json:"config_type,omitempty"`
	FPrebuiltReport *string   `json:"prebuilt_report,omitempty"`
	FReportType     *string   `json:"report_type,omitempty"`
	FFields         *[]string `json:"fields,omitempty"`
	FDimensions     *[]string `json:"dimensions,omitempty"`
	FMetrics        *[]string `json:"metrics,omitempty"`
	FSegments       *[]string `json:"segments,omitempty"`
	FFilter         *string   `json:"filter,omitempty"`
}

func NewConnectorConfigReports() *ConnectorConfigReports {
	return &ConnectorConfigReports{}
}

func (r *ConnectorConfigReports) Table(value string) *ConnectorConfigReports {
	r.FTable = &value
	return r
}

func (r *ConnectorConfigReports) ConfigType(value string) *ConnectorConfigReports {
	r.FConfigType = &value
	return r
}

func (r *ConnectorConfigReports) PrebuiltReport(value string) *ConnectorConfigReports {
	r.FPrebuiltReport = &value
	return r
}

func (r *ConnectorConfigReports) ReportType(value string) *ConnectorConfigReports {
	r.FReportType = &value
	return r
}

func (r *ConnectorConfigReports) Fields(value []string) *ConnectorConfigReports {
	r.FFields = &value
	return r
}

func (r *ConnectorConfigReports) Dimensions(value []string) *ConnectorConfigReports {
	r.FDimensions = &value
	return r
}

func (r *ConnectorConfigReports) Metrics(value []string) *ConnectorConfigReports {
	r.FMetrics = &value
	return r
}

func (r *ConnectorConfigReports) Segments(value []string) *ConnectorConfigReports {
	r.FSegments = &value
	return r
}

func (r *ConnectorConfigReports) Filter(value string) *ConnectorConfigReports {
	r.FFilter = &value
	return r
}
