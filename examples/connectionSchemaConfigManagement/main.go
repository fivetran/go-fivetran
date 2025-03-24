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
	connectionId := "connectionId"

	fivetran.Debug(true)

	client := fivetran.New(apiKey, apiSecret)

	svc := client.NewConnectionSchemaUpdateService()

	table := fivetran.NewConnectionSchemaConfigTable().Enabled(true).Column(columnName, fivetran.NewConnectionSchemaConfigColumn().Enabled(true))
	schema := fivetran.NewConnectionSchemaConfigSchema().Enabled(true).Table(tableName, table)

	svc.Schema(schemaName, schema)

	value, err := svc.ConnectionID(connectionId).Do(context.Background())
	if err != nil {
		fmt.Printf("%+v\n", value)
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", value)
}
