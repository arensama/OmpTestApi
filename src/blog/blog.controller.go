package blog

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/arensama/testapi/src/model"
	"github.com/gorilla/mux"
)

type BlogController struct {
	router *mux.Router
	// userService *user.UserService
	blogService *BlogService
}

type Blog struct {
}

func Init(blogService *BlogService) *BlogController {
	c := BlogController{
		router:      mux.NewRouter(),
		blogService: blogService,
	}
	c.router.HandleFunc("/private/blog", c.listBlogs).Methods("GET")
	c.router.HandleFunc("/private/blog/self", c.userBlogs).Methods("GET")
	c.router.HandleFunc("/private/blog", c.createBlog).Methods("POST")
	return &c
}
func (c *BlogController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.router.ServeHTTP(w, r)
}
func (c *BlogController) listBlogs(w http.ResponseWriter, r *http.Request) {
	req_user := r.Context().Value("user")
	vars := mux.Vars(r)
	limit, err := strconv.Atoi(vars["limit"])
	if err != nil {
		limit = 10
	}
	page, err := strconv.Atoi(vars["page"])
	if err != nil {
		page = 1
	}
	blogs, err := c.blogService.BlogLists(limit, page, req_user.(model.User))
	if err != nil {
		http.Error(w, "Failed to retrieve blogs", http.StatusInternalServerError)
		return
	}
	fmt.Println("get blog")
	json.NewEncoder(w).Encode(blogs)
}
func (c *BlogController) userBlogs(w http.ResponseWriter, r *http.Request) {

	req_user := r.Context().Value("user")
	blogs, err := c.blogService.UserBlogs(req_user.(model.User))
	if err != nil {
		http.Error(w, "Failed to retrieve blogs", http.StatusInternalServerError)
		return
	}
	fmt.Println("get blog/self")
	json.NewEncoder(w).Encode(blogs)
}
func (c *BlogController) createBlog(w http.ResponseWriter, r *http.Request) {
	req_user := r.Context().Value("user")

	var blog model.Blog
	if err := json.NewDecoder(r.Body).Decode(&blog); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	createdBlog, err := c.blogService.CreateBlog(blog.Title, blog.Body, req_user.(model.User))
	if err != nil {
		http.Error(w, "Failed to create blog", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdBlog)
}
