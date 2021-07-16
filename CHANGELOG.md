# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased](https://github.com/fivetran/go-fivetran/compare/v0.2.0...HEAD)

## [0.2.0](https://github.com/fivetran/go-fivetran/releases/tag/v0.1.0...v0.2.0) - 2021-07-16

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
