package lib

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTPayload struct {
	Payload interface{}
	ExpHour int64 // expiration in hour
}

type JWTDataPayload struct {
	Email    string `json:"email"`
	UserCode string `json:"userCode"`
	Token    string `json:"token"`
}

func JwtSign(payload JWTPayload) string {
	// Get verification key

	signedKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(AppConfig.App.PrivateKey))
	if err != nil {
		log.Fatalln(err)
	}

	// Define time expiration
	timeNow := time.Now()
	timeSubtract := time.Duration(payload.ExpHour)
	expDate := timeNow.Add(time.Hour * timeSubtract).UnixMilli()

	// Claim JWT Prop
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"data": payload.Payload,
		"aud":  "graph.teras.work",
		"iss":  "graph.teras.work",
		"exp":  expDate,
	})

	// Sign the JWT
	tokenString, err := token.SignedString(signedKey)

	if err != nil {
		log.Fatalln(err)
	}

	return tokenString
}
