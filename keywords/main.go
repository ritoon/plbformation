package main

import (
	"fmt"
	"time"

	"plbformation/keywords/model"
)

// `break`   `default`   `func`   `interface`   `select`
// `case`   `defer`   `go`   `map`   `struct`
// `chan`   `else`   `goto`   `package`   `switch`
// `const`   `fallthrough`   `if`   `range`   `type`
// `continue`   `for`   `import`   `return`   `var`

func main() {
	var test int = 10
	var u model.User
	u.ID = 1
	u.DBFields.CreatedAt = "2024-06-01"
	u.DBFields.UpdatedAt = "2024-06-02"
	u.Name = "Alice"
	u.Age = 30

	u.Save()
	fmt.Println("Hello, World!", test)

	if u.IsAdult() {
		fmt.Printf("%s is an adult.\n", u.Name)
	} else {
		fmt.Printf("%s is not an adult.\n", u.Name)
	}

	result := u.Dig(2 * time.Second)
	fmt.Println(result)

	community := make(map[string]*model.User)
	community[u.Name] = &u

	for name, user := range community {
		fmt.Printf("User in community: %s, Age: %d\n", name, user.Age)
	}

}
