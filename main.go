package main

import (
	"api-rest-go/database"
	"api-rest-go/routes"
	"fmt"
)

func main() {

	database.ConnectaComBancoDeDados()

	fmt.Println("Iniciando o servidor Rest...")

	routes.HandlerRequest()
}
