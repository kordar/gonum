package vector

import (
	"github.com/kordar/gonum/src/common"
	"log"
	"testing"
)

func TestVector_Print(t *testing.T) {
	vectorV := NewVectorV[int, int](5, map[int]int{
		1: 3, 3: 4, 5: 23,
	})
	log.Println(vectorV.Data())
	vectorV.Print()
	vectorV.ToMatrix().Print()

	// ---------------------------
	vectorH := NewVectorH[int, int](5, map[int]int{
		1: 3, 3: 4, 5: 23,
	})
	log.Println(vectorH.Data())
	vectorH.Print()
	vectorH.ToMatrix().Print()
}

func TestBaseVector_Plus(t *testing.T) {
	vectorV := NewVectorV[int, int](5, map[int]int{
		1: 3, 3: 4, 5: 23,
	})
	vectorV2 := NewVectorV[int, int](5, map[int]int{
		1: -1, 3: 4, 5: 23, 2: 8,
	})
	vectorV.Print()
	vectorV2.Print()
	vectorV.Plus(vectorV2)
	vectorV.Print()
	vectorV.Multiplier(2)
	vectorV.Print()
}

func TestNewGroup(t *testing.T) {

	// ==================================
	group := NewGroup[int, int](common.VERTICAL, 5, []*Vector[int, int]{
		NewVectorV(5, map[int]int{1: 3, 2: 4}),
		NewVectorV(5, map[int]int{1: 3, 2: 4, 5: 9}),
		NewVectorV(5, map[int]int{1: 3, 2: 4, 3: 7}),
	})
	group.ToMatrix().Print()

	// ==================================
	group2 := NewGroup[int, int](common.HORIZONTAL, 5, []*Vector[int, int]{
		NewVectorH(5, map[int]int{1: 3, 2: 4}),
		NewVectorH(5, map[int]int{1: 3, 2: 4, 5: 9}),
		NewVectorH(5, map[int]int{1: 3, 2: 4, 3: 7}),
	})
	group2.Add(NewVectorH(5, map[int]int{1: 3, 2: 3, 4: 4}))
	group2.ToMatrix().Print()
}

func TestGroup_Relevant(t *testing.T) {
	group := NewGroup[int, int](common.HORIZONTAL, 3, []*Vector[int, int]{
		NewVectorV(3, map[int]int{1: 1, 2: 2, 3: -1}),
		NewVectorV(3, map[int]int{1: 2, 2: -3, 3: 1}),
		NewVectorV(3, map[int]int{1: 4, 2: 1, 3: -1}),
	})
	log.Println("线性相关：", group.Relevant())
}
