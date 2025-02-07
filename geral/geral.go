package geral

import (
	"curso3/database"
	"curso3/routes"
	"fmt"
)

func Geral() {
	database.ConectaDB()
	fmt.Println("Iniciando o Server")
	routes.HandleRequest()
}
