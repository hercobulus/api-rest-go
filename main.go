package main

import (
	"api-rest-go/models"
	"api-rest-go/routes"
	"fmt"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Nome: "Nome 1", Historia: "História 1"},
		{Nome: "Nome 2", Historia: "História 2"},
	}

	fmt.Println("Iniciando o servidor Rest...")

	routes.HandlerRequest()
}
