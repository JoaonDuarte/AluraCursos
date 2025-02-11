package main

import (
	"curso3/database"
	"curso3/routes"
)

func main() {

	database.ConectaDB()
	routes.HandleRequest()

}
