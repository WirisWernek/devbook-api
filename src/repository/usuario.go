package repository

import (
	"database/sql"
	"devbook-api/src/models"
)

// UsuariosRepository representa um repository de usuarios
type UsuariosRepository struct {
	db *sql.DB
}

// NewRepositoryUsuarios cria um repositório de usuários
func NewRepositoryUsuarios(db *sql.DB) *UsuariosRepository {
	return &UsuariosRepository{db}
}

// Insert insere um usuário no banco de dados
func (repositorio UsuariosRepository) Insert(usuario models.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare("INSERT INTO usuarios (nome, nick, email, senha) VALUES($1, $2, $3, $4) RETURNING id")

	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	lastID := 0
	erro = statement.QueryRow(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha).Scan(&lastID)

	if erro != nil {
		return 0, erro
	}

	return uint64(lastID), nil
}
