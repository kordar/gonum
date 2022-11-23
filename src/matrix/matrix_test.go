package matrix

import (
	"gonum/src/matrix/det"
	"log"
	"testing"
)

func TestMatrix_Clone(t *testing.T) {
	matrix := NewMatrix[int, int](2, 2, map[int][]int{
		1: {-1, 1},
		2: {2, -2},
	})
	// matrix2.SaveLine(1, -1, 1)
	// matrix2.SaveLine(2, 2, -2)
	line := SaveLine(matrix, 1, 0, 0)
	matrix.SaveLine(1, 0, 0)
	PrintAll(matrix)
	PrintAll(line)
}

func TestSegmentMatrix(t *testing.T) {

}

func TestToDeterminant(t *testing.T) {
	matrix := NewMatrix[int, int](4, 4, map[int][]int{
		1: {-1, 1, 3, 5},
		2: {2, 2, 5, 6},
		3: {7, 6, 2, 1},
		4: {3, 3, 8, 9},
	})
	determinant := matrix.ToDeterminant()
	log.Println(determinant.V())
	target := determinant.Remainder(2, 1)
	det.PrintAll(target)
	log.Println(target.V())
}

func TestAdjointMatrix(t *testing.T) {
	matrix := NewMatrix[int, int](2, 2, map[int][]int{
		//1: {1, 0},
		//2: {-1, 1},
		//1: {1, -2},
		//2: {3, 1},
		1: {2, -3},
		2: {1, 4},
	})
	adjointMatrix := matrix.AdjointMatrix()
	PrintAll(adjointMatrix)
}

func TestRank(t *testing.T) {
	matrix := NewMatrix[int, int](3, 4, map[int][]int{
		1: {2, -3, 8, 2},
		2: {2, 12, -2, 12},
		3: {1, 3, 1, 4},
	})
	PrintAll(matrix)
	matrices := SubMatrices(matrix, 2, 2)
	for _, m := range matrices {
		PrintAll(m)
		log.Println(m.ToDeterminant().V())
	}

	log.Println("rank = ", matrix.Rank())
}

func TestInverseMatrix(t *testing.T) {
	matrix := NewMatrix[int, float64](2, 2, map[int][]float64{
		1: {2, 5},
		2: {-3, -7},
	})
	inverseMatrix := matrix.InverseMatrix()
	if inverseMatrix != nil {
		PrintAll(inverseMatrix)
		matrix.Multiply(inverseMatrix)
		PrintAll(matrix)
	}
}
