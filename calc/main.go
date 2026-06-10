package main

import (
	"errors"
	"math"
)

var (
	ErrDivisionByZero = errors.New("division by zero")
	ErrOverflow       = errors.New("integer overflow")
)

// Créer un nouveau type appelé myInt de type int32
type myInt int32

// Créer les méthodes suivantes :
// Divide : retourner la division avec un nombre n de type int passé en paramètre
func (m myInt) Divide(n int) (myInt, error) {
	if n == 0 {
		return 0, ErrDivisionByZero
	}
	return m / myInt(n), nil
}

// Add : retourner la valeur ajouté par n de type int passé en paramètre
func (m myInt) Add(n int) (myInt, error) {
	res := int64(m) + int64(n)
	if res > math.MaxInt32 || res < math.MinInt32 {
		return 0, ErrOverflow
	}
	return myInt(res), nil
}

// Sub : retourner la valeur soustraite avec n toujours passé en paramètre
func (m myInt) Sub(n int) (myInt, error) {
	res := int64(m) - int64(n)
	if res > math.MaxInt32 || res < math.MinInt32 {
		return 0, ErrOverflow
	}
	return myInt(res), nil
}

// Multiply : retourner la valeur multiplié des deux paramètres de type int en myInt
func (m myInt) Multiply(n int) (myInt, error) {
	res := int64(m) * int64(n)
	if res > math.MaxInt32 || res < math.MinInt32 {
		return 0, ErrOverflow
	}
	return myInt(res), nil
}

// Créer un nouveau type appelé myInt de type int32
type myIntP int32

// Créer les méthodes suivantes :
// Divide : retourner la division avec un nombre n de type int passé en paramètre
func (m *myIntP) Divide(n int) {
	*m = *m / myIntP(n)
}

// Add : retourner la valeur ajouté par n de type int passé en paramètre
func (m *myIntP) Add(n int) {
	*m = *m + myIntP(n)
}

// Sub : retourner la valeur soustraite avec n toujours passé en paramètre
func (m *myIntP) Sub(n int) {
	*m = *m - myIntP(n)
}

// Multiply : retourner la valeur multiplié des deux paramètres de type int en myInt
func (m *myIntP) Multiply(n int) {
	*m = *m * myIntP(n)
}
