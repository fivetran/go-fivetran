package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/fivetran/go-fivetran"
)

func main() {
	apiKey := os.Getenv("FIVETRAN_APIKEY")
	apiSecret := os.Getenv("FIVETRAN_APISECRET")

	tableName := "table"
	columnName := "column"
	schemaName := "schema"
	connectorId := "connectorId"

	fivetran.Debug(true)

	client := fivetran.New(apiKey, apiSecret)

	svc := client.NewConnectorSchemaUpdateService()

	table := fivetran.NewConnectorSchemaConfigTable().Enabled(true).Column(columnName, fivetran.NewConnectorSchemaConfigColumn().Enabled(true))
	schema := fivetran.NewConnectorSchemaConfigSchema().Enabled(true).Table(tableName, table)

	svc.Schema(schemaName, schema)

	value, err := svc.ConnectorID(connectorId).Do(context.Background())
	if err != nil {
		fmt.Printf("%+v\n", value)
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", value)
}
