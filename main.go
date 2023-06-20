package main

import (
	"fmt"

	struct_configuration "github.com/jalal-akbar/anthonygg-tutorial/struct-configuration"
)

func main() {
	// Struct Configuration Pattern
	fmt.Println("The Most Efficient Struct Configuration Pattern For Golang")

	server := struct_configuration.NewServer(struct_configuration.WitTLS, struct_configuration.WithMaxConn(99))
	fmt.Printf("%+v \n", server)
}
