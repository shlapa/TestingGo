package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// With Int
func TestSumInt(t *testing.T) {
	assert.Equal(t, 15, Sum(5, 10))
}

func TestMinusInt(t *testing.T) {
	assert.Equal(t, 5, Minus(10, 5))
}

func TestDivisionInt(t *testing.T) {
	assert.Equal(t, 15, Division(150, 10))
}

func TestMultiplicationInt(t *testing.T) {
	assert.Equal(t, 15, Multiplication(5, 3))
}

// With Double
func TestSumIntDouble(t *testing.T) {
	assert.Equal(t, 10, Sum(5.5, 4.5))
}

func TestMinusIntDouble(t *testing.T) {
	assert.Equal(t, 5.7, Minus(10, 4.3))
}

func TestDivisionIntDouble(t *testing.T) {
	assert.Equal(t, 2.2500000000000004, Division(5.4, 2.4))
}

func TestMultiplicationIntDouble(t *testing.T) {
	assert.Equal(t, 12.96, Multiplication(5.4, 2.4))
}

// Mod DivisionMod Test
func TestDivisionMod(t *testing.T) {
	assert.Equal(t, 0, ModDivision(10, 5))
}
