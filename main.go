package main

import (
	"api-rest-go/routes"
	"fmt"
)

func main() {
	fmt.Println("Iniciando o servidor Rest...")
	routes.HandlerRequest()
}
