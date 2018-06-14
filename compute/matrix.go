package compute

type Matrix struct {
	Data [][]float64
}

func NewMatrix(row []float64) *Matrix {
	ret := new(Matrix)
	ret.AddRow(row)
	return ret
}

func (m *Matrix) AddRow(row []float64) *Matrix {
	m.Data = append(m.Data, row)
	return m
}

func (m *Matrix) GetNbRows() int {
	return len(m.Data)
}

func (m *Matrix) GetNbCols() int {
	return len(m.Data[0])
}

func (m *Matrix) GetRow(i int) *Matrix {
	ret := new(Matrix)
	ret.AddRow(m.Data[i])
	return ret
}

func (m *Matrix) GetCol(j int) *Matrix {
	ret := new(Matrix)
	col := make([]float64, 0)
	for i := 0; i < m.GetNbRows(); i++ {
		col = append(col, m.Data[i][j])
	}
	ret.AddRow(col)
	return ret
}

func (m *Matrix) Copy() *Matrix {
	ret := Matrix{}
	copy(m.Data, ret.Data)
	return &ret
}

func (m *Matrix) Equals(m2 *Matrix) bool {
	if len(m.Data) != len(m2.Data) {
		return false
	}

	for i := 0; i < len(m.Data); i++ {
		for j := 0; j < len(m.Data[i]); j++ {
			if m.Data[i][j] != m2.Data[i][j] {
				return false
			}
		}
	}

	return true
}

func (m *Matrix) Transpose() *Matrix {
	ret := Matrix{}

	for j := 0; j < len(m.Data[0]); j++ {
		col := make([]float64, 0, len(m.Data[j]))
		for i := 0; i < len(m.Data); i++ {
			col = append(col, m.Data[i][j])
		}
		ret.AddRow(col)
	}

	return &ret
}

func (m1 *Matrix) Product(m2 *Matrix) *Matrix {
	if m1.GetNbCols() != m2.GetNbRows() {
		return nil
	}

	matProd := new(Matrix)

	for i := 0; i < m1.GetNbRows(); i++ {
		curRow := m1.GetRow(i)
		rowProd := make([]float64, 0)

		for j := 0; j < m2.GetNbCols(); j++ {
			curCol := m2.GetCol(j)
			val := mult(curRow.Data[0], curCol.Data[0])
			rowProd = append(rowProd, val)
		}

		matProd.AddRow(rowProd)
	}

	return matProd
}

func (m *Matrix) ScalarMultiply(s float64) *Matrix {
	ret := new(Matrix)
	for i := 0; i < m.GetNbRows(); i++ {
		row := make([]float64, 0)

		for j := 0; j < m.GetNbCols(); j++ {
			row = append(row, s*m.Data[i][j])
		}

		ret.AddRow(row)
	}

	return ret
}

func mult(a1, a2 []float64) float64 {
	ret := 0.0

	for i := 0; i < len(a1); i++ {
		ret += a1[i] * a2[i]
	}

	return ret
}
