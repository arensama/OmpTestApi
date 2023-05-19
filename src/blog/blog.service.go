package blog

import "time"

type BlogService struct {
	blogs []Blog
}

func ServiceInit() *BlogService {
	return &BlogService{}
}

func (s *BlogService) ListBlogs() ([]Blog, error) {
	return s.blogs, nil
}

func (s *BlogService) CreateBlog(title, body string) (Blog, error) {
	id := len(s.blogs) + 1
	blug := Blog{
		ID:        id,
		Title:     title,
		Body:      body,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	s.blogs = append(s.blogs, blug)
	return blug, nil
}
