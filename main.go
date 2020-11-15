package main

import (
	"fmt"

	"github.com/snowypainter/Matrix/matrix"
)

func main() {
	/*
		WARN : variables are not a car!
	*/

	m := matrix.NewMatrix(3, 3)
	m.Reset([][]float64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	m2 := matrix.NewMatrix(3, 3)
	m2.Reset([][]float64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})

	m3, err := m.Plus(m2)
	fmt.Println("m + m2 = ", m3)

	m4, err := m.Minus(m2)
	fmt.Println("m - m2 = ", m4)

	m5, err := m.Multiply(m2)
	fmt.Println("m * m2 = ", m5)

	m6 := m.ScalarMultiply(5)
	fmt.Println("|m| * |5| = ", m6)

	fmt.Println("m.Transpose() = ", m.Transpose())

	if err != nil {
		return
	}

	return
}
