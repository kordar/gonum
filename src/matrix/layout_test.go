package matrix

import (
	"log"
	"testing"
)

func TestNewLayout(t *testing.T) {
	matrix := NewMatrix[int, int](3, 4, map[int][]int{
		1: {1, 2, 3, 4},
		2: {21, 22, 23, 24},
		3: {31, 32, 33, 34},
	})
	PrintAll(matrix)
	layout := NewLayout[int, int](matrix, []int{}, []int{1})
	log.Println(layout)
	sections := layout.Sections
	for _, b := range sections {
		PrintBlockMatrix(b)
	}
	toMatrix := layout.ToMatrix()
	PrintAll(toMatrix)
}
