package testingMock

import (
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockRand struct{ mock.Mock }

func newMockRund() *mockRand { return &mockRand{} }

var mockr mockRand

func TestFirstRandMock(t *testing.T) {

	var myperson:= MyPersonalStruct{
		person: Person{
			name:        "John Doe",
			age:         30,
			id:          2,
			randEffects: RandNumberGenerator.getRandEffects(),
		},
		f64Num:  35.5,
		message: "Hello World",
	}

	mockr := newMockRund()
	mockr.On("getRandEffects", 10)

	//quo:=
}
