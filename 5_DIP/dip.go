package main

import (
	"errors"
	"fmt"
)

// Dependency Inversion Principle
// Dependa de abstrações e não de implementações
// Módulos de alto nível não devem depender de módulos de baixo nível
// Ambos devem depender de abastrações

// domain layer
type User struct {
	ID    int
	Name  string
	Email string
}

type UserRepository interface {
	Insert(user User) (int, error)
	GetByID(id int) (*User, error)
}

// infrastructure layer
type UserDatabaseRepository struct {
	db []User
}

func NewUserDatabaseRepository() UserRepository {
	return &UserDatabaseRepository{}
}

func (r *UserDatabaseRepository) Insert(user User) (int, error) {
	user.ID = len(r.db) + 1
	r.db = append(r.db, user)
	return user.ID, nil
}

func (r *UserDatabaseRepository) GetByID(id int) (*User, error) {
	for _, u := range r.db {
		if u.ID == id {
			return &u, nil
		}
	}
	return nil, errors.New("not found")
}

// application layer
type Service interface {
	SendRegistrationEmail(userID int)
}

type EmailService struct {
	repository UserRepository
}

func NewEmailService(repository UserRepository) Service {
	return &EmailService{
		repository: repository,
	}
}

func (s *EmailService) SendRegistrationEmail(userID int) {
	user, err := s.repository.GetByID(userID)
	if err != nil {
		return
	}
	fmt.Println(fmt.Sprintf("Send e-mail to %s", user.Email))
	// send email
}

func main() {
	repository := NewUserDatabaseRepository()
	emailService := NewEmailService(repository)

	id, err := repository.Insert(User{Name: "Alan", Email: "alan@mercadolivre.com"})
	if err == nil {
		emailService.SendRegistrationEmail(id)
	}
}
