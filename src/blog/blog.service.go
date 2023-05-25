package blog

import (
	"errors"
	"time"

	"github.com/arensama/testapi/src/user"
)

type BlogService struct {
	userService *user.UserService
	blogs       []Blog
}

func paginate(arr []Blog, pageSize int, pageNumber int) []Blog {
	startIndex := (pageNumber - 1) * pageSize
	endIndex := pageNumber * pageSize

	if startIndex >= len(arr) {
		return []Blog{}
	}

	if endIndex > len(arr) {
		endIndex = len(arr)
	}

	return arr[startIndex:endIndex]
}

func ServiceInit(userService *user.UserService) *BlogService {
	return &BlogService{
		userService: userService,
	}
}

func (s *BlogService) ListBlogs(limit, page int, req_user user.User) ([]Blog, error) {
	result := paginate(s.blogs, limit, page)
	return result, nil
}

func (s *BlogService) UserBlogs(req_user user.User) ([]Blog, error) {
	userInstance, err := s.userService.GetUserById(req_user.ID)
	if err != nil {
		return []Blog{}, errors.New("user not found")
	}

	var result []Blog

	for _, i := range userInstance.Blogs {
		result = append(result, s.blogs[i])
	}

	return result, nil
}

func (s *BlogService) CreateBlog(title, body string, req_user user.User) (Blog, error) {
	// userInstance, err := s.userService.GetUserById(req_user.ID)
	id := len(s.blogs) + 1
	blogInstance := Blog{
		ID:        id,
		Title:     title,
		Body:      body,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	s.blogs = append(s.blogs, blogInstance)
	s.userService.AddBlogToUser(req_user.ID, id)
	return blogInstance, nil
}
