package middlewares

import (
	"devbook-api/src/auth"
	"devbook-api/src/response"
	"log"
	"net/http"
)

// Autenticar verifica se o usuário que fez a requisição esta autenticado
func Autenticar(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if erro := auth.ValidarToken(r); erro != nil {
			response.Erro(w, http.StatusUnauthorized, erro)
			return
		}
		next(w, r)
	}
}

// Logger registra as requisições no terminal
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\t%s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}
