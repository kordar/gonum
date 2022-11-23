package matrix

import (
	"github.com/kordar/gonum/src/common"
)

// Plus 矩阵相加
func Plus[K common.Key, V common.Value](a *Matrix[K, V], b *Matrix[K, V]) *Matrix[K, V] {
	target := a.Clone()
	target.Plus(b)
	return target
}

// Minus 矩阵相减
func Minus[K common.Key, V common.Value](a *Matrix[K, V], b *Matrix[K, V]) *Matrix[K, V] {
	target := a.Clone()
	target.Minus(b)
	return target
}

// Multiplier 矩阵数乘
func Multiplier[K common.Key, V common.Value](matrix *Matrix[K, V], v float64) *Matrix[K, V] {
	target := matrix.Clone()
	target.Multiplier(v)
	return target
}

// Multiply 矩阵相乘
func Multiply[K common.Key, V common.Value](matrix *Matrix[K, V]) *Matrix[K, V] {
	target := matrix.Clone()
	target.Multiply(matrix)
	return target
}

// Power 幂运算
func Power[K common.Key, V common.Value](matrix *Matrix[K, V], n int) *Matrix[K, V] {
	target := matrix.Clone()
	target.Power(n)
	return target
}

// Transpose 矩阵转置
func Transpose[K common.Key, V common.Value](matrix *Matrix[K, V]) *Matrix[K, V] {
	target := matrix.Clone()
	target.Transpose()
	return target
}
