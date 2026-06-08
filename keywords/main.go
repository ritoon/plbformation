package main

import "fmt"

func main() {
	var test int = 10
	var u User
	u.ID = 1
	u.dbFields.CreatedAt = "2024-06-01"
	u.dbFields.UpdatedAt = "2024-06-02"
	u.Name = "Alice"
	u.Age = 30

	u.Save()
	fmt.Println("Hello, World!", test)
}

type User struct {
	ID       int
	dbFields DBFields
	Name     string
	Age      int
}

type DBFields struct {
	ID        int
	CreatedAt string
	UpdatedAt string
}

func (u User) Save() {
	fmt.Printf("Saving user: %s\n", u.Name)
}
