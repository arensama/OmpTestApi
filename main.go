package main

import (
	"log"
	"net/http"

	"github.com/arensama/testapi/src/auth"
	"github.com/arensama/testapi/src/blog"
	"github.com/arensama/testapi/src/db"
	"github.com/arensama/testapi/src/middlewares"
	"github.com/arensama/testapi/src/user"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var DB = db.Init()

var userService = user.ServiceInit(DB)
var UserController = user.Init(userService)
var authService = auth.ServiceInit(userService)
var AuthController = auth.Init(authService)
var blogService = blog.ServiceInit(userService, DB)
var BlogController = blog.Init(blogService)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	// Create a new router
	router := mux.NewRouter()

	router.PathPrefix("/auth").Handler(AuthController)

	private := router.PathPrefix("/private").Subrouter()
	private.Use(middlewares.AuthMiddleware)

	private.PathPrefix("/user").Handler(UserController)
	private.PathPrefix("/blog").Handler(BlogController)

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", router))
}

// private.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
// 	// Get the user ID from the request URL parameters
// 	// vars := mux.Vars(r)
// 	// user := vars["user"]
// 	fmt.Println("User profile for user", r.Context().Value("user"))
// })
