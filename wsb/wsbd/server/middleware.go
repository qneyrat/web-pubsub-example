package server

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

const pubKeyPath = "config/jwt/public.pem"

var verifyKey *rsa.PublicKey

func init() {
	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatal(err)
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		log.Fatal(err)
	}
}

func jwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "No token found!", http.StatusUnauthorized)

			return
		}

		authHeaderParts := strings.Split(authHeader, " ")
		if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
			http.Error(w, "Authorization header format must be Bearer {token}!", http.StatusUnauthorized)

			return
		}
		authToken := authHeaderParts[1]
		jwtToken, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
			return verifyKey, nil
		})

		switch err.(type) {
		case nil:
			if !jwtToken.Valid {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintln(w, "Invalid Token!")

				return
			}

			log.Printf("Token:%+v\n", jwtToken)
		case *jwt.ValidationError:
			vErr := err.(*jwt.ValidationError)

			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintln(w, "Token Expired, get a new one.")

				return

			default:
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintln(w, "Error while Parsing Token!")
				log.Printf("ValidationError error: %+v\n", vErr.Errors)

				return
			}

		default:
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "Error while Parsing Token!")
			log.Printf("Token parse error: %v\n", err)

			return
		}

		if claims, ok := jwtToken.Claims.(jwt.MapClaims); ok && jwtToken.Valid {
			log.Printf("username: %s\n", claims["username"])
		}

		next.ServeHTTP(w, r)
	})
}
