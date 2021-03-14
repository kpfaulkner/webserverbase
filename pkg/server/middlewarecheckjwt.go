package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"github.com/dgrijalva/jwt-go"
)

type Exception struct {
	Message string `json:"message"`
}

func (s *Server) checkJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("Error with JWT")
					}
					// get secret key from SOMEWHERE... will hardcode for now.
					return []byte("mysecretkey"), nil
				})

				if err != nil {
					json.NewEncoder(w).Encode(Exception{Message: err.Error()})
					return
				}
				if token.Valid {
					ctx := context.WithValue(r.Context(), "decoded", token.Claims)
					next(w, r.WithContext(ctx))
				} else {
					json.NewEncoder(w).Encode(Exception{Message: "Invalid Authorization token"})
				}
			}
		} else {
			json.NewEncoder(w).Encode(Exception{Message: "Missing Authorization header"})
		}
	}
}
