package security

import "golang.org/x/crypto/bcrypt"

// Hash gera o hash de uma string e o devolve
func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

// VeficarSenha compara a senha com o hash e retorna se elas são iguais
func VerificarSenha(senha string, senhaHash string) error {
	return bcrypt.CompareHashAndPassword([]byte(senhaHash), []byte(senha))
}
