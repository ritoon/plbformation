package model

import (
	"fmt"
	"time"
)

type User struct {
	ID       int
	DBFields DBFields
	Name     string
	Age      int
}

func NewUser(name string, age int) *User {
	return &User{
		Name: name,
		Age:  age,
	}
}

type DBFields struct {
	ID        int
	CreatedAt string
	UpdatedAt string
}

func (u User) Save() {
	fmt.Printf("Saving user: %s\n", u.Name)
}

func (u User) IsAdult() bool {
	return u.Age >= 18
}

func (u User) Dig(t time.Duration) string {
	time.Sleep(t)
	return fmt.Sprintf("%s dug for %v", u.Name, t)
}
