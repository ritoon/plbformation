package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	err := A()
	if err != nil {
		if errors.Is(err, ErrCustom) {
			log.Printf("Custom error occurred: %v", err)
		} else {
			log.Printf("An error occurred: %v", err)
		}

	}
}

func A() error {
	err := B()
	if err != nil {
		return fmt.Errorf("error in A function: %w", err)
	}

	return nil
}

func B() error {
	err := C()
	if err != nil {
		return fmt.Errorf("error in B function: %w", err)
	}
	return nil
}

func C() error {
	return ErrCustom
}

var (
	ErrCustom = errors.New("this is a custom error")
)
