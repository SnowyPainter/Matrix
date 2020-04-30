package main

import (
	"fmt"

	"github.com/snowypainter/Matrix/matrix"
)

func main() {
	m := matrix.NewMatrix(2, 3)
	m.SetRow(0, []float64{1, 2})
	m.SetRow(1, []float64{3, 4})
	m.SetRow(2, []float64{5, 6})

	//m.Reset([][]float64{{1, 2}, {3, 4}, {5, 6}})

	m2 := matrix.NewMatrix(2, 2)
	m2.SetRow(0, []float64{1, 3})
	m2.SetRow(1, []float64{2, 4})

	fmt.Println(m)
	fmt.Println(m.Transpose())

	return
}
