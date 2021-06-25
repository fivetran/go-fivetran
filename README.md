# Fivetran SDK for Go

`go-fivetran` is the official Fivetran SDK for the Go programming language. It uses the [Fivetran REST API](https://fivetran.com/docs/rest-api) v1.

Checkout our [CHANGELOG](CHANGELOG.md) for information about the latest bug fixes, updates, and features added to the SDK. 



## API List

### User Management API

REST API Endpoint | Service
--- | ---
[List all Users](https://fivetran.com/docs/rest-api/users#listallusers) | UsersListService
[Retrieve user details](https://fivetran.com/docs/rest-api/users#retrieveuserdetails) | UserDetailsService
[Invite a user](https://fivetran.com/docs/rest-api/users#inviteauser) | UserInviteService 
[Modify a user](https://fivetran.com/docs/rest-api/users#modifyauser) | UserModifyService
[Delete a user](https://fivetran.com/docs/rest-api/users#deleteauser) | UserDeleteService

### Group Management API

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

### Destination Management API

REST API Endpoint | Service/Config
--- | ---
[Create a destination](https://fivetran.com/docs/rest-api/destinations#createadestination) | DestinationCreateService
[Retrieve destination details](https://fivetran.com/docs/rest-api/destinations#retrievedestinationdetails) | DestinationDetailsService
[Modify a destination](https://fivetran.com/docs/rest-api/destinations#modifyadestination) | DestinationModifyService
[Run destination setup tests](https://fivetran.com/docs/rest-api/destinations#rundestinationsetuptests) | DestinationSetupTestsService
[Delete a destination](https://fivetran.com/docs/rest-api/destinations#deleteadestination) | DestinationDeleteService
[Destination Config](https://fivetran.com/docs/rest-api/destinations/config) | DestinationConfig 

### Connector Management API

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
[Connector Config](https://fivetran.com/docs/rest-api/connectors/config) | ConnectorConfig, ConnectorConfigReports, ConnectorConfigProjectCredentials, ConnectorConfigCustomTables
[Connector Auth](https://fivetran.com/docs/rest-api/connectors) | ConnectorAuth, ConnectorAuthClientAccess


- Certificate Management API: Approve a connector certificate, Approve a connector fingerprint, Approve a destination certificate, Approve a destination fingerprint.
