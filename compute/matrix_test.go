package compute

import (
	"testing"
)

func TestEquals(t *testing.T) {
	mat1 := new(Matrix)
	mat1.AddRow([]float64{1, 2, 3})
	mat1.AddRow([]float64{4, 5, 6})
	mat1.AddRow([]float64{7, 8, 9})
	mat2 := new(Matrix)
	mat2.AddRow([]float64{1, 2, 3})
	mat2.AddRow([]float64{4, 5, 6})
	mat2.AddRow([]float64{7, 8, 9})

	if !mat1.Equals(mat2) || !mat2.Equals(mat1) {
		t.Errorf("Both matrices are not equals, mat1=%+v, mat2=%+v", mat1, mat2)
	}
}

func TestTranspose(t *testing.T) {
	mat := new(Matrix)
	mat.AddRow([]float64{1.0, 2.0})
	mat.AddRow([]float64{3.0, 4.0})
	mat.AddRow([]float64{5.0, 6.0})

	mat = mat.Transpose()

	matRes := new(Matrix)
	matRes.AddRow([]float64{1.0, 3.0, 5.0})
	matRes.AddRow([]float64{2.0, 4.0, 6.0})

	if len(mat.Data) != 2 {
		t.Fatalf("Incorrect number of rows, mat=%+v, matRes=%+v", mat, matRes)
	}

	if len(mat.Data[0]) != 3 || len(mat.Data[1]) != 3 {
		t.Fatalf("Incorrect number of cols, mat=%+v, matRes=%+v", mat, matRes)
	}

	for i := 0; i < len(mat.Data); i++ {
		for j := 0; j < len(mat.Data[i]); j++ {
			if mat.Data[i][j] != matRes.Data[i][j] {
				t.Fatalf("Error in transpose matrix content, want: %+v, got: %+v", matRes, mat)
			}
		}
	}
}

func TestGetRow(t *testing.T) {
	mat1 := new(Matrix)
	mat1.AddRow([]float64{1, 2, 3})
	mat1.AddRow([]float64{5, 6, 7})

	if !mat1.GetRow(0).Equals(NewMatrix([]float64{1, 2, 3})) ||
		!mat1.GetRow(1).Equals(NewMatrix([]float64{5, 6, 7})) {
		t.Errorf("Incorrect GetRow function")
	}

}

func TestGetCol(t *testing.T) {
	mat1 := new(Matrix)
	mat1.AddRow([]float64{1, 2, 3})
	mat1.AddRow([]float64{5, 6, 7})

	if !mat1.GetCol(0).Equals(NewMatrix([]float64{1, 5})) ||
		!mat1.GetCol(1).Equals(NewMatrix([]float64{2, 6})) ||
		!mat1.GetCol(2).Equals(NewMatrix([]float64{3, 7})) {
		t.Errorf("Incorrect GetCol function")
	}
}

func TestProduct(t *testing.T) {
	// First test
	mat1 := new(Matrix)
	mat1.AddRow([]float64{11, 3})
	mat1.AddRow([]float64{7, 11})

	mat2 := new(Matrix)
	mat2.AddRow([]float64{8, 0, 1})
	mat2.AddRow([]float64{0, 3, 5})

	matExpected := new(Matrix)
	matExpected.AddRow([]float64{88, 9, 26})
	matExpected.AddRow([]float64{56, 33, 62})

	matProduct := mat1.Product(mat2)
	if !matProduct.Equals(matExpected) {
		t.Errorf("Matrix product is incorrect, want: %+v, got: %+v", matExpected, matProduct)
	}

	// Second test
	mat1 = new(Matrix)
	mat1.AddRow([]float64{1, 2, 3})
	mat1.AddRow([]float64{4, 5, 6})

	mat2 = new(Matrix)
	mat2.AddRow([]float64{7, 8})
	mat2.AddRow([]float64{9, 10})
	mat2.AddRow([]float64{11, 12})

	matExpected = new(Matrix)
	matExpected.AddRow([]float64{58, 64})
	matExpected.AddRow([]float64{139, 154})

	matProduct = mat1.Product(mat2)
	if !matProduct.Equals(matExpected) {
		t.Errorf("Matrix product is incorrect, want: %+v, got: %+v", matExpected, matProduct)
	}

	// Third test
	mat1 = new(Matrix)
	mat1.AddRow([]float64{4, 8})
	mat1.AddRow([]float64{0, 2})
	mat1.AddRow([]float64{1, 6})

	mat2 = new(Matrix)
	mat2.AddRow([]float64{5, 2})
	mat2.AddRow([]float64{9, 4})

	matExpected = new(Matrix)
	matExpected.AddRow([]float64{92, 40})
	matExpected.AddRow([]float64{18, 8})
	matExpected.AddRow([]float64{59, 26})

	matProduct = mat1.Product(mat2)
	if !matProduct.Equals(matExpected) {
		t.Errorf("Matrix product is incorrect, want: %+v, got: %+v", matExpected, matProduct)
	}

	// Fourth test
	mat1 = new(Matrix)
	mat1.AddRow([]float64{1, 4, 6, 10})
	mat1.AddRow([]float64{2, 7, 5, 3})

	mat2 = new(Matrix)
	mat2.AddRow([]float64{1, 4, 6, 10})
	mat2.AddRow([]float64{2, 7, 5, 3})
	mat2.AddRow([]float64{9, 0, 11, 8})

	matProduct = mat1.Product(mat2)
	if matProduct != nil {
		t.Errorf("Matrix product is incorrect, want: nil, got: %+v", matProduct)
	}

	// Fifth test - scalar product as a matrix
	mat1 = new(Matrix)
	mat1.AddRow([]float64{1, 4})
	mat1.AddRow([]float64{2, 7})

	mat2 = new(Matrix)
	mat2.AddRow([]float64{3})

	matProduct = mat1.Product(mat2)
	if matProduct != nil {
		t.Errorf("Matrix product is incorrect, want: nil, got: %+v", matProduct)
	}
}

func TestScalarMultiply(t *testing.T) {
	mat1 := new(Matrix)
	mat1.AddRow([]float64{5, 2, 11})
	mat1.AddRow([]float64{9, 4, 14})

	matExpected := new(Matrix)
	matExpected.AddRow([]float64{15, 6, 33})
	matExpected.AddRow([]float64{27, 12, 42})

	matRes := mat1.ScalarMultiply(3.0)

	if !matRes.Equals(matExpected) {
		t.Errorf("Incorrect ScalarMultiply, want: %+v, got: %+v", matExpected, matRes)
	}
}
