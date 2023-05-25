package middlewares

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/arensama/testapi/src/user"
	"github.com/golang-jwt/jwt/v5"
)

type getUserByIDFunc func(id int) (user.User, error)

func validateToken(tokenString string) (user.User, error) {
	// Parse the token string and verify the signature
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// TODO: Load the secret key from a secure location
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return user.User{}, errors.New("invalid token")
	}

	// Get the claims from the token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return user.User{}, errors.New("invalid token claims")
	}
	// Check if the token has expired

	// p, _ := claims.GetExpirationTime()
	// expireTime := *p
	// fmt.Println("time", expireTime)
	// if time.Now().Unix() > int64(expireTime.Second()) {
	// 	return user.User{}, errors.New("token has expired")
	// }
	userInstance := user.User{}
	userInstance.ID = int(claims["id"].(float64))
	userInstance.Name = string(claims["name"].(string))
	userInstance.Email = string(claims["email"].(string))
	return userInstance, nil
}
func AuthMiddleware(next http.Handler) http.Handler {
	fmt.Println("auth middle")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the token from the Authorization header
		token := r.Header.Get("Authorization")

		// Check if the token is valid and belongs to a valid user
		userInstance, err := validateToken(token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"error": "Unauthorized"})
			return
		}
		fmt.Println("authed", userInstance)
		// Set the user in the context for later use
		ctx := context.WithValue(r.Context(), "user", userInstance)
		// Call the next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
