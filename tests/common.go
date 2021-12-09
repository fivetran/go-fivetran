package tests

import (
	"github.com/fivetran/go-fivetran"
)

func GetClients() map[string]*fivetran.Client {
	apiKey := "_moonbeam_acc_accountworthy_api_key"
	apiSecret := "_moonbeam_acc_accountworthy_api_secret"
	//fivetran.Debug(true)

	clients := make(map[string]*fivetran.Client)

	versions := [...]string{"v1", "v2"};
	for _, version := range versions {
		client := fivetran.New(apiKey, apiSecret);
		client.BaseURL("http://localhost:8001/" + version);
		clients[version] = client;
	}
	return clients;
}





