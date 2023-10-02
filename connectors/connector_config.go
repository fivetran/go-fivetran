package connectors

import "github.com/fivetran/go-fivetran/utils"

// ConnectorConfig builds Connector Management, Connector Config.
// Ref. https://fivetran.com/docs/rest-api/connectors/config
type ConnectorConfig struct {
	adobeAnalyticsConfigurations     []*ConnectorConfigAdobeAnalyticsConfiguration
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
	reportSuites                     []string
	elements                         []string
	metrics                          []string
	dateGranularity                  *string
	timeframeMonths                  *string
	source                           *string
	s3Bucket                         *string
	s3RoleArn                        *string
	absConnectionString              *string
	absContainerName                 *string
	folderId                         *string
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
	isKeypair                        *bool
	advertisables                    []string
	reportType                       *string
	dimensions                       []string
	schemaPrefix                     *string
	apiKey                           *string
	externalID                       *string
	roleArn                          *string
	bucket                           *string
	prefix                           *string
	pattern                          *string
	pat                              *string
	fileType                         *string
	compression                      *string
	onError                          *string
	appendFileOption                 *string
	archivePattern                   *string
	nullSequence                     *string
	delimiter                        *string
	escapeChar                       *string
	skipBefore                       *int
	skipAfter                        *int
	secretsList                      []*FunctionSecret
	projectCredentials               []*ConnectorConfigProjectCredentials
	authMode                         *string
	username                         *string
	password                         *string
	certificate                      *string
	selectedExports                  []string
	consumerGroup                    *string
	servers                          *string
	messageType                      *string
	syncType                         *string
	securityProtocol                 *string
	apps                             []string
	salesAccounts                    []string
	financeAccounts                  []string
	appSyncMode                      *string
	salesAccountSyncMode             *string
	financeAccountSyncMode           *string
	pemCertificate                   *string
	accessKeyID                      *string
	secretKey                        *string
	homeFolder                       *string
	syncDataLocker                   *bool
	projects                         []string
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
	userProfiles                     []string
	reportConfigurationIDs           []string
	enableAllDimensionCombinations   *bool
	instance                         *string
	awsRegionCode                    *string
	accounts                         []string
	fields                           []string
	breakdowns                       []string
	actionBreakdowns                 []string
	aggregation                      *string
	configType                       *string
	prebuiltReport                   *string
	actionReportTime                 *string
	clickAttributionWindow           *string
	viewAttributionWindow            *string
	customTables                     []*ConnectorConfigCustomTables
	pages                            []string
	subdomain                        *string
	host                             *string
	port                             *int
	user                             *string
	isSecure                         *bool
	repositories                     []string
	useWebhooks                      *bool
	dimensionAttributes              []string
	columns                          []string
	networkCode                      *string
	customerID                       *string
	managerAccounts                  []string
	reports                          []*ConnectorConfigReports
	conversionWindowSize             *int
	profiles                         []string
	projectID                        *string
	datasetID                        *string
	bucketName                       *string
	functionTrigger                  *string
	configMethod                     *string
	queryID                          *string
	updateConfigOnEachSync           *bool
	siteURLs                         []string
	path                             *string
	onPremise                        *bool
	accessToken                      *string
	viewThroughAttributionWindowSize *string
	postClickAttributionWindowSize   *string
	useAPIKeys                       *bool
	apiKeys                          []string
	endpoint                         *string
	identity                         *string
	apiQuota                         *int
	domainName                       *string
	resourceURL                      *string
	apiSecret                        *string
	hosts                            []string
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
	advertisers                      []string
	engagementAttributionWindow      *string
	conversionReportTime             *string
	domain                           *string
	updateMethod                     *string
	replicationSlot                  *string
	publicationName                  *string
	dataCenter                       *string
	apiToken                         *string
	subDomain                        *string
	testTableName                    *string
	shop                             *string
	organizations                    []string
	swipeAttributionWindow           *string
	apiAccessToken                   *string
	accountIDs                       []string
	sid                              *string
	secret                           *string
	oauthToken                       *string
	oauthTokenSecret                 *string
	consumerKey                      *string
	consumerSecret                   *string
	key                              *string
	advertisersID                    []string
	syncFormat                       *string
	bucketService                    *string
	userName                         *string
	reportURL                        *string
	uniqueID                         *string
	authType                         *string
	isNewPackage                     *bool
	connectionType                   *string
	isMultiEntityFeatureEnabled      *bool
	alwaysEncrypted                  *bool
	apiType                          *string
	baseUrl                          *string
	entityId                         *string
	soapUri                          *string
	userId                           *string
	encryptionKey                    *string
	euRegion                         *bool
	tokenKey                         *string
	tokenSecret                      *string
	shareURL                         *string
}

type connectorConfigRequest struct {
	AdobeAnalyticsConfigurations     []*connectorConfigAdobeAnalyticsConfigurationRequest `json:"adobe_analytics_configurations,omitempty"`
	Schema                           *string                                              `json:"schema,omitempty"`
	Table                            *string                                              `json:"table,omitempty"`
	SheetID                          *string                                              `json:"sheet_id,omitempty"`
	NamedRange                       *string                                              `json:"named_range,omitempty"`
	ClientID                         *string                                              `json:"client_id,omitempty"`
	ClientSecret                     *string                                              `json:"client_secret,omitempty"`
	TechnicalAccountID               *string                                              `json:"technical_account_id,omitempty"`
	OrganizationID                   *string                                              `json:"organization_id,omitempty"`
	PrivateKey                       *string                                              `json:"private_key,omitempty"`
	SyncMode                         *string                                              `json:"sync_mode,omitempty"`
	ReportSuites                     []string                                             `json:"report_suites,omitempty"`
	Elements                         []string                                             `json:"elements,omitempty"`
	Metrics                          []string                                             `json:"metrics,omitempty"`
	DateGranularity                  *string                                              `json:"date_granularity,omitempty"`
	TimeframeMonths                  *string                                              `json:"timeframe_months,omitempty"`
	Source                           *string                                              `json:"source,omitempty"`
	S3Bucket                         *string                                              `json:"s3bucket,omitempty"`
	S3RoleArn                        *string                                              `json:"s3role_arn,omitempty"`
	ABSConnectionString              *string                                              `json:"abs_connection_string,omitempty"`
	ABSContainerName                 *string                                              `json:"abs_container_name,omitempty"`
	FolderId                         *string                                              `json:"folder_id,omitempty"`
	FTPHost                          *string                                              `json:"ftp_host,omitempty"`
	FTPPort                          *int                                                 `json:"ftp_port,omitempty"`
	FTPUser                          *string                                              `json:"ftp_user,omitempty"`
	FTPPassword                      *string                                              `json:"ftp_password,omitempty"`
	IsFTPS                           *bool                                                `json:"is_ftps,omitempty"`
	SFTPHost                         *string                                              `json:"sftp_host,omitempty"`
	SFTPPort                         *int                                                 `json:"sftp_port,omitempty"`
	SFTPUser                         *string                                              `json:"sftp_user,omitempty"`
	SFTPPassword                     *string                                              `json:"sftp_password,omitempty"`
	SFTPIsKeyPair                    *bool                                                `json:"sftp_is_key_pair,omitempty"`
	IsKeypair                        *bool                                                `json:"is_keypair,omitempty"`
	Advertisables                    []string                                             `json:"advertisables,omitempty"`
	ReportType                       *string                                              `json:"report_type,omitempty"`
	Dimensions                       []string                                             `json:"dimensions,omitempty"`
	SchemaPrefix                     *string                                              `json:"schema_prefix,omitempty"`
	APIKey                           *string                                              `json:"api_key,omitempty"`
	ExternalID                       *string                                              `json:"external_id,omitempty"`
	RoleArn                          *string                                              `json:"role_arn,omitempty"`
	Bucket                           *string                                              `json:"bucket,omitempty"`
	Prefix                           *string                                              `json:"prefix,omitempty"`
	Pattern                          *string                                              `json:"pattern,omitempty"`
	PAT                              *string                                              `json:"pat,omitempty"`
	FileType                         *string                                              `json:"file_type,omitempty"`
	Compression                      *string                                              `json:"compression,omitempty"`
	OnError                          *string                                              `json:"on_error,omitempty"`
	AppendFileOption                 *string                                              `json:"append_file_option,omitempty"`
	ArchivePattern                   *string                                              `json:"archive_pattern,omitempty"`
	NullSequence                     *string                                              `json:"null_sequence,omitempty"`
	Delimiter                        *string                                              `json:"delimiter,omitempty"`
	EscapeChar                       *string                                              `json:"escape_char,omitempty"`
	SkipBefore                       *int                                                 `json:"skip_before,omitempty"`
	SkipAfter                        *int                                                 `json:"skip_after,omitempty"`
	SecretsList                      []*functionSecretRequest                             `json:"secrets_list,omitempty"`
	ProjectCredentials               []*connectorConfigProjectCredentialsRequest          `json:"project_credentials,omitempty"`
	AuthMode                         *string                                              `json:"auth_mode,omitempty"`
	Username                         *string                                              `json:"username,omitempty"`
	Password                         *string                                              `json:"password,omitempty"`
	Certificate                      *string                                              `json:"certificate,omitempty"`
	SelectedExports                  []string                                             `json:"selected_exports,omitempty"`
	ConsumerGroup                    *string                                              `json:"consumer_group,omitempty"`
	Servers                          *string                                              `json:"servers,omitempty"`
	MessageType                      *string                                              `json:"message_type,omitempty"`
	SyncType                         *string                                              `json:"sync_type,omitempty"`
	SecurityProtocol                 *string                                              `json:"security_protocol,omitempty"`
	Apps                             []string                                             `json:"apps,omitempty"`
	SalesAccounts                    []string                                             `json:"sales_accounts,omitempty"`
	FinanceAccounts                  []string                                             `json:"finance_accounts,omitempty"`
	AppSyncMode                      *string                                              `json:"app_sync_mode,omitempty"`
	SalesAccountSyncMode             *string                                              `json:"sales_account_sync_mode,omitempty"`
	FinanceAccountSyncMode           *string                                              `json:"finance_account_sync_mode,omitempty"`
	PEMCertificate                   *string                                              `json:"pem_certificate,omitempty"`
	AccessKeyID                      *string                                              `json:"access_key_id,omitempty"`
	SecretKey                        *string                                              `json:"secret_key,omitempty"`
	HomeFolder                       *string                                              `json:"home_folder,omitempty"`
	SyncDataLocker                   *bool                                                `json:"sync_data_locker,omitempty"`
	Projects                         []string                                             `json:"projects,omitempty"`
	Function                         *string                                              `json:"function,omitempty"`
	Region                           *string                                              `json:"region,omitempty"`
	Secrets                          *string                                              `json:"secrets,omitempty"`
	ContainerName                    *string                                              `json:"container_name,omitempty"`
	ConnectionString                 *string                                              `json:"connection_string,omitempty"`
	FunctionApp                      *string                                              `json:"function_app,omitempty"`
	FunctionName                     *string                                              `json:"function_name,omitempty"`
	FunctionKey                      *string                                              `json:"function_key,omitempty"`
	PublicKey                        *string                                              `json:"public_key,omitempty"`
	MerchantID                       *string                                              `json:"merchant_id,omitempty"`
	APIURL                           *string                                              `json:"api_url,omitempty"`
	CloudStorageType                 *string                                              `json:"cloud_storage_type,omitempty"`
	S3ExternalID                     *string                                              `json:"s3external_id,omitempty"`
	S3Folder                         *string                                              `json:"s3folder,omitempty"`
	GCSBucket                        *string                                              `json:"gcs_bucket,omitempty"`
	GCSFolder                        *string                                              `json:"gcs_folder,omitempty"`
	UserProfiles                     []string                                             `json:"user_profiles,omitempty"`
	ReportConfigurationIDs           []string                                             `json:"report_configuration_ids,omitempty"`
	EnableAllDimensionCombinations   *bool                                                `json:"enable_all_dimension_combinations,omitempty"`
	Instance                         *string                                              `json:"instance,omitempty"`
	AWSRegionCode                    *string                                              `json:"aws_region_code,omitempty"`
	Accounts                         []string                                             `json:"accounts,omitempty"`
	Fields                           []string                                             `json:"fields,omitempty"`
	Breakdowns                       []string                                             `json:"breakdowns,omitempty"`
	ActionBreakdowns                 []string                                             `json:"action_breakdowns,omitempty"`
	Aggregation                      *string                                              `json:"aggregation,omitempty"`
	ConfigType                       *string                                              `json:"config_type,omitempty"`
	PrebuiltReport                   *string                                              `json:"prebuilt_report,omitempty"`
	ActionReportTime                 *string                                              `json:"action_report_time,omitempty"`
	ClickAttributionWindow           *string                                              `json:"click_attribution_window,omitempty"`
	ViewAttributionWindow            *string                                              `json:"view_attribution_window,omitempty"`
	CustomTables                     []*connectorConfigCustomTablesRequest                `json:"custom_tables,omitempty"`
	Pages                            []string                                             `json:"pages,omitempty"`
	Subdomain                        *string                                              `json:"subdomain,omitempty"`
	Host                             *string                                              `json:"host,omitempty"`
	Port                             *int                                                 `json:"port,omitempty"`
	User                             *string                                              `json:"user,omitempty"`
	IsSecure                         *bool                                                `json:"is_secure,omitempty"`
	Repositories                     []string                                             `json:"repositories,omitempty"`
	UseWebhooks                      *bool                                                `json:"use_webhooks,omitempty"`
	DimensionAttributes              []string                                             `json:"dimension_attributes,omitempty"`
	Columns                          []string                                             `json:"columns,omitempty"`
	NetworkCode                      *string                                              `json:"network_code,omitempty"`
	CustomerID                       *string                                              `json:"customer_id,omitempty"`
	ManagerAccounts                  []string                                             `json:"manager_accounts,omitempty"`
	Reports                          []*connectorConfigReportsRequest                     `json:"reports,omitempty"`
	ConversionWindowSize             *int                                                 `json:"conversion_window_size,omitempty"`
	Profiles                         []string                                             `json:"profiles,omitempty"`
	ProjectID                        *string                                              `json:"project_id,omitempty"`
	DatasetID                        *string                                              `json:"dataset_id,omitempty"`
	BucketName                       *string                                              `json:"bucket_name,omitempty"`
	FunctionTrigger                  *string                                              `json:"function_trigger,omitempty"`
	ConfigMethod                     *string                                              `json:"config_method,omitempty"`
	QueryID                          *string                                              `json:"query_id,omitempty"`
	UpdateConfigOnEachSync           *bool                                                `json:"update_config_on_each_sync,omitempty"`
	SiteURLs                         []string                                             `json:"site_urls,omitempty"`
	Path                             *string                                              `json:"path,omitempty"`
	OnPremise                        *bool                                                `json:"on_premise,omitempty"`
	AccessToken                      *string                                              `json:"access_token,omitempty"`
	ViewThroughAttributionWindowSize *string                                              `json:"view_through_attribution_window_size,omitempty"`
	PostClickAttributionWindowSize   *string                                              `json:"post_click_attribution_window_size,omitempty"`
	UseAPIKeys                       *bool                                                `json:"use_api_keys,omitempty"`
	APIKeys                          []string                                             `json:"api_keys,omitempty"`
	Endpoint                         *string                                              `json:"endpoint,omitempty"`
	Identity                         *string                                              `json:"identity,omitempty"`
	APIQuota                         *int                                                 `json:"api_quota,omitempty"`
	DomainName                       *string                                              `json:"domain_name,omitempty"`
	ResourceURL                      *string                                              `json:"resource_url,omitempty"`
	APISecret                        *string                                              `json:"api_secret,omitempty"`
	Hosts                            []string                                             `json:"hosts,omitempty"`
	TunnelHost                       *string                                              `json:"tunnel_host,omitempty"`
	TunnelPort                       *int                                                 `json:"tunnel_port,omitempty"`
	TunnelUser                       *string                                              `json:"tunnel_user,omitempty"`
	Database                         *string                                              `json:"database,omitempty"`
	Datasource                       *string                                              `json:"datasource,omitempty"`
	Account                          *string                                              `json:"account,omitempty"`
	Role                             *string                                              `json:"role,omitempty"`
	Email                            *string                                              `json:"email,omitempty"`
	AccountID                        *string                                              `json:"account_id,omitempty"`
	ServerURL                        *string                                              `json:"server_url,omitempty"`
	UserKey                          *string                                              `json:"user_key,omitempty"`
	APIVersion                       *string                                              `json:"api_version,omitempty"`
	DailyAPICallLimit                *int                                                 `json:"daily_api_call_limit,omitempty"`
	TimeZone                         *string                                              `json:"time_zone,omitempty"`
	IntegrationKey                   *string                                              `json:"integration_key,omitempty"`
	Advertisers                      []string                                             `json:"advertisers,omitempty"`
	EngagementAttributionWindow      *string                                              `json:"engagement_attribution_window,omitempty"`
	ConversionReportTime             *string                                              `json:"conversion_report_time,omitempty"`
	Domain                           *string                                              `json:"domain,omitempty"`
	UpdateMethod                     *string                                              `json:"update_method,omitempty"`
	ReplicationSlot                  *string                                              `json:"replication_slot,omitempty"`
	PublicationName                  *string                                              `json:"publication_name,omitempty"`
	DataCenter                       *string                                              `json:"data_center,omitempty"`
	APIToken                         *string                                              `json:"api_token,omitempty"`
	SubDomain                        *string                                              `json:"sub_domain,omitempty"`
	TestTableName                    *string                                              `json:"test_table_name,omitempty"`
	Shop                             *string                                              `json:"shop,omitempty"`
	Organizations                    []string                                             `json:"organizations,omitempty"`
	SwipeAttributionWindow           *string                                              `json:"swipe_attribution_window,omitempty"`
	APIAccessToken                   *string                                              `json:"api_access_token,omitempty"`
	AccountIDs                       []string                                             `json:"account_ids,omitempty"`
	SID                              *string                                              `json:"sid,omitempty"`
	Secret                           *string                                              `json:"secret,omitempty"`
	OauthToken                       *string                                              `json:"oauth_token,omitempty"`
	OauthTokenSecret                 *string                                              `json:"oauth_token_secret,omitempty"`
	ConsumerKey                      *string                                              `json:"consumer_key,omitempty"`
	ConsumerSecret                   *string                                              `json:"consumer_secret,omitempty"`
	Key                              *string                                              `json:"key,omitempty"`
	AdvertisersID                    []string                                             `json:"advertisers_id,omitempty"`
	SyncFormat                       *string                                              `json:"sync_format,omitempty"`
	BucketService                    *string                                              `json:"bucket_service,omitempty"`
	UserName                         *string                                              `json:"user_name,omitempty"`
	ReportURL                        *string                                              `json:"report_url,omitempty"`
	UniqueID                         *string                                              `json:"unique_id,omitempty"`
	AuthType                         *string                                              `json:"auth_type,omitempty"`
	IsNewPackage                     *bool                                                `json:"is_new_package,omitempty"`
	ConnectionType                   *string                                              `json:"connection_type,omitempty"`
	IsMultiEntityFeatureEnabled      *bool                                                `json:"is_multi_entity_feature_enabled,omitempty"`
	AlwaysEncrypted                  *bool                                                `json:"always_encrypted,omitempty"`
	ApiType                          *string                                              `json:"api_type,omitempty"`
	BaseUrl                          *string                                              `json:"base_url,omitempty"`
	EntityId                         *string                                              `json:"entity_id,omitempty"`
	SoapUri                          *string                                              `json:"soap_uri,omitempty"`
	UserId                           *string                                              `json:"user_id,omitempty"`
	EncryptionKey                    *string                                              `json:"encryption_key,omitempty"`
	EuRegion                         *bool                                                `json:"eu_region,omitempty"`
	TokenKey                         *string                                              `json:"token_key,omitempty"`
	TokenSecret                      *string                                              `json:"token_secret,omitempty"`
	ShareURL                         *string                                              `json:"share_url,omitempty"`
}

type ConnectorConfigResponse struct {
	AdobeAnalyticsConfigurations     []ConnectorConfigAdobeAnalyticsConfigurationResponse `json:"adobe_analytics_configurations"`
	Schema                           string                                               `json:"schema"`
	Table                            string                                               `json:"table"`
	SheetID                          string                                               `json:"sheet_id"`
	NamedRange                       string                                               `json:"named_range"`
	ClientID                         string                                               `json:"client_id"`
	ClientSecret                     string                                               `json:"client_secret"`
	TechnicalAccountID               string                                               `json:"technical_account_id"`
	OrganizationID                   string                                               `json:"organization_id"`
	PrivateKey                       string                                               `json:"private_key"`
	SyncMode                         string                                               `json:"sync_mode"`
	ReportSuites                     []string                                             `json:"report_suites"`
	Elements                         []string                                             `json:"elements"`
	Metrics                          []string                                             `json:"metrics"`
	DateGranularity                  string                                               `json:"date_granularity"`
	TimeframeMonths                  string                                               `json:"timeframe_months"`
	Source                           string                                               `json:"source"`
	S3Bucket                         string                                               `json:"s3bucket"`
	S3RoleArn                        string                                               `json:"s3role_arn"`
	ABSConnectionString              string                                               `json:"abs_connection_string"`
	ABSContainerName                 string                                               `json:"abs_container_name"`
	FolderId                         string                                               `json:"folder_id"`
	FTPHost                          string                                               `json:"ftp_host"`
	FTPPort                          *int                                                 `json:"ftp_port"`
	FTPUser                          string                                               `json:"ftp_user"`
	FTPPassword                      string                                               `json:"ftp_password"`
	IsFTPS                           *bool                                                `json:"is_ftps"`
	SFTPHost                         string                                               `json:"sftp_host"`
	SFTPPort                         *int                                                 `json:"sftp_port"`
	SFTPUser                         string                                               `json:"sftp_user"`
	SFTPPassword                     string                                               `json:"sftp_password"`
	SFTPIsKeyPair                    *bool                                                `json:"sftp_is_key_pair"`
	IsKeypair                        *bool                                                `json:"is_keypair"`
	Advertisables                    []string                                             `json:"advertisables"`
	ReportType                       string                                               `json:"report_type"`
	Dimensions                       []string                                             `json:"dimensions"`
	SchemaPrefix                     string                                               `json:"schema_prefix"`
	APIKey                           string                                               `json:"api_key"`
	ExternalID                       string                                               `json:"external_id"`
	RoleArn                          string                                               `json:"role_arn"`
	Bucket                           string                                               `json:"bucket"`
	Prefix                           string                                               `json:"prefix"`
	Pattern                          string                                               `json:"pattern"`
	PAT                              string                                               `json:"pat"`
	FileType                         string                                               `json:"file_type"`
	Compression                      string                                               `json:"compression"`
	OnError                          string                                               `json:"on_error"`
	AppendFileOption                 string                                               `json:"append_file_option"`
	ArchivePattern                   string                                               `json:"archive_pattern"`
	NullSequence                     string                                               `json:"null_sequence"`
	Delimiter                        string                                               `json:"delimiter"`
	EscapeChar                       string                                               `json:"escape_char"`
	SkipBefore                       *int                                                 `json:"skip_before"`
	SkipAfter                        *int                                                 `json:"skip_after"`
	SecretsList                      []FunctionSecretResponse                             `json:"secrets_list"`
	ProjectCredentials               []ConnectorConfigProjectCredentialsResponse          `json:"project_credentials"`
	AuthMode                         string                                               `json:"auth_mode"`
	Username                         string                                               `json:"username"`
	Password                         string                                               `json:"password"`
	Certificate                      string                                               `json:"certificate"`
	SelectedExports                  []string                                             `json:"selected_exports"`
	ConsumerGroup                    string                                               `json:"consumer_group"`
	Servers                          string                                               `json:"servers"`
	MessageType                      string                                               `json:"message_type"`
	SyncType                         string                                               `json:"sync_type"`
	SecurityProtocol                 string                                               `json:"security_protocol"`
	Apps                             []string                                             `json:"apps"`
	SalesAccounts                    []string                                             `json:"sales_accounts"`
	FinanceAccounts                  []string                                             `json:"finance_accounts"`
	AppSyncMode                      string                                               `json:"app_sync_mode"`
	SalesAccountSyncMode             string                                               `json:"sales_account_sync_mode"`
	FinanceAccountSyncMode           string                                               `json:"finance_account_sync_mode"`
	PEMCertificate                   string                                               `json:"pem_certificate"`
	AccessKeyID                      string                                               `json:"access_key_id"`
	SecretKey                        string                                               `json:"secret_key"`
	HomeFolder                       string                                               `json:"home_folder"`
	SyncDataLocker                   *bool                                                `json:"sync_data_locker"`
	Projects                         []string                                             `json:"projects"`
	Function                         string                                               `json:"function"`
	Region                           string                                               `json:"region"`
	Secrets                          string                                               `json:"secrets"`
	ContainerName                    string                                               `json:"container_name"`
	ConnectionString                 string                                               `json:"connection_string"`
	FunctionApp                      string                                               `json:"function_app"`
	FunctionName                     string                                               `json:"function_name"`
	FunctionKey                      string                                               `json:"function_key"`
	PublicKey                        string                                               `json:"public_key"`
	MerchantID                       string                                               `json:"merchant_id"`
	APIURL                           string                                               `json:"api_url"`
	CloudStorageType                 string                                               `json:"cloud_storage_type"`
	S3ExternalID                     string                                               `json:"s3external_id"`
	S3Folder                         string                                               `json:"s3folder"`
	GCSBucket                        string                                               `json:"gcs_bucket"`
	GCSFolder                        string                                               `json:"gcs_folder"`
	UserProfiles                     []string                                             `json:"user_profiles"`
	ReportConfigurationIDs           []string                                             `json:"report_configuration_ids"`
	EnableAllDimensionCombinations   *bool                                                `json:"enable_all_dimension_combinations"`
	Instance                         string                                               `json:"instance"`
	AWSRegionCode                    string                                               `json:"aws_region_code"`
	Accounts                         []string                                             `json:"accounts"`
	Fields                           []string                                             `json:"fields"`
	Breakdowns                       []string                                             `json:"breakdowns"`
	ActionBreakdowns                 []string                                             `json:"action_breakdowns"`
	Aggregation                      string                                               `json:"aggregation"`
	ConfigType                       string                                               `json:"config_type"`
	PrebuiltReport                   string                                               `json:"prebuilt_report"`
	ActionReportTime                 string                                               `json:"action_report_time"`
	ClickAttributionWindow           string                                               `json:"click_attribution_window"`
	ViewAttributionWindow            string                                               `json:"view_attribution_window"`
	CustomTables                     []ConnectorConfigCustomTablesResponse                `json:"custom_tables"`
	Pages                            []string                                             `json:"pages"`
	Subdomain                        string                                               `json:"subdomain"`
	Host                             string                                               `json:"host"`
	Port                             *int                                                 `json:"port"`
	User                             string                                               `json:"user"`
	IsSecure                         *bool                                                `json:"is_secure"`
	Repositories                     []string                                             `json:"repositories"`
	UseWebhooks                      *bool                                                `json:"use_webhooks"`
	DimensionAttributes              []string                                             `json:"dimension_attributes"`
	Columns                          []string                                             `json:"columns"`
	NetworkCode                      string                                               `json:"network_code"`
	CustomerID                       string                                               `json:"customer_id"`
	ManagerAccounts                  []string                                             `json:"manager_accounts"`
	Reports                          []ConnectorConfigReportsResponse                     `json:"reports"`
	ConversionWindowSize             *int                                                 `json:"conversion_window_size"`
	Profiles                         []string                                             `json:"profiles"`
	ProjectID                        string                                               `json:"project_id"`
	DatasetID                        string                                               `json:"dataset_id"`
	BucketName                       string                                               `json:"bucket_name"`
	FunctionTrigger                  string                                               `json:"function_trigger"`
	ConfigMethod                     string                                               `json:"config_method"`
	QueryID                          string                                               `json:"query_id"`
	UpdateConfigOnEachSync           *bool                                                `json:"update_config_on_each_sync"`
	SiteURLs                         []string                                             `json:"site_urls"`
	Path                             string                                               `json:"path"`
	OnPremise                        *bool                                                `json:"on_premise"`
	AccessToken                      string                                               `json:"access_token"`
	ViewThroughAttributionWindowSize string                                               `json:"view_through_attribution_window_size"`
	PostClickAttributionWindowSize   string                                               `json:"post_click_attribution_window_size"`
	UseAPIKeys                       *bool                                                `json:"use_api_keys"`
	APIKeys                          []string                                             `json:"api_keys"`
	Endpoint                         string                                               `json:"endpoint"`
	Identity                         string                                               `json:"identity"`
	APIQuota                         *int                                                 `json:"api_quota"`
	DomainName                       string                                               `json:"domain_name"`
	ResourceURL                      string                                               `json:"resource_url"`
	APISecret                        string                                               `json:"api_secret"`
	Hosts                            []string                                             `json:"hosts"`
	TunnelHost                       string                                               `json:"tunnel_host"`
	TunnelPort                       *int                                                 `json:"tunnel_port"`
	TunnelUser                       string                                               `json:"tunnel_user"`
	Database                         string                                               `json:"database"`
	Datasource                       string                                               `json:"datasource"`
	Account                          string                                               `json:"account"`
	Role                             string                                               `json:"role"`
	Email                            string                                               `json:"email"`
	AccountID                        string                                               `json:"account_id"`
	ServerURL                        string                                               `json:"server_url"`
	UserKey                          string                                               `json:"user_key"`
	APIVersion                       string                                               `json:"api_version"`
	DailyAPICallLimit                *int                                                 `json:"daily_api_call_limit"`
	TimeZone                         string                                               `json:"time_zone"`
	IntegrationKey                   string                                               `json:"integration_key"`
	Advertisers                      []string                                             `json:"advertisers"`
	EngagementAttributionWindow      string                                               `json:"engagement_attribution_window"`
	ConversionReportTime             string                                               `json:"conversion_report_time"`
	Domain                           string                                               `json:"domain"`
	UpdateMethod                     string                                               `json:"update_method"`
	ReplicationSlot                  string                                               `json:"replication_slot"`
	PublicationName                  string                                               `json:"publication_name"`
	DataCenter                       string                                               `json:"data_center"`
	APIToken                         string                                               `json:"api_token"`
	SubDomain                        string                                               `json:"sub_domain"`
	TestTableName                    string                                               `json:"test_table_name"`
	Shop                             string                                               `json:"shop"`
	Organizations                    []string                                             `json:"organizations"`
	SwipeAttributionWindow           string                                               `json:"swipe_attribution_window"`
	APIAccessToken                   string                                               `json:"api_access_token"`
	AccountIDs                       []string                                             `json:"account_ids"`
	SID                              string                                               `json:"sid"`
	Secret                           string                                               `json:"secret"`
	OauthToken                       string                                               `json:"oauth_token"`
	OauthTokenSecret                 string                                               `json:"oauth_token_secret"`
	ConsumerKey                      string                                               `json:"consumer_key"`
	ConsumerSecret                   string                                               `json:"consumer_secret"`
	Key                              string                                               `json:"key"`
	AdvertisersID                    []string                                             `json:"advertisers_id"`
	SyncFormat                       string                                               `json:"sync_format"`
	BucketService                    string                                               `json:"bucket_service"`
	UserName                         string                                               `json:"user_name"`
	ReportURL                        string                                               `json:"report_url"`
	UniqueID                         string                                               `json:"unique_id"`
	AuthType                         string                                               `json:"auth_type"`
	LatestVersion                    string                                               `json:"latest_version"`
	AuthorizationMethod              string                                               `json:"authorization_method"`
	ServiceVersion                   string                                               `json:"service_version"`
	LastSyncedChangesUtc             string                                               `json:"last_synced_changes__utc_"`
	IsNewPackage                     *bool                                                `json:"is_new_package"`
	ConnectionType                   string                                               `json:"connection_type"`
	IsMultiEntityFeatureEnabled      *bool                                                `json:"is_multi_entity_feature_enabled"`
	AlwaysEncrypted                  *bool                                                `json:"always_encrypted"`
	ApiType                          string                                               `json:"api_type"`
	BaseUrl                          string                                               `json:"base_url"`
	EntityId                         string                                               `json:"entity_id"`
	SoapUri                          string                                               `json:"soap_uri"`
	UserId                           string                                               `json:"user_id"`
	EncryptionKey                    string                                               `json:"encryption_key"`
	EuRegion                         *bool                                                `json:"eu_region"`
	TokenKey                         string                                               `json:"token_key"`
	TokenSecret                      string                                               `json:"token_secret"`
	ShareURL                         string                                               `json:"share_url"`
}

func (cc *ConnectorConfig) Merge(customConfig *map[string]interface{}) (*map[string]interface{}, error) {
	err := utils.MergeIntoMap(cc.Request(), customConfig)
	if err != nil {
		return nil, err
	}
	return customConfig, nil
}

func (cc *ConnectorConfig) Request() *connectorConfigRequest {
	var projectCredentials []*connectorConfigProjectCredentialsRequest
	if cc.projectCredentials != nil {
		for _, pc := range cc.projectCredentials {
			projectCredentials = append(projectCredentials, pc.request())
		}
	}

	var customTables []*connectorConfigCustomTablesRequest
	if cc.customTables != nil {
		for _, ct := range cc.customTables {
			customTables = append(customTables, ct.request())
		}
	}

	var reports []*connectorConfigReportsRequest
	if cc.reports != nil {
		for _, r := range cc.reports {
			reports = append(reports, r.request())
		}
	}

	var adobeAnalyticsConfigurations []*connectorConfigAdobeAnalyticsConfigurationRequest
	if cc.adobeAnalyticsConfigurations != nil {
		for _, r := range cc.adobeAnalyticsConfigurations {
			adobeAnalyticsConfigurations = append(adobeAnalyticsConfigurations, r.request())
		}
	}

	var functionSecrets []*functionSecretRequest
	if cc.secretsList != nil {
		for _, s := range cc.secretsList {
			functionSecrets = append(functionSecrets, s.request())
		}
	}

	return &connectorConfigRequest{
		AdobeAnalyticsConfigurations:     adobeAnalyticsConfigurations,
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
		FolderId:                         cc.folderId,
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
		IsKeypair:                        cc.isKeypair,
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
		PAT:                              cc.pat,
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
		SecretsList:                      functionSecrets,
		ProjectCredentials:               projectCredentials,
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
		CustomTables:                     customTables,
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
		Reports:                          reports,
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
		PublicationName:                  cc.publicationName,
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
		AuthType:                         cc.authType,
		IsNewPackage:                     cc.isNewPackage,
		ConnectionType:                   cc.connectionType,
		IsMultiEntityFeatureEnabled:      cc.isMultiEntityFeatureEnabled,
		AlwaysEncrypted:                  cc.alwaysEncrypted,
		ApiType:                          cc.apiType,
		BaseUrl:                          cc.baseUrl,
		EntityId:                         cc.entityId,
		SoapUri:                          cc.soapUri,
		UserId:                           cc.userId,
		EncryptionKey:                    cc.encryptionKey,
		EuRegion:                         cc.euRegion,
		TokenKey:                         cc.tokenKey,
		TokenSecret:                      cc.tokenSecret,
		ShareURL:                         cc.shareURL,
	}
}

func (cc *ConnectorConfig) Schema(value string) *ConnectorConfig {
	cc.schema = &value
	return cc
}

func (cc *ConnectorConfig) Table(value string) *ConnectorConfig {
	cc.table = &value
	return cc
}

func (cc *ConnectorConfig) SheetID(value string) *ConnectorConfig {
	cc.sheetID = &value
	return cc
}

func (cc *ConnectorConfig) NamedRange(value string) *ConnectorConfig {
	cc.namedRange = &value
	return cc
}

func (cc *ConnectorConfig) ClientID(value string) *ConnectorConfig {
	cc.clientID = &value
	return cc
}

func (cc *ConnectorConfig) ClientSecret(value string) *ConnectorConfig {
	cc.clientSecret = &value
	return cc
}

func (cc *ConnectorConfig) TechnicalAccountID(value string) *ConnectorConfig {
	cc.technicalAccountID = &value
	return cc
}

func (cc *ConnectorConfig) OrganizationID(value string) *ConnectorConfig {
	cc.organizationID = &value
	return cc
}

func (cc *ConnectorConfig) PrivateKey(value string) *ConnectorConfig {
	cc.privateKey = &value
	return cc
}

func (cc *ConnectorConfig) SyncMode(value string) *ConnectorConfig {
	cc.syncMode = &value
	return cc
}

func (cc *ConnectorConfig) ReportSuites(value []string) *ConnectorConfig {
	cc.reportSuites = value
	return cc
}

func (cc *ConnectorConfig) Elements(value []string) *ConnectorConfig {
	cc.elements = value
	return cc
}

func (cc *ConnectorConfig) Metrics(value []string) *ConnectorConfig {
	cc.metrics = value
	return cc
}

func (cc *ConnectorConfig) DateGranularity(value string) *ConnectorConfig {
	cc.dateGranularity = &value
	return cc
}

func (cc *ConnectorConfig) TimeframeMonths(value string) *ConnectorConfig {
	cc.timeframeMonths = &value
	return cc
}

func (cc *ConnectorConfig) Source(value string) *ConnectorConfig {
	cc.source = &value
	return cc
}

func (cc *ConnectorConfig) S3Bucket(value string) *ConnectorConfig {
	cc.s3Bucket = &value
	return cc
}

func (cc *ConnectorConfig) S3RoleArn(value string) *ConnectorConfig {
	cc.s3RoleArn = &value
	return cc
}

func (cc *ConnectorConfig) ABSConnectionString(value string) *ConnectorConfig {
	cc.absConnectionString = &value
	return cc
}

func (cc *ConnectorConfig) ABSContainerName(value string) *ConnectorConfig {
	cc.absContainerName = &value
	return cc
}

func (cc *ConnectorConfig) FolderId(value string) *ConnectorConfig {
	cc.folderId = &value
	return cc
}

func (cc *ConnectorConfig) FTPHost(value string) *ConnectorConfig {
	cc.ftpHost = &value
	return cc
}

func (cc *ConnectorConfig) FTPPort(value int) *ConnectorConfig {
	cc.ftpPort = &value
	return cc
}

func (cc *ConnectorConfig) FTPUser(value string) *ConnectorConfig {
	cc.ftpUser = &value
	return cc
}

func (cc *ConnectorConfig) FTPPassword(value string) *ConnectorConfig {
	cc.ftpPassword = &value
	return cc
}

func (cc *ConnectorConfig) IsFTPS(value bool) *ConnectorConfig {
	cc.isFTPS = &value
	return cc
}

func (cc *ConnectorConfig) SFTPHost(value string) *ConnectorConfig {
	cc.sFTPHost = &value
	return cc
}

func (cc *ConnectorConfig) SFTPPort(value int) *ConnectorConfig {
	cc.sFTPPort = &value
	return cc
}

func (cc *ConnectorConfig) SFTPUser(value string) *ConnectorConfig {
	cc.sFTPUser = &value
	return cc
}

func (cc *ConnectorConfig) SFTPPassword(value string) *ConnectorConfig {
	cc.sFTPPassword = &value
	return cc
}

func (cc *ConnectorConfig) SFTPIsKeyPair(value bool) *ConnectorConfig {
	cc.sFTPIsKeyPair = &value
	return cc
}

func (cc *ConnectorConfig) IsKeypair(value bool) *ConnectorConfig {
	cc.isKeypair = &value
	return cc
}

func (cc *ConnectorConfig) Advertisables(value []string) *ConnectorConfig {
	cc.advertisables = value
	return cc
}

func (cc *ConnectorConfig) ReportType(value string) *ConnectorConfig {
	cc.reportType = &value
	return cc
}

func (cc *ConnectorConfig) Dimensions(value []string) *ConnectorConfig {
	cc.dimensions = value
	return cc
}

func (cc *ConnectorConfig) SchemaPrefix(value string) *ConnectorConfig {
	cc.schemaPrefix = &value
	return cc
}

func (cc *ConnectorConfig) APIKey(value string) *ConnectorConfig {
	cc.apiKey = &value
	return cc
}

func (cc *ConnectorConfig) ExternalID(value string) *ConnectorConfig {
	cc.externalID = &value
	return cc
}

func (cc *ConnectorConfig) RoleArn(value string) *ConnectorConfig {
	cc.roleArn = &value
	return cc
}

func (cc *ConnectorConfig) Bucket(value string) *ConnectorConfig {
	cc.bucket = &value
	return cc
}

func (cc *ConnectorConfig) Prefix(value string) *ConnectorConfig {
	cc.prefix = &value
	return cc
}

func (cc *ConnectorConfig) Pattern(value string) *ConnectorConfig {
	cc.pattern = &value
	return cc
}

func (cc *ConnectorConfig) PAT(value string) *ConnectorConfig {
	cc.pat = &value
	return cc
}

func (cc *ConnectorConfig) FileType(value string) *ConnectorConfig {
	cc.fileType = &value
	return cc
}

func (cc *ConnectorConfig) Compression(value string) *ConnectorConfig {
	cc.compression = &value
	return cc
}

func (cc *ConnectorConfig) OnError(value string) *ConnectorConfig {
	cc.onError = &value
	return cc
}

func (cc *ConnectorConfig) AppendFileOption(value string) *ConnectorConfig {
	cc.appendFileOption = &value
	return cc
}

func (cc *ConnectorConfig) ArchivePattern(value string) *ConnectorConfig {
	cc.archivePattern = &value
	return cc
}

func (cc *ConnectorConfig) NullSequence(value string) *ConnectorConfig {
	cc.nullSequence = &value
	return cc
}

func (cc *ConnectorConfig) Delimiter(value string) *ConnectorConfig {
	cc.delimiter = &value
	return cc
}

func (cc *ConnectorConfig) EscapeChar(value string) *ConnectorConfig {
	cc.escapeChar = &value
	return cc
}

func (cc *ConnectorConfig) SkipBefore(value int) *ConnectorConfig {
	cc.skipBefore = &value
	return cc
}

func (cc *ConnectorConfig) SkipAfter(value int) *ConnectorConfig {
	cc.skipAfter = &value
	return cc
}

func (cc *ConnectorConfig) SecretsList(value []*FunctionSecret) *ConnectorConfig {
	cc.secretsList = value
	return cc
}

func (cc *ConnectorConfig) ProjectCredentials(value []*ConnectorConfigProjectCredentials) *ConnectorConfig {
	cc.projectCredentials = value
	return cc
}

func (cc *ConnectorConfig) AuthMode(value string) *ConnectorConfig {
	cc.authMode = &value
	return cc
}

func (cc *ConnectorConfig) Username(value string) *ConnectorConfig {
	cc.username = &value
	return cc
}

func (cc *ConnectorConfig) Password(value string) *ConnectorConfig {
	cc.password = &value
	return cc
}

func (cc *ConnectorConfig) Certificate(value string) *ConnectorConfig {
	cc.certificate = &value
	return cc
}

func (cc *ConnectorConfig) SelectedExports(value []string) *ConnectorConfig {
	cc.selectedExports = value
	return cc
}

func (cc *ConnectorConfig) ConsumerGroup(value string) *ConnectorConfig {
	cc.consumerGroup = &value
	return cc
}

func (cc *ConnectorConfig) Servers(value string) *ConnectorConfig {
	cc.servers = &value
	return cc
}

func (cc *ConnectorConfig) MessageType(value string) *ConnectorConfig {
	cc.messageType = &value
	return cc
}

func (cc *ConnectorConfig) SyncType(value string) *ConnectorConfig {
	cc.syncType = &value
	return cc
}

func (cc *ConnectorConfig) SecurityProtocol(value string) *ConnectorConfig {
	cc.securityProtocol = &value
	return cc
}

func (cc *ConnectorConfig) Apps(value []string) *ConnectorConfig {
	cc.apps = value
	return cc
}

func (cc *ConnectorConfig) SalesAccounts(value []string) *ConnectorConfig {
	cc.salesAccounts = value
	return cc
}

func (cc *ConnectorConfig) FinanceAccounts(value []string) *ConnectorConfig {
	cc.financeAccounts = value
	return cc
}

func (cc *ConnectorConfig) AppSyncMode(value string) *ConnectorConfig {
	cc.appSyncMode = &value
	return cc
}

func (cc *ConnectorConfig) SalesAccountSyncMode(value string) *ConnectorConfig {
	cc.salesAccountSyncMode = &value
	return cc
}

func (cc *ConnectorConfig) FinanceAccountSyncMode(value string) *ConnectorConfig {
	cc.financeAccountSyncMode = &value
	return cc
}

func (cc *ConnectorConfig) PEMCertificate(value string) *ConnectorConfig {
	cc.pemCertificate = &value
	return cc
}

func (cc *ConnectorConfig) AccessKeyID(value string) *ConnectorConfig {
	cc.accessKeyID = &value
	return cc
}

func (cc *ConnectorConfig) SecretKey(value string) *ConnectorConfig {
	cc.secretKey = &value
	return cc
}

func (cc *ConnectorConfig) HomeFolder(value string) *ConnectorConfig {
	cc.homeFolder = &value
	return cc
}

func (cc *ConnectorConfig) SyncDataLocker(value bool) *ConnectorConfig {
	cc.syncDataLocker = &value
	return cc
}

func (cc *ConnectorConfig) Projects(value []string) *ConnectorConfig {
	cc.projects = value
	return cc
}

func (cc *ConnectorConfig) Function(value string) *ConnectorConfig {
	cc.function = &value
	return cc
}

func (cc *ConnectorConfig) Region(value string) *ConnectorConfig {
	cc.region = &value
	return cc
}

func (cc *ConnectorConfig) Secrets(value string) *ConnectorConfig {
	cc.secrets = &value
	return cc
}

func (cc *ConnectorConfig) ContainerName(value string) *ConnectorConfig {
	cc.containerName = &value
	return cc
}

func (cc *ConnectorConfig) ConnectionString(value string) *ConnectorConfig {
	cc.connectionString = &value
	return cc
}

func (cc *ConnectorConfig) FunctionApp(value string) *ConnectorConfig {
	cc.functionApp = &value
	return cc
}

func (cc *ConnectorConfig) FunctionName(value string) *ConnectorConfig {
	cc.functionName = &value
	return cc
}

func (cc *ConnectorConfig) FunctionKey(value string) *ConnectorConfig {
	cc.functionKey = &value
	return cc
}

func (cc *ConnectorConfig) PublicKey(value string) *ConnectorConfig {
	cc.publicKey = &value
	return cc
}

func (cc *ConnectorConfig) MerchantID(value string) *ConnectorConfig {
	cc.merchantID = &value
	return cc
}

func (cc *ConnectorConfig) APIURL(value string) *ConnectorConfig {
	cc.apiURL = &value
	return cc
}

func (cc *ConnectorConfig) CloudStorageType(value string) *ConnectorConfig {
	cc.cloudStorageType = &value
	return cc
}

func (cc *ConnectorConfig) S3ExternalID(value string) *ConnectorConfig {
	cc.s3ExternalID = &value
	return cc
}

func (cc *ConnectorConfig) S3Folder(value string) *ConnectorConfig {
	cc.s3Folder = &value
	return cc
}

func (cc *ConnectorConfig) GCSBucket(value string) *ConnectorConfig {
	cc.gcsBucket = &value
	return cc
}

func (cc *ConnectorConfig) GCSFolder(value string) *ConnectorConfig {
	cc.gcsFolder = &value
	return cc
}

func (cc *ConnectorConfig) UserProfiles(value []string) *ConnectorConfig {
	cc.userProfiles = value
	return cc
}

func (cc *ConnectorConfig) ReportConfigurationIDs(value []string) *ConnectorConfig {
	cc.reportConfigurationIDs = value
	return cc
}

func (cc *ConnectorConfig) EnableAllDimensionCombinations(value bool) *ConnectorConfig {
	cc.enableAllDimensionCombinations = &value
	return cc
}

func (cc *ConnectorConfig) Instance(value string) *ConnectorConfig {
	cc.instance = &value
	return cc
}

func (cc *ConnectorConfig) AWSRegionCode(value string) *ConnectorConfig {
	cc.awsRegionCode = &value
	return cc
}

func (cc *ConnectorConfig) Accounts(value []string) *ConnectorConfig {
	cc.accounts = value
	return cc
}

func (cc *ConnectorConfig) Fields(value []string) *ConnectorConfig {
	cc.fields = value
	return cc
}

func (cc *ConnectorConfig) Breakdowns(value []string) *ConnectorConfig {
	cc.breakdowns = value
	return cc
}

func (cc *ConnectorConfig) ActionBreakdowns(value []string) *ConnectorConfig {
	cc.actionBreakdowns = value
	return cc
}

func (cc *ConnectorConfig) Aggregation(value string) *ConnectorConfig {
	cc.aggregation = &value
	return cc
}

func (cc *ConnectorConfig) ConfigType(value string) *ConnectorConfig {
	cc.configType = &value
	return cc
}

func (cc *ConnectorConfig) PrebuiltReport(value string) *ConnectorConfig {
	cc.prebuiltReport = &value
	return cc
}

func (cc *ConnectorConfig) ActionReportTime(value string) *ConnectorConfig {
	cc.actionReportTime = &value
	return cc
}

func (cc *ConnectorConfig) ClickAttributionWindow(value string) *ConnectorConfig {
	cc.clickAttributionWindow = &value
	return cc
}

func (cc *ConnectorConfig) ViewAttributionWindow(value string) *ConnectorConfig {
	cc.viewAttributionWindow = &value
	return cc
}

func (cc *ConnectorConfig) CustomTables(value []*ConnectorConfigCustomTables) *ConnectorConfig {
	cc.customTables = value
	return cc
}

func (cc *ConnectorConfig) Pages(value []string) *ConnectorConfig {
	cc.pages = value
	return cc
}

func (cc *ConnectorConfig) Subdomain(value string) *ConnectorConfig {
	cc.subdomain = &value
	return cc
}

func (cc *ConnectorConfig) Host(value string) *ConnectorConfig {
	cc.host = &value
	return cc
}

func (cc *ConnectorConfig) Port(value int) *ConnectorConfig {
	cc.port = &value
	return cc
}

func (cc *ConnectorConfig) User(value string) *ConnectorConfig {
	cc.user = &value
	return cc
}

func (cc *ConnectorConfig) IsSecure(value bool) *ConnectorConfig {
	cc.isSecure = &value
	return cc
}

func (cc *ConnectorConfig) Repositories(value []string) *ConnectorConfig {
	cc.repositories = value
	return cc
}

func (cc *ConnectorConfig) UseWebhooks(value bool) *ConnectorConfig {
	cc.useWebhooks = &value
	return cc
}

func (cc *ConnectorConfig) DimensionAttributes(value []string) *ConnectorConfig {
	cc.dimensionAttributes = value
	return cc
}

func (cc *ConnectorConfig) Columns(value []string) *ConnectorConfig {
	cc.columns = value
	return cc
}

func (cc *ConnectorConfig) NetworkCode(value string) *ConnectorConfig {
	cc.networkCode = &value
	return cc
}

func (cc *ConnectorConfig) CustomerID(value string) *ConnectorConfig {
	cc.customerID = &value
	return cc
}

func (cc *ConnectorConfig) ManagerAccounts(value []string) *ConnectorConfig {
	cc.managerAccounts = value
	return cc
}

func (cc *ConnectorConfig) Reports(value []*ConnectorConfigReports) *ConnectorConfig {
	cc.reports = value
	return cc
}

func (cc *ConnectorConfig) AdobeAnalyticsConfigurations(value []*ConnectorConfigAdobeAnalyticsConfiguration) *ConnectorConfig {
	cc.adobeAnalyticsConfigurations = value
	return cc
}

func (cc *ConnectorConfig) ConversionWindowSize(value int) *ConnectorConfig {
	cc.conversionWindowSize = &value
	return cc
}

func (cc *ConnectorConfig) Profiles(value []string) *ConnectorConfig {
	cc.profiles = value
	return cc
}

func (cc *ConnectorConfig) ProjectID(value string) *ConnectorConfig {
	cc.projectID = &value
	return cc
}

func (cc *ConnectorConfig) DatasetID(value string) *ConnectorConfig {
	cc.datasetID = &value
	return cc
}

func (cc *ConnectorConfig) BucketName(value string) *ConnectorConfig {
	cc.bucketName = &value
	return cc
}

func (cc *ConnectorConfig) FunctionTrigger(value string) *ConnectorConfig {
	cc.functionTrigger = &value
	return cc
}

func (cc *ConnectorConfig) ConfigMethod(value string) *ConnectorConfig {
	cc.configMethod = &value
	return cc
}

func (cc *ConnectorConfig) QueryID(value string) *ConnectorConfig {
	cc.queryID = &value
	return cc
}

func (cc *ConnectorConfig) UpdateConfigOnEachSync(value bool) *ConnectorConfig {
	cc.updateConfigOnEachSync = &value
	return cc
}

func (cc *ConnectorConfig) SiteURLs(value []string) *ConnectorConfig {
	cc.siteURLs = value
	return cc
}

func (cc *ConnectorConfig) Path(value string) *ConnectorConfig {
	cc.path = &value
	return cc
}

func (cc *ConnectorConfig) OnPremise(value bool) *ConnectorConfig {
	cc.onPremise = &value
	return cc
}

func (cc *ConnectorConfig) AccessToken(value string) *ConnectorConfig {
	cc.accessToken = &value
	return cc
}

func (cc *ConnectorConfig) ViewThroughAttributionWindowSize(value string) *ConnectorConfig {
	cc.viewThroughAttributionWindowSize = &value
	return cc
}

func (cc *ConnectorConfig) PostClickAttributionWindowSize(value string) *ConnectorConfig {
	cc.postClickAttributionWindowSize = &value
	return cc
}

func (cc *ConnectorConfig) UseAPIKeys(value bool) *ConnectorConfig {
	cc.useAPIKeys = &value
	return cc
}

func (cc *ConnectorConfig) APIKeys(value []string) *ConnectorConfig {
	cc.apiKeys = value
	return cc
}

func (cc *ConnectorConfig) Endpoint(value string) *ConnectorConfig {
	cc.endpoint = &value
	return cc
}

func (cc *ConnectorConfig) Identity(value string) *ConnectorConfig {
	cc.identity = &value
	return cc
}

func (cc *ConnectorConfig) APIQuota(value int) *ConnectorConfig {
	cc.apiQuota = &value
	return cc
}

func (cc *ConnectorConfig) DomainName(value string) *ConnectorConfig {
	cc.domainName = &value
	return cc
}

func (cc *ConnectorConfig) ResourceURL(value string) *ConnectorConfig {
	cc.resourceURL = &value
	return cc
}

func (cc *ConnectorConfig) APISecret(value string) *ConnectorConfig {
	cc.apiSecret = &value
	return cc
}

func (cc *ConnectorConfig) Hosts(value []string) *ConnectorConfig {
	cc.hosts = value
	return cc
}

func (cc *ConnectorConfig) TunnelHost(value string) *ConnectorConfig {
	cc.tunnelHost = &value
	return cc
}

func (cc *ConnectorConfig) TunnelPort(value int) *ConnectorConfig {
	cc.tunnelPort = &value
	return cc
}

func (cc *ConnectorConfig) TunnelUser(value string) *ConnectorConfig {
	cc.tunnelUser = &value
	return cc
}

func (cc *ConnectorConfig) Database(value string) *ConnectorConfig {
	cc.database = &value
	return cc
}

func (cc *ConnectorConfig) Datasource(value string) *ConnectorConfig {
	cc.datasource = &value
	return cc
}

func (cc *ConnectorConfig) Account(value string) *ConnectorConfig {
	cc.account = &value
	return cc
}

func (cc *ConnectorConfig) Role(value string) *ConnectorConfig {
	cc.role = &value
	return cc
}

func (cc *ConnectorConfig) Email(value string) *ConnectorConfig {
	cc.email = &value
	return cc
}

func (cc *ConnectorConfig) AccountID(value string) *ConnectorConfig {
	cc.accountID = &value
	return cc
}

func (cc *ConnectorConfig) ServerURL(value string) *ConnectorConfig {
	cc.serverURL = &value
	return cc
}

func (cc *ConnectorConfig) UserKey(value string) *ConnectorConfig {
	cc.userKey = &value
	return cc
}

func (cc *ConnectorConfig) APIVersion(value string) *ConnectorConfig {
	cc.apiVersion = &value
	return cc
}

func (cc *ConnectorConfig) DailyAPICallLimit(value int) *ConnectorConfig {
	cc.dailyAPICallLimit = &value
	return cc
}

func (cc *ConnectorConfig) TimeZone(value string) *ConnectorConfig {
	cc.timeZone = &value
	return cc
}

func (cc *ConnectorConfig) IntegrationKey(value string) *ConnectorConfig {
	cc.integrationKey = &value
	return cc
}

func (cc *ConnectorConfig) Advertisers(value []string) *ConnectorConfig {
	cc.advertisers = value
	return cc
}

func (cc *ConnectorConfig) EngagementAttributionWindow(value string) *ConnectorConfig {
	cc.engagementAttributionWindow = &value
	return cc
}

func (cc *ConnectorConfig) ConversionReportTime(value string) *ConnectorConfig {
	cc.conversionReportTime = &value
	return cc
}

func (cc *ConnectorConfig) Domain(value string) *ConnectorConfig {
	cc.domain = &value
	return cc
}

func (cc *ConnectorConfig) UpdateMethod(value string) *ConnectorConfig {
	cc.updateMethod = &value
	return cc
}

func (cc *ConnectorConfig) ReplicationSlot(value string) *ConnectorConfig {
	cc.replicationSlot = &value
	return cc
}

func (cc *ConnectorConfig) PublicationName(value string) *ConnectorConfig {
	cc.publicationName = &value
	return cc
}

func (cc *ConnectorConfig) DataCenter(value string) *ConnectorConfig {
	cc.dataCenter = &value
	return cc
}

func (cc *ConnectorConfig) APIToken(value string) *ConnectorConfig {
	cc.apiToken = &value
	return cc
}

func (cc *ConnectorConfig) SubDomain(value string) *ConnectorConfig {
	cc.subDomain = &value
	return cc
}

func (cc *ConnectorConfig) TestTableName(value string) *ConnectorConfig {
	cc.testTableName = &value
	return cc
}

func (cc *ConnectorConfig) Shop(value string) *ConnectorConfig {
	cc.shop = &value
	return cc
}

func (cc *ConnectorConfig) Organizations(value []string) *ConnectorConfig {
	cc.organizations = value
	return cc
}

func (cc *ConnectorConfig) SwipeAttributionWindow(value string) *ConnectorConfig {
	cc.swipeAttributionWindow = &value
	return cc
}

func (cc *ConnectorConfig) APIAccessToken(value string) *ConnectorConfig {
	cc.apiAccessToken = &value
	return cc
}

func (cc *ConnectorConfig) AccountIDs(value []string) *ConnectorConfig {
	cc.accountIDs = value
	return cc
}

func (cc *ConnectorConfig) SID(value string) *ConnectorConfig {
	cc.sid = &value
	return cc
}

func (cc *ConnectorConfig) Secret(value string) *ConnectorConfig {
	cc.secret = &value
	return cc
}

func (cc *ConnectorConfig) OauthToken(value string) *ConnectorConfig {
	cc.oauthToken = &value
	return cc
}

func (cc *ConnectorConfig) OauthTokenSecret(value string) *ConnectorConfig {
	cc.oauthTokenSecret = &value
	return cc
}

func (cc *ConnectorConfig) ConsumerKey(value string) *ConnectorConfig {
	cc.consumerKey = &value
	return cc
}

func (cc *ConnectorConfig) ConsumerSecret(value string) *ConnectorConfig {
	cc.consumerSecret = &value
	return cc
}

func (cc *ConnectorConfig) Key(value string) *ConnectorConfig {
	cc.key = &value
	return cc
}

func (cc *ConnectorConfig) AdvertisersID(value []string) *ConnectorConfig {
	cc.advertisersID = value
	return cc
}

func (cc *ConnectorConfig) SyncFormat(value string) *ConnectorConfig {
	cc.syncFormat = &value
	return cc
}

func (cc *ConnectorConfig) BucketService(value string) *ConnectorConfig {
	cc.bucketService = &value
	return cc
}

func (cc *ConnectorConfig) UserName(value string) *ConnectorConfig {
	cc.userName = &value
	return cc
}

func (cc *ConnectorConfig) ReportURL(value string) *ConnectorConfig {
	cc.reportURL = &value
	return cc
}

func (cc *ConnectorConfig) UniqueID(value string) *ConnectorConfig {
	cc.uniqueID = &value
	return cc
}

func (cc *ConnectorConfig) AuthType(value string) *ConnectorConfig {
	cc.authType = &value
	return cc
}

func (cc *ConnectorConfig) IsNewPackage(value bool) *ConnectorConfig {
	cc.isNewPackage = &value
	return cc
}

func (cc *ConnectorConfig) ConnectionType(value string) *ConnectorConfig {
	cc.connectionType = &value
	return cc
}

func (cc *ConnectorConfig) IsMultiEntityFeatureEnabled(value bool) *ConnectorConfig {
	cc.isMultiEntityFeatureEnabled = &value
	return cc
}

func (cc *ConnectorConfig) AlwaysEncrypted(value bool) *ConnectorConfig {
	cc.alwaysEncrypted = &value
	return cc
}

func (cc *ConnectorConfig) ApiType(value string) *ConnectorConfig {
	cc.apiType = &value
	return cc
}

func (cc *ConnectorConfig) BaseUrl(value string) *ConnectorConfig {
	cc.baseUrl = &value
	return cc
}

func (cc *ConnectorConfig) EntityId(value string) *ConnectorConfig {
	cc.entityId = &value
	return cc
}

func (cc *ConnectorConfig) SoapUri(value string) *ConnectorConfig {
	cc.soapUri = &value
	return cc
}

func (cc *ConnectorConfig) UserId(value string) *ConnectorConfig {
	cc.userId = &value
	return cc
}

func (cc *ConnectorConfig) EncryptionKey(value string) *ConnectorConfig {
	cc.encryptionKey = &value
	return cc
}

func (cc *ConnectorConfig) EuRegion(value bool) *ConnectorConfig {
	cc.euRegion = &value
	return cc
}

func (cc *ConnectorConfig) TokenKey(value string) *ConnectorConfig {
	cc.tokenKey = &value
	return cc
}

func (cc *ConnectorConfig) TokenSecret(value string) *ConnectorConfig {
	cc.tokenSecret = &value
	return cc
}

func (cc *ConnectorConfig) ShareURL(value string) *ConnectorConfig {
	cc.shareURL = &value
	return cc
}
