package controllers

import (
	"devbook-api/src/banco"
	"devbook-api/src/models"
	"devbook-api/src/repository"
	"devbook-api/src/response"
	"devbook-api/src/security"
	"encoding/json"
	"io"
	"net/http"
)

// Login é responsável por autenticar um usuário na aplicação
func Login(w http.ResponseWriter, r *http.Request) {
	body, erro := io.ReadAll(r.Body)

	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario models.Usuario
	if erro = json.Unmarshal(body, &usuario); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repository.NewRepositoryUsuarios(db)
	usuarioBanco, erro := repositorio.GetByEmail(usuario.Email)

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = security.VerificarSenha(usuario.Senha, usuarioBanco.Senha); erro != nil {
		response.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	response.JSON(w, http.StatusOK, usuarioBanco)
}
