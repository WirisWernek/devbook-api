package router

import (
	"devbook-api/src/router/rotas"

	"github.com/gorilla/mux"
)

// Gerar vai retornar o router com as rotas configuradas
func Gerar() *mux.Router {
	router := mux.NewRouter()
	return rotas.Configurar(router)
}
