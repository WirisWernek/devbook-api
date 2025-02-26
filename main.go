package main

import (
	"devbook-api/src/router"
	"fmt"
	"net/http"
)

func main() {
	router := router.Gerar()

	fmt.Println("Rodando API")

	http.ListenAndServe(":5000", router)

}
