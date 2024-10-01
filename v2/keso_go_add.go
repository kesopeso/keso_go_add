// Package keso_go_add is a simple go package with a single function.
// Its purpose is to learn go module management system.
package keso_go_add

import (
	"golang.org/x/exp/constraints"
)

// Number is a common interface for float and integer number types
type Number interface {
	constraints.Float | constraints.Integer
}

// Add sums up two numers.
//
// It takes two numbers as input parameters.
// You can learn more about addition mathematical function [here](https://www.mathsisfun.com/numbers/addition.html).
func Add[T Number](a, b T) T {
	return a + b
}
