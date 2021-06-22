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

	resp, err := fivetran.New(apiKey, apiSecret).NewUsersList().Limit(5).Cursor("nextCursor").Do(context.Background())
	if err != nil {
		fmt.Printf("%+v\n", resp)
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", resp)
}
