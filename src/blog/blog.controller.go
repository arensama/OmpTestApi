package blog

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type BlogController struct {
	router *mux.Router
	// userService *user.UserService
	blogService *BlogService
}
type Blog struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func Init(blogService *BlogService) *BlogController {
	c := BlogController{
		router:      mux.NewRouter(),
		blogService: blogService,
	}
	c.router.HandleFunc("/private/blog", c.listBlogs).Methods("GET")
	c.router.HandleFunc("/private/blog", c.createBlog).Methods("POST")
	return &c
}
func (c *BlogController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.router.ServeHTTP(w, r)
}
func (c *BlogController) listBlogs(w http.ResponseWriter, r *http.Request) {
	blogs, err := c.blogService.ListBlogs()
	if err != nil {
		http.Error(w, "Failed to retrieve blogs", http.StatusInternalServerError)
		return
	}
	fmt.Println("get blog")
	json.NewEncoder(w).Encode(blogs)
}
func (c *BlogController) createBlog(w http.ResponseWriter, r *http.Request) {
	var blog Blog
	if err := json.NewDecoder(r.Body).Decode(&blog); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	createdBlog, err := c.blogService.CreateBlog(blog.Title, blog.Body)
	if err != nil {
		http.Error(w, "Failed to create blog", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdBlog)
}
