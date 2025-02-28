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

// func init() {
// 	chave := make([]byte, 64)
// 	if _, erro := rand.Read(chave); erro != nil {
// 		log.Fatal(erro)
// 	}
// 	stringBase64 := base64.StdEncoding.EncodeToString(chave)
// 	fmt.Print(stringBase64)
// }
