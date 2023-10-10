package password

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/kx-boutique/app/auth/service/internal/conf"
	"time"
)

func GenerateJWT(c *conf.App_Jwt, userId uuid.UUID, userEmail string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	now := time.Now()
	claims := token.Claims.(jwt.MapClaims)
	claims["iss"] = "kx-boutique-auth-service"
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()
	claims["exp"] = now.Add(c.Expiration.AsDuration()).Unix()
	claims["sub"] = userEmail
	claims["user_id"] = userId.String()

	tokenString, err := token.SignedString([]byte(c.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
