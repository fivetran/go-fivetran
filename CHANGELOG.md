# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased](https://github.com/fivetran/go-fivetran/compare/v0.5.4...HEAD)

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
