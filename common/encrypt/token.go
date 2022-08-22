package encrypt

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
)

func Token(secretKey string, expires int64, iat int64, userId int) (string, error) {
	logx.Infof("access token: %s", secretKey, "expires: %d", expires, iat, userId)
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + expires
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
