package main

import (
	"curso3/database"
	"curso3/routes"
	"fmt" 
)

func main() {

	database.ConectaDB()
	fmt.Println("Iniciando o Server")
	routes.HandleRequest()

}
