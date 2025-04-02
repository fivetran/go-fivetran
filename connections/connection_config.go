package connections

import "github.com/fivetran/go-fivetran/utils"

type ConnectionConfig struct {
	adobeAnalyticsConfigurations     []*ConnectionConfigAdobeAnalyticsConfiguration
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
	projectCredentials               []*ConnectionConfigProjectCredentials
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
	customTables                     []*ConnectionConfigCustomTables
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
	reports                          []*ConnectionConfigReports
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

type connectionConfigRequest struct {
	AdobeAnalyticsConfigurations     []*connectionConfigAdobeAnalyticsConfigurationRequest `json:"adobe_analytics_configurations,omitempty"`
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
	ProjectCredentials               []*connectionConfigProjectCredentialsRequest          `json:"project_credentials,omitempty"`
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
	CustomTables                     []*connectionConfigCustomTablesRequest                `json:"custom_tables,omitempty"`
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
	Reports                          []*connectionConfigReportsRequest                     `json:"reports,omitempty"`
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

type ConnectionConfigResponse struct {
	AdobeAnalyticsConfigurations     []ConnectionConfigAdobeAnalyticsConfigurationResponse `json:"adobe_analytics_configurations"`
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
	ProjectCredentials               []ConnectionConfigProjectCredentialsResponse          `json:"project_credentials"`
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
	CustomTables                     []ConnectionConfigCustomTablesResponse                `json:"custom_tables"`
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
	Reports                          []ConnectionConfigReportsResponse                     `json:"reports"`
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

func (cc *ConnectionConfig) Merge(customConfig *map[string]interface{}) (*map[string]interface{}, error) {
	err := utils.MergeIntoMap(cc.Request(), customConfig)
	if err != nil {
		return nil, err
	}
	return customConfig, nil
}

func (cc *ConnectionConfig) Request() *connectionConfigRequest {
	var projectCredentials []*connectionConfigProjectCredentialsRequest
	if cc.projectCredentials != nil {
		for _, pc := range cc.projectCredentials {
			projectCredentials = append(projectCredentials, pc.request())
		}
	}

	var customTables []*connectionConfigCustomTablesRequest
	if cc.customTables != nil {
		for _, ct := range cc.customTables {
			customTables = append(customTables, ct.request())
		}
	}

	var reports []*connectionConfigReportsRequest
	if cc.reports != nil {
		for _, r := range cc.reports {
			reports = append(reports, r.request())
		}
	}

	var adobeAnalyticsConfigurations []*connectionConfigAdobeAnalyticsConfigurationRequest
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

	return &connectionConfigRequest{
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

func (cc *ConnectionConfig) Schema(value string) *ConnectionConfig {
	cc.schema = &value
	return cc
}

func (cc *ConnectionConfig) Table(value string) *ConnectionConfig {
	cc.table = &value
	return cc
}

func (cc *ConnectionConfig) SheetID(value string) *ConnectionConfig {
	cc.sheetID = &value
	return cc
}

func (cc *ConnectionConfig) NamedRange(value string) *ConnectionConfig {
	cc.namedRange = &value
	return cc
}

func (cc *ConnectionConfig) ClientID(value string) *ConnectionConfig {
	cc.clientID = &value
	return cc
}

func (cc *ConnectionConfig) ClientSecret(value string) *ConnectionConfig {
	cc.clientSecret = &value
	return cc
}

func (cc *ConnectionConfig) TechnicalAccountID(value string) *ConnectionConfig {
	cc.technicalAccountID = &value
	return cc
}

func (cc *ConnectionConfig) OrganizationID(value string) *ConnectionConfig {
	cc.organizationID = &value
	return cc
}

func (cc *ConnectionConfig) PrivateKey(value string) *ConnectionConfig {
	cc.privateKey = &value
	return cc
}

func (cc *ConnectionConfig) SyncMode(value string) *ConnectionConfig {
	cc.syncMode = &value
	return cc
}

func (cc *ConnectionConfig) ReportSuites(value []string) *ConnectionConfig {
	cc.reportSuites = value
	return cc
}

func (cc *ConnectionConfig) Elements(value []string) *ConnectionConfig {
	cc.elements = value
	return cc
}

func (cc *ConnectionConfig) Metrics(value []string) *ConnectionConfig {
	cc.metrics = value
	return cc
}

func (cc *ConnectionConfig) DateGranularity(value string) *ConnectionConfig {
	cc.dateGranularity = &value
	return cc
}

func (cc *ConnectionConfig) TimeframeMonths(value string) *ConnectionConfig {
	cc.timeframeMonths = &value
	return cc
}

func (cc *ConnectionConfig) Source(value string) *ConnectionConfig {
	cc.source = &value
	return cc
}

func (cc *ConnectionConfig) S3Bucket(value string) *ConnectionConfig {
	cc.s3Bucket = &value
	return cc
}

func (cc *ConnectionConfig) S3RoleArn(value string) *ConnectionConfig {
	cc.s3RoleArn = &value
	return cc
}

func (cc *ConnectionConfig) ABSConnectionString(value string) *ConnectionConfig {
	cc.absConnectionString = &value
	return cc
}

func (cc *ConnectionConfig) ABSContainerName(value string) *ConnectionConfig {
	cc.absContainerName = &value
	return cc
}

func (cc *ConnectionConfig) FolderId(value string) *ConnectionConfig {
	cc.folderId = &value
	return cc
}

func (cc *ConnectionConfig) FTPHost(value string) *ConnectionConfig {
	cc.ftpHost = &value
	return cc
}

func (cc *ConnectionConfig) FTPPort(value int) *ConnectionConfig {
	cc.ftpPort = &value
	return cc
}

func (cc *ConnectionConfig) FTPUser(value string) *ConnectionConfig {
	cc.ftpUser = &value
	return cc
}

func (cc *ConnectionConfig) FTPPassword(value string) *ConnectionConfig {
	cc.ftpPassword = &value
	return cc
}

func (cc *ConnectionConfig) IsFTPS(value bool) *ConnectionConfig {
	cc.isFTPS = &value
	return cc
}

func (cc *ConnectionConfig) SFTPHost(value string) *ConnectionConfig {
	cc.sFTPHost = &value
	return cc
}

func (cc *ConnectionConfig) SFTPPort(value int) *ConnectionConfig {
	cc.sFTPPort = &value
	return cc
}

func (cc *ConnectionConfig) SFTPUser(value string) *ConnectionConfig {
	cc.sFTPUser = &value
	return cc
}

func (cc *ConnectionConfig) SFTPPassword(value string) *ConnectionConfig {
	cc.sFTPPassword = &value
	return cc
}

func (cc *ConnectionConfig) SFTPIsKeyPair(value bool) *ConnectionConfig {
	cc.sFTPIsKeyPair = &value
	return cc
}

func (cc *ConnectionConfig) IsKeypair(value bool) *ConnectionConfig {
	cc.isKeypair = &value
	return cc
}

func (cc *ConnectionConfig) Advertisables(value []string) *ConnectionConfig {
	cc.advertisables = value
	return cc
}

func (cc *ConnectionConfig) ReportType(value string) *ConnectionConfig {
	cc.reportType = &value
	return cc
}

func (cc *ConnectionConfig) Dimensions(value []string) *ConnectionConfig {
	cc.dimensions = value
	return cc
}

func (cc *ConnectionConfig) SchemaPrefix(value string) *ConnectionConfig {
	cc.schemaPrefix = &value
	return cc
}

func (cc *ConnectionConfig) APIKey(value string) *ConnectionConfig {
	cc.apiKey = &value
	return cc
}

func (cc *ConnectionConfig) ExternalID(value string) *ConnectionConfig {
	cc.externalID = &value
	return cc
}

func (cc *ConnectionConfig) RoleArn(value string) *ConnectionConfig {
	cc.roleArn = &value
	return cc
}

func (cc *ConnectionConfig) Bucket(value string) *ConnectionConfig {
	cc.bucket = &value
	return cc
}

func (cc *ConnectionConfig) Prefix(value string) *ConnectionConfig {
	cc.prefix = &value
	return cc
}

func (cc *ConnectionConfig) Pattern(value string) *ConnectionConfig {
	cc.pattern = &value
	return cc
}

func (cc *ConnectionConfig) PAT(value string) *ConnectionConfig {
	cc.pat = &value
	return cc
}

func (cc *ConnectionConfig) FileType(value string) *ConnectionConfig {
	cc.fileType = &value
	return cc
}

func (cc *ConnectionConfig) Compression(value string) *ConnectionConfig {
	cc.compression = &value
	return cc
}

func (cc *ConnectionConfig) OnError(value string) *ConnectionConfig {
	cc.onError = &value
	return cc
}

func (cc *ConnectionConfig) AppendFileOption(value string) *ConnectionConfig {
	cc.appendFileOption = &value
	return cc
}

func (cc *ConnectionConfig) ArchivePattern(value string) *ConnectionConfig {
	cc.archivePattern = &value
	return cc
}

func (cc *ConnectionConfig) NullSequence(value string) *ConnectionConfig {
	cc.nullSequence = &value
	return cc
}

func (cc *ConnectionConfig) Delimiter(value string) *ConnectionConfig {
	cc.delimiter = &value
	return cc
}

func (cc *ConnectionConfig) EscapeChar(value string) *ConnectionConfig {
	cc.escapeChar = &value
	return cc
}

func (cc *ConnectionConfig) SkipBefore(value int) *ConnectionConfig {
	cc.skipBefore = &value
	return cc
}

func (cc *ConnectionConfig) SkipAfter(value int) *ConnectionConfig {
	cc.skipAfter = &value
	return cc
}

func (cc *ConnectionConfig) SecretsList(value []*FunctionSecret) *ConnectionConfig {
	cc.secretsList = value
	return cc
}

func (cc *ConnectionConfig) ProjectCredentials(value []*ConnectionConfigProjectCredentials) *ConnectionConfig {
	cc.projectCredentials = value
	return cc
}

func (cc *ConnectionConfig) AuthMode(value string) *ConnectionConfig {
	cc.authMode = &value
	return cc
}

func (cc *ConnectionConfig) Username(value string) *ConnectionConfig {
	cc.username = &value
	return cc
}

func (cc *ConnectionConfig) Password(value string) *ConnectionConfig {
	cc.password = &value
	return cc
}

func (cc *ConnectionConfig) Certificate(value string) *ConnectionConfig {
	cc.certificate = &value
	return cc
}

func (cc *ConnectionConfig) SelectedExports(value []string) *ConnectionConfig {
	cc.selectedExports = value
	return cc
}

func (cc *ConnectionConfig) ConsumerGroup(value string) *ConnectionConfig {
	cc.consumerGroup = &value
	return cc
}

func (cc *ConnectionConfig) Servers(value string) *ConnectionConfig {
	cc.servers = &value
	return cc
}

func (cc *ConnectionConfig) MessageType(value string) *ConnectionConfig {
	cc.messageType = &value
	return cc
}

func (cc *ConnectionConfig) SyncType(value string) *ConnectionConfig {
	cc.syncType = &value
	return cc
}

func (cc *ConnectionConfig) SecurityProtocol(value string) *ConnectionConfig {
	cc.securityProtocol = &value
	return cc
}

func (cc *ConnectionConfig) Apps(value []string) *ConnectionConfig {
	cc.apps = value
	return cc
}

func (cc *ConnectionConfig) SalesAccounts(value []string) *ConnectionConfig {
	cc.salesAccounts = value
	return cc
}

func (cc *ConnectionConfig) FinanceAccounts(value []string) *ConnectionConfig {
	cc.financeAccounts = value
	return cc
}

func (cc *ConnectionConfig) AppSyncMode(value string) *ConnectionConfig {
	cc.appSyncMode = &value
	return cc
}

func (cc *ConnectionConfig) SalesAccountSyncMode(value string) *ConnectionConfig {
	cc.salesAccountSyncMode = &value
	return cc
}

func (cc *ConnectionConfig) FinanceAccountSyncMode(value string) *ConnectionConfig {
	cc.financeAccountSyncMode = &value
	return cc
}

func (cc *ConnectionConfig) PEMCertificate(value string) *ConnectionConfig {
	cc.pemCertificate = &value
	return cc
}

func (cc *ConnectionConfig) AccessKeyID(value string) *ConnectionConfig {
	cc.accessKeyID = &value
	return cc
}

func (cc *ConnectionConfig) SecretKey(value string) *ConnectionConfig {
	cc.secretKey = &value
	return cc
}

func (cc *ConnectionConfig) HomeFolder(value string) *ConnectionConfig {
	cc.homeFolder = &value
	return cc
}

func (cc *ConnectionConfig) SyncDataLocker(value bool) *ConnectionConfig {
	cc.syncDataLocker = &value
	return cc
}

func (cc *ConnectionConfig) Projects(value []string) *ConnectionConfig {
	cc.projects = value
	return cc
}

func (cc *ConnectionConfig) Function(value string) *ConnectionConfig {
	cc.function = &value
	return cc
}

func (cc *ConnectionConfig) Region(value string) *ConnectionConfig {
	cc.region = &value
	return cc
}

func (cc *ConnectionConfig) Secrets(value string) *ConnectionConfig {
	cc.secrets = &value
	return cc
}

func (cc *ConnectionConfig) ContainerName(value string) *ConnectionConfig {
	cc.containerName = &value
	return cc
}

func (cc *ConnectionConfig) ConnectionString(value string) *ConnectionConfig {
	cc.connectionString = &value
	return cc
}

func (cc *ConnectionConfig) FunctionApp(value string) *ConnectionConfig {
	cc.functionApp = &value
	return cc
}

func (cc *ConnectionConfig) FunctionName(value string) *ConnectionConfig {
	cc.functionName = &value
	return cc
}

func (cc *ConnectionConfig) FunctionKey(value string) *ConnectionConfig {
	cc.functionKey = &value
	return cc
}

func (cc *ConnectionConfig) PublicKey(value string) *ConnectionConfig {
	cc.publicKey = &value
	return cc
}

func (cc *ConnectionConfig) MerchantID(value string) *ConnectionConfig {
	cc.merchantID = &value
	return cc
}

func (cc *ConnectionConfig) APIURL(value string) *ConnectionConfig {
	cc.apiURL = &value
	return cc
}

func (cc *ConnectionConfig) CloudStorageType(value string) *ConnectionConfig {
	cc.cloudStorageType = &value
	return cc
}

func (cc *ConnectionConfig) S3ExternalID(value string) *ConnectionConfig {
	cc.s3ExternalID = &value
	return cc
}

func (cc *ConnectionConfig) S3Folder(value string) *ConnectionConfig {
	cc.s3Folder = &value
	return cc
}

func (cc *ConnectionConfig) GCSBucket(value string) *ConnectionConfig {
	cc.gcsBucket = &value
	return cc
}

func (cc *ConnectionConfig) GCSFolder(value string) *ConnectionConfig {
	cc.gcsFolder = &value
	return cc
}

func (cc *ConnectionConfig) UserProfiles(value []string) *ConnectionConfig {
	cc.userProfiles = value
	return cc
}

func (cc *ConnectionConfig) ReportConfigurationIDs(value []string) *ConnectionConfig {
	cc.reportConfigurationIDs = value
	return cc
}

func (cc *ConnectionConfig) EnableAllDimensionCombinations(value bool) *ConnectionConfig {
	cc.enableAllDimensionCombinations = &value
	return cc
}

func (cc *ConnectionConfig) Instance(value string) *ConnectionConfig {
	cc.instance = &value
	return cc
}

func (cc *ConnectionConfig) AWSRegionCode(value string) *ConnectionConfig {
	cc.awsRegionCode = &value
	return cc
}

func (cc *ConnectionConfig) Accounts(value []string) *ConnectionConfig {
	cc.accounts = value
	return cc
}

func (cc *ConnectionConfig) Fields(value []string) *ConnectionConfig {
	cc.fields = value
	return cc
}

func (cc *ConnectionConfig) Breakdowns(value []string) *ConnectionConfig {
	cc.breakdowns = value
	return cc
}

func (cc *ConnectionConfig) ActionBreakdowns(value []string) *ConnectionConfig {
	cc.actionBreakdowns = value
	return cc
}

func (cc *ConnectionConfig) Aggregation(value string) *ConnectionConfig {
	cc.aggregation = &value
	return cc
}

func (cc *ConnectionConfig) ConfigType(value string) *ConnectionConfig {
	cc.configType = &value
	return cc
}

func (cc *ConnectionConfig) PrebuiltReport(value string) *ConnectionConfig {
	cc.prebuiltReport = &value
	return cc
}

func (cc *ConnectionConfig) ActionReportTime(value string) *ConnectionConfig {
	cc.actionReportTime = &value
	return cc
}

func (cc *ConnectionConfig) ClickAttributionWindow(value string) *ConnectionConfig {
	cc.clickAttributionWindow = &value
	return cc
}

func (cc *ConnectionConfig) ViewAttributionWindow(value string) *ConnectionConfig {
	cc.viewAttributionWindow = &value
	return cc
}

func (cc *ConnectionConfig) CustomTables(value []*ConnectionConfigCustomTables) *ConnectionConfig {
	cc.customTables = value
	return cc
}

func (cc *ConnectionConfig) Pages(value []string) *ConnectionConfig {
	cc.pages = value
	return cc
}

func (cc *ConnectionConfig) Subdomain(value string) *ConnectionConfig {
	cc.subdomain = &value
	return cc
}

func (cc *ConnectionConfig) Host(value string) *ConnectionConfig {
	cc.host = &value
	return cc
}

func (cc *ConnectionConfig) Port(value int) *ConnectionConfig {
	cc.port = &value
	return cc
}

func (cc *ConnectionConfig) User(value string) *ConnectionConfig {
	cc.user = &value
	return cc
}

func (cc *ConnectionConfig) IsSecure(value bool) *ConnectionConfig {
	cc.isSecure = &value
	return cc
}

func (cc *ConnectionConfig) Repositories(value []string) *ConnectionConfig {
	cc.repositories = value
	return cc
}

func (cc *ConnectionConfig) UseWebhooks(value bool) *ConnectionConfig {
	cc.useWebhooks = &value
	return cc
}

func (cc *ConnectionConfig) DimensionAttributes(value []string) *ConnectionConfig {
	cc.dimensionAttributes = value
	return cc
}

func (cc *ConnectionConfig) Columns(value []string) *ConnectionConfig {
	cc.columns = value
	return cc
}

func (cc *ConnectionConfig) NetworkCode(value string) *ConnectionConfig {
	cc.networkCode = &value
	return cc
}

func (cc *ConnectionConfig) CustomerID(value string) *ConnectionConfig {
	cc.customerID = &value
	return cc
}

func (cc *ConnectionConfig) ManagerAccounts(value []string) *ConnectionConfig {
	cc.managerAccounts = value
	return cc
}

func (cc *ConnectionConfig) Reports(value []*ConnectionConfigReports) *ConnectionConfig {
	cc.reports = value
	return cc
}

func (cc *ConnectionConfig) AdobeAnalyticsConfigurations(value []*ConnectionConfigAdobeAnalyticsConfiguration) *ConnectionConfig {
	cc.adobeAnalyticsConfigurations = value
	return cc
}

func (cc *ConnectionConfig) ConversionWindowSize(value int) *ConnectionConfig {
	cc.conversionWindowSize = &value
	return cc
}

func (cc *ConnectionConfig) Profiles(value []string) *ConnectionConfig {
	cc.profiles = value
	return cc
}

func (cc *ConnectionConfig) ProjectID(value string) *ConnectionConfig {
	cc.projectID = &value
	return cc
}

func (cc *ConnectionConfig) DatasetID(value string) *ConnectionConfig {
	cc.datasetID = &value
	return cc
}

func (cc *ConnectionConfig) BucketName(value string) *ConnectionConfig {
	cc.bucketName = &value
	return cc
}

func (cc *ConnectionConfig) FunctionTrigger(value string) *ConnectionConfig {
	cc.functionTrigger = &value
	return cc
}

func (cc *ConnectionConfig) ConfigMethod(value string) *ConnectionConfig {
	cc.configMethod = &value
	return cc
}

func (cc *ConnectionConfig) QueryID(value string) *ConnectionConfig {
	cc.queryID = &value
	return cc
}

func (cc *ConnectionConfig) UpdateConfigOnEachSync(value bool) *ConnectionConfig {
	cc.updateConfigOnEachSync = &value
	return cc
}

func (cc *ConnectionConfig) SiteURLs(value []string) *ConnectionConfig {
	cc.siteURLs = value
	return cc
}

func (cc *ConnectionConfig) Path(value string) *ConnectionConfig {
	cc.path = &value
	return cc
}

func (cc *ConnectionConfig) OnPremise(value bool) *ConnectionConfig {
	cc.onPremise = &value
	return cc
}

func (cc *ConnectionConfig) AccessToken(value string) *ConnectionConfig {
	cc.accessToken = &value
	return cc
}

func (cc *ConnectionConfig) ViewThroughAttributionWindowSize(value string) *ConnectionConfig {
	cc.viewThroughAttributionWindowSize = &value
	return cc
}

func (cc *ConnectionConfig) PostClickAttributionWindowSize(value string) *ConnectionConfig {
	cc.postClickAttributionWindowSize = &value
	return cc
}

func (cc *ConnectionConfig) UseAPIKeys(value bool) *ConnectionConfig {
	cc.useAPIKeys = &value
	return cc
}

func (cc *ConnectionConfig) APIKeys(value []string) *ConnectionConfig {
	cc.apiKeys = value
	return cc
}

func (cc *ConnectionConfig) Endpoint(value string) *ConnectionConfig {
	cc.endpoint = &value
	return cc
}

func (cc *ConnectionConfig) Identity(value string) *ConnectionConfig {
	cc.identity = &value
	return cc
}

func (cc *ConnectionConfig) APIQuota(value int) *ConnectionConfig {
	cc.apiQuota = &value
	return cc
}

func (cc *ConnectionConfig) DomainName(value string) *ConnectionConfig {
	cc.domainName = &value
	return cc
}

func (cc *ConnectionConfig) ResourceURL(value string) *ConnectionConfig {
	cc.resourceURL = &value
	return cc
}

func (cc *ConnectionConfig) APISecret(value string) *ConnectionConfig {
	cc.apiSecret = &value
	return cc
}

func (cc *ConnectionConfig) Hosts(value []string) *ConnectionConfig {
	cc.hosts = value
	return cc
}

func (cc *ConnectionConfig) TunnelHost(value string) *ConnectionConfig {
	cc.tunnelHost = &value
	return cc
}

func (cc *ConnectionConfig) TunnelPort(value int) *ConnectionConfig {
	cc.tunnelPort = &value
	return cc
}

func (cc *ConnectionConfig) TunnelUser(value string) *ConnectionConfig {
	cc.tunnelUser = &value
	return cc
}

func (cc *ConnectionConfig) Database(value string) *ConnectionConfig {
	cc.database = &value
	return cc
}

func (cc *ConnectionConfig) Datasource(value string) *ConnectionConfig {
	cc.datasource = &value
	return cc
}

func (cc *ConnectionConfig) Account(value string) *ConnectionConfig {
	cc.account = &value
	return cc
}

func (cc *ConnectionConfig) Role(value string) *ConnectionConfig {
	cc.role = &value
	return cc
}

func (cc *ConnectionConfig) Email(value string) *ConnectionConfig {
	cc.email = &value
	return cc
}

func (cc *ConnectionConfig) AccountID(value string) *ConnectionConfig {
	cc.accountID = &value
	return cc
}

func (cc *ConnectionConfig) ServerURL(value string) *ConnectionConfig {
	cc.serverURL = &value
	return cc
}

func (cc *ConnectionConfig) UserKey(value string) *ConnectionConfig {
	cc.userKey = &value
	return cc
}

func (cc *ConnectionConfig) APIVersion(value string) *ConnectionConfig {
	cc.apiVersion = &value
	return cc
}

func (cc *ConnectionConfig) DailyAPICallLimit(value int) *ConnectionConfig {
	cc.dailyAPICallLimit = &value
	return cc
}

func (cc *ConnectionConfig) TimeZone(value string) *ConnectionConfig {
	cc.timeZone = &value
	return cc
}

func (cc *ConnectionConfig) IntegrationKey(value string) *ConnectionConfig {
	cc.integrationKey = &value
	return cc
}

func (cc *ConnectionConfig) Advertisers(value []string) *ConnectionConfig {
	cc.advertisers = value
	return cc
}

func (cc *ConnectionConfig) EngagementAttributionWindow(value string) *ConnectionConfig {
	cc.engagementAttributionWindow = &value
	return cc
}

func (cc *ConnectionConfig) ConversionReportTime(value string) *ConnectionConfig {
	cc.conversionReportTime = &value
	return cc
}

func (cc *ConnectionConfig) Domain(value string) *ConnectionConfig {
	cc.domain = &value
	return cc
}

func (cc *ConnectionConfig) UpdateMethod(value string) *ConnectionConfig {
	cc.updateMethod = &value
	return cc
}

func (cc *ConnectionConfig) ReplicationSlot(value string) *ConnectionConfig {
	cc.replicationSlot = &value
	return cc
}

func (cc *ConnectionConfig) PublicationName(value string) *ConnectionConfig {
	cc.publicationName = &value
	return cc
}

func (cc *ConnectionConfig) DataCenter(value string) *ConnectionConfig {
	cc.dataCenter = &value
	return cc
}

func (cc *ConnectionConfig) APIToken(value string) *ConnectionConfig {
	cc.apiToken = &value
	return cc
}

func (cc *ConnectionConfig) SubDomain(value string) *ConnectionConfig {
	cc.subDomain = &value
	return cc
}

func (cc *ConnectionConfig) TestTableName(value string) *ConnectionConfig {
	cc.testTableName = &value
	return cc
}

func (cc *ConnectionConfig) Shop(value string) *ConnectionConfig {
	cc.shop = &value
	return cc
}

func (cc *ConnectionConfig) Organizations(value []string) *ConnectionConfig {
	cc.organizations = value
	return cc
}

func (cc *ConnectionConfig) SwipeAttributionWindow(value string) *ConnectionConfig {
	cc.swipeAttributionWindow = &value
	return cc
}

func (cc *ConnectionConfig) APIAccessToken(value string) *ConnectionConfig {
	cc.apiAccessToken = &value
	return cc
}

func (cc *ConnectionConfig) AccountIDs(value []string) *ConnectionConfig {
	cc.accountIDs = value
	return cc
}

func (cc *ConnectionConfig) SID(value string) *ConnectionConfig {
	cc.sid = &value
	return cc
}

func (cc *ConnectionConfig) Secret(value string) *ConnectionConfig {
	cc.secret = &value
	return cc
}

func (cc *ConnectionConfig) OauthToken(value string) *ConnectionConfig {
	cc.oauthToken = &value
	return cc
}

func (cc *ConnectionConfig) OauthTokenSecret(value string) *ConnectionConfig {
	cc.oauthTokenSecret = &value
	return cc
}

func (cc *ConnectionConfig) ConsumerKey(value string) *ConnectionConfig {
	cc.consumerKey = &value
	return cc
}

func (cc *ConnectionConfig) ConsumerSecret(value string) *ConnectionConfig {
	cc.consumerSecret = &value
	return cc
}

func (cc *ConnectionConfig) Key(value string) *ConnectionConfig {
	cc.key = &value
	return cc
}

func (cc *ConnectionConfig) AdvertisersID(value []string) *ConnectionConfig {
	cc.advertisersID = value
	return cc
}

func (cc *ConnectionConfig) SyncFormat(value string) *ConnectionConfig {
	cc.syncFormat = &value
	return cc
}

func (cc *ConnectionConfig) BucketService(value string) *ConnectionConfig {
	cc.bucketService = &value
	return cc
}

func (cc *ConnectionConfig) UserName(value string) *ConnectionConfig {
	cc.userName = &value
	return cc
}

func (cc *ConnectionConfig) ReportURL(value string) *ConnectionConfig {
	cc.reportURL = &value
	return cc
}

func (cc *ConnectionConfig) UniqueID(value string) *ConnectionConfig {
	cc.uniqueID = &value
	return cc
}

func (cc *ConnectionConfig) AuthType(value string) *ConnectionConfig {
	cc.authType = &value
	return cc
}

func (cc *ConnectionConfig) IsNewPackage(value bool) *ConnectionConfig {
	cc.isNewPackage = &value
	return cc
}

func (cc *ConnectionConfig) ConnectionType(value string) *ConnectionConfig {
	cc.connectionType = &value
	return cc
}

func (cc *ConnectionConfig) IsMultiEntityFeatureEnabled(value bool) *ConnectionConfig {
	cc.isMultiEntityFeatureEnabled = &value
	return cc
}

func (cc *ConnectionConfig) AlwaysEncrypted(value bool) *ConnectionConfig {
	cc.alwaysEncrypted = &value
	return cc
}

func (cc *ConnectionConfig) ApiType(value string) *ConnectionConfig {
	cc.apiType = &value
	return cc
}

func (cc *ConnectionConfig) BaseUrl(value string) *ConnectionConfig {
	cc.baseUrl = &value
	return cc
}

func (cc *ConnectionConfig) EntityId(value string) *ConnectionConfig {
	cc.entityId = &value
	return cc
}

func (cc *ConnectionConfig) SoapUri(value string) *ConnectionConfig {
	cc.soapUri = &value
	return cc
}

func (cc *ConnectionConfig) UserId(value string) *ConnectionConfig {
	cc.userId = &value
	return cc
}

func (cc *ConnectionConfig) EncryptionKey(value string) *ConnectionConfig {
	cc.encryptionKey = &value
	return cc
}

func (cc *ConnectionConfig) EuRegion(value bool) *ConnectionConfig {
	cc.euRegion = &value
	return cc
}

func (cc *ConnectionConfig) TokenKey(value string) *ConnectionConfig {
	cc.tokenKey = &value
	return cc
}

func (cc *ConnectionConfig) TokenSecret(value string) *ConnectionConfig {
	cc.tokenSecret = &value
	return cc
}

func (cc *ConnectionConfig) ShareURL(value string) *ConnectionConfig {
	cc.shareURL = &value
	return cc
}
