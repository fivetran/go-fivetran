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
	fivetran.Debug(true)

	client := fivetran.New(apiKey, apiSecret)

	svc := client.NewGroupModify()

	value, err := svc.GroupID("flavorful_guesses").Name("My_new_GRoUP_name").Do(context.Background())
	checkErr(err, value)

	fmt.Printf("%+v\n", value)
}

func checkErr(err error, value interface{}) {
	if err != nil {
		fmt.Printf("%+v\n", value)
		log.Fatal(err)
	}
}
