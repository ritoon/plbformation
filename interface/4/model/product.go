package model

import "fmt"

type Product struct {
	ID    int
	Name  string
	Price float64
}

func (p Product) String() string {
	return fmt.Sprintf("Product: %s, Price: %.2f", p.Name, p.Price)
}
