package blog

import (
	"github.com/arensama/testapi/src/db"
	"github.com/arensama/testapi/src/model"
	"github.com/arensama/testapi/src/user"
)

// func (s *BlogService) AddBlogToUser(userId, blogId uint) (user.User, error) {
// 	db := s.db.Db
// 	var userInstance user.User
// 	err := db.Where(userId).First(&userInstance)
// 	if err.Error != nil {
// 		return user.User{}, errors.New("cant create user ")
// 	}
// 	userblogInstance := Usermodel.Blog{
// 		UserID: userId,
// 		BlogID: blogId,
// 	}
// 	userInstance.UserBlogs = append(userInstance.UserBlogs, user.UserBlog(userblogInstance))
// 	err = db.Save(&userInstance)
// 	return userInstance, err.Error
// }

type BlogService struct {
	userService *user.UserService
	db          *db.DB
}

func ServiceInit(userService *user.UserService, db *db.DB) *BlogService {
	db.Migrate(model.Blog{})
	return &BlogService{
		userService: userService,
		db:          db,
	}
}

func (s *BlogService) BlogLists(limit, page int, req_user model.User) ([]model.Blog, error) {
	db := s.db.Db
	var blogs []model.Blog
	err := db.Preload("User").Limit(limit).Offset((page - 1) * limit).Find(&blogs)
	if err.Error != nil {
		return []model.Blog{}, err.Error
	}
	return blogs, nil
}

func (s *BlogService) CreateBlog(title, body string, req_user model.User) (model.Blog, error) {
	db := s.db.Db
	blogInstance := model.Blog{
		Title:  title,
		Body:   body,
		UserID: req_user.ID,
		User:   req_user,
	}
	err := db.Create(&blogInstance)
	if err.Error != nil {
		return model.Blog{}, err.Error
	}
	// s.AddBlogToUser(req_user.ID, blogInstance.ID)
	return blogInstance, nil
}

func (s *BlogService) UserBlogs(req_user model.User) ([]model.Blog, error) {
	// db := s.db.Db
	// var userInstance user.User
	// // var blogI Blog
	// // _ = db.Preload("User").Where("id = ?", req_user.ID).First(&blogI).Error
	// err := db.Preload("Blogs").Where("id = ?", req_user.ID).First(&userInstance).Error
	// if err != nil {
	// 	return []interfaces.BlogInterface{}, err
	// }
	return []model.Blog{}, nil
}
