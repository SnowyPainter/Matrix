package matrix

import (
	"errors"
)

//ErrMatrixesNotCoincided error
var ErrMatrixesNotCoincided error = errors.New("Two Matrixes are Not Coincided")

//ErrMatrixesNotBeMultiplicable error
var ErrMatrixesNotBeMultiplicable error = errors.New("Two Matrixes Row and Col not same")

//ElementAddHandler handles custom + or -
type ElementAddHandler func(float64, float64) float64

func checkCoincide(m1 *Matrix, m2 *Matrix) bool {
	if m1.Row != m2.Row || m1.Column != m2.Column {
		return false
	}
	return true
}
func checkMultiplicable(m1 *Matrix, m2 *Matrix) bool {
	if m1.Row != m2.Column {
		return false
	}

	return true
}

//Add sum each matrix
func (m *Matrix) add(matrix *Matrix, handler ElementAddHandler) *Matrix {
	result := NewMatrix(m.Row, m.Column)
	for i := 0; i < m.Column; i++ {
		rowSum := make([]float64, m.Row)
		r1 := m.GetRow(i)
		r2 := matrix.GetRow(i)
		for j := 0; j < m.Row; j++ {
			rowSum[j] = handler(r1[j], r2[j])
		}
		result.SetRow(i, rowSum)
	}
	return result
}

//Plus Add two matrixes '+'
func (m *Matrix) Plus(mat *Matrix) (*Matrix, error) {
	if !checkCoincide(m, mat) {
		return nil, ErrMatrixesNotCoincided
	}

	return m.add(mat, func(a float64, b float64) float64 {
		return a + b
	}), nil
}

//Minus Add two matrixes '-'
func (m *Matrix) Minus(mat *Matrix) (*Matrix, error) {
	if !checkCoincide(m, mat) {
		return nil, ErrMatrixesNotCoincided
	}

	return m.add(mat, func(a float64, b float64) float64 {
		return a - b
	}), nil
}

//ScalarMultiply multiply constant value
func (m *Matrix) ScalarMultiply(c float64) *Matrix {
	result := NewMatrix(m.Row, m.Column)
	for i := 0; i < m.Column; i++ {
		rowSum := make([]float64, m.Row)
		for j, v := range m.GetRow(i) {
			rowSum[j] = v * c
		}
		result.SetRow(i, rowSum)
	}
	return result
}

//Multiply Multiply each matrixes
func (m *Matrix) Multiply(mat *Matrix) (*Matrix, error) {
	if !checkMultiplicable(m, mat) {
		return nil, ErrMatrixesNotBeMultiplicable
	}

	result := NewMatrix(mat.Row, m.Column)

	for i := 0; i < m.Column; i++ {
		tmp := make([]float64, m.Row)
		row := m.GetRow(i)
		for j := 0; j < mat.Row; j++ {
			for k, n := range mat.GetColumn(j) {
				tmp[j] += row[k] * n
			}
		}
		result.SetRow(i, tmp)
	}

	return result, nil
}

//Transpose return transpose matrix
func (m *Matrix) Transpose() (matrix *Matrix) {
	matrix = NewMatrix(m.Column, m.Row)

	for i := 0; i < m.Row; i++ { //transposed matrix's col == original.row
		matrix.SetRow(i, m.GetColumn(i))
	}

	return matrix
}
