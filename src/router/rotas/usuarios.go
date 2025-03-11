package rotas

import (
	"devbook-api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		URI:                "/usuarios",
		Metodo:             http.MethodPost,
		Funcao:             controllers.InsertUsuario,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuarios",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUsuarios,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuarios/{usuarioID}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.GetByIdUsuario,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuarios/{usuarioID}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.UpdateUsuario,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuarios/{usuarioID}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeleteUsuario,
		RequerAutenticacao: true,
	},
}
