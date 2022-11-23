package matrix

import (
	"github.com/kordar/gonum/src/common"
	"log"
)

type TransformMatrix[K common.Key, V common.Value] struct {
	*BaseMatrix[K, V]
}

// MultiplierRowPlusToRow 初等行变换
func (tm *TransformMatrix[K, V]) MultiplierRowPlusToRow(p float64, r K, tr K) {
	if r > tm.Row() || tr > tm.Row() {
		log.Fatalln("操作行数异常")
	}
	for c := K(1); c <= tm.Col(); c++ {
		v := tm.Get(r, c)
		v2 := tm.Get(tr, c)
		tm.Save(tr, c, v2+V(p*float64(v)))
	}
}

// ExchangeRow 交换行
func (tm *TransformMatrix[K, V]) ExchangeRow(r K, tr K) {
	if r > tm.Row() || tr > tm.Row() {
		log.Fatalln("操作行数异常")
	}
	for c := K(1); c <= tm.Col(); c++ {
		v := tm.Get(r, c)
		v2 := tm.Get(tr, c)
		tm.Save(tr, c, v)
		tm.Save(r, c, v2)
	}
}

// ExchangeCol 交换列
func (tm *TransformMatrix[K, V]) ExchangeCol(c K, tc K) {
	if c > tm.Col() || tc > tm.Col() {
		log.Fatalln("操作列数异常")
	}
	for r := K(1); r <= tm.Row(); r++ {
		v := tm.Get(r, c)
		v2 := tm.Get(r, tc)
		tm.Save(r, tc, v)
		tm.Save(r, c, v2)
	}
}

// MultiplierRow 初等变换乘数行
func (tm *TransformMatrix[K, V]) MultiplierRow(p float64, r K) {
	if r > tm.Row() {
		log.Fatalln("操作行数异常")
	}
	for c := K(1); c <= tm.Col(); c++ {
		v := tm.Get(r, c)
		tm.Save(r, c, V(p*float64(v)))
	}
}
