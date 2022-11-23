package vector

import (
	"fmt"
	"github.com/kordar/gonum/src/common"
	"github.com/kordar/gonum/src/matrix"
)

func PrintAll[K common.Key, V common.Value](vec *Vector[K, V]) {
	fmt.Println("***************************")
	fmt.Println(fmt.Sprintf("** vector, row:%v, col:%v", vec.Row(), vec.Col()))
	fmt.Println("***************************")
	Print(vec, 1, 1, vec.Row(), vec.Col())
}

func Print[K common.Key, V common.Value](vec *Vector[K, V], sRow K, sCol K, eRow K, eCol K) {
	common.OutOfBounds(sRow, sCol, vec.Row(), vec.Col())
	common.OutOfBounds(eRow, eCol, vec.Row(), vec.Col())
	for r := sRow; r <= eRow; r++ {
		for c := sCol; c <= eCol; c++ {
			if vec.Direction() == common.VERTICAL {
				fmt.Printf("%v", vec.Get(r))
			}
			if vec.Direction() == common.HORIZONTAL {
				fmt.Printf("%v\t", vec.Get(c))
			}
		}
		fmt.Print("\n")
	}
}

func ToMatrix[K common.Key, V common.Value](vec *Vector[K, V]) *matrix.Matrix[K, V] {
	m := matrix.NewMatrix[K, V](vec.Row(), vec.Col(), nil)
	if vec.Direction() == common.HORIZONTAL {
		for k, v := range vec.data {
			m.Save(1, k, v)
		}
	} else if vec.Direction() == common.VERTICAL {
		for k, v := range vec.data {
			m.Save(k, 1, v)
		}
	}
	return m
}
