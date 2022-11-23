package matrix

import (
	"github.com/kordar/gonum/src/common"
	"log"
)

type BaseMatrix[K common.Key, V common.Value] struct {
	row  K
	col  K
	data map[K]V
}

func (m *BaseMatrix[K, V]) SetData(data map[K]V) {
	m.data = data
}

func (m *BaseMatrix[K, V]) Data() map[K]V {
	return m.data
}

func (m *BaseMatrix[K, V]) Put(k K, v V) {
	if v == V(0) {
		delete(m.data, k)
	} else {
		m.data[k] = v
	}
}

func (m *BaseMatrix[K, V]) Row() K {
	return m.row
}

func (m *BaseMatrix[K, V]) Col() K {
	return m.col
}

// GetSerialByNumber 通过下标获取行列序号
func (m *BaseMatrix[K, V]) GetSerialByNumber(number K) (i K, j K) {
	return common.GetSerialByKey(number, m.col)
}

// GetNumberBySerial 通过行列序号获取下标
func (m *BaseMatrix[K, V]) GetNumberBySerial(row K, col K) K {
	return common.GetKeyBySerial(row, col, m.col)
}

func (m *BaseMatrix[K, V]) S(i K, j K, v V, ckOutBound bool) {
	if ckOutBound {
		common.OutOfBounds(i, j, m.Row(), m.Col())
	}
	key := m.GetNumberBySerial(i, j)
	if v == V(0) {
		delete(m.data, key)
	} else {
		m.data[key] = v
	}
}

func (m *BaseMatrix[K, V]) Get(i K, j K) V {
	common.OutOfBounds(i, j, m.Row(), m.Col())
	key := m.GetNumberBySerial(i, j)
	return m.data[key]
}

func (m *BaseMatrix[K, V]) Clone() *BaseMatrix[K, V] {
	mm := make(map[K]V)
	for k, v := range m.Data() {
		mm[k] = v
	}
	return &BaseMatrix[K, V]{m.row, m.col, mm}
}

// Load 加载
func (m *BaseMatrix[K, V]) Load(matrix *BaseMatrix[K, V]) {
	m.data = matrix.Data()
	m.col = matrix.Col()
	m.row = matrix.Row()
}

// ----------------------

// IsZero 是否零矩阵
func (m *BaseMatrix[K, V]) IsZero() bool {
	return IsZero[K, V](m)
}

// IsSquare 是否方阵
func (m *BaseMatrix[K, V]) IsSquare() bool {
	return IsSquare[K, V](m)
}

// IsUpperTriangleSquare 是否上三角矩阵
func (m *BaseMatrix[K, V]) IsUpperTriangleSquare() bool {
	return IsUpperTriangleSquare[K, V](m)
}

// IsLowerTriangleSquare 是否下三角矩阵
func (m *BaseMatrix[K, V]) IsLowerTriangleSquare() bool {
	return IsLowerTriangleSquare[K, V](m)
}

// IsDiag 是否对角阵
func (m *BaseMatrix[K, V]) IsDiag() bool {
	return IsDiag[K, V](m)
}

// QuantityMatrix 数量矩阵
func (m *BaseMatrix[K, V]) QuantityMatrix() (V, bool) {
	return QuantityMatrix[K, V](m)
}

// UnitMatrix 单位矩阵
func (m *BaseMatrix[K, V]) UnitMatrix() (V, bool) {
	return UnitMatrix[K, V](m)
}

// ************* 矩阵更新操作 ***************

func (m *BaseMatrix[K, V]) Save(i K, j K, v V) {
	m.S(i, j, v, true)
}

func (m *BaseMatrix[K, V]) SaveLine(i K, v ...V) {
	for col, val := range v {
		m.S(i, K(col)+1, val, true)
	}
}

func (m *BaseMatrix[K, V]) SaveLines(data map[K][]V) {
	for row, vs := range data {
		for col, v := range vs {
			m.S(row, K(col)+1, v, true)
		}
	}
}

func (m *BaseMatrix[K, V]) Equal(matrix *BaseMatrix[K, V]) bool {
	return Equal[K, V](m, matrix)
}

func (m *BaseMatrix[K, V]) Similar(matrix *BaseMatrix[K, V]) bool {
	return Similar(m, matrix)
}

// ************** 矩阵基本运算，矩阵相加、相减、数乘、相乘 ******************

// Plus 矩阵相加
func (m *BaseMatrix[K, V]) Plus(matrix *BaseMatrix[K, V]) {
	if !m.Similar(matrix) {
		log.Panicln("dissimilar matrices cannot be added")
	}
	if matrix.IsZero() {
		return
	}
	targetData := matrix.Data()
	for i, value := range targetData {
		v := m.data[i] + value
		m.Put(i, v)
	}
}

// Minus 矩阵相减
func (m *BaseMatrix[K, V]) Minus(matrix *BaseMatrix[K, V]) {
	if !m.Similar(matrix) {
		log.Panicln("dissimilar matrices cannot be minus")
	}
	if matrix.IsZero() {
		return
	}
	targetData := matrix.Data()
	for i, value := range targetData {
		v := m.data[i] - value
		m.Put(i, v)
	}
}

// Multiplier 矩阵数乘
func (m *BaseMatrix[K, V]) Multiplier(v float64) {
	if V(v) == V(0) {
		m.data = map[K]V{}
		return
	}
	for i, value := range m.data {
		m.data[i] = V(v * float64(value))
	}
}

// Multiply 矩阵相乘
func (m *BaseMatrix[K, V]) Multiply(matrix *BaseMatrix[K, V]) {
	m.Load(BaseMultiply(m, matrix))
}

// Power 幂运算
func (m *BaseMatrix[K, V]) Power(n int) {
	for i := 1; i < n; i++ {
		m.Multiply(m)
	}
}

// Transpose 矩阵转置
func (m *BaseMatrix[K, V]) Transpose() {
	m.Load(BaseTranspose(m))
}
