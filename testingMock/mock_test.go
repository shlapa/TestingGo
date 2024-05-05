package testingMock

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
)

type randNumberGenerator interface {
	GetRandEffects() int
}

type mockRand struct{ mock.Mock }

func newMockRand() *mockRand { return &mockRand{} }

var mockr *mockRand

type MyPersonalStruct struct {
	person  string
	f64Num  float64
	message string
}

type Person struct {
	name        string
	age         int
	id          int
	randEffects int
}

func (m *mockRand) GetRandEffects() int {
	args := m.Called()
	return args.Int(0)
}

func TestFirstRandMock(t *testing.T) {
	mockr = newMockRand()
	mockr.On("GetRandEffects").Return(6)

	person := Person{
		name:        "John Doe",
		age:         30,
		id:          2,
		randEffects: mockr.GetRandEffects(), // Здесь мы используем GetRandEffects напрямую из мок-объекта
	}

	mypersonMock := MyPersonalStruct{
		person:  "John Doe", // Здесь используем строку, как требуется
		f64Num:  35.5,
		message: "Hello World",
	}

	// Проверяем, что метод был вызван один раз
	mockr.AssertExpectations(t)

	// Проверяем, что randEffects равен 6
	if person.randEffects != 6 {
		t.Errorf("Expected 6, got %d", person.randEffects)
	}
	
	if(CheckEvenUsers.)

	fmt.Println(mypersonMock)
}
