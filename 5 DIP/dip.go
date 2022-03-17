package main

import (
	"errors"
	"fmt"
)

// Dependency Inversion Principle
// Módulos de alto nível não devem depender de módulos de baixo nível
// Ambos devem depender de abastrações

// domain layer
type User struct {
	ID uint
}

type UserRepository interface {
	Insert(user UserDB) (uint, error)
	GetByID(id uint) (*User, error)
}

// infrastructure layer
type UserDB struct {
	ID uint `json:"id"`
}

func (u *UserDB) ToUser() *User {
	return &User{
		ID: u.ID,
	}
}

func (u *User) ToUserDB() *UserDB {
	return &UserDB{
		ID: u.ID,
	}
}

type UserDatabaseRepository struct {
	db []UserDB
}

/*
type AnotherRepository struct {

}
type AnotherRepository struct {
}
*/

func NewUserDatabaseRepository() UserRepository {
	return &UserDatabaseRepository{}
}

func (r *UserDatabaseRepository) Insert(user UserDB) (uint, error) {
	//add
	return user.ID, nil
}

func (r *UserDatabaseRepository) GetByID(id uint) (*User, error) {
	for _, u := range r.db {
		if u.ID == id {
			return u.ToUser(), nil
		}
	}
	return nil, errors.New("not found")
}

// application layer
type Service interface {
	SendRegistrationEmail(userID uint) error
}

type EmailService struct {
	repository UserRepository
}

func NewEmailService(repository UserRepository) Service {
	return &EmailService{
		repository: repository,
	}
}

func (s *EmailService) SendRegistrationEmail(userID uint) error {
	user, err := s.repository.GetByID(userID)
	if err != nil {
		return err
	}
	fmt.Println(user)
	// send email
	return nil
}

func main() {
	user := User{}

	repository := NewUserDatabaseRepository()
	id, _ := repository.Insert(*user.ToUserDB())

	service := NewEmailService(repository)
	_ = service.SendRegistrationEmail(id)
}
