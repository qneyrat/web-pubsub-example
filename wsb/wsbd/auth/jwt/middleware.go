package jwt

import (
	"chat-example/wsb/wsbd/auth"
	"context"
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

var verifyKey *rsa.PublicKey

func init() {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}

	pubKeyPath := os.Getenv("JWT_KEY")
	log.Println(pubKeyPath)
	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatal(err)
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		log.Fatal(err)
	}
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		keys, ok := r.URL.Query()["token"]
		if !ok || len(keys) < 1 {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, "Url Param 'token' is missing!")
			return
		}

		jwtToken, err := jwt.Parse(keys[0], func(token *jwt.Token) (interface{}, error) { return verifyKey, nil })
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, "Invalid Token!")
			return
		}

		if !jwtToken.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		claims, ok := jwtToken.Claims.(jwt.MapClaims)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(
			r.Context(),
			auth.SessionContextKey,
			auth.Session{Identifier: claims["username"].(string)},
		)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
