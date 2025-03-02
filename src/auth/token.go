package auth

import (
	"devbook-api/src/config"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// GerarToken gera o token JWT para um usuário com suas permissões
func GerarToken(usuarioID uint64) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["usuarioID"] = usuarioID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
	return token.SignedString(config.SecretKey)
}
