package fivetran

type ConnectorConfig struct {
	FSchema                           *string                              `json:"schema,omitempty"`
	FTable                            *string                              `json:"table,omitempty"`
	FSheetID                          *string                              `json:"sheet_id,omitempty"`
	FNamedRange                       *string                              `json:"named_range,omitempty"`
	FClientID                         *string                              `json:"client_id,omitempty"`
	FClientSecret                     *string                              `json:"client_secret,omitempty"`
	FTechnicalAccountID               *string                              `json:"technical_account_id,omitempty"`
	FOrganizationID                   *string                              `json:"organization_id,omitempty"`
	FPrivateKey                       *string                              `json:"private_key,omitempty"`
	FSyncMode                         *string                              `json:"sync_mode,omitempty"`
	FReportSuites                     *[]string                            `json:"report_suites,omitempty"`
	FElements                         *[]string                            `json:"elements,omitempty"`
	FMetrics                          *[]string                            `json:"metrics,omitempty"`
	FDateGranularity                  *string                              `json:"date_granularity,omitempty"`
	FTimeframeMonths                  *string                              `json:"timeframe_months,omitempty"`
	FSource                           *string                              `json:"source,omitempty"`
	FS3Bucket                         *string                              `json:"s3bucket,omitempty"`
	FS3RoleArn                        *string                              `json:"s3role_arn,omitempty"`
	FAbsConnectionString              *string                              `json:"abs_connection_string,omitempty"`
	FAbsContainerName                 *string                              `json:"abs_container_name,omitempty"`
	FFtpHost                          *string                              `json:"ftp_host,omitempty"`
	FFtpPort                          *int                                 `json:"ftp_port,omitempty"`
	FFtpUser                          *string                              `json:"ftp_user,omitempty"`
	FFtpPassword                      *string                              `json:"ftp_password,omitempty"`
	FIsFtps                           *bool                                `json:"is_ftps,omitempty"`
	FSftpHost                         *string                              `json:"sftp_host,omitempty"`
	FSftpPort                         *int                                 `json:"sftp_port,omitempty"`
	FSftpUser                         *string                              `json:"sftp_user,omitempty"`
	FSftpPassword                     *string                              `json:"sftp_password,omitempty"`
	FSftpIsKeyPair                    *bool                                `json:"sftp_is_key_pair,omitempty"`
	FAdvertisables                    *[]string                            `json:"advertisables,omitempty"`
	FReportType                       *string                              `json:"report_type,omitempty"`
	FDimensions                       *[]string                            `json:"dimensions,omitempty"`
	FSchemaPrefix                     *string                              `json:"schema_prefix,omitempty"`
	FAPIKey                           *string                              `json:"api_key,omitempty"`
	FExternalID                       *string                              `json:"external_id,omitempty"`
	FRoleArn                          *string                              `json:"role_arn,omitempty"`
	FBucket                           *string                              `json:"bucket,omitempty"`
	FPrefix                           *string                              `json:"prefix,omitempty"`
	FPattern                          *string                              `json:"pattern,omitempty"`
	FFileType                         *string                              `json:"file_type,omitempty"`
	FCompression                      *string                              `json:"compression,omitempty"`
	FOnError                          *string                              `json:"on_error,omitempty"`
	FAppendFileOption                 *string                              `json:"append_file_option,omitempty"`
	FArchivePattern                   *string                              `json:"archive_pattern,omitempty"`
	FNullSequence                     *string                              `json:"null_sequence,omitempty"`
	FDelimiter                        *string                              `json:"delimiter,omitempty"`
	FEscapeChar                       *string                              `json:"escape_char,omitempty"`
	FSkipBefore                       *string                              `json:"skip_before,omitempty"`
	FSkipAfter                        *string                              `json:"skip_after,omitempty"`
	FProjectCredentials               *[]ConnectorConfigProjectCredentials `json:"project_credentials,omitempty"`
	FAuthMode                         *string                              `json:"auth_mode,omitempty"`
	FUsername                         *string                              `json:"username,omitempty"`
	FPassword                         *string                              `json:"password,omitempty"`
	FCertificate                      *string                              `json:"certificate,omitempty"`
	FSelectedExports                  *[]string                            `json:"selected_exports,omitempty"`
	FConsumerGroup                    *string                              `json:"consumer_group,omitempty"`
	FServers                          *string                              `json:"servers,omitempty"`
	FMessageType                      *string                              `json:"message_type,omitempty"`
	FSyncType                         *string                              `json:"sync_type,omitempty"`
	FSecurityProtocol                 *string                              `json:"security_protocol,omitempty"`
	FApps                             *[]string                            `json:"apps,omitempty"`
	FSalesAccounts                    *[]string                            `json:"sales_accounts,omitempty"`
	FFinanceAccounts                  *[]string                            `json:"finance_accounts,omitempty"`
	FAppSyncMode                      *string                              `json:"app_sync_mode,omitempty"`
	FSalesAccountSyncMode             *string                              `json:"sales_account_sync_mode,omitempty"`
	FFinanceAccountSyncMode           *string                              `json:"finance_account_sync_mode,omitempty"`
	FPemCertificate                   *string                              `json:"pem_certificate,omitempty"`
	FAccessKeyID                      *string                              `json:"access_key_id,omitempty"`
	FSecretKey                        *string                              `json:"secret_key,omitempty"`
	FHomeFolder                       *string                              `json:"home_folder,omitempty"`
	FSyncDataLocker                   *bool                                `json:"sync_data_locker,omitempty"`
	FProjects                         *[]string                            `json:"projects,omitempty"`
	FFunction                         *string                              `json:"function,omitempty"`
	FRegion                           *string                              `json:"region,omitempty"`
	FSecrets                          *string                              `json:"secrets,omitempty"`
	FContainerName                    *string                              `json:"container_name,omitempty"`
	FConnectionString                 *string                              `json:"connection_string,omitempty"`
	FFunctionApp                      *string                              `json:"function_app,omitempty"`
	FFunctionName                     *string                              `json:"function_name,omitempty"`
	FFunctionKey                      *string                              `json:"function_key,omitempty"`
	FPublicKey                        *string                              `json:"public_key,omitempty"`
	FMerchantID                       *string                              `json:"merchant_id,omitempty"`
	FAPIURL                           *string                              `json:"api_url,omitempty"`
	FCloudStorageType                 *string                              `json:"cloud_storage_type,omitempty"`
	FS3ExternalID                     *string                              `json:"s3external_id,omitempty"`
	FS3Folder                         *string                              `json:"s3folder,omitempty"`
	FGcsBucket                        *string                              `json:"gcs_bucket,omitempty"`
	FGcsFolder                        *string                              `json:"gcs_folder,omitempty"`
	FUserProfiles                     *[]string                            `json:"user_profiles,omitempty"` // continue here
	FReportConfigurationIds           *[]string                            `json:"report_configuration_ids,omitempty"`
	FEnableAllDimensionCombinations   *bool                                `json:"enable_all_dimension_combinations,omitempty"`
	FInstance                         *string                              `json:"instance,omitempty"`
	FAwsRegionCode                    *string                              `json:"aws_region_code,omitempty"`
	FAccounts                         *[]string                            `json:"accounts,omitempty"`
	FFields                           *[]string                            `json:"fields,omitempty"`
	FBreakdowns                       *[]string                            `json:"breakdowns,omitempty"`
	FActionBreakdowns                 *[]string                            `json:"action_breakdowns,omitempty"`
	FAggregation                      *string                              `json:"aggregation,omitempty"`
	FConfigType                       *string                              `json:"config_type,omitempty"`
	FPrebuiltReport                   *string                              `json:"prebuilt_report,omitempty"`
	FActionReportTime                 *string                              `json:"action_report_time,omitempty"`
	FClickAttributionWindow           *string                              `json:"click_attribution_window,omitempty"`
	FViewAttributionWindow            *string                              `json:"view_attribution_window,omitempty"`
	FCustomTables                     *[]ConnectorConfigCustomTables       `json:"custom_tables,omitempty"`
	FPages                            *[]string                            `json:"pages,omitempty"`
	FSubdomain                        *string                              `json:"subdomain,omitempty"`
	FHost                             *string                              `json:"host,omitempty"`
	FPort                             *int                                 `json:"port,omitempty"` // Fport changed to int to support Postgres; should be type string / interface
	FUser                             *string                              `json:"user,omitempty"`
	FIsSecure                         *string                              `json:"is_secure,omitempty"`
	FRepositories                     *[]string                            `json:"repositories,omitempty"`
	FUseWebhooks                      *bool                                `json:"use_webhooks,omitempty"`
	FDimensionAttributes              *[]string                            `json:"dimension_attributes,omitempty"`
	FColumns                          *[]string                            `json:"columns,omitempty"`
	FNetworkCode                      *string                              `json:"network_code,omitempty"`
	FCustomerID                       *string                              `json:"customer_id,omitempty"`
	FManagerAccounts                  *[]string                            `json:"manager_accounts,omitempty"`
	FReports                          *[]ConnectorConfigReports            `json:"reports,omitempty"`
	FConversionWindowSize             *int                                 `json:"conversion_window_size,omitempty"`
	FProfiles                         *[]string                            `json:"profiles,omitempty"`
	FProjectID                        *string                              `json:"project_id,omitempty"`
	FDatasetID                        *string                              `json:"dataset_id,omitempty"`
	FBucketName                       *string                              `json:"bucket_name,omitempty"`
	FFunctionTrigger                  *string                              `json:"function_trigger,omitempty"`
	FConfigMethod                     *string                              `json:"config_method,omitempty"`
	FQueryID                          *string                              `json:"query_id,omitempty"`
	FUpdateConfigOnEachSync           *bool                                `json:"update_config_on_each_sync,omitempty"`
	FSiteUrls                         *[]string                            `json:"site_urls,omitempty"`
	FPath                             *string                              `json:"path,omitempty"`
	FOnPremise                        *bool                                `json:"on_premise,omitempty"`
	FAccessToken                      *string                              `json:"access_token,omitempty"`
	FViewThroughAttributionWindowSize *string                              `json:"view_through_attribution_window_size,omitempty"`
	FPostClickAttributionWindowSize   *string                              `json:"post_click_attribution_window_size,omitempty"`
	FUseAPIKeys                       *string                              `json:"use_api_keys,omitempty"`
	FAPIKeys                          *string                              `json:"api_keys,omitempty"`
	FEndpoint                         *string                              `json:"endpoint,omitempty"`
	FIdentity                         *string                              `json:"identity,omitempty"`
	FAPIQuota                         *int                                 `json:"api_quota,omitempty"`
	FDomainName                       *string                              `json:"domain_name,omitempty"`
	FResourceURL                      *string                              `json:"resource_url,omitempty"`
	FAPISecret                        *string                              `json:"api_secret,omitempty"`
	FHosts                            *[]string                            `json:"hosts,omitempty"`
	FTunnelHost                       *string                              `json:"tunnel_host,omitempty"`
	FTunnelPort                       *int                                 `json:"tunnel_port,omitempty"`
	FTunnelUser                       *string                              `json:"tunnel_user,omitempty"`
	FDatabase                         *string                              `json:"database,omitempty"`
	FDatasource                       *string                              `json:"datasource,omitempty"`
	FAccount                          *string                              `json:"account,omitempty"`
	FRole                             *string                              `json:"role,omitempty"`
	FEmail                            *string                              `json:"email,omitempty"`
	FAccountID                        *string                              `json:"account_id,omitempty"`
	FServerURL                        *string                              `json:"server_url,omitempty"`
	FUserKey                          *string                              `json:"user_key,omitempty"`
	FAPIVersion                       *string                              `json:"api_version,omitempty"`
	FDailyAPICallLimit                *int                                 `json:"daily_api_call_limit,omitempty"`
	FTimeZone                         *string                              `json:"time_zone,omitempty"`
	FIntegrationKey                   *string                              `json:"integration_key,omitempty"`
	FAdvertisers                      *[]string                            `json:"advertisers,omitempty"`
	FEngagementAttributionWindow      *string                              `json:"engagement_attribution_window,omitempty"`
	FConversionReportTime             *string                              `json:"conversion_report_time,omitempty"`
	FDomain                           *string                              `json:"domain,omitempty"`
	FUpdateMethod                     *string                              `json:"update_method,omitempty"`
	FReplicationSlot                  *string                              `json:"replication_slot,omitempty"`
	FDataCenter                       *string                              `json:"data_center,omitempty"`
	FAPIToken                         *string                              `json:"api_token,omitempty"`
	FSubDomain                        *string                              `json:"sub_domain,omitempty"`
	FTestTableName                    *string                              `json:"test_table_name,omitempty"`
	FShop                             *string                              `json:"shop,omitempty"`
	FOrganizations                    *[]string                            `json:"organizations,omitempty"`
	FSwipeAttributionWindow           *string                              `json:"swipe_attribution_window,omitempty"`
	FAPIAccessToken                   *string                              `json:"api_access_token,omitempty"`
	FAccountIds                       *string                              `json:"account_ids,omitempty"`
	FSid                              *string                              `json:"sid,omitempty"`
	FSecret                           *string                              `json:"secret,omitempty"`
	FOauthToken                       *string                              `json:"oauth_token,omitempty"`
	FOauthTokenSecret                 *string                              `json:"oauth_token_secret,omitempty"`
	FConsumerKey                      *string                              `json:"consumer_key,omitempty"`
	FConsumerSecret                   *string                              `json:"consumer_secret,omitempty"`
	FKey                              *string                              `json:"key,omitempty"`
	FAdvertisersID                    *[]string                            `json:"advertisers_id,omitempty"`
	FSyncFormat                       *string                              `json:"sync_format,omitempty"`
	FBucketService                    *string                              `json:"bucket_service,omitempty"`
	FUserName                         *string                              `json:"user_name,omitempty"`
	FReportURL                        *string                              `json:"report_url,omitempty"`
	FUniqueID                         *string                              `json:"unique_id,omitempty"`
	// FPort                             *int                                 `json:"port,omitempty"`     // splunk, postgresql
	// FAccounts                         *[]int                               `json:"accounts,omitempty"` // tiktok ads
	LatestVersion        *string `json:"latest_version,omitempty"`
	AuthorizationMethod  *string `json:"authorization_method,omitempty"`
	ServiceVersion       *string `json:"service_version,omitempty"`
	LastSyncedChangesUtc *string `json:"last_synced_changes__utc_,omitempty"`
}

func NewConnectorConfig() *ConnectorConfig {
	return &ConnectorConfig{}
}

func (cc *ConnectorConfig) Schema(value string) *ConnectorConfig {
	cc.FSchema = &value
	return cc
}

func (cc *ConnectorConfig) Table(value string) *ConnectorConfig {
	cc.FTable = &value
	return cc
}

func (cc *ConnectorConfig) SheetID(value string) *ConnectorConfig {
	cc.FSheetID = &value
	return cc
}

func (cc *ConnectorConfig) NamedRange(value string) *ConnectorConfig {
	cc.FNamedRange = &value
	return cc
}

func (cc *ConnectorConfig) ClientID(value string) *ConnectorConfig {
	cc.FClientID = &value
	return cc
}

func (cc *ConnectorConfig) ClientSecret(value string) *ConnectorConfig {
	cc.FClientSecret = &value
	return cc
}

func (cc *ConnectorConfig) TechnicalAccountID(value string) *ConnectorConfig {
	cc.FTechnicalAccountID = &value
	return cc
}

func (cc *ConnectorConfig) OrganizationID(value string) *ConnectorConfig {
	cc.FOrganizationID = &value
	return cc
}

func (cc *ConnectorConfig) PrivateKey(value string) *ConnectorConfig {
	cc.FPrivateKey = &value
	return cc
}

func (cc *ConnectorConfig) SyncMode(value string) *ConnectorConfig {
	cc.FSyncMode = &value
	return cc
}

func (cc *ConnectorConfig) ReportSuites(value []string) *ConnectorConfig {
	cc.FReportSuites = &value
	return cc
}

func (cc *ConnectorConfig) Elements(value []string) *ConnectorConfig {
	cc.FElements = &value
	return cc
}

func (cc *ConnectorConfig) Metrics(value []string) *ConnectorConfig {
	cc.FMetrics = &value
	return cc
}

func (cc *ConnectorConfig) DateGranularity(value string) *ConnectorConfig {
	cc.FDateGranularity = &value
	return cc
}

func (cc *ConnectorConfig) TimeframeMonths(value string) *ConnectorConfig {
	cc.FTimeframeMonths = &value
	return cc
}

func (cc *ConnectorConfig) Source(value string) *ConnectorConfig {
	cc.FSource = &value
	return cc
}

func (cc *ConnectorConfig) S3Bucket(value string) *ConnectorConfig {
	cc.FS3Bucket = &value
	return cc
}

func (cc *ConnectorConfig) S3RoleArn(value string) *ConnectorConfig {
	cc.FS3RoleArn = &value
	return cc
}

func (cc *ConnectorConfig) AbsConnectionString(value string) *ConnectorConfig {
	cc.FAbsConnectionString = &value
	return cc
}

func (cc *ConnectorConfig) AbsContainerName(value string) *ConnectorConfig {
	cc.FAbsContainerName = &value
	return cc
}

func (cc *ConnectorConfig) FtpHost(value string) *ConnectorConfig {
	cc.FFtpHost = &value
	return cc
}

func (cc *ConnectorConfig) FtpPort(value int) *ConnectorConfig {
	cc.FFtpPort = &value
	return cc
}

func (cc *ConnectorConfig) FtpUser(value string) *ConnectorConfig {
	cc.FFtpUser = &value
	return cc
}

func (cc *ConnectorConfig) FtpPassword(value string) *ConnectorConfig {
	cc.FFtpPassword = &value
	return cc
}

func (cc *ConnectorConfig) IsFtps(value bool) *ConnectorConfig {
	cc.FIsFtps = &value
	return cc
}

func (cc *ConnectorConfig) SftpHost(value string) *ConnectorConfig {
	cc.FSftpHost = &value
	return cc
}

func (cc *ConnectorConfig) SftpPort(value int) *ConnectorConfig {
	cc.FSftpPort = &value
	return cc
}

func (cc *ConnectorConfig) SftpUser(value string) *ConnectorConfig {
	cc.FSftpUser = &value
	return cc
}

func (cc *ConnectorConfig) SftpPassword(value string) *ConnectorConfig {
	cc.FSftpPassword = &value
	return cc
}

func (cc *ConnectorConfig) SftpIsKeyPair(value bool) *ConnectorConfig {
	cc.FSftpIsKeyPair = &value
	return cc
}

func (cc *ConnectorConfig) Advertisables(value []string) *ConnectorConfig {
	cc.FAdvertisables = &value
	return cc
}

func (cc *ConnectorConfig) ReportType(value string) *ConnectorConfig {
	cc.FReportType = &value
	return cc
}

func (cc *ConnectorConfig) Dimensions(value []string) *ConnectorConfig {
	cc.FDimensions = &value
	return cc
}

func (cc *ConnectorConfig) SchemaPrefix(value string) *ConnectorConfig {
	cc.FSchemaPrefix = &value
	return cc
}

func (cc *ConnectorConfig) APIKey(value string) *ConnectorConfig {
	cc.FAPIKey = &value
	return cc
}

func (cc *ConnectorConfig) ExternalID(value string) *ConnectorConfig {
	cc.FExternalID = &value
	return cc
}

func (cc *ConnectorConfig) RoleArn(value string) *ConnectorConfig {
	cc.FRoleArn = &value
	return cc
}

func (cc *ConnectorConfig) Bucket(value string) *ConnectorConfig {
	cc.FBucket = &value
	return cc
}

func (cc *ConnectorConfig) Prefix(value string) *ConnectorConfig {
	cc.FPrefix = &value
	return cc
}

func (cc *ConnectorConfig) Pattern(value string) *ConnectorConfig {
	cc.FPattern = &value
	return cc
}

func (cc *ConnectorConfig) FileType(value string) *ConnectorConfig {
	cc.FFileType = &value
	return cc
}

func (cc *ConnectorConfig) Compression(value string) *ConnectorConfig {
	cc.FCompression = &value
	return cc
}

func (cc *ConnectorConfig) OnError(value string) *ConnectorConfig {
	cc.FOnError = &value
	return cc
}

func (cc *ConnectorConfig) AppendFileOption(value string) *ConnectorConfig {
	cc.FAppendFileOption = &value
	return cc
}

func (cc *ConnectorConfig) ArchivePattern(value string) *ConnectorConfig {
	cc.FArchivePattern = &value
	return cc
}

func (cc *ConnectorConfig) NullSequence(value string) *ConnectorConfig {
	cc.FNullSequence = &value
	return cc
}

func (cc *ConnectorConfig) Delimiter(value string) *ConnectorConfig {
	cc.FDelimiter = &value
	return cc
}

func (cc *ConnectorConfig) EscapeChar(value string) *ConnectorConfig {
	cc.FEscapeChar = &value
	return cc
}

func (cc *ConnectorConfig) SkipBefore(value string) *ConnectorConfig {
	cc.FSkipBefore = &value
	return cc
}

func (cc *ConnectorConfig) SkipAfter(value string) *ConnectorConfig {
	cc.FSkipAfter = &value
	return cc
}

func (cc *ConnectorConfig) ProjectCredentials(value []ConnectorConfigProjectCredentials) *ConnectorConfig {
	cc.FProjectCredentials = &value
	return cc
}

func (cc *ConnectorConfig) AuthMode(value string) *ConnectorConfig {
	cc.FAuthMode = &value
	return cc
}

func (cc *ConnectorConfig) Username(value string) *ConnectorConfig {
	cc.FUsername = &value
	return cc
}

func (cc *ConnectorConfig) Password(value string) *ConnectorConfig {
	cc.FPassword = &value
	return cc
}

func (cc *ConnectorConfig) Certificate(value string) *ConnectorConfig {
	cc.FCertificate = &value
	return cc
}

func (cc *ConnectorConfig) SelectedExports(value []string) *ConnectorConfig {
	cc.FSelectedExports = &value
	return cc
}

func (cc *ConnectorConfig) ConsumerGroup(value string) *ConnectorConfig {
	cc.FConsumerGroup = &value
	return cc
}

func (cc *ConnectorConfig) Servers(value string) *ConnectorConfig {
	cc.FServers = &value
	return cc
}

func (cc *ConnectorConfig) MessageType(value string) *ConnectorConfig {
	cc.FMessageType = &value
	return cc
}

func (cc *ConnectorConfig) SyncType(value string) *ConnectorConfig {
	cc.FSyncType = &value
	return cc
}

func (cc *ConnectorConfig) SecurityProtocol(value string) *ConnectorConfig {
	cc.FSecurityProtocol = &value
	return cc
}

func (cc *ConnectorConfig) Apps(value []string) *ConnectorConfig {
	cc.FApps = &value
	return cc
}

func (cc *ConnectorConfig) SalesAccounts(value []string) *ConnectorConfig {
	cc.FSalesAccounts = &value
	return cc
}

func (cc *ConnectorConfig) FinanceAccounts(value []string) *ConnectorConfig {
	cc.FFinanceAccounts = &value
	return cc
}

func (cc *ConnectorConfig) AppSyncMode(value string) *ConnectorConfig {
	cc.FAppSyncMode = &value
	return cc
}

func (cc *ConnectorConfig) SalesAccountSyncMode(value string) *ConnectorConfig {
	cc.FSalesAccountSyncMode = &value
	return cc
}

func (cc *ConnectorConfig) FinanceAccountSyncMode(value string) *ConnectorConfig {
	cc.FFinanceAccountSyncMode = &value
	return cc
}

func (cc *ConnectorConfig) PemCertificate(value string) *ConnectorConfig {
	cc.FPemCertificate = &value
	return cc
}

func (cc *ConnectorConfig) AccessKeyID(value string) *ConnectorConfig {
	cc.FAccessKeyID = &value
	return cc
}

func (cc *ConnectorConfig) SecretKey(value string) *ConnectorConfig {
	cc.FSecretKey = &value
	return cc
}

func (cc *ConnectorConfig) HomeFolder(value string) *ConnectorConfig {
	cc.FHomeFolder = &value
	return cc
}

func (cc *ConnectorConfig) SyncDataLocker(value bool) *ConnectorConfig {
	cc.FSyncDataLocker = &value
	return cc
}

func (cc *ConnectorConfig) Projects(value []string) *ConnectorConfig {
	cc.FProjects = &value
	return cc
}

func (cc *ConnectorConfig) Function(value string) *ConnectorConfig {
	cc.FFunction = &value
	return cc
}

func (cc *ConnectorConfig) Region(value string) *ConnectorConfig {
	cc.FRegion = &value
	return cc
}

func (cc *ConnectorConfig) Secrets(value string) *ConnectorConfig {
	cc.FSecrets = &value
	return cc
}

func (cc *ConnectorConfig) ContainerName(value string) *ConnectorConfig {
	cc.FContainerName = &value
	return cc
}

func (cc *ConnectorConfig) ConnectionString(value string) *ConnectorConfig {
	cc.FConnectionString = &value
	return cc
}

func (cc *ConnectorConfig) FunctionApp(value string) *ConnectorConfig {
	cc.FFunctionApp = &value
	return cc
}

func (cc *ConnectorConfig) FunctionName(value string) *ConnectorConfig {
	cc.FFunctionName = &value
	return cc
}

func (cc *ConnectorConfig) FunctionKey(value string) *ConnectorConfig {
	cc.FFunctionKey = &value
	return cc
}

func (cc *ConnectorConfig) PublicKey(value string) *ConnectorConfig {
	cc.FPublicKey = &value
	return cc
}

func (cc *ConnectorConfig) MerchantID(value string) *ConnectorConfig {
	cc.FMerchantID = &value
	return cc
}

func (cc *ConnectorConfig) APIURL(value string) *ConnectorConfig {
	cc.FAPIURL = &value
	return cc
}

func (cc *ConnectorConfig) CloudStorageType(value string) *ConnectorConfig {
	cc.FCloudStorageType = &value
	return cc
}

func (cc *ConnectorConfig) S3ExternalID(value string) *ConnectorConfig {
	cc.FS3ExternalID = &value
	return cc
}

func (cc *ConnectorConfig) S3Folder(value string) *ConnectorConfig {
	cc.FS3Folder = &value
	return cc
}

func (cc *ConnectorConfig) GcsBucket(value string) *ConnectorConfig {
	cc.FGcsFolder = &value
	return cc
}

func (cc *ConnectorConfig) UserProfiles(value []string) *ConnectorConfig {
	cc.FUserProfiles = &value
	return cc
}

func (cc *ConnectorConfig) ReportConfigurationIds(value []string) *ConnectorConfig {
	cc.FReportConfigurationIds = &value
	return cc
}

func (cc *ConnectorConfig) EnableAllDimensionCombinations(value bool) *ConnectorConfig {
	cc.FEnableAllDimensionCombinations = &value
	return cc
}

func (cc *ConnectorConfig) Instance(value string) *ConnectorConfig {
	cc.FInstance = &value
	return cc
}

func (cc *ConnectorConfig) AwsRegionCode(value string) *ConnectorConfig {
	cc.FAwsRegionCode = &value
	return cc
}

func (cc *ConnectorConfig) Accounts(value []string) *ConnectorConfig {
	cc.FAccounts = &value
	return cc
}

func (cc *ConnectorConfig) Fields(value []string) *ConnectorConfig {
	cc.FFields = &value
	return cc
}

func (cc *ConnectorConfig) Breakdowns(value []string) *ConnectorConfig {
	cc.FBreakdowns = &value
	return cc
}

func (cc *ConnectorConfig) ActionBreakdowns(value []string) *ConnectorConfig {
	cc.FActionBreakdowns = &value
	return cc
}

func (cc *ConnectorConfig) Aggregation(value string) *ConnectorConfig {
	cc.FAggregation = &value
	return cc
}

func (cc *ConnectorConfig) ConfigType(value string) *ConnectorConfig {
	cc.FConfigType = &value
	return cc
}

func (cc *ConnectorConfig) PrebuiltReport(value string) *ConnectorConfig {
	cc.FPrebuiltReport = &value
	return cc
}

func (cc *ConnectorConfig) ActionReportTime(value string) *ConnectorConfig {
	cc.FActionReportTime = &value
	return cc
}

func (cc *ConnectorConfig) ClickAttributionWindow(value string) *ConnectorConfig {
	cc.FClickAttributionWindow = &value
	return cc
}

func (cc *ConnectorConfig) ViewAttributionWindow(value string) *ConnectorConfig {
	cc.FViewAttributionWindow = &value
	return cc
}

func (cc *ConnectorConfig) CustomTables(value []ConnectorConfigCustomTables) *ConnectorConfig {
	cc.FCustomTables = &value
	return cc
}

func (cc *ConnectorConfig) Pages(value []string) *ConnectorConfig {
	cc.FPages = &value
	return cc
}

func (cc *ConnectorConfig) Subdomain(value string) *ConnectorConfig {
	cc.FSubdomain = &value
	return cc
}

func (cc *ConnectorConfig) Host(value string) *ConnectorConfig {
	cc.FHost = &value
	return cc
}

func (cc *ConnectorConfig) Port(value int) *ConnectorConfig {
	cc.FPort = &value
	return cc
}

func (cc *ConnectorConfig) User(value string) *ConnectorConfig {
	cc.FUser = &value
	return cc
}

func (cc *ConnectorConfig) IsSecure(value string) *ConnectorConfig {
	cc.FIsSecure = &value
	return cc
}

func (cc *ConnectorConfig) Repositories(value []string) *ConnectorConfig {
	cc.FRepositories = &value
	return cc
}

func (cc *ConnectorConfig) UseWebhooks(value bool) *ConnectorConfig {
	cc.FUseWebhooks = &value
	return cc
}

func (cc *ConnectorConfig) DimensionAttributes(value []string) *ConnectorConfig {
	cc.FDimensionAttributes = &value
	return cc
}

func (cc *ConnectorConfig) Columns(value []string) *ConnectorConfig {
	cc.FColumns = &value
	return cc
}

func (cc *ConnectorConfig) NetworkCode(value string) *ConnectorConfig {
	cc.FNetworkCode = &value
	return cc
}

func (cc *ConnectorConfig) CustomerID(value string) *ConnectorConfig {
	cc.FCustomerID = &value
	return cc
}

func (cc *ConnectorConfig) ManagerAccounts(value []string) *ConnectorConfig {
	cc.FManagerAccounts = &value
	return cc
}

func (cc *ConnectorConfig) Reports(value []ConnectorConfigReports) *ConnectorConfig {
	cc.FReports = &value
	return cc
}

func (cc *ConnectorConfig) ConversionWindowSize(value int) *ConnectorConfig {
	cc.FConversionWindowSize = &value
	return cc
}

func (cc *ConnectorConfig) Profiles(value []string) *ConnectorConfig {
	cc.FProfiles = &value
	return cc
}

func (cc *ConnectorConfig) ProjectID(value string) *ConnectorConfig {
	cc.FProjectID = &value
	return cc
}

func (cc *ConnectorConfig) DatasetID(value string) *ConnectorConfig {
	cc.FDatasetID = &value
	return cc
}

func (cc *ConnectorConfig) BucketName(value string) *ConnectorConfig {
	cc.FBucketName = &value
	return cc
}

func (cc *ConnectorConfig) FunctionTrigger(value string) *ConnectorConfig {
	cc.FFunctionTrigger = &value
	return cc
}

func (cc *ConnectorConfig) ConfigMethod(value string) *ConnectorConfig {
	cc.FConfigMethod = &value
	return cc
}

func (cc *ConnectorConfig) QueryID(value string) *ConnectorConfig {
	cc.FQueryID = &value
	return cc
}

func (cc *ConnectorConfig) UpdateConfigOnEachSync(value bool) *ConnectorConfig {
	cc.FUpdateConfigOnEachSync = &value
	return cc
}

func (cc *ConnectorConfig) SiteUrls(value []string) *ConnectorConfig {
	cc.FSiteUrls = &value
	return cc
}

func (cc *ConnectorConfig) Path(value string) *ConnectorConfig {
	cc.FPath = &value
	return cc
}

func (cc *ConnectorConfig) OnPremise(value bool) *ConnectorConfig {
	cc.FOnPremise = &value
	return cc
}

func (cc *ConnectorConfig) AccessToken(value string) *ConnectorConfig {
	cc.FAccessToken = &value
	return cc
}

func (cc *ConnectorConfig) ViewThroughAttributionWindowSize(value string) *ConnectorConfig {
	cc.FViewThroughAttributionWindowSize = &value
	return cc
}

func (cc *ConnectorConfig) PostClickAttributionWindowSize(value string) *ConnectorConfig {
	cc.FPostClickAttributionWindowSize = &value
	return cc
}

func (cc *ConnectorConfig) UseAPIKeys(value string) *ConnectorConfig {
	cc.FUseAPIKeys = &value
	return cc
}

func (cc *ConnectorConfig) APIKeys(value string) *ConnectorConfig {
	cc.FAPIKeys = &value
	return cc
}

func (cc *ConnectorConfig) Endpoint(value string) *ConnectorConfig {
	cc.FEndpoint = &value
	return cc
}

func (cc *ConnectorConfig) Identity(value string) *ConnectorConfig {
	cc.FIdentity = &value
	return cc
}

func (cc *ConnectorConfig) APIQuota(value int) *ConnectorConfig {
	cc.FAPIQuota = &value
	return cc
}

func (cc *ConnectorConfig) DomainName(value string) *ConnectorConfig {
	cc.FDomainName = &value
	return cc
}

func (cc *ConnectorConfig) ResourceURL(value string) *ConnectorConfig {
	cc.FResourceURL = &value
	return cc
}

func (cc *ConnectorConfig) APISecret(value string) *ConnectorConfig {
	cc.FAPISecret = &value
	return cc
}

func (cc *ConnectorConfig) Hosts(value []string) *ConnectorConfig {
	cc.FHosts = &value
	return cc
}

func (cc *ConnectorConfig) TunnelHost(value string) *ConnectorConfig {
	cc.FTunnelHost = &value
	return cc
}

func (cc *ConnectorConfig) TunnelPort(value int) *ConnectorConfig {
	cc.FTunnelPort = &value
	return cc
}

func (cc *ConnectorConfig) TunnelUser(value string) *ConnectorConfig {
	cc.FTunnelUser = &value
	return cc
}

func (cc *ConnectorConfig) Database(value string) *ConnectorConfig {
	cc.FDatabase = &value
	return cc
}

func (cc *ConnectorConfig) Datasource(value string) *ConnectorConfig {
	cc.FDatasource = &value
	return cc
}

func (cc *ConnectorConfig) Account(value string) *ConnectorConfig {
	cc.FAccount = &value
	return cc
}

func (cc *ConnectorConfig) Role(value string) *ConnectorConfig {
	cc.FRole = &value
	return cc
}

func (cc *ConnectorConfig) Email(value string) *ConnectorConfig {
	cc.FEmail = &value
	return cc
}

func (cc *ConnectorConfig) AccountID(value string) *ConnectorConfig {
	cc.FAccountID = &value
	return cc
}

func (cc *ConnectorConfig) ServerURL(value string) *ConnectorConfig {
	cc.FServerURL = &value
	return cc
}

func (cc *ConnectorConfig) UserKey(value string) *ConnectorConfig {
	cc.FUserKey = &value
	return cc
}

func (cc *ConnectorConfig) APIVersion(value string) *ConnectorConfig {
	cc.FAPIVersion = &value
	return cc
}

func (cc *ConnectorConfig) DailyAPICallLimit(value int) *ConnectorConfig {
	cc.FDailyAPICallLimit = &value
	return cc
}

func (cc *ConnectorConfig) TimeZone(value string) *ConnectorConfig {
	cc.FTimeZone = &value
	return cc
}

func (cc *ConnectorConfig) IntegrationKey(value string) *ConnectorConfig {
	cc.FIntegrationKey = &value
	return cc
}

func (cc *ConnectorConfig) Advertisers(value []string) *ConnectorConfig {
	cc.FAdvertisers = &value
	return cc
}

func (cc *ConnectorConfig) EngagementAttributionWindow(value string) *ConnectorConfig {
	cc.FEngagementAttributionWindow = &value
	return cc
}

func (cc *ConnectorConfig) ConversionReportTime(value string) *ConnectorConfig {
	cc.FConversionReportTime = &value
	return cc
}

func (cc *ConnectorConfig) Domain(value string) *ConnectorConfig {
	cc.FDomain = &value
	return cc
}

func (cc *ConnectorConfig) UpdateMethod(value string) *ConnectorConfig {
	cc.FUpdateMethod = &value
	return cc
}

func (cc *ConnectorConfig) ReplicationSlot(value string) *ConnectorConfig {
	cc.FReplicationSlot = &value
	return cc
}

func (cc *ConnectorConfig) DataCenter(value string) *ConnectorConfig {
	cc.FDataCenter = &value
	return cc
}

func (cc *ConnectorConfig) APIToken(value string) *ConnectorConfig {
	cc.FAPIToken = &value
	return cc
}

func (cc *ConnectorConfig) SubDomain(value string) *ConnectorConfig {
	cc.FSubDomain = &value
	return cc
}

func (cc *ConnectorConfig) TestTableName(value string) *ConnectorConfig {
	cc.FTestTableName = &value
	return cc
}

func (cc *ConnectorConfig) Shop(value string) *ConnectorConfig {
	cc.FShop = &value
	return cc
}

func (cc *ConnectorConfig) Organizations(value []string) *ConnectorConfig {
	cc.FOrganizations = &value
	return cc
}

func (cc *ConnectorConfig) SwipeAttributionWindow(value string) *ConnectorConfig {
	cc.FSwipeAttributionWindow = &value
	return cc
}

func (cc *ConnectorConfig) APIAccessToken(value string) *ConnectorConfig {
	cc.FAPIAccessToken = &value
	return cc
}

func (cc *ConnectorConfig) AccountIds(value string) *ConnectorConfig {
	cc.FAccountIds = &value
	return cc
}

func (cc *ConnectorConfig) Sid(value string) *ConnectorConfig {
	cc.FSid = &value
	return cc
}

func (cc *ConnectorConfig) Secret(value string) *ConnectorConfig {
	cc.FSecret = &value
	return cc
}

func (cc *ConnectorConfig) OauthToken(value string) *ConnectorConfig {
	cc.FOauthToken = &value
	return cc
}

func (cc *ConnectorConfig) OauthTokenSecret(value string) *ConnectorConfig {
	cc.FOauthTokenSecret = &value
	return cc
}

func (cc *ConnectorConfig) ConsumerKey(value string) *ConnectorConfig {
	cc.FConsumerKey = &value
	return cc
}

func (cc *ConnectorConfig) ConsumerSecret(value string) *ConnectorConfig {
	cc.FConsumerSecret = &value
	return cc
}

func (cc *ConnectorConfig) Key(value string) *ConnectorConfig {
	cc.FKey = &value
	return cc
}

func (cc *ConnectorConfig) AdvertisersID(value []string) *ConnectorConfig {
	cc.FAdvertisersID = &value
	return cc
}

func (cc *ConnectorConfig) SyncFormat(value string) *ConnectorConfig {
	cc.FSyncFormat = &value
	return cc
}

func (cc *ConnectorConfig) BucketService(value string) *ConnectorConfig {
	cc.FBucketService = &value
	return cc
}

func (cc *ConnectorConfig) UserName(value string) *ConnectorConfig {
	cc.FUserName = &value
	return cc
}

func (cc *ConnectorConfig) ReportURL(value string) *ConnectorConfig {
	cc.FReportURL = &value
	return cc
}

func (cc *ConnectorConfig) UniqueID(value string) *ConnectorConfig {
	cc.FUniqueID = &value
	return cc
}
