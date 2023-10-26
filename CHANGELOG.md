# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased](https://github.com/fivetran/go-fivetran/compare/v0.7.13...HEAD)

## [0.7.13](https://github.com/fivetran/go-fivetran/compare/v0.7.12...v0.7.13)

## Added

Added fields to support Azure Data Lake Storage:
- `DestinationConfig.storageAccountName`
- `DestinationConfig.containerName`
- `DestinationConfig.tenantId`
- `DestinationConfig.clientId`
- `DestinationConfig.secretValue`

## [0.7.12](https://github.com/fivetran/go-fivetran/compare/v0.7.11...v0.7.12)

## Added

Supported the following Fivetran API endpoints:
- [List all approved certificates for connector](https://fivetran.com/docs/rest-api/certificates#listallapprovedcertificatesforconnector)
- [List all approved fingerprints for connector](https://fivetran.com/docs/rest-api/certificates#listallapprovedfingerprintsforconnector)
- [Retrieve a connector certificate details](https://fivetran.com/docs/rest-api/certificates#retrieveaconnectorcertificatedetails)
- [Retrieve a connector fingerprint details](https://fivetran.com/docs/rest-api/certificates#retrieveaconnectorfingerprintdetails)
- [Revoke a connector certificate](https://fivetran.com/docs/rest-api/certificates#revokeaconnectorcertificate)
- [Revoke a connector fingerprint](https://fivetran.com/docs/rest-api/certificates#revokeaconnectorfingerprint)
- [List all approved certificates for destination](https://fivetran.com/docs/rest-api/certificates#listallapprovedcertificatesfordestination)
- [List all approved fingerprints for destination](https://fivetran.com/docs/rest-api/certificates#listallapprovedfingerprintsfordestination)
- [Retrieve a destination certificate details](https://fivetran.com/docs/rest-api/certificates#retrieveadestinationcertificatedetails)
- [Retrieve a destination fingerprint details](https://fivetran.com/docs/rest-api/certificates#retrieveadestinationfingerprintdetails)
- [Revoke a destination certificate](https://fivetran.com/docs/rest-api/certificates#revokeadestinationcertificate)
- [Revoke a destination fingerprint](https://fivetran.com/docs/rest-api/certificates#revokeadestinationfingerprint)
- [Modify a connector database schema config](https://fivetran.com/docs/rest-api/connectors#modifyaconnectordatabaseschemaconfig)
- [Modify a connector table config](https://fivetran.com/docs/rest-api/connectors#modifyaconnectortableconfig)
- [Modify a connector column config](https://fivetran.com/docs/rest-api/connectors#modifyaconnectorcolumnconfig)
- [Connect Card](https://fivetran.com/docs/rest-api/connectors/connect-card)

## Updated
Extended response object for the following endpoints:
- [Approve a connector certificate](https://fivetran.com/docs/rest-api/certificates#approveaconnectorcertificate)
- [Approve a connector fingerprint](https://fivetran.com/docs/rest-api/certificates#approveaconnectorfingerprint)
- [Approve a destination certificate](https://fivetran.com/docs/rest-api/certificates#approveadestinationcertificate)
- [Approve a destination fingerprint](https://fivetran.com/docs/rest-api/certificates#approveadestinationfingerprint)

## [0.7.11](https://github.com/fivetran/go-fivetran/compare/v0.7.10...v0.7.11)

## Added

Supported the following Fivetran API endpoints:
- Team Management API: [List all teams](https://fivetran.com/docs/rest-api/teams#listallteams)
- Team Management API: [Retrieve team details](https://fivetran.com/docs/rest-api/teams#retrieveteamdetails)
- Team Management API: [Create a team](https://fivetran.com/docs/rest-api/teams#createateam)
- Team Management API: [Modify a team](https://fivetran.com/docs/rest-api/teams#modifyateam)
- Team Management API: [Delete a team role in the account](https://fivetran.com/docs/rest-api/teams#deleteteamroleinaccount)
- Team Management API: [Delete a team](https://fivetran.com/docs/rest-api/teams#deleteateam)
- Team Management API User memberships: [List all user memberships](https://fivetran.com/docs/rest-api/teams#listallusermemberships)
- Team Management API User memberships: [Retrieve user membership](https://fivetran.com/docs/rest-api/teams#retrieveusermembershipinateam)
- Team Management API User memberships: [Add a user to a team](https://fivetran.com/docs/rest-api/teams#addausertoateam)
- Team Management API User memberships: [Modify a user membership](https://fivetran.com/docs/rest-api/teams#modifyausermembership)
- Team Management API User memberships: [Delete a user from a team](https://fivetran.com/docs/rest-api/teams#deleteauserfromateam)
- Team Management API Connector memberships: [List all connector memberships](https://fivetran.com/docs/rest-api/teams#listallconnectormemberships)
- Team Management API Connector memberships: [Retrieve connector membership](https://fivetran.com/docs/rest-api/teams#retrieveconnectormembership)
- Team Management API Connector memberships: [Add connector membership](https://fivetran.com/docs/rest-api/teams#addconnectormembership)
- Team Management API Connector memberships: [Update connector membership](https://fivetran.com/docs/rest-api/teams#updateconnectormembership)
- Team Management API Connector memberships: [Delete connector membership](https://fivetran.com/docs/rest-api/teams#deleteconnectormembership)
- Team Management API Group memberships: [List all group memberships](https://fivetran.com/docs/rest-api/teams#listallgroupmemberships)
- Team Management API Group memberships: [Retrieve group membership](https://fivetran.com/docs/rest-api/teams#retrievegroupmembership)
- Team Management API Group memberships: [Add group membership](https://fivetran.com/docs/rest-api/teams#addgroupmembership)
- Team Management API Group memberships: [Update group membership](https://fivetran.com/docs/rest-api/teams#updategroupmembership)
- Team Management API Group memberships: [Delete group membership](https://fivetran.com/docs/rest-api/teams#deletegroupmembership)

## [0.7.10](https://github.com/fivetran/go-fivetran/compare/v0.7.9...v0.7.10)

## Fixed

- Fixed export NewWebhookTest method
- Fixed markup in the README file

## [0.7.9](https://github.com/fivetran/go-fivetran/compare/v0.7.8...v0.7.9)

## Added
- Support for dbt Project status in responses.

## [0.7.8](https://github.com/fivetran/go-fivetran/compare/v0.7.7...v0.7.8)

## Fixed
- Dbt Project setters renamed for consistency ProjectID -> DbtProjectID for all endpoints
- Fixed path for delete dbt project 

## [0.7.7](https://github.com/fivetran/go-fivetran/compare/v0.7.6...v0.7.7)

## Added
Supported the following Fivetran API endpoints:
- [Retrieve dbt model details](https://fivetran.com/docs/rest-api/dbt-transformation-management#retrievedbtprojectmodeldetails)
- [Retrieve dbt models list](https://fivetran.com/docs/rest-api/dbt-transformation-management#retrievedbtprojectmodels)
- [Create dbt project](https://fivetran.com/docs/rest-api/dbt-transformation-management#createdbtproject)
- [Retrieve dbt projects list](https://fivetran.com/docs/rest-api/dbt-transformation-management#retrievedbtprojects)
- [Retrieve dbt project details](https://fivetran.com/docs/rest-api/dbt-transformation-management#retrievedbtprojectdetails)
- [Update dbt project](https://fivetran.com/docs/rest-api/dbt-transformation-management#modifydbtproject)
- [Delete dbt project](https://fivetran.com/docs/rest-api/dbt-transformation-management#deletedbtproject)
- [Retrieve schema metadata](https://fivetran.com/docs/rest-api/metadata#retrieveschemametadata)
- [Retrieve table metadata](https://fivetran.com/docs/rest-api/metadata#retrievetablemetadata)
- [Retrieve column metadata](https://fivetran.com/docs/rest-api/metadata#retrievecolumnmetadata)
- [Create account webhook](https://fivetran.com/docs/rest-api/webhooks#createaccountwebhook)
- [Create group webhook](https://fivetran.com/docs/rest-api/webhooks#creategroupwebhook)
- [Retrieve webhook details](https://fivetran.com/docs/rest-api/webhooks#retrievewebhookdetails)
- [Update webhook](https://fivetran.com/docs/rest-api/webhooks#updatewebhook)
- [Delete webhook](https://fivetran.com/docs/rest-api/webhooks#deletewebhook)
- [Retrieve the list of webhooks](https://fivetran.com/docs/rest-api/webhooks#retrievethelistofwebhooks)
- [Test webhook](https://fivetran.com/docs/rest-api/webhooks#testwebhook)
- [List all roles](https://fivetran.com/docs/rest-api/roles#listallroles)

## [0.7.6](https://github.com/fivetran/go-fivetran/compare/v0.7.5...v0.7.6)

## Fixes 
- DBT Transformations: `paused` field update issue.

## [0.7.5](https://github.com/fivetran/go-fivetran/compare/v0.7.4...v0.7.5)

## Fixed
- DBT Transformations: `paused` field supported.

## Added
Supported the following Fivetran API endpoints:
- [Create a Log Service](https://fivetran.com/docs/rest-api/log-service-management#createalogservice)
- [Retrieve Log Service Details](https://fivetran.com/docs/rest-api/log-service-management#retrievelogservicedetails)
- [Update a Log Service](https://fivetran.com/docs/rest-api/log-service-management#updatealogservice)
- [Delete a Log Service](https://fivetran.com/docs/rest-api/log-service-management#deletealogservice)
- [Run Log Service Setup Tests](https://fivetran.com/docs/rest-api/log-service-management#runlogservicesetuptests)

## [0.7.4](https://github.com/fivetran/go-fivetran/compare/v0.7.3...v0.7.4)

## Added
- Automatic rate-limiting errors handling
- DBT Transformations API support: create, update, delete, get details

## [0.7.3](https://github.com/fivetran/go-fivetran/compare/v0.7.2...v0.7.3)

## Added
- `DestinationConfigResponse.FivetranRoleArn` missing field added (S3 Data Lake)
- `DestinationConfigResponse.PrefixPath` missing field added (S3 Data Lake)
- `DestinationConfigResponse.Region` missing field added (S3 Data Lake)

## [0.7.2](https://github.com/fivetran/go-fivetran/compare/v0.7.1...v0.7.2)

Fix user/picture set to null issue (#55)
* delete methods for picture and phone
* tests
* refactoring of the picture and phone JSON marshalling + full test
* refactor - move nullableString to common file

## [0.7.1](https://github.com/fivetran/go-fivetran/compare/v0.7.0...v0.7.1) - 2022-12-14

## Fixed
- Connector response should be deserialized even if response code doesn't match expected to provide exact error that API returned.

## [0.7.0](https://github.com/fivetran/go-fivetran/compare/v0.6.10...v0.7.0) - 2022-12-14

## Added
- `ConnectorConfig.ShareUrl` missing field added
- `ConnectorConfig.IsKeypair` missing field added
- `ConnectorConfig.SecretsList` missing field added
- New approach that allows to pass connector configuration as raw map[string]interface{} 
    - `Connector<Operation>Service.ConfigCustom(config map[string]interface{})` method
    - `Connector<Operation>Service.AuthCustom(auth map[string]interface{})` method
    - `ConnectorCustom<Operation>Response` types 
    - `NewConnector<Operation>Service.DoCustom()` methods

## [0.6.10](https://github.com/fivetran/go-fivetran/compare/v0.6.9...v0.6.10) - 2022-11-24

## Fixed
- `ConnectorConfigResponse.UseAPIKeys` wrong type `string` -> `bool`
- `ConnectorConfigResponse.IsSecure` wrong type `string` -> `bool`
- `ConnectorConfigResponse.SkipBefore` wrong type `int` -> `*int`
- `ConnectorConfigResponse.SkipAfter` wrong type `int` -> `*int`

## Added
- `DestinationConfigResponse.Catalog` missing field added (Databricks)

## [0.6.9](https://github.com/fivetran/go-fivetran/compare/v0.6.8...v0.6.9) - 2022-10-04

## Fixed
- `DestinationConfigResponse.IsPrivateKeyEncrypted` wrong type

## [0.6.8](https://github.com/fivetran/go-fivetran/compare/v0.6.7...v0.6.8) - 2022-09-13

## Added
- `DestinationConfigResponse.Role` missing field added (Snowflake)
- `DestinationConfigResponse.IsPrivateKeyEncrypted` missing field added (Snowflake)
- `DestinationConfigResponse.Passphrase` missing field added (Snowflake)

## [0.6.7](https://github.com/fivetran/go-fivetran/compare/v0.6.6...v0.6.7) - 2022-08-29

## Added
- `ConnectorSchemaConfigTable.SyncMode` field that allows to switch table sync mode

## [0.6.6](https://github.com/fivetran/go-fivetran/compare/v0.6.5...v0.6.6) - 2022-08-15

## Fixed
- `DestinationConfigResponse.Location` missing field added (field is used by BQ destination as `data_set_location` field)

## [0.6.5](https://github.com/fivetran/go-fivetran/compare/v0.6.4...v0.6.5) - 2022-08-15

## Fixed
- `DestinationConfigResponse.PublicKey` missing field added (field is readonly and represented only in response)
- `DestinationConfigResponse.PrivateKey` missing field added

## [0.6.4](https://github.com/fivetran/go-fivetran/compare/v0.6.3...v0.6.4) - 2022-07-27

## Added
Mock HttpClient class with a unit test example

## Fixed
- `ConnectorSchemaDetailsResponse.Data.Schemas` type changed: 
    - Old `map[string]ConnectorSchemaConfigSchemaResponse`
    - New `map[string]*ConnectorSchemaConfigSchemaResponse`

## [0.6.3](https://github.com/fivetran/go-fivetran/compare/v0.6.2...v0.6.3) - 2022-07-20

## Fixed
- `ConnectorSchemaConfigTableResponse.EnabledPatchSettings` missing field added
- `ConnectorSchemaConfigTableResponse.NameInDestination` missing field added
- `ConnectorSchemaConfigSchemaResponse.NameInDestination` missing field added

## [0.6.2](https://github.com/fivetran/go-fivetran/compare/v0.6.1...v0.6.2) - 2022-07-20

## Fixed
- `ConnectorConfig.TokenKey` missing field added
- `ConnectorConfig.TokenSecret` missing field added

## [0.6.1](https://github.com/fivetran/go-fivetran/compare/v0.6.0...v0.6.1) - 2022-07-15

## Fixed
- `ConnectorSchemaConfigTableResponse` type accesibility

## [0.6.0](https://github.com/fivetran/go-fivetran/compare/v0.5.11...v0.6.0) - 2022-07-08

## Added
Supported the following Fivetran API endpoints:
- [Retrieve a Connector Schema Config](https://fivetran.com/docs/rest-api/connectors#retrieveaconnectorschemaconfig)
- [Reload a Connector Schema Config](https://fivetran.com/docs/rest-api/connectors#reloadaconnectorschemaconfig)
- [Modify a Connector Schema Config](https://fivetran.com/docs/rest-api/connectors#modifyaconnectorschemaconfig)

## [0.5.11](https://github.com/fivetran/go-fivetran/compare/v0.5.10...v0.5.11) - 2022-07-05

## Fixed
- `ConnectorConfig.PAT` missing field added (Personal Access Token for github connector)

## [0.5.10](https://github.com/fivetran/go-fivetran/compare/v0.5.9...v0.5.10) - 2022-06-16

## Fixed
- `ConnectorConfig.EuRegion` missing field added

## [0.5.9](https://github.com/fivetran/go-fivetran/compare/v0.5.8...v0.5.9) - 2022-06-09

## Fixed
- `ConnectorConfig.PublicationName` missing field added

## [0.5.8](https://github.com/fivetran/go-fivetran/compare/v0.5.7...v0.5.8) - 2022-05-24

## Fixed 
- `DestinationConfigResponse.CreateExternalTables` field type updated

## [0.5.7](https://github.com/fivetran/go-fivetran/compare/v0.5.6...v0.5.7) - 2022-05-13

## Fixed
- `ConnectorConfig.SkipBefore` field transformed to type `int`
- `ConnectorConfig.SkipAfter`  field transformed to type `int`

## [0.5.6](https://github.com/fivetran/go-fivetran/compare/v0.5.5...v0.5.6) - 2022-04-26

## Fixed
- `ConnectorConfig.APIKeys` field transformed to type `[]string`
- `ConnectorConfig.AccountIds` field transfromed to type `[]string`

## [0.5.5](https://github.com/fivetran/go-fivetran/compare/v0.5.4...v0.5.5) - 2022-04-20

## Fixed
- Added `folder_id` missed field to connector config.

## [0.5.4](https://github.com/fivetran/go-fivetran/compare/v0.5.3...v0.5.4) - 2022-02-25

## Fixed
- Added `base_url` missed field to connector config.
- Added `entity_id` missed field to connector config.
- Added `soap_uri` missed field to connector config.
- Added `user_id` missed field to connector config.
- Added `encryption_key` missed field to connector config.

## [0.5.3](https://github.com/fivetran/go-fivetran/compare/v0.5.2...v0.5.3) - 2022-02-21

## Fixed
- Added `api_type` missed field to connector config.

## [0.5.2](https://github.com/fivetran/go-fivetran/compare/v0.5.1...v0.5.2) - 2022-01-31

## Fixed
- Added `is_multi_entity_feature_enabled` missed field to connector config.
- Added `always_encrypted` missed field to connector config.

## [0.5.1](https://github.com/fivetran/go-fivetran/compare/v0.5.0...v0.5.1) - 2022-01-31

## Fixed
- Used `connection_type` key in destination config responses with v2 accept header for consistency.
- Added `connection_type` missed field to connector config.

## [0.5.0](https://github.com/fivetran/go-fivetran/compare/v0.4.0...v0.5.0) - 2022-01-24

## Added
The following fields were added to user resource responses
- UserDetailsResponse.Data.Role - RoleName for user role in account
- UserInviteResponse.Data.Role - RoleName for user role in account
- UserModifyResponse.Data.Role - RoleName for user role in account

## [0.4.0](https://github.com/fivetran/go-fivetran/compare/v0.3.1...v0.4.0) - 2022-01-18

## Added
- E2E tests. 
To be sure that SDK is stable we have added e2e tests which are triggered on each pull request to the main branch. Each e2e test has `E2E` suffix and located in a corresponding `_test.go` file.
- GitHub actions workflow to run tests on a testing account.

## Fixed
- `ConnectorConfigRequest.IsNewPackage` missing field added
- `ConnectorConfigRequest.AdobeAnalyticsConfigurations` missing field added

## [0.3.1](https://github.com/fivetran/go-fivetran/compare/v0.3.0...v0.3.1) - 2021-12-08

## Fixed
- `DestinationConfigRequest.ClusterId` missing field added.
- `DestinationConfigRequest.ClusterRegion` missing field added.

## [0.3.0](https://github.com/fivetran/go-fivetran/compare/v0.2.2...v0.3.0) - 2021-11-10

### Added
- `CustomUserAgent` method for overriding User-Agent header in http-responses (for applications that uses SDK)

## [0.2.2](https://github.com/fivetran/go-fivetran/compare/v0.2.1...v0.2.2) - 2021-09-22

## Fixed
- `DestinationConfigRequest.SecretKey` missing field added.

## [0.2.1](https://github.com/fivetran/go-fivetran/compare/v0.2.0...v0.2.1) - 2021-07-27

## Fixed
- `ConnectorConfigResponse.Port` type is now *int as the response type has been fixed in the REST API v2.

## [0.2.0](https://github.com/fivetran/go-fivetran/compare/v0.1.0...v0.2.0) - 2021-07-16

### Added
- `UsersListResponse.Role`
- `GroupListUsersResponse.Role`
- `ConnectorConfig.AuthType`
- `ConnectorCreateService.SyncFrequency`
- `ConnectorCreateService.DailySyncTime`
- `ConnectorCreateService.PauseAfterTrial`
- `ConnectorCreateResponse.Data.Paused`
- `ConnectorCreateResponse.Data.DailySyncTime`
- `ConnectorCreateResponse.Data.PauseAfterTrial`
- `ConnectorDetailsResponse.Data.Paused`
- `ConnectorDetailsResponse.Data.PauseAfterTrial`
- `ConnectorDetailsResponse.Data.DailySyncTime`
- `ConnectorModifyService.PauseAfterTrial` 

### Changed
- `ConnectorCreateService`, `ConnectorDetailsService`, `ConnectorModifyService`, and `ConnectorSetupTestsService` are now using REST API v2.
- All `int` and `bool` fields of all response types are now `*int` and `*bool`. 

### Removed
- Removed the unnecessary `ConnectorsSourceMetadataResponse.LinkToErd` JSON annotation `omitempty`.

### Fixed
- `DestinationConfigResponse` field `ConnectionType` has changed to `ConnectionMethod` to adequate to the REST API response.

## [0.1.0](https://github.com/fivetran/go-fivetran/releases/tag/v0.1.0) - 2021-07-05

Initial release. 

### Added

- User Management API: List all Users, Retrieve user details, Invite a user, Modify a user, Delete a user.
- Group Management API: Create a group, List all groups, Retrieve group details, Modify a group, List all connectors within a group, List all users within a group, Add a user to a group, Remove a user from a group, Delete a group.
- Certificate Management API: Approve a connector certificate, Approve a connector fingerprint, Approve a destination certificate, Approve a destination fingerprint.
- Destination Management API: Create a destination, Retrieve destination details, Modify a destination, Run destination setup tests, Delete a destination, Destination Config.
- Connector Management API: Retrieve source metadata, Create a connector, Retrieve connector details, Modify a connector, Sync connector data, Re-sync connector table data, Run connector setup tests, Delete a connector, Connector Config, Connector Auth.
