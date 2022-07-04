package main

import (
	"encoding/json"
	"github.com/cristalhq/jwt/v3"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func ValidateToken(token string) (*jwt.RegisteredClaims, error) {
	key := []byte(os.Getenv("JWT_SECRET"))
	verifier, err := jwt.NewVerifierHS(jwt.HS256, key)
	if err != nil {
		return nil, err
	}

	// parse and verify a token
	newToken, err := jwt.ParseAndVerify([]byte(token), verifier)
	if err != nil {
		return nil, err
	}

	// get Registered claims
	var newClaims jwt.RegisteredClaims
	err = json.Unmarshal(newToken.RawClaims(), &newClaims)
	if err != nil {
		return nil, err
	}

	return &newClaims, nil
}

func Authorized() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Authorization")
		claims, err := ValidateToken(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized - route protected",
			})
			return
		}
		// set claims to context
		ctx.Set("email", claims.Subject)

		ctx.Next()
	}
}
