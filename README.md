# Fivetran SDK for Go

`go-fivetran` is the official Fivetran SDK for the Go programming language.

Checkout our [CHANGELOG](CHANGELOG.md) for information about the latest bug fixes, updates, and features added to the SDK. 

Make sure you read the Fivetran REST API [documentation](https://fivetran.com/docs/rest-api) before using the SDK.

## Installation

```
go get github.com/fivetran/go-fivetran
```

## Importing

```go
import (
    "github.com/fivetran/go-fivetran"
)
```

## Getting started



## API List

The following [Fivetran REST API](https://fivetran.com/docs/rest-api) v1 endpoints are implemented by the Fivetran SDK for Go: 

### [User Management API](https://fivetran.com/docs/rest-api/users)

REST API Endpoint | Service
--- | ---
[List all Users](https://fivetran.com/docs/rest-api/users#listallusers) | UsersListService
[Retrieve user details](https://fivetran.com/docs/rest-api/users#retrieveuserdetails) | UserDetailsService
[Invite a user](https://fivetran.com/docs/rest-api/users#inviteauser) | UserInviteService 
[Modify a user](https://fivetran.com/docs/rest-api/users#modifyauser) | UserModifyService
[Delete a user](https://fivetran.com/docs/rest-api/users#deleteauser) | UserDeleteService

### [Group Management API](https://fivetran.com/docs/rest-api/groups)

REST API Endpoint | Service
--- | ---
[Create a group](https://fivetran.com/docs/rest-api/groups#createagroup) | GroupCreateService
[List all groups](https://fivetran.com/docs/rest-api/groups#listallgroups) | GroupsListService
[Retrieve group details](https://fivetran.com/docs/rest-api/groups#retrievegroupdetails) | GroupDetailsService
[Modify a group](https://fivetran.com/docs/rest-api/groups#modifyagroup) | GroupModifyService
[List all connectors within a group](https://fivetran.com/docs/rest-api/groups#listallconnectorswithinagroup) | GroupListConnectorsService
[List all users within a group](https://fivetran.com/docs/rest-api/groups#listalluserswithinagroup) | GroupListUsersService
[Add a user to a group](https://fivetran.com/docs/rest-api/groups#addausertoagroup) | GroupAddUserService
[Remove a user from a group](https://fivetran.com/docs/rest-api/groups#removeauserfromagroup) | GroupRemoveUserService
[Delete a group](https://fivetran.com/docs/rest-api/groups#deleteagroup) | GroupDeleteService

### [Destination Management API](https://fivetran.com/docs/rest-api/destinations)

REST API Endpoint | Service/Config
--- | ---
[Create a destination](https://fivetran.com/docs/rest-api/destinations#createadestination) | DestinationCreateService
[Retrieve destination details](https://fivetran.com/docs/rest-api/destinations#retrievedestinationdetails) | DestinationDetailsService
[Modify a destination](https://fivetran.com/docs/rest-api/destinations#modifyadestination) | DestinationModifyService
[Run destination setup tests](https://fivetran.com/docs/rest-api/destinations#rundestinationsetuptests) | DestinationSetupTestsService
[Delete a destination](https://fivetran.com/docs/rest-api/destinations#deleteadestination) | DestinationDeleteService
[Destination Config](https://fivetran.com/docs/rest-api/destinations/config) | DestinationConfig 

### [Connector Management API](https://fivetran.com/docs/rest-api/connectors)

REST API Endpoint | Service/Config/Auth
--- | ---
[Retrieve source metadata](https://fivetran.com/docs/rest-api/connectors#retrievesourcemetadata) | ConnectorsSourceMetadataService
[Create a connector](https://fivetran.com/docs/rest-api/connectors#createaconnector) | ConnectorCreateService
[Retrieve connector details](https://fivetran.com/docs/rest-api/connectors#retrieveconnectordetails) | ConnectorDetailsService
[Modify a connector](https://fivetran.com/docs/rest-api/connectors#modifyaconnector) | ConnectorModifyService
[Sync connector data](https://fivetran.com/docs/rest-api/connectors#syncconnectordata) | ConnectorSyncService
[Re-sync connector table data](https://fivetran.com/docs/rest-api/connectors#resyncconnectortabledata) | ConnectorReSyncTableService
[Run connector setup tests](https://fivetran.com/docs/rest-api/connectors#runconnectorsetuptests) | ConnectorSetupTestsService
[Delete a connector](https://fivetran.com/docs/rest-api/connectors#deleteaconnector) | ConnectorDeleteService
[Retrieve a connector schema config](https://fivetran.com/docs/rest-api/connectors#retrieveaconnectorschemaconfig) | not implemented
[Retrieve source table columns config](https://fivetran.com/docs/rest-api/connectors#retrievesourcetablecolumnsconfig) | not implemented
[Reload a connector schema config](https://fivetran.com/docs/rest-api/connectors#reloadaconnectorschemaconfig) | not implemented
[Modify a connector schema config](https://fivetran.com/docs/rest-api/connectors#modifyaconnectorschemaconfig) | not implemented
[Modify a connector database schema config](https://fivetran.com/docs/rest-api/connectors#modifyaconnectordatabaseschemaconfig) | not implemented
[Modify a connector table config](https://fivetran.com/docs/rest-api/connectors#modifyaconnectortableconfig) | not implemented
[Modify a connector column config](https://fivetran.com/docs/rest-api/connectors#modifyaconnectorcolumnconfig) | not implemented
[Connector Config](https://fivetran.com/docs/rest-api/connectors/config) | ConnectorConfig<br> ConnectorConfigReports<br> ConnectorConfigProjectCredentials<br> ConnectorConfigCustomTables
[Connector Auth](https://fivetran.com/docs/rest-api/connectors) | ConnectorAuth<br> ConnectorAuthClientAccess
[Connect Card](https://fivetran.com/docs/rest-api/connectors/connect-card) | not implemented

### [Certificate Management API](https://fivetran.com/docs/rest-api/certificates)
REST API Endpoint | Service
--- | ---
[Approve a connector certificate](https://fivetran.com/docs/rest-api/certificates#approveaconnectorcertificate) | CertificateConnectorCertificateApproveService
[Approve a connector fingerprint](https://fivetran.com/docs/rest-api/certificates#approveaconnectorfingerprint) | CertificateConnectorFingerprintApproveService
[Approve a destination certificate](https://fivetran.com/docs/rest-api/certificates#approveadestinationcertificate) | CertificateDestinationCertificateApproveService
[Approve a destination fingerprint](https://fivetran.com/docs/rest-api/certificates#approveadestinationfingerprint) | CertificateDestinationFingerprintApproveService
