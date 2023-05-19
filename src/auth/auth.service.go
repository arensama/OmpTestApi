package auth

import (
	"os"
	"time"

	"github.com/arensama/testapi/src/user"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userService *user.UserService
}

func checkPassword(password, hashedPassword string) error {
	// Compare the password with the hashed password
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}

	return nil
}
func GenerateToken(user user.User, secretKey string, expirationTime time.Time) (string, error) {
	// Create a new token object
	token := jwt.New(jwt.SigningMethodHS256)

	// Set the claims for the token
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["exp"] = expirationTime.Unix()

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
func ServiceInit(userService *user.UserService) *AuthService {
	return &AuthService{
		userService: userService,
	}
}

type SigninRes struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

func (s *AuthService) Signin(email, password string) (SigninRes, error) {
	userInstance, err := s.userService.GetUserByEmail(email)
	if err := checkPassword(password, userInstance.Password); err != nil {
		return SigninRes{}, err
	}
	token, _ := GenerateToken(userInstance, os.Getenv("JWT_SECRET"), time.Now().AddDate(0, 0, 7))
	res := SigninRes{
		Email: userInstance.Email,
		Token: token,
	}
	return res, err
}
func (s *AuthService) Signup() (any, error) {

	return true, nil
}
