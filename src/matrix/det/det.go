package det

import (
	"gonum/src/common"
)

type Determinant[K common.Key, V common.Value] struct {
	*BaseDeterminant[K, V]
}

func NewDeterminant[K common.Key, V common.Value](n K, data map[K][]V) *Determinant[K, V] {
	base := &BaseDeterminant[K, V]{row: n, col: n, symbol: 1, data: make(map[K]V)}
	determinant := &Determinant[K, V]{base}
	if data != nil {
		determinant.SaveLines(data)
	}
	return determinant
}

func (m *Determinant[K, V]) Clone() *Determinant[K, V] {
	base := m.BaseDeterminant.Clone()
	return &Determinant[K, V]{base}
}

// Load 加载
func (m *Determinant[K, V]) Load(determinant *Determinant[K, V]) {
	m.BaseDeterminant.Load(determinant.BaseDeterminant)
}

// Remainder 余子式
func (m *Determinant[K, V]) Remainder(i K, j K) *Determinant[K, V] {
	return &Determinant[K, V]{m.BaseDeterminant.Remainder(i, j)}
}

// AlgebraicRemainder 代数余子式
func (m *Determinant[K, V]) AlgebraicRemainder(i K, j K) *Determinant[K, V] {
	return &Determinant[K, V]{m.BaseDeterminant.AlgebraicRemainder(i, j)}
}

func (m *Determinant[K, V]) Print() {
	PrintAll(m)
}
