package model

import "fmt"

type User struct {
	ID   int
	Name string
	Age  int
}

func (u User) String() string {
	return fmt.Sprintf("User: %s, Age: %d", u.Name, u.Age)
}
