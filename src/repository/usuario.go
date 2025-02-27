package repository

import (
	"database/sql"
	"devbook-api/src/models"
	"fmt"
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

// BuscarUsuarios busca os usuários com determinado nome ou nick
func (repositorio UsuariosRepository) BuscarUsuarios(nomeOrNick string) ([]models.Usuario, error) {
	nomeOrNick = fmt.Sprintf("%%%s%%", nomeOrNick)
	statement, erro := repositorio.db.Prepare("SELECT u.id, u.nome, u.nick, u.email, criado_em FROM usuarios u WHERE u.nome like $1 or u.nick like $1")

	if erro != nil {
		return nil, erro
	}

	defer statement.Close()

	linhas, erro := statement.Query(nomeOrNick)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []models.Usuario
	for linhas.Next() {
		var usuario models.Usuario
		if erro = linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriadoEm); erro != nil {
			return nil, erro
		}
		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil

}
