package controllers

import (
	"devbook-api/src/banco"
	"devbook-api/src/models"
	"devbook-api/src/repository"
	"devbook-api/src/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// InsertUsuarios cadastra um usuario
func InsertUsuario(w http.ResponseWriter, r *http.Request) {
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

	if erro = usuario.Preparar("cadastro"); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	usuarioRepository := repository.NewRepositoryUsuarios(db)
	usuario.ID, erro = usuarioRepository.Insert(usuario)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	response.JSON(w, http.StatusCreated, usuario)

}

// GetAllUsuarios Busca todos os usuarios cadastrados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	nomeOrNick := strings.ToLower(r.URL.Query().Get("usuario"))

	db, erro := banco.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repository.NewRepositoryUsuarios(db)
	usuarios, erro := repositorio.BuscarUsuarios(nomeOrNick)

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusOK, usuarios)
}

// GetAllUsuarios Busca todos os usuarios cadastrados
func GetAllUsuarios(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Teste getAll")
}

// GetByIdUsuario Busca um usuário pelo se id
func GetByIdUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioID"], 10, 64)

	if erro != nil {
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
	usuario, erro := repositorio.GetByIdUsuario(usuarioID)

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusOK, usuario)
}

// UpdateUsuario Atualiza os dados de um Usuario
func UpdateUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioID"], 10, 64)

	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

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

	if erro = usuario.Preparar("edicao"); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repository.NewRepositoryUsuarios(db)
	if erro = repositorio.UpdateUsuario(usuarioID, usuario); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)

}

// DeleteUsuario Exclui um usuario
func DeleteUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioID"], 10, 64)

	if erro != nil {
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
	if erro = repositorio.DeleteUsuario(usuarioID); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
