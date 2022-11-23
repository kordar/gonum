package det

import (
	"gonum/src/common"
)

// IsZero 是否空行列式
func IsZero[K common.Key, V common.Value](determinant *BaseDeterminant[K, V]) bool {
	return len(determinant.Data()) == 0
}

// IsSquare 是否方阵，行列式必须为方阵
func IsSquare[K common.Key, V common.Value](determinant *BaseDeterminant[K, V]) bool {
	return determinant.Row() == determinant.Col()
}

// IsUpperTriangleSquare 是否上三角形
func IsUpperTriangleSquare[K common.Key, V common.Value](determinant *BaseDeterminant[K, V]) bool {
	if IsZero(determinant) || !IsSquare(determinant) {
		return false
	}
	hasUpper := false
	for k, v := range determinant.Data() {
		if v == 0 {
			continue
		}
		i, j := determinant.GetSerialByNumber(k)
		if i > j {
			return false
		} else if i < j {
			hasUpper = true
		}
	}
	return hasUpper
}

// IsLowerTriangleSquare 是否下三角形
func IsLowerTriangleSquare[K common.Key, V common.Value](determinant *BaseDeterminant[K, V]) bool {
	if IsZero(determinant) || !IsSquare(determinant) {
		return false
	}
	hasLower := false
	for k, v := range determinant.Data() {
		i, j := determinant.GetSerialByNumber(k)
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

// IsDiag 是否对角形
func IsDiag[K common.Key, V common.Value](determinant *BaseDeterminant[K, V]) bool {
	if IsZero(determinant) || !IsSquare(determinant) {
		return false
	}
	for k, v := range determinant.Data() {
		i, j := determinant.GetSerialByNumber(k)
		if i != j && v != 0 {
			return false
		}
	}
	return true
}

// EqualShape 判断行列式是否相等（形状）
func EqualShape[K common.Key, V common.Value](a *BaseDeterminant[K, V], b *BaseDeterminant[K, V]) bool {
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

// Equal 判断行列式是否相等（值相等）
func Equal[K common.Key, V common.Value](a *BaseDeterminant[K, V], b *BaseDeterminant[K, V]) bool {
	return a.V() == b.V()
}

// Remainder 余子式
func Remainder[K common.Key, V common.Value](determinant *BaseDeterminant[K, V], i K, j K, symbol V) *BaseDeterminant[K, V] {
	common.OutOfBounds(i, j, determinant.Row(), determinant.Col())
	if determinant.Row() == 1 || determinant.Col() == 1 {
		return nil
	}
	target := &BaseDeterminant[K, V]{symbol, determinant.Row() - 1, determinant.Col() - 1, make(map[K]V)}
	for k, v := range determinant.Data() {
		i2, j2 := determinant.GetSerialByNumber(k)
		if i2 == i || j2 == j {
			continue
		}
		if j2 > j {
			j2 = j2 - 1
		}
		if i2 > i {
			i2 = i2 - 1
		}
		target.Save(i2, j2, v)
	}
	return target
}

// AlgebraicRemainder 代数余子式
func AlgebraicRemainder[K common.Key, V common.Value](determinant *BaseDeterminant[K, V], i K, j K) *BaseDeterminant[K, V] {
	k := 1
	if (i+j)%2 != 0 {
		k = -1
	}
	return Remainder(determinant, i, j, V(k))
}

// BaseTranspose 将矩阵转置为一个新的矩阵
func BaseTranspose[K common.Key, V common.Value](determinant *BaseDeterminant[K, V]) *BaseDeterminant[K, V] {
	target := &BaseDeterminant[K, V]{determinant.Symbol(), determinant.Col(), determinant.Row(), make(map[K]V)}
	if determinant.IsZero() {
		return target
	}
	for k, v := range determinant.Data() {
		i, j := common.GetSerialByKey(k, determinant.Col())
		// fmt.Printf("k = %v, i = %v, j = %v, col = %v, mod= %v\n", k, i, j, matrix2.Col(), k/matrix2.Col())
		target.Save(j, i, v)
	}
	return target
}

func Calculate[K common.Key, V common.Value](determinant *BaseDeterminant[K, V]) V {
	res := V(0)
	for j := K(1); j <= determinant.Col(); j++ {
		v := determinant.Get(1, j)
		if v != 0 {
			// 获取代数余子式
			target := AlgebraicRemainder(determinant, 1, j)
			if target == nil {
				return determinant.Get(1, 1)
			}

			res += target.Symbol() * v * Calculate(target)

			//log.Println(target, p, j, res)
		}
	}

	return res
}
