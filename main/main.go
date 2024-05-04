package main

import (
	"fmt"
	"math/rand"
)

// region Testify
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

func ModDivision(a int, b int) (result int) {
	result = a % b
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

// region Mock
type randNumberGenerator interface {
	getRandEffects() int
}

var RandNumberGenerator = standardRand{}

type Person struct {
	name        string
	age         int
	id          int
	randEffects int
}

type MyPersonalStruct struct {
	person  Person
	f64Num  float64
	message string
}

type standardRand struct{}

func (standard standardRand) getRandEffects() (res int) {
	return rand.Intn(10)
}

func CheckEvenUsers(personalStruct MyPersonalStruct) (check bool) {
	return (personalStruct.person.id/2 != 0)
}

// endregion
func main() {
	myperson := MyPersonalStruct{
		person: Person{
			name:        "John Doe",
			age:         30,
			id:          0,
			randEffects: RandNumberGenerator.getRandEffects(),
		},
		f64Num:  35.5,
		message: "Hello World",
	}
	fmt.Print(CheckEvenUsers(myperson), " RandEffect:", myperson.person.randEffects, "\n")

	myperson = MyPersonalStruct{
		person: Person{
			name:        "John Doe",
			age:         30,
			id:          2,
			randEffects: RandNumberGenerator.getRandEffects(),
		},
		f64Num:  35.5,
		message: "Hello World",
	}
	fmt.Print(CheckEvenUsers(myperson), " RandEffect:", myperson.person.randEffects, "\n")

	mypersonSec := &myperson
	myperson.person.id = 0
	fmt.Print(CheckEvenUsers(*mypersonSec), "\n")
	fmt.Print(CheckEvenUsers(myperson), "\n")
}
