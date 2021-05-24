package main

import (
	"fmt"

	"github.com/fivetran/go-fivetran"
)

func main() {
	dc := fivetran.NewDestinationConfig().
		Host("myhostname.com").
		Port(123).Database("myDatabase").
		User("myUser").
		Password("myPassword")

	fmt.Printf("%+v\n", dc)
}
