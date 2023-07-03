package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/arensama/testapi/src/model"
	"github.com/gorilla/mux"
)

type UserController struct {
	router      *mux.Router
	userService *UserService
}

type User struct {
}

func Init(userService *UserService) *UserController {
	c := UserController{
		router:      mux.NewRouter(),
		userService: userService,
	}
	c.router.HandleFunc("/private/user", c.listUsers).Methods("GET")
	c.router.HandleFunc("/private/user/{id}", c.getUser).Methods("GET")
	c.router.HandleFunc("/private/user", c.createUser).Methods("POST")
	c.router.HandleFunc("/private/user/{id}", c.updateUser).Methods("PUT")
	c.router.HandleFunc("/private/user/{id}", c.deleteUser).Methods("DELETE")
	return &c
}

func (c *UserController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.router.ServeHTTP(w, r)
}

func (c *UserController) listUsers(w http.ResponseWriter, r *http.Request) {
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 10
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 1
	}
	users, err := c.userService.UserLists(limit, page)
	if err != nil {
		http.Error(w, "Failed to retrieve users", http.StatusInternalServerError)
		return
	}
	fmt.Println("get user")
	json.NewEncoder(w).Encode(users)
}

func (c *UserController) getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	user, err := c.userService.GetUser(uint(id))
	if err != nil {
		http.NotFound(w, r)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (c *UserController) createUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	createdUser, err := c.userService.CreateUser(user.Name, user.Surname, user.Email, user.Password)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
}

func (c *UserController) updateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	var updatedUser model.User
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	updatedUser.ID = uint(id)
	user, err := c.userService.UpdateUser(updatedUser)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (c *UserController) deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	if err := c.userService.DeleteUser(uint(id)); err != nil {
		http.NotFound(w, r)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
