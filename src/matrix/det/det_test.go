package det

import (
	"log"
	"testing"
)

func TestCalculate2(t *testing.T) {
	determinant := NewDeterminant[int, int](2, nil)
	determinant.SaveLines(map[int][]int{
		1: {3, 1},
		2: {2, -1},
	})
	PrintAll(determinant)
	log.Println(determinant.V())
}

func TestCalculate(t *testing.T) {

	determinant := NewDeterminant[int, int](4, nil)
	determinant.SaveLines(map[int][]int{
		1: {7, 1, 1, 1},
		2: {1, 4, 1, 1},
		3: {1, 1, -2, 1},
		4: {1, 1, 1, -5},
	})
	PrintAll(determinant)
	log.Println(determinant.V())
}
