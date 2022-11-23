package matrix

import (
	"github.com/kordar/gonum/src/common"
)

// MultiplierRowPlusToRow 矩阵初等变换，p->系数, r->系数行, tr->目标行
func MultiplierRowPlusToRow[K common.Key, V common.Value](matrix *Matrix[K, V], p float64, r K, tr K) *Matrix[K, V] {
	target := matrix.Clone()
	target.MultiplierRowPlusToRow(p, r, tr)
	return target
}

// ExchangeRow 交换行
func ExchangeRow[K common.Key, V common.Value](matrix *Matrix[K, V], r K, tr K) *Matrix[K, V] {
	target := matrix.Clone()
	target.ExchangeRow(r, tr)
	return target
}

// ExchangeCol 交换列
func ExchangeCol[K common.Key, V common.Value](matrix *Matrix[K, V], c K, tc K) *Matrix[K, V] {
	target := matrix.Clone()
	target.ExchangeCol(c, tc)
	return target
}

// MultiplierRow 行数数乘
func MultiplierRow[K common.Key, V common.Value](matrix *Matrix[K, V], p float64, r K) *Matrix[K, V] {
	target := matrix.Clone()
	target.MultiplierRow(p, r)
	return target
}
