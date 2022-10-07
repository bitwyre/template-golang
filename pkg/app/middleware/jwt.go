package middleware

import (
	"crypto/rsa"
	"fmt"
	"net/http"
	"strings"

	"github.com/bitwyre/template-golang/pkg/helper"
	"github.com/bitwyre/template-golang/pkg/lib"
	"github.com/bitwyre/template-golang/pkg/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// RSAKey Store RSA Public & Private key into struct
// Fetch the key during runtime, preventing from re-retrieving from each request
type RSAKey struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

var tracer opentracing.Tracer

func JWTMiddleware() gin.HandlerFunc {
	var rsaKey = getRSAKey()
	return func(c *gin.Context) {
		helper.Block{
			Try: func() {
				bearerToken := c.GetHeader("authorization")
				parsedToken := parseToken(bearerToken)
				validateToken(parsedToken, rsaKey.publicKey)

				// TODO Get UUID from DB

				// TODO Check is allow to trade

				c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
				c.Next()
			},
			Catch: func(e helper.Exception) {
				logrus.Errorf("Caught %v", e)
				helper.HttpErrorResponse(http.StatusUnauthorized, model.BaseErrorResponseSchema{
					Code:    "UNAUTHORIZED",
					Message: "Invalid authentication header",
				}, c)
				c.Abort()
			},
			Finally: nil,
		}.Do()
	}
}

// Parse JWT Token from Request Header
// Return string to be used by validateToken() later
func parseToken(bearerToken string) string {
	splitToken := strings.Split(bearerToken, " ")
	if bearerToken == "" || len(splitToken) < 2 {
		helper.Throw("Invalid JWT Format")
	}

	return splitToken[1]
}

// getRSAKey Get RSA Public & Private Key
// Return RSAKey struct to be used by validateToken() later
func getRSAKey() *RSAKey {
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(lib.AppConfig.App.PublicKey))
	if err != nil {
		helper.Throw("Cannot read JWT private key")
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(lib.AppConfig.App.PrivateKey))
	if err != nil {
		helper.Throw("Cannot read JWT private key")
	}

	return &RSAKey{
		privateKey: privateKey,
		publicKey:  publicKey,
	}
}

// validateToken Do Validate JWT Token
func validateToken(jwtToken string, publicKey *rsa.PublicKey) {
	token, err := jwt.Parse(jwtToken, func(tok *jwt.Token) (interface{}, error) {
		if _, ok := tok.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", tok.Header["alg"])
		}
		return publicKey, nil
	})

	if err != nil {
		helper.Throw(err)
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		helper.Throw(ok)
	}
}
