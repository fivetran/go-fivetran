package fivetran

type ConnectorConfigCustomTables struct {
	FTableName              *string   `json:"table_name,omitempty"`
	FConfigType             *string   `json:"config_type,omitempty"`
	FFields                 *[]string `json:"fields,omitempty"`
	FBreakdowns             *[]string `json:"breakdowns,omitempty"`
	FActionBreakdowns       *[]string `json:"action_breakdowns,omitempty"`
	FAggregation            *string   `json:"aggregation,omitempty"`
	FActionReportTime       *string   `json:"action_report_time,omitempty"`
	FClickAttributionWindow *string   `json:"click_attribution_window,omitempty"`
	FViewAttributionWindow  *string   `json:"view_attribution_window,omitempty"`
	FPrebuiltReportName     *string   `json:"prebuilt_report_name,omitempty"`
}

func NewConnectorConfigCustomTables() *ConnectorConfigCustomTables {
	return &ConnectorConfigCustomTables{}
}

func (ct *ConnectorConfigCustomTables) TableName(value string) *ConnectorConfigCustomTables {
	ct.FTableName = &value
	return ct
}

func (ct *ConnectorConfigCustomTables) ConfigType(value string) *ConnectorConfigCustomTables {
	ct.FConfigType = &value
	return ct
}

func (ct *ConnectorConfigCustomTables) Fields(value []string) *ConnectorConfigCustomTables {
	ct.FFields = &value
	return ct
}

func (ct *ConnectorConfigCustomTables) Breakdowns(value []string) *ConnectorConfigCustomTables {
	ct.FBreakdowns = &value
	return ct
}

func (ct *ConnectorConfigCustomTables) ActionBreakdowns(value []string) *ConnectorConfigCustomTables {
	ct.FActionBreakdowns = &value
	return ct
}

func (ct *ConnectorConfigCustomTables) Aggregation(value string) *ConnectorConfigCustomTables {
	ct.FAggregation = &value
	return ct
}

func (ct *ConnectorConfigCustomTables) ActionReportTime(value string) *ConnectorConfigCustomTables {
	ct.FActionReportTime = &value
	return ct
}

func (ct *ConnectorConfigCustomTables) ClickAttributionWindow(value string) *ConnectorConfigCustomTables {
	ct.FClickAttributionWindow = &value
	return ct
}

func (ct *ConnectorConfigCustomTables) ViewAttributionWindow(value string) *ConnectorConfigCustomTables {
	ct.FViewAttributionWindow = &value
	return ct
}

func (ct *ConnectorConfigCustomTables) PrebuiltReportName(value string) *ConnectorConfigCustomTables {
	ct.FPrebuiltReportName = &value
	return ct
}
