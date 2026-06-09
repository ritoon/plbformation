package main

// Créer un nouveau type appelé myInt de type int32
type myInt int32

// Créer les méthodes suivantes :
// Divide : retourner la division avec un nombre n de type int passé en paramètre
func (m myInt) Divide(n int) myInt {
	return m / myInt(n)
}

// Add : retourner la valeur ajouté par n de type int passé en paramètre
func (m myInt) Add(n int) myInt {
	return m + myInt(n)
}

// Sub : retourner la valeur soustraite avec n toujours passé en paramètre
func (m myInt) Sub(n int) myInt {
	return m - myInt(n)
}

// Multiply : retourner la valeur multiplié des deux paramètres de type int en myInt
func (m myInt) Multiply(n int) myInt {
	return m * myInt(n)
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
