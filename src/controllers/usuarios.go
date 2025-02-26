package controllers

import (
	"fmt"
	"net/http"
)

// InsertUsuarios cadastra um usuario
func InsertUsuario(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Teste insert")
}

// GetAllUsuarios Busca todos os usuarios cadastrados
func GetAllUsuarios(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Teste getAll")
}

// GetByIdUsuario Busca um usuário pelo se id
func GetByIdUsuario(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Teste getById")
}

// UpdateUsuario Atualiza os dados de um Usuario
func UpdateUsuario(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Teste update")
}

// DeleteUsuario Exclui um usuario
func DeleteUsuario(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Teste delete")
}
