package vector

import (
	"fmt"
	"github.com/kordar/gonum/src/common"
	"log"
)

type BaseVector[K common.Key, V common.Value] struct {
	row  K
	col  K
	data map[K]V
}

func (vec *BaseVector[K, V]) Row() K {
	return vec.row
}

func (vec *BaseVector[K, V]) Col() K {
	return vec.col
}

func NewBaseVector[K common.Key, V common.Value](row K, col K, data map[K]V) *BaseVector[K, V] {
	return &BaseVector[K, V]{data: data, row: row, col: col}
}

func (vec *BaseVector[K, V]) SetData(data map[K]V) {
	vec.data = data
}

func (vec *BaseVector[K, V]) S(i K, v V, ckOutBound bool) {
	if ckOutBound {
		common.OutOfBoundsMax(i, vec.Size())
	}
	if v == V(0) {
		delete(vec.data, i)
	} else {
		vec.data[i] = v
	}
}

func (vec *BaseVector[K, V]) Size() K {
	return common.Max(vec.row, vec.col)
}

func (vec *BaseVector[K, V]) Data() map[K]V {
	return vec.data
}

func (vec *BaseVector[K, V]) Similar(t *BaseVector[K, V]) bool {
	return Similar(vec, t)
}

func (vec *BaseVector[K, V]) Equal(t *BaseVector[K, V]) bool {
	return Equal[K, V](vec, t)
}

// IsZero 是否零向量
func (vec *BaseVector[K, V]) IsZero() bool {
	return IsZero[K, V](vec)
}

/****************************
* 向量操作
*****************************/

func (vec *BaseVector[K, V]) Get(i K) V {
	if i > vec.Size() {
		log.Panicln(fmt.Sprintf("out of bounds, max index = %d", vec.Size()))
	}
	return vec.data[i]
}

// Plus 向量相加
func (vec *BaseVector[K, V]) Plus(t *BaseVector[K, V]) {
	if !vec.Similar(t) {
		log.Panicln("dissimilar vector cannot be added")
	}
	if t.IsZero() {
		return
	}
	targetData := t.Data()
	for i, value := range targetData {
		v := vec.data[i] + value
		vec.S(i, v, false)
	}
}

// Minus 向量相减
func (vec *BaseVector[K, V]) Minus(t *BaseVector[K, V]) {
	if !vec.Similar(t) {
		log.Panicln("dissimilar vector cannot be minus")
	}
	if t.IsZero() {
		return
	}
	targetData := t.Data()
	for i, value := range targetData {
		v := vec.data[i] - value
		vec.S(i, v, false)
	}
}

// Multiplier 向量数乘
func (vec *BaseVector[K, V]) Multiplier(v float64) {
	if V(v) == V(0) {
		vec.data = map[K]V{}
		return
	}
	for i, value := range vec.data {
		vec.data[i] = V(v * float64(value))
	}
}

func (vec *BaseVector[K, V]) Clone() *BaseVector[K, V] {
	return NewBaseVector[K, V](vec.Row(), vec.Col(), vec.Data())
}
