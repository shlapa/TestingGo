package main

import "fmt"

//region Testify
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

//endregion

//region Mock
type Person struct {
	name string
	id   int
}

type MyPersonalStruct struct {
	person  Person
	f64Num  float64
	message string
}

func CheckEvenUsers(personalStruct MyPersonalStruct) (check bool) {
	if personalStruct.person.id/2 != 0 {

	}
}

//endregion
func main() {
	fmt.Print(Sum(3, 4))
}
