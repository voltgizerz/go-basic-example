package main

import "fmt"

func toString[T any](value T) string {
	return fmt.Sprintf("%v", value)
}

func isEqual[T comparable](value1 T, value2 T) bool {
	return value1 == value2
}
