package connections

type ConnectionConfigCustomTables struct {
	tableName              *string
	configType             *string
	fields                 []string
	breakdowns             []string
	actionBreakdowns       []string
	aggregation            *string
	actionReportTime       *string
	clickAttributionWindow *string
	viewAttributionWindow  *string
	prebuiltReportName     *string
}

type connectionConfigCustomTablesRequest struct {
	TableName              *string  `json:"table_name,omitempty"`
	ConfigType             *string  `json:"config_type,omitempty"`
	Fields                 []string `json:"fields,omitempty"`
	Breakdowns             []string `json:"breakdowns,omitempty"`
	ActionBreakdowns       []string `json:"action_breakdowns,omitempty"`
	Aggregation            *string  `json:"aggregation,omitempty"`
	ActionReportTime       *string  `json:"action_report_time,omitempty"`
	ClickAttributionWindow *string  `json:"click_attribution_window,omitempty"`
	ViewAttributionWindow  *string  `json:"view_attribution_window,omitempty"`
	PrebuiltReportName     *string  `json:"prebuilt_report_name,omitempty"`
}

type ConnectionConfigCustomTablesResponse struct {
	TableName              string   `json:"table_name"`
	ConfigType             string   `json:"config_type"`
	Fields                 []string `json:"fields"`
	Breakdowns             []string `json:"breakdowns"`
	ActionBreakdowns       []string `json:"action_breakdowns"`
	Aggregation            string   `json:"aggregation"`
	ActionReportTime       string   `json:"action_report_time"`
	ClickAttributionWindow string   `json:"click_attribution_window"`
	ViewAttributionWindow  string   `json:"view_attribution_window"`
	PrebuiltReportName     string   `json:"prebuilt_report_name"`
}

func (ct *ConnectionConfigCustomTables) request() *connectionConfigCustomTablesRequest {
	var tableName *string
	if ct.tableName != nil {
		tableName = ct.tableName
	}

	var configType *string
	if ct.configType != nil {
		configType = ct.configType
	}

	var aggregation *string
	if ct.aggregation != nil {
		aggregation = ct.aggregation
	}

	var actionReportTime *string
	if ct.actionReportTime != nil {
		actionReportTime = ct.actionReportTime
	}

	var clickAttributionWindow *string
	if ct.clickAttributionWindow != nil {
		clickAttributionWindow = ct.clickAttributionWindow
	}

	var viewAttributionWindow *string
	if ct.viewAttributionWindow != nil {
		viewAttributionWindow = ct.viewAttributionWindow
	}

	var prebuiltReportName *string
	if ct.prebuiltReportName != nil {
		prebuiltReportName = ct.prebuiltReportName
	}

	return &connectionConfigCustomTablesRequest{
		TableName:              tableName,
		ConfigType:             configType,
		Fields:                 ct.fields,
		Breakdowns:             ct.breakdowns,
		ActionBreakdowns:       ct.actionBreakdowns,
		Aggregation:            aggregation,
		ActionReportTime:       actionReportTime,
		ClickAttributionWindow: clickAttributionWindow,
		ViewAttributionWindow:  viewAttributionWindow,
		PrebuiltReportName:     prebuiltReportName,
	}
}

func (ct *ConnectionConfigCustomTables) TableName(value string) *ConnectionConfigCustomTables {
	ct.tableName = &value
	return ct
}

func (ct *ConnectionConfigCustomTables) ConfigType(value string) *ConnectionConfigCustomTables {
	ct.configType = &value
	return ct
}

func (ct *ConnectionConfigCustomTables) Fields(value []string) *ConnectionConfigCustomTables {
	ct.fields = value
	return ct
}

func (ct *ConnectionConfigCustomTables) Breakdowns(value []string) *ConnectionConfigCustomTables {
	ct.breakdowns = value
	return ct
}

func (ct *ConnectionConfigCustomTables) ActionBreakdowns(value []string) *ConnectionConfigCustomTables {
	ct.actionBreakdowns = value
	return ct
}

func (ct *ConnectionConfigCustomTables) Aggregation(value string) *ConnectionConfigCustomTables {
	ct.aggregation = &value
	return ct
}

func (ct *ConnectionConfigCustomTables) ActionReportTime(value string) *ConnectionConfigCustomTables {
	ct.actionReportTime = &value
	return ct
}

func (ct *ConnectionConfigCustomTables) ClickAttributionWindow(value string) *ConnectionConfigCustomTables {
	ct.clickAttributionWindow = &value
	return ct
}

func (ct *ConnectionConfigCustomTables) ViewAttributionWindow(value string) *ConnectionConfigCustomTables {
	ct.viewAttributionWindow = &value
	return ct
}

func (ct *ConnectionConfigCustomTables) PrebuiltReportName(value string) *ConnectionConfigCustomTables {
	ct.prebuiltReportName = &value
	return ct
}
