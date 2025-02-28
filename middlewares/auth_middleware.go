package middlewares

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type BstackCustomClaims struct {
	Username string `json:"username"`
	NanoId string `json:"nanoId"`
	jwt.RegisteredClaims
}

func CheckUserAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jwtSecret := os.Getenv("JWT_SECRET")

		cookie, err := r.Cookie("bstack_token")
		if err != nil {
			log.Printf("Auth failed: %v", err.Error())
			next.ServeHTTP(w, r)
			return
		}
		
		token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(jwtSecret), nil
		})
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			ctx := context.WithValue(r.Context(), "username", claims["username"]);
			ctx = context.WithValue(ctx, "nanoId", claims["nanoId"]);
			newRequest := r.WithContext(ctx)

			next.ServeHTTP(w, newRequest)
			return
		} else {
			next.ServeHTTP(w, r)
			return
		}
	})
}

