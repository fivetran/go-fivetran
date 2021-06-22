package fivetran

type connectorConfig struct {
	schema                           *string
	table                            *string
	sheetID                          *string
	namedRange                       *string
	clientID                         *string
	clientSecret                     *string
	technicalAccountID               *string
	organizationID                   *string
	privateKey                       *string
	syncMode                         *string
	reportSuites                     *[]string
	elements                         *[]string
	metrics                          *[]string
	dateGranularity                  *string
	timeframeMonths                  *string
	source                           *string
	s3Bucket                         *string
	s3RoleArn                        *string
	absConnectionString              *string
	absContainerName                 *string
	ftpHost                          *string
	ftpPort                          *int
	ftpUser                          *string
	ftpPassword                      *string
	isFTPS                           *bool
	sFTPHost                         *string
	sFTPPort                         *int
	sFTPUser                         *string
	sFTPPassword                     *string
	sFTPIsKeyPair                    *bool
	advertisables                    *[]string
	reportType                       *string
	dimensions                       *[]string
	schemaPrefix                     *string
	apiKey                           *string
	externalID                       *string
	roleArn                          *string
	bucket                           *string
	prefix                           *string
	pattern                          *string
	fileType                         *string
	compression                      *string
	onError                          *string
	appendFileOption                 *string
	archivePattern                   *string
	nullSequence                     *string
	delimiter                        *string
	escapeChar                       *string
	skipBefore                       *string
	skipAfter                        *string
	projectCredentials               *[]*ConnectorConfigProjectCredentials
	authMode                         *string
	username                         *string
	password                         *string
	certificate                      *string
	selectedExports                  *[]string
	consumerGroup                    *string
	servers                          *string
	messageType                      *string
	syncType                         *string
	securityProtocol                 *string
	apps                             *[]string
	salesAccounts                    *[]string
	financeAccounts                  *[]string
	appSyncMode                      *string
	salesAccountSyncMode             *string
	financeAccountSyncMode           *string
	pemCertificate                   *string
	accessKeyID                      *string
	secretKey                        *string
	homeFolder                       *string
	syncDataLocker                   *bool
	projects                         *[]string
	function                         *string
	region                           *string
	secrets                          *string
	containerName                    *string
	connectionString                 *string
	functionApp                      *string
	functionName                     *string
	functionKey                      *string
	publicKey                        *string
	merchantID                       *string
	apiURL                           *string
	cloudStorageType                 *string
	s3ExternalID                     *string
	s3Folder                         *string
	gcsBucket                        *string
	gcsFolder                        *string
	userProfiles                     *[]string
	reportConfigurationIDs           *[]string
	enableAllDimensionCombinations   *bool
	instance                         *string
	awsRegionCode                    *string
	accounts                         *[]string
	fields                           *[]string
	breakdowns                       *[]string
	actionBreakdowns                 *[]string
	aggregation                      *string
	configType                       *string
	prebuiltReport                   *string
	actionReportTime                 *string
	clickAttributionWindow           *string
	viewAttributionWindow            *string
	customTables                     *[]*ConnectorConfigCustomTables
	pages                            *[]string
	subdomain                        *string
	host                             *string
	port                             *int
	user                             *string
	isSecure                         *string
	repositories                     *[]string
	useWebhooks                      *bool
	dimensionAttributes              *[]string
	columns                          *[]string
	networkCode                      *string
	customerID                       *string
	managerAccounts                  *[]string
	reports                          *[]*ConnectorConfigReports
	conversionWindowSize             *int
	profiles                         *[]string
	projectID                        *string
	datasetID                        *string
	bucketName                       *string
	functionTrigger                  *string
	configMethod                     *string
	queryID                          *string
	updateConfigOnEachSync           *bool
	siteURLs                         *[]string
	path                             *string
	onPremise                        *bool
	accessToken                      *string
	viewThroughAttributionWindowSize *string
	postClickAttributionWindowSize   *string
	useAPIKeys                       *string
	apiKeys                          *string
	endpoint                         *string
	identity                         *string
	apiQuota                         *int
	domainName                       *string
	resourceURL                      *string
	apiSecret                        *string
	hosts                            *[]string
	tunnelHost                       *string
	tunnelPort                       *int
	tunnelUser                       *string
	database                         *string
	datasource                       *string
	account                          *string
	role                             *string
	email                            *string
	accountID                        *string
	serverURL                        *string
	userKey                          *string
	apiVersion                       *string
	dailyAPICallLimit                *int
	timeZone                         *string
	integrationKey                   *string
	advertisers                      *[]string
	engagementAttributionWindow      *string
	conversionReportTime             *string
	domain                           *string
	updateMethod                     *string
	replicationSlot                  *string
	dataCenter                       *string
	apiToken                         *string
	subDomain                        *string
	testTableName                    *string
	shop                             *string
	organizations                    *[]string
	swipeAttributionWindow           *string
	apiAccessToken                   *string
	accountIDs                       *string
	sid                              *string
	secret                           *string
	oauthToken                       *string
	oauthTokenSecret                 *string
	consumerKey                      *string
	consumerSecret                   *string
	key                              *string
	advertisersID                    *[]string
	syncFormat                       *string
	bucketService                    *string
	userName                         *string
	reportURL                        *string
	uniqueID                         *string
}

type connectorConfigRequest struct {
	Schema                           *string                                      `json:"schema,omitempty"`
	Table                            *string                                      `json:"table,omitempty"`
	SheetID                          *string                                      `json:"sheet_id,omitempty"`
	NamedRange                       *string                                      `json:"named_range,omitempty"`
	ClientID                         *string                                      `json:"client_id,omitempty"`
	ClientSecret                     *string                                      `json:"client_secret,omitempty"`
	TechnicalAccountID               *string                                      `json:"technical_account_id,omitempty"`
	OrganizationID                   *string                                      `json:"organization_id,omitempty"`
	PrivateKey                       *string                                      `json:"private_key,omitempty"`
	SyncMode                         *string                                      `json:"sync_mode,omitempty"`
	ReportSuites                     *[]string                                    `json:"report_suites,omitempty"`
	Elements                         *[]string                                    `json:"elements,omitempty"`
	Metrics                          *[]string                                    `json:"metrics,omitempty"`
	DateGranularity                  *string                                      `json:"date_granularity,omitempty"`
	TimeframeMonths                  *string                                      `json:"timeframe_months,omitempty"`
	Source                           *string                                      `json:"source,omitempty"`
	S3Bucket                         *string                                      `json:"s3bucket,omitempty"`
	S3RoleArn                        *string                                      `json:"s3role_arn,omitempty"`
	ABSConnectionString              *string                                      `json:"abs_connection_string,omitempty"`
	ABSContainerName                 *string                                      `json:"abs_container_name,omitempty"`
	FTPHost                          *string                                      `json:"ftp_host,omitempty"`
	FTPPort                          *int                                         `json:"ftp_port,omitempty"`
	FTPUser                          *string                                      `json:"ftp_user,omitempty"`
	FTPPassword                      *string                                      `json:"ftp_password,omitempty"`
	IsFTPS                           *bool                                        `json:"is_ftps,omitempty"`
	SFTPHost                         *string                                      `json:"sftp_host,omitempty"`
	SFTPPort                         *int                                         `json:"sftp_port,omitempty"`
	SFTPUser                         *string                                      `json:"sftp_user,omitempty"`
	SFTPPassword                     *string                                      `json:"sftp_password,omitempty"`
	SFTPIsKeyPair                    *bool                                        `json:"sftp_is_key_pair,omitempty"`
	Advertisables                    *[]string                                    `json:"advertisables,omitempty"`
	ReportType                       *string                                      `json:"report_type,omitempty"`
	Dimensions                       *[]string                                    `json:"dimensions,omitempty"`
	SchemaPrefix                     *string                                      `json:"schema_prefix,omitempty"`
	APIKey                           *string                                      `json:"api_key,omitempty"`
	ExternalID                       *string                                      `json:"external_id,omitempty"`
	RoleArn                          *string                                      `json:"role_arn,omitempty"`
	Bucket                           *string                                      `json:"bucket,omitempty"`
	Prefix                           *string                                      `json:"prefix,omitempty"`
	Pattern                          *string                                      `json:"pattern,omitempty"`
	FileType                         *string                                      `json:"file_type,omitempty"`
	Compression                      *string                                      `json:"compression,omitempty"`
	OnError                          *string                                      `json:"on_error,omitempty"`
	AppendFileOption                 *string                                      `json:"append_file_option,omitempty"`
	ArchivePattern                   *string                                      `json:"archive_pattern,omitempty"`
	NullSequence                     *string                                      `json:"null_sequence,omitempty"`
	Delimiter                        *string                                      `json:"delimiter,omitempty"`
	EscapeChar                       *string                                      `json:"escape_char,omitempty"`
	SkipBefore                       *string                                      `json:"skip_before,omitempty"`
	SkipAfter                        *string                                      `json:"skip_after,omitempty"`
	ProjectCredentials               *[]*connectorConfigProjectCredentialsRequest `json:"project_credentials,omitempty"`
	AuthMode                         *string                                      `json:"auth_mode,omitempty"`
	Username                         *string                                      `json:"username,omitempty"`
	Password                         *string                                      `json:"password,omitempty"`
	Certificate                      *string                                      `json:"certificate,omitempty"`
	SelectedExports                  *[]string                                    `json:"selected_exports,omitempty"`
	ConsumerGroup                    *string                                      `json:"consumer_group,omitempty"`
	Servers                          *string                                      `json:"servers,omitempty"`
	MessageType                      *string                                      `json:"message_type,omitempty"`
	SyncType                         *string                                      `json:"sync_type,omitempty"`
	SecurityProtocol                 *string                                      `json:"security_protocol,omitempty"`
	Apps                             *[]string                                    `json:"apps,omitempty"`
	SalesAccounts                    *[]string                                    `json:"sales_accounts,omitempty"`
	FinanceAccounts                  *[]string                                    `json:"finance_accounts,omitempty"`
	AppSyncMode                      *string                                      `json:"app_sync_mode,omitempty"`
	SalesAccountSyncMode             *string                                      `json:"sales_account_sync_mode,omitempty"`
	FinanceAccountSyncMode           *string                                      `json:"finance_account_sync_mode,omitempty"`
	PEMCertificate                   *string                                      `json:"pem_certificate,omitempty"`
	AccessKeyID                      *string                                      `json:"access_key_id,omitempty"`
	SecretKey                        *string                                      `json:"secret_key,omitempty"`
	HomeFolder                       *string                                      `json:"home_folder,omitempty"`
	SyncDataLocker                   *bool                                        `json:"sync_data_locker,omitempty"`
	Projects                         *[]string                                    `json:"projects,omitempty"`
	Function                         *string                                      `json:"function,omitempty"`
	Region                           *string                                      `json:"region,omitempty"`
	Secrets                          *string                                      `json:"secrets,omitempty"`
	ContainerName                    *string                                      `json:"container_name,omitempty"`
	ConnectionString                 *string                                      `json:"connection_string,omitempty"`
	FunctionApp                      *string                                      `json:"function_app,omitempty"`
	FunctionName                     *string                                      `json:"function_name,omitempty"`
	FunctionKey                      *string                                      `json:"function_key,omitempty"`
	PublicKey                        *string                                      `json:"public_key,omitempty"`
	MerchantID                       *string                                      `json:"merchant_id,omitempty"`
	APIURL                           *string                                      `json:"api_url,omitempty"`
	CloudStorageType                 *string                                      `json:"cloud_storage_type,omitempty"`
	S3ExternalID                     *string                                      `json:"s3external_id,omitempty"`
	S3Folder                         *string                                      `json:"s3folder,omitempty"`
	GCSBucket                        *string                                      `json:"gcs_bucket,omitempty"`
	GCSFolder                        *string                                      `json:"gcs_folder,omitempty"`
	UserProfiles                     *[]string                                    `json:"user_profiles,omitempty"`
	ReportConfigurationIDs           *[]string                                    `json:"report_configuration_ids,omitempty"`
	EnableAllDimensionCombinations   *bool                                        `json:"enable_all_dimension_combinations,omitempty"`
	Instance                         *string                                      `json:"instance,omitempty"`
	AWSRegionCode                    *string                                      `json:"aws_region_code,omitempty"`
	Accounts                         *[]string                                    `json:"accounts,omitempty"`
	Fields                           *[]string                                    `json:"fields,omitempty"`
	Breakdowns                       *[]string                                    `json:"breakdowns,omitempty"`
	ActionBreakdowns                 *[]string                                    `json:"action_breakdowns,omitempty"`
	Aggregation                      *string                                      `json:"aggregation,omitempty"`
	ConfigType                       *string                                      `json:"config_type,omitempty"`
	PrebuiltReport                   *string                                      `json:"prebuilt_report,omitempty"`
	ActionReportTime                 *string                                      `json:"action_report_time,omitempty"`
	ClickAttributionWindow           *string                                      `json:"click_attribution_window,omitempty"`
	ViewAttributionWindow            *string                                      `json:"view_attribution_window,omitempty"`
	CustomTables                     *[]*connectorConfigCustomTablesRequest       `json:"custom_tables,omitempty"`
	Pages                            *[]string                                    `json:"pages,omitempty"`
	Subdomain                        *string                                      `json:"subdomain,omitempty"`
	Host                             *string                                      `json:"host,omitempty"`
	Port                             *int                                         `json:"port,omitempty"`
	User                             *string                                      `json:"user,omitempty"`
	IsSecure                         *string                                      `json:"is_secure,omitempty"`
	Repositories                     *[]string                                    `json:"repositories,omitempty"`
	UseWebhooks                      *bool                                        `json:"use_webhooks,omitempty"`
	DimensionAttributes              *[]string                                    `json:"dimension_attributes,omitempty"`
	Columns                          *[]string                                    `json:"columns,omitempty"`
	NetworkCode                      *string                                      `json:"network_code,omitempty"`
	CustomerID                       *string                                      `json:"customer_id,omitempty"`
	ManagerAccounts                  *[]string                                    `json:"manager_accounts,omitempty"`
	Reports                          *[]*connectorConfigReportsRequest            `json:"reports,omitempty"`
	ConversionWindowSize             *int                                         `json:"conversion_window_size,omitempty"`
	Profiles                         *[]string                                    `json:"profiles,omitempty"`
	ProjectID                        *string                                      `json:"project_id,omitempty"`
	DatasetID                        *string                                      `json:"dataset_id,omitempty"`
	BucketName                       *string                                      `json:"bucket_name,omitempty"`
	FunctionTrigger                  *string                                      `json:"function_trigger,omitempty"`
	ConfigMethod                     *string                                      `json:"config_method,omitempty"`
	QueryID                          *string                                      `json:"query_id,omitempty"`
	UpdateConfigOnEachSync           *bool                                        `json:"update_config_on_each_sync,omitempty"`
	SiteURLs                         *[]string                                    `json:"site_urls,omitempty"`
	Path                             *string                                      `json:"path,omitempty"`
	OnPremise                        *bool                                        `json:"on_premise,omitempty"`
	AccessToken                      *string                                      `json:"access_token,omitempty"`
	ViewThroughAttributionWindowSize *string                                      `json:"view_through_attribution_window_size,omitempty"`
	PostClickAttributionWindowSize   *string                                      `json:"post_click_attribution_window_size,omitempty"`
	UseAPIKeys                       *string                                      `json:"use_api_keys,omitempty"`
	APIKeys                          *string                                      `json:"api_keys,omitempty"`
	Endpoint                         *string                                      `json:"endpoint,omitempty"`
	Identity                         *string                                      `json:"identity,omitempty"`
	APIQuota                         *int                                         `json:"api_quota,omitempty"`
	DomainName                       *string                                      `json:"domain_name,omitempty"`
	ResourceURL                      *string                                      `json:"resource_url,omitempty"`
	APISecret                        *string                                      `json:"api_secret,omitempty"`
	Hosts                            *[]string                                    `json:"hosts,omitempty"`
	TunnelHost                       *string                                      `json:"tunnel_host,omitempty"`
	TunnelPort                       *int                                         `json:"tunnel_port,omitempty"`
	TunnelUser                       *string                                      `json:"tunnel_user,omitempty"`
	Database                         *string                                      `json:"database,omitempty"`
	Datasource                       *string                                      `json:"datasource,omitempty"`
	Account                          *string                                      `json:"account,omitempty"`
	Role                             *string                                      `json:"role,omitempty"`
	Email                            *string                                      `json:"email,omitempty"`
	AccountID                        *string                                      `json:"account_id,omitempty"`
	ServerURL                        *string                                      `json:"server_url,omitempty"`
	UserKey                          *string                                      `json:"user_key,omitempty"`
	APIVersion                       *string                                      `json:"api_version,omitempty"`
	DailyAPICallLimit                *int                                         `json:"daily_api_call_limit,omitempty"`
	TimeZone                         *string                                      `json:"time_zone,omitempty"`
	IntegrationKey                   *string                                      `json:"integration_key,omitempty"`
	Advertisers                      *[]string                                    `json:"advertisers,omitempty"`
	EngagementAttributionWindow      *string                                      `json:"engagement_attribution_window,omitempty"`
	ConversionReportTime             *string                                      `json:"conversion_report_time,omitempty"`
	Domain                           *string                                      `json:"domain,omitempty"`
	UpdateMethod                     *string                                      `json:"update_method,omitempty"`
	ReplicationSlot                  *string                                      `json:"replication_slot,omitempty"`
	DataCenter                       *string                                      `json:"data_center,omitempty"`
	APIToken                         *string                                      `json:"api_token,omitempty"`
	SubDomain                        *string                                      `json:"sub_domain,omitempty"`
	TestTableName                    *string                                      `json:"test_table_name,omitempty"`
	Shop                             *string                                      `json:"shop,omitempty"`
	Organizations                    *[]string                                    `json:"organizations,omitempty"`
	SwipeAttributionWindow           *string                                      `json:"swipe_attribution_window,omitempty"`
	APIAccessToken                   *string                                      `json:"api_access_token,omitempty"`
	AccountIDs                       *string                                      `json:"account_ids,omitempty"`
	SID                              *string                                      `json:"sid,omitempty"`
	Secret                           *string                                      `json:"secret,omitempty"`
	OauthToken                       *string                                      `json:"oauth_token,omitempty"`
	OauthTokenSecret                 *string                                      `json:"oauth_token_secret,omitempty"`
	ConsumerKey                      *string                                      `json:"consumer_key,omitempty"`
	ConsumerSecret                   *string                                      `json:"consumer_secret,omitempty"`
	Key                              *string                                      `json:"key,omitempty"`
	AdvertisersID                    *[]string                                    `json:"advertisers_id,omitempty"`
	SyncFormat                       *string                                      `json:"sync_format,omitempty"`
	BucketService                    *string                                      `json:"bucket_service,omitempty"`
	UserName                         *string                                      `json:"user_name,omitempty"`
	ReportURL                        *string                                      `json:"report_url,omitempty"`
	UniqueID                         *string                                      `json:"unique_id,omitempty"`
}

type ConnectorConfigResponse struct {
	Schema                           string                                      `json:"schema"`
	Table                            string                                      `json:"table"`
	SheetID                          string                                      `json:"sheet_id"`
	NamedRange                       string                                      `json:"named_range"`
	ClientID                         string                                      `json:"client_id"`
	ClientSecret                     string                                      `json:"client_secret"`
	TechnicalAccountID               string                                      `json:"technical_account_id"`
	OrganizationID                   string                                      `json:"organization_id"`
	PrivateKey                       string                                      `json:"private_key"`
	SyncMode                         string                                      `json:"sync_mode"`
	ReportSuites                     []string                                    `json:"report_suites"`
	Elements                         []string                                    `json:"elements"`
	Metrics                          []string                                    `json:"metrics"`
	DateGranularity                  string                                      `json:"date_granularity"`
	TimeframeMonths                  string                                      `json:"timeframe_months"`
	Source                           string                                      `json:"source"`
	S3Bucket                         string                                      `json:"s3bucket"`
	S3RoleArn                        string                                      `json:"s3role_arn"`
	ABSConnectionString              string                                      `json:"abs_connection_string"`
	ABSContainerName                 string                                      `json:"abs_container_name"`
	FTPHost                          string                                      `json:"ftp_host"`
	FTPPort                          int                                         `json:"ftp_port"`
	FTPUser                          string                                      `json:"ftp_user"`
	FTPPassword                      string                                      `json:"ftp_password"`
	IsFTPS                           bool                                        `json:"is_ftps"`
	SFTPHost                         string                                      `json:"sftp_host"`
	SFTPPort                         int                                         `json:"sftp_port"`
	SFTPUser                         string                                      `json:"sftp_user"`
	SFTPPassword                     string                                      `json:"sftp_password"`
	SFTPIsKeyPair                    bool                                        `json:"sftp_is_key_pair"`
	Advertisables                    []string                                    `json:"advertisables"`
	ReportType                       string                                      `json:"report_type"`
	Dimensions                       []string                                    `json:"dimensions"`
	SchemaPrefix                     string                                      `json:"schema_prefix"`
	APIKey                           string                                      `json:"api_key"`
	ExternalID                       string                                      `json:"external_id"`
	RoleArn                          string                                      `json:"role_arn"`
	Bucket                           string                                      `json:"bucket"`
	Prefix                           string                                      `json:"prefix"`
	Pattern                          string                                      `json:"pattern"`
	FileType                         string                                      `json:"file_type"`
	Compression                      string                                      `json:"compression"`
	OnError                          string                                      `json:"on_error"`
	AppendFileOption                 string                                      `json:"append_file_option"`
	ArchivePattern                   string                                      `json:"archive_pattern"`
	NullSequence                     string                                      `json:"null_sequence"`
	Delimiter                        string                                      `json:"delimiter"`
	EscapeChar                       string                                      `json:"escape_char"`
	SkipBefore                       string                                      `json:"skip_before"`
	SkipAfter                        string                                      `json:"skip_after"`
	ProjectCredentials               []ConnectorConfigProjectCredentialsResponse `json:"project_credentials"`
	AuthMode                         string                                      `json:"auth_mode"`
	Username                         string                                      `json:"username"`
	Password                         string                                      `json:"password"`
	Certificate                      string                                      `json:"certificate"`
	SelectedExports                  []string                                    `json:"selected_exports"`
	ConsumerGroup                    string                                      `json:"consumer_group"`
	Servers                          string                                      `json:"servers"`
	MessageType                      string                                      `json:"message_type"`
	SyncType                         string                                      `json:"sync_type"`
	SecurityProtocol                 string                                      `json:"security_protocol"`
	Apps                             []string                                    `json:"apps"`
	SalesAccounts                    []string                                    `json:"sales_accounts"`
	FinanceAccounts                  []string                                    `json:"finance_accounts"`
	AppSyncMode                      string                                      `json:"app_sync_mode"`
	SalesAccountSyncMode             string                                      `json:"sales_account_sync_mode"`
	FinanceAccountSyncMode           string                                      `json:"finance_account_sync_mode"`
	PEMCertificate                   string                                      `json:"pem_certificate"`
	AccessKeyID                      string                                      `json:"access_key_id"`
	SecretKey                        string                                      `json:"secret_key"`
	HomeFolder                       string                                      `json:"home_folder"`
	SyncDataLocker                   bool                                        `json:"sync_data_locker"`
	Projects                         []string                                    `json:"projects"`
	Function                         string                                      `json:"function"`
	Region                           string                                      `json:"region"`
	Secrets                          string                                      `json:"secrets"`
	ContainerName                    string                                      `json:"container_name"`
	ConnectionString                 string                                      `json:"connection_string"`
	FunctionApp                      string                                      `json:"function_app"`
	FunctionName                     string                                      `json:"function_name"`
	FunctionKey                      string                                      `json:"function_key"`
	PublicKey                        string                                      `json:"public_key"`
	MerchantID                       string                                      `json:"merchant_id"`
	APIURL                           string                                      `json:"api_url"`
	CloudStorageType                 string                                      `json:"cloud_storage_type"`
	S3ExternalID                     string                                      `json:"s3external_id"`
	S3Folder                         string                                      `json:"s3folder"`
	GCSBucket                        string                                      `json:"gcs_bucket"`
	GCSFolder                        string                                      `json:"gcs_folder"`
	UserProfiles                     []string                                    `json:"user_profiles"`
	ReportConfigurationIDs           []string                                    `json:"report_configuration_ids"`
	EnableAllDimensionCombinations   bool                                        `json:"enable_all_dimension_combinations"`
	Instance                         string                                      `json:"instance"`
	AWSRegionCode                    string                                      `json:"aws_region_code"`
	Accounts                         []string                                    `json:"accounts"`
	Fields                           []string                                    `json:"fields"`
	Breakdowns                       []string                                    `json:"breakdowns"`
	ActionBreakdowns                 []string                                    `json:"action_breakdowns"`
	Aggregation                      string                                      `json:"aggregation"`
	ConfigType                       string                                      `json:"config_type"`
	PrebuiltReport                   string                                      `json:"prebuilt_report"`
	ActionReportTime                 string                                      `json:"action_report_time"`
	ClickAttributionWindow           string                                      `json:"click_attribution_window"`
	ViewAttributionWindow            string                                      `json:"view_attribution_window"`
	CustomTables                     []ConnectorConfigCustomTablesResponse       `json:"custom_tables"`
	Pages                            []string                                    `json:"pages"`
	Subdomain                        string                                      `json:"subdomain"`
	Host                             string                                      `json:"host"`
	Port                             string                                      `json:"port"`
	User                             string                                      `json:"user"`
	IsSecure                         string                                      `json:"is_secure"`
	Repositories                     []string                                    `json:"repositories"`
	UseWebhooks                      bool                                        `json:"use_webhooks"`
	DimensionAttributes              []string                                    `json:"dimension_attributes"`
	Columns                          []string                                    `json:"columns"`
	NetworkCode                      string                                      `json:"network_code"`
	CustomerID                       string                                      `json:"customer_id"`
	ManagerAccounts                  []string                                    `json:"manager_accounts"`
	Reports                          []ConnectorConfigReportsResponse            `json:"reports"`
	ConversionWindowSize             int                                         `json:"conversion_window_size"`
	Profiles                         []string                                    `json:"profiles"`
	ProjectID                        string                                      `json:"project_id"`
	DatasetID                        string                                      `json:"dataset_id"`
	BucketName                       string                                      `json:"bucket_name"`
	FunctionTrigger                  string                                      `json:"function_trigger"`
	ConfigMethod                     string                                      `json:"config_method"`
	QueryID                          string                                      `json:"query_id"`
	UpdateConfigOnEachSync           bool                                        `json:"update_config_on_each_sync"`
	SiteURLs                         []string                                    `json:"site_urls"`
	Path                             string                                      `json:"path"`
	OnPremise                        bool                                        `json:"on_premise"`
	AccessToken                      string                                      `json:"access_token"`
	ViewThroughAttributionWindowSize string                                      `json:"view_through_attribution_window_size"`
	PostClickAttributionWindowSize   string                                      `json:"post_click_attribution_window_size"`
	UseAPIKeys                       string                                      `json:"use_api_keys"`
	APIKeys                          string                                      `json:"api_keys"`
	Endpoint                         string                                      `json:"endpoint"`
	Identity                         string                                      `json:"identity"`
	APIQuota                         int                                         `json:"api_quota"`
	DomainName                       string                                      `json:"domain_name"`
	ResourceURL                      string                                      `json:"resource_url"`
	APISecret                        string                                      `json:"api_secret"`
	Hosts                            []string                                    `json:"hosts"`
	TunnelHost                       string                                      `json:"tunnel_host"`
	TunnelPort                       int                                         `json:"tunnel_port"`
	TunnelUser                       string                                      `json:"tunnel_user"`
	Database                         string                                      `json:"database"`
	Datasource                       string                                      `json:"datasource"`
	Account                          string                                      `json:"account"`
	Role                             string                                      `json:"role"`
	Email                            string                                      `json:"email"`
	AccountID                        string                                      `json:"account_id"`
	ServerURL                        string                                      `json:"server_url"`
	UserKey                          string                                      `json:"user_key"`
	APIVersion                       string                                      `json:"api_version"`
	DailyAPICallLimit                int                                         `json:"daily_api_call_limit"`
	TimeZone                         string                                      `json:"time_zone"`
	IntegrationKey                   string                                      `json:"integration_key"`
	Advertisers                      []string                                    `json:"advertisers"`
	EngagementAttributionWindow      string                                      `json:"engagement_attribution_window"`
	ConversionReportTime             string                                      `json:"conversion_report_time"`
	Domain                           string                                      `json:"domain"`
	UpdateMethod                     string                                      `json:"update_method"`
	ReplicationSlot                  string                                      `json:"replication_slot"`
	DataCenter                       string                                      `json:"data_center"`
	APIToken                         string                                      `json:"api_token"`
	SubDomain                        string                                      `json:"sub_domain"`
	TestTableName                    string                                      `json:"test_table_name"`
	Shop                             string                                      `json:"shop"`
	Organizations                    []string                                    `json:"organizations"`
	SwipeAttributionWindow           string                                      `json:"swipe_attribution_window"`
	APIAccessToken                   string                                      `json:"api_access_token"`
	AccountIDs                       string                                      `json:"account_ids"`
	SID                              string                                      `json:"sid"`
	Secret                           string                                      `json:"secret"`
	OauthToken                       string                                      `json:"oauth_token"`
	OauthTokenSecret                 string                                      `json:"oauth_token_secret"`
	ConsumerKey                      string                                      `json:"consumer_key"`
	ConsumerSecret                   string                                      `json:"consumer_secret"`
	Key                              string                                      `json:"key"`
	AdvertisersID                    []string                                    `json:"advertisers_id"`
	SyncFormat                       string                                      `json:"sync_format"`
	BucketService                    string                                      `json:"bucket_service"`
	UserName                         string                                      `json:"user_name"`
	ReportURL                        string                                      `json:"report_url"`
	UniqueID                         string                                      `json:"unique_id"`
	LatestVersion                    string                                      `json:"latest_version"`
	AuthorizationMethod              string                                      `json:"authorization_method"`
	ServiceVersion                   string                                      `json:"service_version"`
	LastSyncedChangesUtc             string                                      `json:"last_synced_changes__utc_"`
}

func NewConnectorConfig() *connectorConfig {
	return &connectorConfig{}
}

func (cc *connectorConfig) request() *connectorConfigRequest {
	var projectCredentials []*connectorConfigProjectCredentialsRequest
	var projectCredentialsP *[]*connectorConfigProjectCredentialsRequest
	if cc.projectCredentials != nil {
		for _, pc := range *cc.projectCredentials {
			projectCredentials = append(projectCredentials, pc.request())
		}
	}
	if len(projectCredentials) > 0 {
		projectCredentialsP = &projectCredentials
	}

	var customTables []*connectorConfigCustomTablesRequest
	var customTablesP *[]*connectorConfigCustomTablesRequest
	if cc.customTables != nil {
		for _, ct := range *cc.customTables {
			customTables = append(customTables, ct.request())
		}
	}
	if len(customTables) > 0 {
		customTablesP = &customTables
	}

	var reports []*connectorConfigReportsRequest
	var reportsP *[]*connectorConfigReportsRequest
	if cc.reports != nil {
		for _, r := range *cc.reports {
			reports = append(reports, r.request())
		}
	}
	if len(reports) > 0 {
		reportsP = &reports
	}

	return &connectorConfigRequest{
		Schema:                           cc.schema,
		Table:                            cc.table,
		SheetID:                          cc.sheetID,
		NamedRange:                       cc.namedRange,
		ClientID:                         cc.clientID,
		ClientSecret:                     cc.clientSecret,
		TechnicalAccountID:               cc.technicalAccountID,
		OrganizationID:                   cc.organizationID,
		PrivateKey:                       cc.privateKey,
		SyncMode:                         cc.syncMode,
		ReportSuites:                     cc.reportSuites,
		Elements:                         cc.elements,
		Metrics:                          cc.metrics,
		DateGranularity:                  cc.dateGranularity,
		TimeframeMonths:                  cc.timeframeMonths,
		Source:                           cc.source,
		S3Bucket:                         cc.s3Bucket,
		S3RoleArn:                        cc.s3RoleArn,
		ABSConnectionString:              cc.absConnectionString,
		ABSContainerName:                 cc.absContainerName,
		FTPHost:                          cc.ftpHost,
		FTPPort:                          cc.ftpPort,
		FTPUser:                          cc.ftpUser,
		FTPPassword:                      cc.ftpPassword,
		IsFTPS:                           cc.isFTPS,
		SFTPHost:                         cc.sFTPHost,
		SFTPPort:                         cc.sFTPPort,
		SFTPUser:                         cc.sFTPUser,
		SFTPPassword:                     cc.sFTPPassword,
		SFTPIsKeyPair:                    cc.sFTPIsKeyPair,
		Advertisables:                    cc.advertisables,
		ReportType:                       cc.reportType,
		Dimensions:                       cc.dimensions,
		SchemaPrefix:                     cc.schemaPrefix,
		APIKey:                           cc.apiKey,
		ExternalID:                       cc.externalID,
		RoleArn:                          cc.roleArn,
		Bucket:                           cc.bucket,
		Prefix:                           cc.prefix,
		Pattern:                          cc.pattern,
		FileType:                         cc.fileType,
		Compression:                      cc.compression,
		OnError:                          cc.onError,
		AppendFileOption:                 cc.appendFileOption,
		ArchivePattern:                   cc.archivePattern,
		NullSequence:                     cc.nullSequence,
		Delimiter:                        cc.delimiter,
		EscapeChar:                       cc.escapeChar,
		SkipBefore:                       cc.skipBefore,
		SkipAfter:                        cc.skipAfter,
		ProjectCredentials:               projectCredentialsP,
		AuthMode:                         cc.authMode,
		Username:                         cc.username,
		Password:                         cc.password,
		Certificate:                      cc.certificate,
		SelectedExports:                  cc.selectedExports,
		ConsumerGroup:                    cc.consumerGroup,
		Servers:                          cc.servers,
		MessageType:                      cc.messageType,
		SyncType:                         cc.syncType,
		SecurityProtocol:                 cc.securityProtocol,
		Apps:                             cc.apps,
		SalesAccounts:                    cc.salesAccounts,
		FinanceAccounts:                  cc.financeAccounts,
		AppSyncMode:                      cc.appSyncMode,
		SalesAccountSyncMode:             cc.salesAccountSyncMode,
		FinanceAccountSyncMode:           cc.financeAccountSyncMode,
		PEMCertificate:                   cc.pemCertificate,
		AccessKeyID:                      cc.accessKeyID,
		SecretKey:                        cc.secretKey,
		HomeFolder:                       cc.homeFolder,
		SyncDataLocker:                   cc.syncDataLocker,
		Projects:                         cc.projects,
		Function:                         cc.function,
		Region:                           cc.region,
		Secrets:                          cc.secrets,
		ContainerName:                    cc.containerName,
		ConnectionString:                 cc.connectionString,
		FunctionApp:                      cc.functionApp,
		FunctionName:                     cc.functionName,
		FunctionKey:                      cc.functionKey,
		PublicKey:                        cc.publicKey,
		MerchantID:                       cc.merchantID,
		APIURL:                           cc.apiURL,
		CloudStorageType:                 cc.cloudStorageType,
		S3ExternalID:                     cc.s3ExternalID,
		S3Folder:                         cc.s3Folder,
		GCSBucket:                        cc.gcsBucket,
		GCSFolder:                        cc.gcsFolder,
		UserProfiles:                     cc.userProfiles,
		ReportConfigurationIDs:           cc.reportConfigurationIDs,
		EnableAllDimensionCombinations:   cc.enableAllDimensionCombinations,
		Instance:                         cc.instance,
		AWSRegionCode:                    cc.awsRegionCode,
		Accounts:                         cc.accounts,
		Fields:                           cc.fields,
		Breakdowns:                       cc.breakdowns,
		ActionBreakdowns:                 cc.actionBreakdowns,
		Aggregation:                      cc.aggregation,
		ConfigType:                       cc.configType,
		PrebuiltReport:                   cc.prebuiltReport,
		ActionReportTime:                 cc.actionReportTime,
		ClickAttributionWindow:           cc.clickAttributionWindow,
		ViewAttributionWindow:            cc.viewAttributionWindow,
		CustomTables:                     customTablesP,
		Pages:                            cc.pages,
		Subdomain:                        cc.subdomain,
		Host:                             cc.host,
		Port:                             cc.port,
		User:                             cc.user,
		IsSecure:                         cc.isSecure,
		Repositories:                     cc.repositories,
		UseWebhooks:                      cc.useWebhooks,
		DimensionAttributes:              cc.dimensionAttributes,
		Columns:                          cc.columns,
		NetworkCode:                      cc.networkCode,
		CustomerID:                       cc.customerID,
		ManagerAccounts:                  cc.managerAccounts,
		Reports:                          reportsP,
		ConversionWindowSize:             cc.conversionWindowSize,
		Profiles:                         cc.profiles,
		ProjectID:                        cc.projectID,
		DatasetID:                        cc.datasetID,
		BucketName:                       cc.bucketName,
		FunctionTrigger:                  cc.functionTrigger,
		ConfigMethod:                     cc.configMethod,
		QueryID:                          cc.queryID,
		UpdateConfigOnEachSync:           cc.updateConfigOnEachSync,
		SiteURLs:                         cc.siteURLs,
		Path:                             cc.path,
		OnPremise:                        cc.onPremise,
		AccessToken:                      cc.accessToken,
		ViewThroughAttributionWindowSize: cc.viewThroughAttributionWindowSize,
		PostClickAttributionWindowSize:   cc.postClickAttributionWindowSize,
		UseAPIKeys:                       cc.useAPIKeys,
		APIKeys:                          cc.apiKeys,
		Endpoint:                         cc.endpoint,
		Identity:                         cc.identity,
		APIQuota:                         cc.apiQuota,
		DomainName:                       cc.domainName,
		ResourceURL:                      cc.resourceURL,
		APISecret:                        cc.apiSecret,
		Hosts:                            cc.hosts,
		TunnelHost:                       cc.tunnelHost,
		TunnelPort:                       cc.tunnelPort,
		TunnelUser:                       cc.tunnelUser,
		Database:                         cc.database,
		Datasource:                       cc.datasource,
		Account:                          cc.account,
		Role:                             cc.role,
		Email:                            cc.email,
		AccountID:                        cc.accountID,
		ServerURL:                        cc.serverURL,
		UserKey:                          cc.userKey,
		APIVersion:                       cc.apiVersion,
		DailyAPICallLimit:                cc.dailyAPICallLimit,
		TimeZone:                         cc.timeZone,
		IntegrationKey:                   cc.integrationKey,
		Advertisers:                      cc.advertisers,
		EngagementAttributionWindow:      cc.engagementAttributionWindow,
		ConversionReportTime:             cc.conversionReportTime,
		Domain:                           cc.domain,
		UpdateMethod:                     cc.updateMethod,
		ReplicationSlot:                  cc.replicationSlot,
		DataCenter:                       cc.dataCenter,
		APIToken:                         cc.apiToken,
		SubDomain:                        cc.subDomain,
		TestTableName:                    cc.testTableName,
		Shop:                             cc.shop,
		Organizations:                    cc.organizations,
		SwipeAttributionWindow:           cc.swipeAttributionWindow,
		APIAccessToken:                   cc.apiAccessToken,
		AccountIDs:                       cc.accountIDs,
		SID:                              cc.sid,
		Secret:                           cc.secret,
		OauthToken:                       cc.oauthToken,
		OauthTokenSecret:                 cc.oauthTokenSecret,
		ConsumerKey:                      cc.consumerKey,
		ConsumerSecret:                   cc.consumerSecret,
		Key:                              cc.key,
		AdvertisersID:                    cc.advertisersID,
		SyncFormat:                       cc.syncFormat,
		BucketService:                    cc.bucketService,
		UserName:                         cc.userName,
		ReportURL:                        cc.reportURL,
		UniqueID:                         cc.uniqueID,
	}
}

func (cc *connectorConfig) Schema(value string) *connectorConfig {
	cc.schema = &value
	return cc
}

func (cc *connectorConfig) Table(value string) *connectorConfig {
	cc.table = &value
	return cc
}

func (cc *connectorConfig) SheetID(value string) *connectorConfig {
	cc.sheetID = &value
	return cc
}

func (cc *connectorConfig) NamedRange(value string) *connectorConfig {
	cc.namedRange = &value
	return cc
}

func (cc *connectorConfig) ClientID(value string) *connectorConfig {
	cc.clientID = &value
	return cc
}

func (cc *connectorConfig) ClientSecret(value string) *connectorConfig {
	cc.clientSecret = &value
	return cc
}

func (cc *connectorConfig) TechnicalAccountID(value string) *connectorConfig {
	cc.technicalAccountID = &value
	return cc
}

func (cc *connectorConfig) OrganizationID(value string) *connectorConfig {
	cc.organizationID = &value
	return cc
}

func (cc *connectorConfig) PrivateKey(value string) *connectorConfig {
	cc.privateKey = &value
	return cc
}

func (cc *connectorConfig) SyncMode(value string) *connectorConfig {
	cc.syncMode = &value
	return cc
}

func (cc *connectorConfig) ReportSuites(value []string) *connectorConfig {
	cc.reportSuites = &value
	return cc
}

func (cc *connectorConfig) Elements(value []string) *connectorConfig {
	cc.elements = &value
	return cc
}

func (cc *connectorConfig) Metrics(value []string) *connectorConfig {
	cc.metrics = &value
	return cc
}

func (cc *connectorConfig) DateGranularity(value string) *connectorConfig {
	cc.dateGranularity = &value
	return cc
}

func (cc *connectorConfig) TimeframeMonths(value string) *connectorConfig {
	cc.timeframeMonths = &value
	return cc
}

func (cc *connectorConfig) Source(value string) *connectorConfig {
	cc.source = &value
	return cc
}

func (cc *connectorConfig) S3Bucket(value string) *connectorConfig {
	cc.s3Bucket = &value
	return cc
}

func (cc *connectorConfig) S3RoleArn(value string) *connectorConfig {
	cc.s3RoleArn = &value
	return cc
}

func (cc *connectorConfig) ABSConnectionString(value string) *connectorConfig {
	cc.absConnectionString = &value
	return cc
}

func (cc *connectorConfig) ABSContainerName(value string) *connectorConfig {
	cc.absContainerName = &value
	return cc
}

func (cc *connectorConfig) FTPHost(value string) *connectorConfig {
	cc.ftpHost = &value
	return cc
}

func (cc *connectorConfig) FTPPort(value int) *connectorConfig {
	cc.ftpPort = &value
	return cc
}

func (cc *connectorConfig) FTPUser(value string) *connectorConfig {
	cc.ftpUser = &value
	return cc
}

func (cc *connectorConfig) FTPPassword(value string) *connectorConfig {
	cc.ftpPassword = &value
	return cc
}

func (cc *connectorConfig) IsFTPS(value bool) *connectorConfig {
	cc.isFTPS = &value
	return cc
}

func (cc *connectorConfig) SFTPHost(value string) *connectorConfig {
	cc.sFTPHost = &value
	return cc
}

func (cc *connectorConfig) SFTPPort(value int) *connectorConfig {
	cc.sFTPPort = &value
	return cc
}

func (cc *connectorConfig) SFTPUser(value string) *connectorConfig {
	cc.sFTPUser = &value
	return cc
}

func (cc *connectorConfig) SFTPPassword(value string) *connectorConfig {
	cc.sFTPPassword = &value
	return cc
}

func (cc *connectorConfig) SFTPIsKeyPair(value bool) *connectorConfig {
	cc.sFTPIsKeyPair = &value
	return cc
}

func (cc *connectorConfig) Advertisables(value []string) *connectorConfig {
	cc.advertisables = &value
	return cc
}

func (cc *connectorConfig) ReportType(value string) *connectorConfig {
	cc.reportType = &value
	return cc
}

func (cc *connectorConfig) Dimensions(value []string) *connectorConfig {
	cc.dimensions = &value
	return cc
}

func (cc *connectorConfig) SchemaPrefix(value string) *connectorConfig {
	cc.schemaPrefix = &value
	return cc
}

func (cc *connectorConfig) APIKey(value string) *connectorConfig {
	cc.apiKey = &value
	return cc
}

func (cc *connectorConfig) ExternalID(value string) *connectorConfig {
	cc.externalID = &value
	return cc
}

func (cc *connectorConfig) RoleArn(value string) *connectorConfig {
	cc.roleArn = &value
	return cc
}

func (cc *connectorConfig) Bucket(value string) *connectorConfig {
	cc.bucket = &value
	return cc
}

func (cc *connectorConfig) Prefix(value string) *connectorConfig {
	cc.prefix = &value
	return cc
}

func (cc *connectorConfig) Pattern(value string) *connectorConfig {
	cc.pattern = &value
	return cc
}

func (cc *connectorConfig) FileType(value string) *connectorConfig {
	cc.fileType = &value
	return cc
}

func (cc *connectorConfig) Compression(value string) *connectorConfig {
	cc.compression = &value
	return cc
}

func (cc *connectorConfig) OnError(value string) *connectorConfig {
	cc.onError = &value
	return cc
}

func (cc *connectorConfig) AppendFileOption(value string) *connectorConfig {
	cc.appendFileOption = &value
	return cc
}

func (cc *connectorConfig) ArchivePattern(value string) *connectorConfig {
	cc.archivePattern = &value
	return cc
}

func (cc *connectorConfig) NullSequence(value string) *connectorConfig {
	cc.nullSequence = &value
	return cc
}

func (cc *connectorConfig) Delimiter(value string) *connectorConfig {
	cc.delimiter = &value
	return cc
}

func (cc *connectorConfig) EscapeChar(value string) *connectorConfig {
	cc.escapeChar = &value
	return cc
}

func (cc *connectorConfig) SkipBefore(value string) *connectorConfig {
	cc.skipBefore = &value
	return cc
}

func (cc *connectorConfig) SkipAfter(value string) *connectorConfig {
	cc.skipAfter = &value
	return cc
}

func (cc *connectorConfig) ProjectCredentials(value *[]*ConnectorConfigProjectCredentials) *connectorConfig {
	cc.projectCredentials = value
	return cc
}

func (cc *connectorConfig) AuthMode(value string) *connectorConfig {
	cc.authMode = &value
	return cc
}

func (cc *connectorConfig) Username(value string) *connectorConfig {
	cc.username = &value
	return cc
}

func (cc *connectorConfig) Password(value string) *connectorConfig {
	cc.password = &value
	return cc
}

func (cc *connectorConfig) Certificate(value string) *connectorConfig {
	cc.certificate = &value
	return cc
}

func (cc *connectorConfig) SelectedExports(value []string) *connectorConfig {
	cc.selectedExports = &value
	return cc
}

func (cc *connectorConfig) ConsumerGroup(value string) *connectorConfig {
	cc.consumerGroup = &value
	return cc
}

func (cc *connectorConfig) Servers(value string) *connectorConfig {
	cc.servers = &value
	return cc
}

func (cc *connectorConfig) MessageType(value string) *connectorConfig {
	cc.messageType = &value
	return cc
}

func (cc *connectorConfig) SyncType(value string) *connectorConfig {
	cc.syncType = &value
	return cc
}

func (cc *connectorConfig) SecurityProtocol(value string) *connectorConfig {
	cc.securityProtocol = &value
	return cc
}

func (cc *connectorConfig) Apps(value []string) *connectorConfig {
	cc.apps = &value
	return cc
}

func (cc *connectorConfig) SalesAccounts(value []string) *connectorConfig {
	cc.salesAccounts = &value
	return cc
}

func (cc *connectorConfig) FinanceAccounts(value []string) *connectorConfig {
	cc.financeAccounts = &value
	return cc
}

func (cc *connectorConfig) AppSyncMode(value string) *connectorConfig {
	cc.appSyncMode = &value
	return cc
}

func (cc *connectorConfig) SalesAccountSyncMode(value string) *connectorConfig {
	cc.salesAccountSyncMode = &value
	return cc
}

func (cc *connectorConfig) FinanceAccountSyncMode(value string) *connectorConfig {
	cc.financeAccountSyncMode = &value
	return cc
}

func (cc *connectorConfig) PEMCertificate(value string) *connectorConfig {
	cc.pemCertificate = &value
	return cc
}

func (cc *connectorConfig) AccessKeyID(value string) *connectorConfig {
	cc.accessKeyID = &value
	return cc
}

func (cc *connectorConfig) SecretKey(value string) *connectorConfig {
	cc.secretKey = &value
	return cc
}

func (cc *connectorConfig) HomeFolder(value string) *connectorConfig {
	cc.homeFolder = &value
	return cc
}

func (cc *connectorConfig) SyncDataLocker(value bool) *connectorConfig {
	cc.syncDataLocker = &value
	return cc
}

func (cc *connectorConfig) Projects(value []string) *connectorConfig {
	cc.projects = &value
	return cc
}

func (cc *connectorConfig) Function(value string) *connectorConfig {
	cc.function = &value
	return cc
}

func (cc *connectorConfig) Region(value string) *connectorConfig {
	cc.region = &value
	return cc
}

func (cc *connectorConfig) Secrets(value string) *connectorConfig {
	cc.secrets = &value
	return cc
}

func (cc *connectorConfig) ContainerName(value string) *connectorConfig {
	cc.containerName = &value
	return cc
}

func (cc *connectorConfig) ConnectionString(value string) *connectorConfig {
	cc.connectionString = &value
	return cc
}

func (cc *connectorConfig) FunctionApp(value string) *connectorConfig {
	cc.functionApp = &value
	return cc
}

func (cc *connectorConfig) FunctionName(value string) *connectorConfig {
	cc.functionName = &value
	return cc
}

func (cc *connectorConfig) FunctionKey(value string) *connectorConfig {
	cc.functionKey = &value
	return cc
}

func (cc *connectorConfig) PublicKey(value string) *connectorConfig {
	cc.publicKey = &value
	return cc
}

func (cc *connectorConfig) MerchantID(value string) *connectorConfig {
	cc.merchantID = &value
	return cc
}

func (cc *connectorConfig) APIURL(value string) *connectorConfig {
	cc.apiURL = &value
	return cc
}

func (cc *connectorConfig) CloudStorageType(value string) *connectorConfig {
	cc.cloudStorageType = &value
	return cc
}

func (cc *connectorConfig) S3ExternalID(value string) *connectorConfig {
	cc.s3ExternalID = &value
	return cc
}

func (cc *connectorConfig) S3Folder(value string) *connectorConfig {
	cc.s3Folder = &value
	return cc
}

func (cc *connectorConfig) GCSBucket(value string) *connectorConfig {
	cc.gcsBucket = &value
	return cc
}

func (cc *connectorConfig) GCSFolder(value string) *connectorConfig {
	cc.gcsFolder = &value
	return cc
}

func (cc *connectorConfig) UserProfiles(value []string) *connectorConfig {
	cc.userProfiles = &value
	return cc
}

func (cc *connectorConfig) ReportConfigurationIDs(value []string) *connectorConfig {
	cc.reportConfigurationIDs = &value
	return cc
}

func (cc *connectorConfig) EnableAllDimensionCombinations(value bool) *connectorConfig {
	cc.enableAllDimensionCombinations = &value
	return cc
}

func (cc *connectorConfig) Instance(value string) *connectorConfig {
	cc.instance = &value
	return cc
}

func (cc *connectorConfig) AWSRegionCode(value string) *connectorConfig {
	cc.awsRegionCode = &value
	return cc
}

func (cc *connectorConfig) Accounts(value []string) *connectorConfig {
	cc.accounts = &value
	return cc
}

func (cc *connectorConfig) Fields(value []string) *connectorConfig {
	cc.fields = &value
	return cc
}

func (cc *connectorConfig) Breakdowns(value []string) *connectorConfig {
	cc.breakdowns = &value
	return cc
}

func (cc *connectorConfig) ActionBreakdowns(value []string) *connectorConfig {
	cc.actionBreakdowns = &value
	return cc
}

func (cc *connectorConfig) Aggregation(value string) *connectorConfig {
	cc.aggregation = &value
	return cc
}

func (cc *connectorConfig) ConfigType(value string) *connectorConfig {
	cc.configType = &value
	return cc
}

func (cc *connectorConfig) PrebuiltReport(value string) *connectorConfig {
	cc.prebuiltReport = &value
	return cc
}

func (cc *connectorConfig) ActionReportTime(value string) *connectorConfig {
	cc.actionReportTime = &value
	return cc
}

func (cc *connectorConfig) ClickAttributionWindow(value string) *connectorConfig {
	cc.clickAttributionWindow = &value
	return cc
}

func (cc *connectorConfig) ViewAttributionWindow(value string) *connectorConfig {
	cc.viewAttributionWindow = &value
	return cc
}

func (cc *connectorConfig) CustomTables(value *[]*ConnectorConfigCustomTables) *connectorConfig {
	cc.customTables = value
	return cc
}

func (cc *connectorConfig) Pages(value []string) *connectorConfig {
	cc.pages = &value
	return cc
}

func (cc *connectorConfig) Subdomain(value string) *connectorConfig {
	cc.subdomain = &value
	return cc
}

func (cc *connectorConfig) Host(value string) *connectorConfig {
	cc.host = &value
	return cc
}

func (cc *connectorConfig) Port(value int) *connectorConfig {
	cc.port = &value
	return cc
}

func (cc *connectorConfig) User(value string) *connectorConfig {
	cc.user = &value
	return cc
}

func (cc *connectorConfig) IsSecure(value string) *connectorConfig {
	cc.isSecure = &value
	return cc
}

func (cc *connectorConfig) Repositories(value []string) *connectorConfig {
	cc.repositories = &value
	return cc
}

func (cc *connectorConfig) UseWebhooks(value bool) *connectorConfig {
	cc.useWebhooks = &value
	return cc
}

func (cc *connectorConfig) DimensionAttributes(value []string) *connectorConfig {
	cc.dimensionAttributes = &value
	return cc
}

func (cc *connectorConfig) Columns(value []string) *connectorConfig {
	cc.columns = &value
	return cc
}

func (cc *connectorConfig) NetworkCode(value string) *connectorConfig {
	cc.networkCode = &value
	return cc
}

func (cc *connectorConfig) CustomerID(value string) *connectorConfig {
	cc.customerID = &value
	return cc
}

func (cc *connectorConfig) ManagerAccounts(value []string) *connectorConfig {
	cc.managerAccounts = &value
	return cc
}

func (cc *connectorConfig) Reports(value *[]*ConnectorConfigReports) *connectorConfig {
	cc.reports = value
	return cc
}

func (cc *connectorConfig) ConversionWindowSize(value int) *connectorConfig {
	cc.conversionWindowSize = &value
	return cc
}

func (cc *connectorConfig) Profiles(value []string) *connectorConfig {
	cc.profiles = &value
	return cc
}

func (cc *connectorConfig) ProjectID(value string) *connectorConfig {
	cc.projectID = &value
	return cc
}

func (cc *connectorConfig) DatasetID(value string) *connectorConfig {
	cc.datasetID = &value
	return cc
}

func (cc *connectorConfig) BucketName(value string) *connectorConfig {
	cc.bucketName = &value
	return cc
}

func (cc *connectorConfig) FunctionTrigger(value string) *connectorConfig {
	cc.functionTrigger = &value
	return cc
}

func (cc *connectorConfig) ConfigMethod(value string) *connectorConfig {
	cc.configMethod = &value
	return cc
}

func (cc *connectorConfig) QueryID(value string) *connectorConfig {
	cc.queryID = &value
	return cc
}

func (cc *connectorConfig) UpdateConfigOnEachSync(value bool) *connectorConfig {
	cc.updateConfigOnEachSync = &value
	return cc
}

func (cc *connectorConfig) SiteURLs(value []string) *connectorConfig {
	cc.siteURLs = &value
	return cc
}

func (cc *connectorConfig) Path(value string) *connectorConfig {
	cc.path = &value
	return cc
}

func (cc *connectorConfig) OnPremise(value bool) *connectorConfig {
	cc.onPremise = &value
	return cc
}

func (cc *connectorConfig) AccessToken(value string) *connectorConfig {
	cc.accessToken = &value
	return cc
}

func (cc *connectorConfig) ViewThroughAttributionWindowSize(value string) *connectorConfig {
	cc.viewThroughAttributionWindowSize = &value
	return cc
}

func (cc *connectorConfig) PostClickAttributionWindowSize(value string) *connectorConfig {
	cc.postClickAttributionWindowSize = &value
	return cc
}

func (cc *connectorConfig) UseAPIKeys(value string) *connectorConfig {
	cc.useAPIKeys = &value
	return cc
}

func (cc *connectorConfig) APIKeys(value string) *connectorConfig {
	cc.apiKeys = &value
	return cc
}

func (cc *connectorConfig) Endpoint(value string) *connectorConfig {
	cc.endpoint = &value
	return cc
}

func (cc *connectorConfig) Identity(value string) *connectorConfig {
	cc.identity = &value
	return cc
}

func (cc *connectorConfig) APIQuota(value int) *connectorConfig {
	cc.apiQuota = &value
	return cc
}

func (cc *connectorConfig) DomainName(value string) *connectorConfig {
	cc.domainName = &value
	return cc
}

func (cc *connectorConfig) ResourceURL(value string) *connectorConfig {
	cc.resourceURL = &value
	return cc
}

func (cc *connectorConfig) APISecret(value string) *connectorConfig {
	cc.apiSecret = &value
	return cc
}

func (cc *connectorConfig) Hosts(value []string) *connectorConfig {
	cc.hosts = &value
	return cc
}

func (cc *connectorConfig) TunnelHost(value string) *connectorConfig {
	cc.tunnelHost = &value
	return cc
}

func (cc *connectorConfig) TunnelPort(value int) *connectorConfig {
	cc.tunnelPort = &value
	return cc
}

func (cc *connectorConfig) TunnelUser(value string) *connectorConfig {
	cc.tunnelUser = &value
	return cc
}

func (cc *connectorConfig) Database(value string) *connectorConfig {
	cc.database = &value
	return cc
}

func (cc *connectorConfig) Datasource(value string) *connectorConfig {
	cc.datasource = &value
	return cc
}

func (cc *connectorConfig) Account(value string) *connectorConfig {
	cc.account = &value
	return cc
}

func (cc *connectorConfig) Role(value string) *connectorConfig {
	cc.role = &value
	return cc
}

func (cc *connectorConfig) Email(value string) *connectorConfig {
	cc.email = &value
	return cc
}

func (cc *connectorConfig) AccountID(value string) *connectorConfig {
	cc.accountID = &value
	return cc
}

func (cc *connectorConfig) ServerURL(value string) *connectorConfig {
	cc.serverURL = &value
	return cc
}

func (cc *connectorConfig) UserKey(value string) *connectorConfig {
	cc.userKey = &value
	return cc
}

func (cc *connectorConfig) APIVersion(value string) *connectorConfig {
	cc.apiVersion = &value
	return cc
}

func (cc *connectorConfig) DailyAPICallLimit(value int) *connectorConfig {
	cc.dailyAPICallLimit = &value
	return cc
}

func (cc *connectorConfig) TimeZone(value string) *connectorConfig {
	cc.timeZone = &value
	return cc
}

func (cc *connectorConfig) IntegrationKey(value string) *connectorConfig {
	cc.integrationKey = &value
	return cc
}

func (cc *connectorConfig) Advertisers(value []string) *connectorConfig {
	cc.advertisers = &value
	return cc
}

func (cc *connectorConfig) EngagementAttributionWindow(value string) *connectorConfig {
	cc.engagementAttributionWindow = &value
	return cc
}

func (cc *connectorConfig) ConversionReportTime(value string) *connectorConfig {
	cc.conversionReportTime = &value
	return cc
}

func (cc *connectorConfig) Domain(value string) *connectorConfig {
	cc.domain = &value
	return cc
}

func (cc *connectorConfig) UpdateMethod(value string) *connectorConfig {
	cc.updateMethod = &value
	return cc
}

func (cc *connectorConfig) ReplicationSlot(value string) *connectorConfig {
	cc.replicationSlot = &value
	return cc
}

func (cc *connectorConfig) DataCenter(value string) *connectorConfig {
	cc.dataCenter = &value
	return cc
}

func (cc *connectorConfig) APIToken(value string) *connectorConfig {
	cc.apiToken = &value
	return cc
}

func (cc *connectorConfig) SubDomain(value string) *connectorConfig {
	cc.subDomain = &value
	return cc
}

func (cc *connectorConfig) TestTableName(value string) *connectorConfig {
	cc.testTableName = &value
	return cc
}

func (cc *connectorConfig) Shop(value string) *connectorConfig {
	cc.shop = &value
	return cc
}

func (cc *connectorConfig) Organizations(value []string) *connectorConfig {
	cc.organizations = &value
	return cc
}

func (cc *connectorConfig) SwipeAttributionWindow(value string) *connectorConfig {
	cc.swipeAttributionWindow = &value
	return cc
}

func (cc *connectorConfig) APIAccessToken(value string) *connectorConfig {
	cc.apiAccessToken = &value
	return cc
}

func (cc *connectorConfig) AccountIDs(value string) *connectorConfig {
	cc.accountIDs = &value
	return cc
}

func (cc *connectorConfig) SID(value string) *connectorConfig {
	cc.sid = &value
	return cc
}

func (cc *connectorConfig) Secret(value string) *connectorConfig {
	cc.secret = &value
	return cc
}

func (cc *connectorConfig) OauthToken(value string) *connectorConfig {
	cc.oauthToken = &value
	return cc
}

func (cc *connectorConfig) OauthTokenSecret(value string) *connectorConfig {
	cc.oauthTokenSecret = &value
	return cc
}

func (cc *connectorConfig) ConsumerKey(value string) *connectorConfig {
	cc.consumerKey = &value
	return cc
}

func (cc *connectorConfig) ConsumerSecret(value string) *connectorConfig {
	cc.consumerSecret = &value
	return cc
}

func (cc *connectorConfig) Key(value string) *connectorConfig {
	cc.key = &value
	return cc
}

func (cc *connectorConfig) AdvertisersID(value []string) *connectorConfig {
	cc.advertisersID = &value
	return cc
}

func (cc *connectorConfig) SyncFormat(value string) *connectorConfig {
	cc.syncFormat = &value
	return cc
}

func (cc *connectorConfig) BucketService(value string) *connectorConfig {
	cc.bucketService = &value
	return cc
}

func (cc *connectorConfig) UserName(value string) *connectorConfig {
	cc.userName = &value
	return cc
}

func (cc *connectorConfig) ReportURL(value string) *connectorConfig {
	cc.reportURL = &value
	return cc
}

func (cc *connectorConfig) UniqueID(value string) *connectorConfig {
	cc.uniqueID = &value
	return cc
}
