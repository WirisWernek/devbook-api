package banco

import (
	"database/sql"
	"devbook-api/src/config"

	_ "github.com/lib/pq" // Driver de Conexão com o PostgreSQL
)

// Conectar abre a conexao com o bando de dados e a retorna
func Conectar() (*sql.DB, error) {
	db, erro := sql.Open("postgres", config.StringConexao)

	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil
}
