package det

import (
	"gonum/src/common"
)

type BaseDeterminant[K common.Key, V common.Value] struct {
	symbol V
	row    K
	col    K
	data   map[K]V
}

func (m *BaseDeterminant[K, V]) Symbol() V {
	return m.symbol
}

func (m *BaseDeterminant[K, V]) SetSymbol(symbol V) {
	m.symbol = symbol
}

func (m *BaseDeterminant[K, V]) SetData(data map[K]V) {
	m.data = data
}

func NewBaseDeterminant[K common.Key, V common.Value](n K) *BaseDeterminant[K, V] {
	return &BaseDeterminant[K, V]{row: n, col: n, symbol: V(1), data: make(map[K]V)}
}

func (m *BaseDeterminant[K, V]) Data() map[K]V {
	return m.data
}

func (m *BaseDeterminant[K, V]) Put(k K, v V) {
	if v == V(0) {
		delete(m.data, k)
	} else {
		m.data[k] = v
	}
}

func (m *BaseDeterminant[K, V]) Row() K {
	return m.row
}

func (m *BaseDeterminant[K, V]) Col() K {
	return m.col
}

// GetSerialByNumber 通过下标获取行列序号
func (m *BaseDeterminant[K, V]) GetSerialByNumber(number K) (i K, j K) {
	return common.GetSerialByKey(number, m.col)
}

// GetNumberBySerial 通过行列序号获取下标
func (m *BaseDeterminant[K, V]) GetNumberBySerial(row K, col K) K {
	return common.GetKeyBySerial(row, col, m.col)
}

func (m *BaseDeterminant[K, V]) S(i K, j K, v V, ckOutBound bool) {
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

// V 计算行列式的值
func (m *BaseDeterminant[K, V]) V() V {
	return Calculate(m) * m.Symbol()
}

func (m *BaseDeterminant[K, V]) Get(i K, j K) V {
	common.OutOfBounds(i, j, m.Row(), m.Col())
	key := m.GetNumberBySerial(i, j)
	return m.data[key]
}

func (m *BaseDeterminant[K, V]) Clone() *BaseDeterminant[K, V] {
	mm := make(map[K]V)
	for k, v := range m.Data() {
		mm[k] = v
	}
	return &BaseDeterminant[K, V]{V(1), m.row, m.col, mm}
}

// Load 加载
func (m *BaseDeterminant[K, V]) Load(determinant *BaseDeterminant[K, V]) {
	m.data = determinant.Data()
	m.col = determinant.Col()
	m.row = determinant.Row()
}

// ----------------------

// IsZero 是否空
func (m *BaseDeterminant[K, V]) IsZero() bool {
	return IsZero[K, V](m)
}

// IsSquare 是否方阵
func (m *BaseDeterminant[K, V]) IsSquare() bool {
	return IsSquare[K, V](m)
}

// IsUpperTriangleSquare 是否上三角形
func (m *BaseDeterminant[K, V]) IsUpperTriangleSquare() bool {
	return IsUpperTriangleSquare[K, V](m)
}

// IsLowerTriangleSquare 是否下三角形
func (m *BaseDeterminant[K, V]) IsLowerTriangleSquare() bool {
	return IsLowerTriangleSquare[K, V](m)
}

// IsDiag 是否对角阵
func (m *BaseDeterminant[K, V]) IsDiag() bool {
	return IsDiag[K, V](m)
}

// ************* 行列式更新操作 ***************

func (m *BaseDeterminant[K, V]) Save(i K, j K, v V) {
	m.S(i, j, v, true)
}

func (m *BaseDeterminant[K, V]) SaveLine(i K, v ...V) {
	for col, val := range v {
		m.S(i, K(col)+1, val, true)
	}
}

func (m *BaseDeterminant[K, V]) SaveLines(data map[K][]V) {
	for row, vs := range data {
		for col, v := range vs {
			m.S(row, K(col)+1, v, true)
		}
	}
}

func (m *BaseDeterminant[K, V]) Equal(determinant *BaseDeterminant[K, V]) bool {
	return Equal[K, V](m, determinant)
}

func (m *BaseDeterminant[K, V]) EqualShape(determinant *BaseDeterminant[K, V]) bool {
	return EqualShape(m, determinant)
}

// ************** 矩阵基本运算，行列式相加、相减、数乘、相乘 ******************

// Plus 行列式相加
func (m *BaseDeterminant[K, V]) Plus(determinant *BaseDeterminant[K, V]) V {
	return m.V() + determinant.V()
}

// Minus 行列式相减
func (m *BaseDeterminant[K, V]) Minus(determinant *BaseDeterminant[K, V]) V {
	return m.V() - determinant.V()
}

// Multiplier 行列式数乘
func (m *BaseDeterminant[K, V]) Multiplier(v float64) V {
	return V(v * float64(m.V()))
}

// Multiply 行列式相乘
func (m *BaseDeterminant[K, V]) Multiply(determinant *BaseDeterminant[K, V]) V {
	return m.V() * determinant.V()
}

// Power 幂运算
func (m *BaseDeterminant[K, V]) Power(n int) {
	for i := 1; i < n; i++ {
		m.Multiply(m)
	}
}

// Transpose 行列式转置
func (m *BaseDeterminant[K, V]) Transpose() {
	m.Load(BaseTranspose(m))
}

func (m *BaseDeterminant[K, V]) Remainder(i K, j K) *BaseDeterminant[K, V] {
	return Remainder(m, i, j, V(1))
}

func (m *BaseDeterminant[K, V]) AlgebraicRemainder(i K, j K) *BaseDeterminant[K, V] {
	return AlgebraicRemainder(m, i, j)
}
