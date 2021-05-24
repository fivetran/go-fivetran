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

	client := fivetran.NewClient(apiKey, apiSecret)

	svc := client.NewGroupAddUserService()

	svc.ID("replying_ministry")
	svc.Email("sometestuser@fivetran.com")
	svc.Role("ReadOnly")

	value, err := svc.Do(context.Background())
	checkErr(err, value)

	fmt.Printf("%+v\n", value)
}

func checkErr(err error, value interface{}) {
	if err != nil {
		fmt.Printf("%+v\n", value)
		log.Fatal(err)
	}
}
