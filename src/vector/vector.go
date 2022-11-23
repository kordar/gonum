package vector

import (
	"github.com/kordar/gonum/src/common"
	"github.com/kordar/gonum/src/matrix"
)

type Vector[K common.Key, V common.Value] struct {
	direction common.VectorType // vertical | horizontal
	*BaseVector[K, V]
}

func (vec *Vector[K, V]) Direction() common.VectorType {
	return vec.direction
}

func NewVectorH[K common.Key, V common.Value](size K, data map[K]V) *Vector[K, V] {
	if data == nil {
		data = make(map[K]V)
	}
	baseVector := NewBaseVector(1, size, data)
	return &Vector[K, V]{common.HORIZONTAL, baseVector}
}

func NewVectorV[K common.Key, V common.Value](size K, data map[K]V) *Vector[K, V] {
	if data == nil {
		data = make(map[K]V)
	}
	baseVector := NewBaseVector(size, 1, data)
	return &Vector[K, V]{common.VERTICAL, baseVector}
}

// Plus 向量相加
func (vec *Vector[K, V]) Plus(matrix *Vector[K, V]) {
	vec.BaseVector.Plus(matrix.BaseVector)
}

// Minus 向量相减
func (vec *Vector[K, V]) Minus(matrix *Vector[K, V]) {
	vec.BaseVector.Minus(matrix.BaseVector)
}

func (vec *Vector[K, V]) Clone() *Vector[K, V] {
	base := vec.BaseVector.Clone()
	return &Vector[K, V]{vec.Direction(), base}
}

func (vec *Vector[K, V]) Print() {
	PrintAll(vec)
}

// ToMatrix 转换向量为矩阵
func (vec *Vector[K, V]) ToMatrix() *matrix.Matrix[K, V] {
	return ToMatrix(vec)
}
