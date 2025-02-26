package main

import (
	"devbook-api/src/config"
	"devbook-api/src/router"
	"fmt"
	"net/http"
)

func main() {
	config.Carregar()

	router := router.Gerar()

	fmt.Println(fmt.Sprintf("Escutando na porta %d", config.Porta))

	http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), router)

}
