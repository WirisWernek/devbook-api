package models

import (
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// Usuario representa um usuário da rede social
type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criadoEm,omitempty"`
}

// Preparar chama os métodos para validas e formatar o usuário recebido
func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	usuario.formatar()
	return nil

}

func (usuario *Usuario) validar(etapa string) error {
	var erros []error

	if usuario.Nome == "" {
		erros = append(erros, errors.New("Nome é obrigatório e não pode estar em branco"))
	}

	if usuario.Nick == "" {
		erros = append(erros, errors.New("Nick é obrigatório e não pode estar em branco"))
	}

	if usuario.Email == "" {
		erros = append(erros, errors.New("Email é obrigatório e não pode estar em branco"))
	}

	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
		erros = append(erros, errors.New("O formato do email inserido é inválido"))
	}

	if etapa == "cadastro" && usuario.Senha == "" {
		erros = append(erros, errors.New("Senha é obrigatória e não pode estar em branco"))
	}

	if len(erros) > 0 {
		return errors.Join(erros...)
	}

	return nil
}

func (usuario *Usuario) formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Email = strings.TrimSpace(usuario.Email)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Senha = strings.TrimSpace(usuario.Senha)
}
