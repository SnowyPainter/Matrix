package matrix

import "fmt"

//Matrix contains Matrix raw
type Matrix struct {
	raw    [][]float64
	Row    int
	Column int
}

//NewMatrix creats new Matrix structure
func NewMatrix(row int, col int) *Matrix {
	r := make([][]float64, col)
	for i := range r {
		r[i] = make([]float64, row)
	}

	return &Matrix{
		raw:    r,
		Row:    row,
		Column: col,
	}
}

//Reset set all data
func (m *Matrix) Reset(data [][]float64) {
	m.raw = data
	m.Row = len(data[0])
	m.Column = len(data)
}

//SetColumn set raw
func (m *Matrix) SetColumn(index int, data []float64) {
	for i, r := range m.raw {
		if i < len(data) {
			r[index] = data[i]
			continue
		}
		r[index] = 0
	}
}

//SetRow set raw
func (m *Matrix) SetRow(index int, data []float64) {
	if len(data) == m.Row {
		m.raw[index] = data
		return
	}

	for i := 0; i < m.Row; i++ {
		if i >= len(data) {
			m.raw[index][i] = 0
			continue
		}
		m.raw[index][i] = data[i]
	}
}

//GetColumn returns column of index
func (m *Matrix) GetColumn(index int) (array []float64) {
	array = make([]float64, m.Column)
	for i, r := range m.raw {
		array[i] = r[index]
	}
	return
}

//GetRow returns row of index
func (m *Matrix) GetRow(index int) []float64 {
	return m.raw[index]
}

func (m *Matrix) String() string {
	arr := ""
	for i := 0; i < m.Column; i++ {
		arr += fmt.Sprintf("  %v\n", m.GetRow(i))
	}
	return fmt.Sprintf("[\n%v]", arr)
}
