package connectors

// ConnectorConfigCustomTables builds Connector Management, Connector Config Custom Tables.
// Ref. https://fivetran.com/docs/rest-api/connectors/config
type ConnectorConfigCustomTables struct {
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

type connectorConfigCustomTablesRequest struct {
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

type ConnectorConfigCustomTablesResponse struct {
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

func (ct *ConnectorConfigCustomTables) request() *connectorConfigCustomTablesRequest {
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

	return &connectorConfigCustomTablesRequest{
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

func (ct *ConnectorConfigCustomTables) TableName(value string) *ConnectorConfigCustomTables {
	ct.tableName = &value
	return ct
}

func (ct *ConnectorConfigCustomTables) ConfigType(value string) *ConnectorConfigCustomTables {
	ct.configType = &value
	return ct
}

func (ct *ConnectorConfigCustomTables) Fields(value []string) *ConnectorConfigCustomTables {
	ct.fields = value
	return ct
}

func (ct *ConnectorConfigCustomTables) Breakdowns(value []string) *ConnectorConfigCustomTables {
	ct.breakdowns = value
	return ct
}

func (ct *ConnectorConfigCustomTables) ActionBreakdowns(value []string) *ConnectorConfigCustomTables {
	ct.actionBreakdowns = value
	return ct
}

func (ct *ConnectorConfigCustomTables) Aggregation(value string) *ConnectorConfigCustomTables {
	ct.aggregation = &value
	return ct
}

func (ct *ConnectorConfigCustomTables) ActionReportTime(value string) *ConnectorConfigCustomTables {
	ct.actionReportTime = &value
	return ct
}

func (ct *ConnectorConfigCustomTables) ClickAttributionWindow(value string) *ConnectorConfigCustomTables {
	ct.clickAttributionWindow = &value
	return ct
}

func (ct *ConnectorConfigCustomTables) ViewAttributionWindow(value string) *ConnectorConfigCustomTables {
	ct.viewAttributionWindow = &value
	return ct
}

func (ct *ConnectorConfigCustomTables) PrebuiltReportName(value string) *ConnectorConfigCustomTables {
	ct.prebuiltReportName = &value
	return ct
}
