package main

import (
	"belajar-golang/cmd/ProductService/databases"
	"belajar-golang/cmd/ProductService/routes"
	"log"
)

func main() {
	if err := databases.Database.Connect(); err != nil {
		log.Fatal(err)
	}

	routes.RunRoute()
}
