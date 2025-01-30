package types

import "time"

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(User) error
}

type User struct {
	ID        int       `json: "id"`
	FirstName string    `json: "firstName"`
	LastName  string    `json: "lastName"`
	Email     string    `json: "email"`
	Password  string    `json: "-"`
	CreatedAt time.Time `json: "createdAt"`
}

type Product struct {
	ID          int       `json: "id"`
	Name        string    `json: "firstName"`
	Description string    `json: "lastName"`
	Image       string    `json: "email"`
	Quantity    string    `json: "-"`
	CreatedAt   time.Time `json: "createdAt"`
}

type Order struct {
	ID        int       `json: "id"`
	UserID    string    `json: "firstName"`
	Total     string    `json: "lastName"`
	Status    string    `json: "email"`
	Address   string    `json: "-"`
	CreatedAt time.Time `json: "createdAt"`
}

type Orders_Item struct {
	ID        int    `json: "id"`
	OrderID   string `json: "firstName"`
	ProductID string `json: "lastName"`
	Quantity  string `json: "email"`
	Price     string `json: "-"`
}

type RegisterUserPayload struct {
	FirstName string `json: "firstName" validate: "required"`
	LastName  string `json: "lastName" validate: "required"`
	Email     string `json: "email" validate: "required,email"`
	Password  string `json: "password" validate: "required,min=5,max=30"`
}
