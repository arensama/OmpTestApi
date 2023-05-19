package middlewares

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/arensama/testapi/src/user"
	"github.com/golang-jwt/jwt/v5"
)

type getUserByIDFunc func(id int) (user.User, error)

func validateToken(tokenString string, getUserById getUserByIDFunc) (user.User, error) {
	// Parse the token string and verify the signature
	fmt.Println("validateToken", tokenString)
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// TODO: Load the secret key from a secure location
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	fmt.Println("validateToken claims", token.Claims)

	if err != nil {
		return user.User{}, errors.New("Invalid token")
	}

	// Get the claims from the token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return user.User{}, errors.New("Invalid token claims")
	}
	fmt.Println("validateToken claims2", claims)
	// Check if the token has expired
	// p, _ := claims.Claims.GetExpirationTime()

	// expireTime := *p
	// fmt.println("time", expireTime)
	// if time.Now().Unix() > int64(expireTime) {
	// 	return nil, errors.New("Token has expired")
	// }

	// Get the user from the database by ID

	// s, ok := userID.(userID)
	// if ok != true {
	// 	fmt.Println("s", s, ok)

	// 	return user.User{}, errors.New("jwt malformed")
	// }
	// // userID, _ = strconv.Atoi(s)
	var userID interface{} = claims["id"]
	userIDStr, ok := userID.(string)
	if !ok {
		fmt.Println("userID is not a string")

	}

	if userIDStr == "" {
		fmt.Println("userID is an empty string")

	}

	userIDInt, err := strconv.Atoi(userIDStr)
	if err != nil {
		fmt.Println("failed to convert userID to int:", err)

	}

	fmt.Println("res", userIDInt)
	// if !ok {
	// 	fmt.Println("userID is not an int")
	// 	return user.User{}, errors.New("User not found")
	// }

	// Print the integer value

	// userInstance, err := getUserById(s)
	// if err != nil {
	return user.User{}, errors.New("User not found")
	// }
	// return userInstance, nil
}
func AuthMiddleware(next http.Handler, getUserById getUserByIDFunc) http.Handler {
	fmt.Println("auth middle")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the token from the Authorization header
		token := r.Header.Get("Authorization")

		// Check if the token is valid and belongs to a valid user
		userInstance, err := validateToken(token, getUserById)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"error": "Unauthorized"})
			return
		}

		// Set the user in the context for later use
		ctx := context.WithValue(r.Context(), "user", userInstance)

		// Call the next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
