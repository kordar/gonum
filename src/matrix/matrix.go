package matrix

import (
	"gonum/src/common"
	"gonum/src/matrix/det"
	"log"
)

type Matrix[K common.Key, V common.Value] struct {
	*BaseMatrix[K, V]
	*TransformMatrix[K, V]
}

func NewMatrix[K common.Key, V common.Value](row K, col K, data map[K][]V) *Matrix[K, V] {
	base := &BaseMatrix[K, V]{row: row, col: col, data: make(map[K]V)}
	matrix := &Matrix[K, V]{base, &TransformMatrix[K, V]{base}}
	if data != nil {
		matrix.SaveLines(data)
	}
	return matrix
}

func NewMatrixE[K common.Key, V common.Value](n K) *Matrix[K, V] {
	base := E[K, V](n)
	return &Matrix[K, V]{base, &TransformMatrix[K, V]{base}}
}

// Plus 矩阵相加
func (m *Matrix[K, V]) Plus(matrix *Matrix[K, V]) {
	m.BaseMatrix.Plus(matrix.BaseMatrix)
}

// Minus 矩阵相减
func (m *Matrix[K, V]) Minus(matrix *Matrix[K, V]) {
	m.BaseMatrix.Minus(matrix.BaseMatrix)
}

// Multiply 矩阵相乘
func (m *Matrix[K, V]) Multiply(matrix *Matrix[K, V]) {
	m.BaseMatrix.Multiply(matrix.BaseMatrix)
}

func (m *Matrix[K, V]) Clone() *Matrix[K, V] {
	base := m.BaseMatrix.Clone()
	return &Matrix[K, V]{base, &TransformMatrix[K, V]{base}}
}

// Load 加载
func (m *Matrix[K, V]) Load(matrix *Matrix[K, V]) {
	m.BaseMatrix.Load(matrix.BaseMatrix)
}

// ToDeterminant 转换为行列式
func (m *Matrix[K, V]) ToDeterminant() *det.Determinant[K, V] {
	return ToDeterminant(m)
}

// AdjointMatrix 伴随矩阵
func (m *Matrix[K, V]) AdjointMatrix() *Matrix[K, V] {
	return AdjointMatrix[K, V](m)
}

// Rank 矩阵的秩
func (m *Matrix[K, V]) Rank() K {
	return Rank(m, 0)
}

// Singular 奇异矩阵
func (m *Matrix[K, V]) Singular() bool {
	if !m.IsSquare() {
		log.Println("singular matrix must be a square")
	}
	return m.ToDeterminant().V() == 0
}

// NonSingular 非奇异矩阵
func (m *Matrix[K, V]) NonSingular() bool {
	if !m.IsSquare() {
		log.Println("non-singular matrix must be a square")
	}
	return m.ToDeterminant().V() > 0
}

// InverseMatrix 逆矩阵
func (m *Matrix[K, V]) InverseMatrix() *Matrix[K, V] {
	return InverseMatrix(m)
}

func (m *Matrix[K, V]) Print() {
	PrintAll(m)
}
