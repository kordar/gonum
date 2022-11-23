package matrix

import (
	"gonum/src/common"
	"log"
)

func E[K common.Key, V common.Value](n K) *BaseMatrix[K, V] {
	base := &BaseMatrix[K, V]{n, n, make(map[K]V)}
	for i := K(1); i <= n; i++ {
		base.Save(i, i, 1)
	}
	return base
}

// IsZero 是否零矩阵
func IsZero[K common.Key, V common.Value](matrix *BaseMatrix[K, V]) bool {
	return len(matrix.Data()) == 0
}

// IsSquare 是否方阵
func IsSquare[K common.Key, V common.Value](matrix *BaseMatrix[K, V]) bool {
	return matrix.Row() == matrix.Col()
}

// IsUpperTriangleSquare 是否上三角矩阵
func IsUpperTriangleSquare[K common.Key, V common.Value](matrix *BaseMatrix[K, V]) bool {
	if IsZero(matrix) || !IsSquare(matrix) {
		return false
	}
	hasUpper := false
	for k, v := range matrix.Data() {
		if v == 0 {
			continue
		}
		i, j := matrix.GetSerialByNumber(k)
		if i > j {
			return false
		} else if i < j {
			hasUpper = true
		}
	}
	return hasUpper
}

// IsLowerTriangleSquare 是否下三角矩阵
func IsLowerTriangleSquare[K common.Key, V common.Value](matrix *BaseMatrix[K, V]) bool {
	if IsZero(matrix) || !IsSquare(matrix) {
		return false
	}
	hasLower := false
	for k, v := range matrix.Data() {
		i, j := matrix.GetSerialByNumber(k)
		if v != 0 {
			if i < j {
				return false
			} else if i > j {
				hasLower = true
			}
		}
	}
	return hasLower
}

// IsDiag 是否对角阵
func IsDiag[K common.Key, V common.Value](matrix *BaseMatrix[K, V]) bool {
	if IsZero(matrix) || !IsSquare(matrix) {
		return false
	}
	for k, v := range matrix.Data() {
		i, j := matrix.GetSerialByNumber(k)
		if i != j && v != 0 {
			return false
		}
	}
	return true
}

// QuantityMatrix 数量矩阵
func QuantityMatrix[K common.Key, V common.Value](matrix *BaseMatrix[K, V]) (V, bool) {
	if !IsDiag(matrix) {
		log.Println("数量矩阵必须为对角阵")
		return 0, false
	}
	tmp := V(0)
	for _, v := range matrix.Data() {
		if tmp == 0 {
			tmp = v
		} else if tmp != v {
			log.Println("非数量矩阵")
			return 0, false
		}
	}
	return tmp, true
}

// UnitMatrix 单位矩阵
func UnitMatrix[K common.Key, V common.Value](matrix *BaseMatrix[K, V]) (V, bool) {
	if !IsDiag(matrix) {
		log.Println("单位矩阵必须为对角阵")
		return 0, false
	}
	for _, v := range matrix.Data() {
		if 1 != v {
			log.Println("非单位矩阵")
			return 0, false
		}
	}
	return 1, true
}

// Similar 判断两个矩阵是否相似
func Similar[K common.Key, V common.Value](a *BaseMatrix[K, V], b *BaseMatrix[K, V]) bool {
	return a.Row() == b.Row() || a.Col() == b.Col()
}

// Equal 判断矩阵是否相等
func Equal[K common.Key, V common.Value](a *BaseMatrix[K, V], b *BaseMatrix[K, V]) bool {
	if !Similar(a, b) {
		return false
	}
	if len(a.Data()) != len(b.Data()) {
		return false
	}
	t := a.Data()
	for k, v := range b.Data() {
		if v != t[k] {
			return false
		}
	}
	return true
}

// BaseTranspose 将矩阵转置为一个新的矩阵
func BaseTranspose[K common.Key, V common.Value](matrix *BaseMatrix[K, V]) *BaseMatrix[K, V] {
	target := &BaseMatrix[K, V]{matrix.Col(), matrix.Row(), make(map[K]V)}
	if matrix.IsZero() {
		return target
	}
	//log.Println(matrix2)
	for k, v := range matrix.Data() {
		i, j := common.GetSerialByKey(k, matrix.Col())
		// fmt.Printf("k = %v, i = %v, j = %v, col = %v, mod= %v\n", k, i, j, matrix2.Col(), k/matrix2.Col())
		target.Save(j, i, v)
	}
	return target
}

// BaseMultiply 两个矩阵相乘返回新矩阵
func BaseMultiply[K common.Key, V common.Value](a *BaseMatrix[K, V], b *BaseMatrix[K, V]) *BaseMatrix[K, V] {
	if a.Col() != b.Row() {
		log.Panicln("只有当第一个矩阵（左边的矩阵）的列数与第二个矩阵（右边的矩阵）的行数相等时，两个矩阵才能相乘．")
	}
	target := &BaseMatrix[K, V]{a.Row(), b.Col(), make(map[K]V)}
	if b.IsZero() {
		return target
	}
	mm := make(map[K]V)
	for i := K(1); i <= a.Row(); i++ {
		for j := K(1); j <= b.Col(); j++ {
			val := V(0)
			for k := K(1); k <= b.Row(); k++ {
				val += a.Get(i, k) * b.Get(k, j)
				//fmt.Printf("%v*%v+", a.Get(i, k), b.Get(k, j))
				//
			}
			if val != 0 {
				key := target.GetNumberBySerial(i, j)
				mm[key] = val
			}
			//fmt.Print("\t")
		}
		//fmt.Print("\n")
	}
	target.data = mm
	return target
}
