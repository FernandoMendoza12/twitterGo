package jwt

import (
	"errors"
	"strings"

	"github.com/FernandoMendoza12/twitterGo/models"
	jwt "github.com/golang-jwt/jwt/v5"
)

var Email string
var IDUsuario string

func ProcesoToken(token string, JWTSign string) (*models.Claim, bool, string, error) {
	miClave := []byte(JWTSign)
	var claims models.Claim

	splitToken := strings.Split(token, "Bearer")

	if len(splitToken) != 2 {
		return &claims, false, string(""), errors.New("formato de token invalido")
	}

	token = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if err == nil {
		// Revisa contra la base de datos
	}

	if !tkn.Valid {
		return &claims, false, string(""), errors.New("token invalido")
	}

	return &claims, false, string(""), err
}
