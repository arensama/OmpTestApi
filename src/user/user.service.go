package user

import (
	"errors"

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
	users []User
}

func ServiceInit() *UserService {
	pass, _ := hashPassword("password")
	return &UserService{
		users: []User{
			{
				ID:       1,
				Name:     "amirreza",
				Surname:  "namazi",
				Email:    "amirreza@gmail.com",
				Password: pass,
			},
		},
	}
}

func (s *UserService) ListUsers() ([]User, error) {
	return s.users, nil
}

func (s *UserService) GetUser(id int) (User, error) {
	for _, user := range s.users {
		if user.ID == id {
			return user, nil
		}
	}
	return User{}, errors.New("user not found")
}
func (s *UserService) GetUserByEmail(email string) (User, error) {
	for _, user := range s.users {
		if user.Email == email {
			return user, nil
		}
	}
	return User{}, errors.New("user not found")
}
func (s *UserService) GetUserById(ID int) (User, error) {
	for _, user := range s.users {
		if user.ID == ID {
			return user, nil
		}
	}
	return User{}, errors.New("user not found")
}
func (s *UserService) CreateUser(name, surname, email, password string) (User, error) {
	id := len(s.users) + 1
	password, _ = hashPassword(password)
	user := User{ID: id, Name: name, Surname: surname, Email: email, Password: password}
	s.users = append(s.users, user)
	return user, nil
}

func (s *UserService) UpdateUser(user User) (User, error) {
	for i, u := range s.users {
		if u.ID == user.ID {
			s.users[i] = user
			return user, nil
		}
	}
	return User{}, errors.New("user not found")
}

func (s *UserService) DeleteUser(id int) error {
	for i, user := range s.users {
		if user.ID == id {
			s.users = append(s.users[:i], s.users[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}
