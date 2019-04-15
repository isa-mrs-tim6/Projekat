package main

import (
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Email string `json:"email"`
	Type  string `json:"type"`
	jwt.StandardClaims
}

func Validate(next http.HandlerFunc, userTypes []string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				// If the cookie is not set, return an unauthorized status
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			// For any other type of error, return a bad request status
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Get the JWT string from the cookie
		tknStr := c.Value

		// Initialize a new instance of `Claims`
		claims := &Claims{}

		// Parse the JWT string and store the result in `claims`.
		// Note that we are passing the key in this method as well. This method will return an error
		// if the token is invalid (if it has expired according to the expiry time we set on sign in),
		// or if the signature does not match
		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Validate user type
		if !(contains(userTypes, claims.Type)) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// All is good, route the request forward
		next(w, r)
	})
}

func getEmail(r *http.Request) string {
	c, _ := r.Cookie("token")
	tknStr := c.Value
	claims := &Claims{}
	_, _ = jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return claims.Email
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
