package auth

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type AuthController struct {
	router *mux.Router
	// userService *user.UserService
	authService *AuthService
}
type Signin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Init(authService *AuthService) *AuthController {
	c := AuthController{
		router:      mux.NewRouter(),
		authService: authService,
	}
	c.router.HandleFunc("/auth/signin", c.singin).Methods("POST")
	c.router.HandleFunc("/auth/signup", c.signup).Methods("POST")
	return &c
}
func (c *AuthController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.router.ServeHTTP(w, r)
}

func (c *AuthController) singin(w http.ResponseWriter, r *http.Request) {
	var credential Signin
	if err := json.NewDecoder(r.Body).Decode(&credential); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	result, err := c.authService.Signin(credential.Email, credential.Password)
	if err != nil {
		http.Error(w, "Failed to retrieve users", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result)
}

type Signup struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Password string `json:"password"`
}

func (c *AuthController) signup(w http.ResponseWriter, r *http.Request) {
	var body Signup
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	result, err := c.authService.Signup(body.Name, body.Surname, body.Email, body.Password)
	if err != nil {
		http.Error(w, "Failed to retrieve users", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result)
}
