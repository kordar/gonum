package matrix

import (
	"fmt"
	"github.com/kordar/gonum/src/common"
	"github.com/kordar/gonum/src/matrix/det"
	"log"
)

// Save 通过行列保存数据
func Save[K common.Key, V common.Value](matrix *Matrix[K, V], i K, j K, v V) *Matrix[K, V] {
	target := matrix.Clone()
	target.Save(i, j, v)
	return target
}

// SaveLine 保存一行
func SaveLine[K common.Key, V common.Value](matrix *Matrix[K, V], i K, v ...V) *Matrix[K, V] {
	target := matrix.Clone()
	target.SaveLine(i, v...)
	return target
}

// SaveLines 保存多行
func SaveLines[K common.Key, V common.Value](matrix *Matrix[K, V], data map[K][]V) *Matrix[K, V] {
	target := matrix.Clone()
	target.SaveLines(data)
	return target
}

// Get 通过行列坐标获取值
func Get[K common.Key, V common.Value](matrix *Matrix[K, V], i K, j K) V {
	return matrix.Get(i, j)
}

func PrintAll[K common.Key, V common.Value](matrix *Matrix[K, V]) {
	fmt.Println("***************************")
	fmt.Println(fmt.Sprintf("** matrix, row:%v, col:%v", matrix.Row(), matrix.Col()))
	fmt.Println("***************************")
	Print(matrix, 1, 1, matrix.Row(), matrix.Col())
}

func Print[K common.Key, V common.Value](matrix *Matrix[K, V], sRow K, sCol K, eRow K, eCol K) {
	common.OutOfBounds(sRow, sCol, matrix.Row(), matrix.Col())
	common.OutOfBounds(eRow, eCol, matrix.Row(), matrix.Col())
	for r := sRow; r <= eRow; r++ {
		for c := sCol; c <= eCol; c++ {
			fmt.Printf("%v\t", matrix.Get(r, c))
		}
		fmt.Print("\n")
	}
}

// AdjointMatrix 伴随矩阵
func AdjointMatrix[K common.Key, V common.Value](matrix *Matrix[K, V]) *Matrix[K, V] {
	if !matrix.IsSquare() {
		log.Panicln("伴随矩阵必须为方阵")
	}

	if matrix.Row() == 1 || matrix.Col() == 1 {
		return nil
	}
	target := NewMatrix[K, V](matrix.Row(), matrix.Col(), nil)
	determinant := matrix.ToDeterminant()
	for i := K(1); i <= matrix.Row(); i++ {
		for j := K(1); j <= matrix.Col(); j++ {
			remainder := determinant.AlgebraicRemainder(i, j)
			target.Save(j, i, remainder.V())
		}
	}
	return target
}

// SubMatrices 获取子矩阵列表
func SubMatrices[K common.Key, V common.Value](matrix *Matrix[K, V], r K, c K) []*Matrix[K, V] {
	rEles := common.CombinationList(matrix.Row(), r)
	cEles := common.CombinationList(matrix.Col(), c)
	data := make([]*Matrix[K, V], 0)
	for _, rEle := range rEles {
		for _, cEle := range cEles {
			data = append(data, SubMatrixBySerial(matrix, rEle, cEle))
		}
	}
	return data
}

// SubMatrixBySerial 通过序号列表获取新矩阵
func SubMatrixBySerial[K common.Key, V common.Value](matrix *Matrix[K, V], ele []K, ele2 []K) *Matrix[K, V] {
	newMatrix := NewMatrix[K, V](K(len(ele)), K(len(ele2)), nil)
	for ki, i := range ele {
		for kj, j := range ele2 {
			v := matrix.Get(i, j)
			newMatrix.Save(K(ki+1), K(kj+1), v)
		}
	}
	return newMatrix
}

// Rank 秩
func Rank[K common.Key, V common.Value](matrix *Matrix[K, V], k K) K {
	if matrix.Row()-k <= 1 || matrix.Col()-k <= 1 {
		return 0
	}
	min := common.Min(matrix.Row()-k, matrix.Col()-k)
	SubMatrices := SubMatrices(matrix, min, min)
	for _, m := range SubMatrices {
		determinant := m.ToDeterminant()
		if determinant.V() != 0 {
			return min
		}
	}
	return Rank(matrix, k+1)
}

// DeterminantToMatrix 行列式转矩阵
func DeterminantToMatrix[K common.Key, V common.Value](determinant *det.Determinant[K, V]) *Matrix[K, V] {
	matrix := NewMatrix[K, V](determinant.Row(), determinant.Col(), nil)
	matrix.SetData(determinant.Data())
	return matrix
}

// ToDeterminant 矩阵转行列式
func ToDeterminant[K common.Key, V common.Value](matrix *Matrix[K, V]) *det.Determinant[K, V] {
	if !matrix.IsSquare() {
		log.Panicln("行列式必须为方阵")
	}
	determinant := det.NewDeterminant[K, V](matrix.Row(), nil)
	determinant.SetData(matrix.Data())
	return determinant
}

// InverseMatrix 逆矩阵
func InverseMatrix[K common.Key, V common.Value](matrix *Matrix[K, V]) *Matrix[K, V] {
	if !matrix.NonSingular() {
		return nil
	}
	determinant := matrix.ToDeterminant()
	adjointMatrix := matrix.AdjointMatrix()
	adjointMatrix.Multiplier(float64(1) / float64(determinant.V()))
	return adjointMatrix
}
