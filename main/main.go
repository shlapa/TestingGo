package main

import "fmt"

type Number interface {
	int | float64
}

func Sum[T Number](a T, b T) (result T) {
	result = a + b
	return result
}

func Minus[T Number](a T, b T) (result T) {
	result = a - b
	return result
}

func Division[T Number](a T, b T) (result T) {
	result = a / b
	return result
}

func Multiplication[T Number](a T, b T) (result T) {
	result = a * b
	return result
}

func main() {
	fmt.Print(Sum(3, 4))
}
