package user

import (
	"fmt"

	"github.com/arensama/testapi/src/db"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	// Generate a salt with a cost of 10
	salt, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}

	// Generate the hashed password with the salt
	hashedPassword := string(salt)
	return hashedPassword, nil
}

type UserService struct {
	db *db.DB
}

func ServiceInit(db *db.DB) *UserService {
	db.Migrate(User{})
	return &UserService{
		db: db,
	}
}

func (s *UserService) UserLists() ([]User, error) {
	db := s.db.Db
	var users []User
	err := db.Find(&users)
	if err.Error != nil {
		return []User{}, err.Error
	}
	return users, nil
}

func (s *UserService) GetUser(id uint) (User, error) {
	db := s.db.Db
	var user User
	err := db.Find(&user)
	if err.Error != nil {
		return User{}, err.Error
	}
	return user, nil

}
func (s *UserService) GetUserByEmail(email string) (User, error) {
	db := s.db.Db
	var userInstance User
	err := db.Where("email = ?", email).First(&userInstance)
	if err.Error != nil {
		return User{}, err.Error
	}
	return userInstance, nil
}
func (s *UserService) GetUserById(ID uint) (User, error) {
	db := s.db.Db
	var user User
	err := db.Find(&user)
	if err.Error != nil {
		return User{}, err.Error
	}
	return user, nil
}
func (s *UserService) CreateUser(name, surname, email, password string) (User, error) {
	db := s.db.Db
	fmt.Println("passworfd", password)
	password, _ = hashPassword(password)
	fmt.Println("passworfd2", password)
	// Create a new user
	user := User{
		Name:     name,
		Surname:  surname,
		Email:    email,
		Password: password,
	}
	err := db.Create(&user)
	if err.Error != nil {
		return User{}, err.Error
	}
	return user, nil
}

func (s *UserService) UpdateUser(user User) (User, error) {
	db := s.db.Db
	var userI User
	err := db.Find(&userI, user.ID)
	if err.Error != nil {
		return User{}, err.Error
	}
	db.Save(&user)
	if err.Error != nil {
		return User{}, err.Error
	}
	return user, nil
}

func (s *UserService) DeleteUser(id uint) error {
	db := s.db.Db
	err := db.Delete(&User{}, id)
	return err.Error
}
